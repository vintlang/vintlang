package bundler

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/module"
	"github.com/vintlang/vintlang/parser"
)

// PackageProcessor handles processing packages from dependency files
type PackageProcessor struct {
	bundle *FileBundle
}

// ProcessedBundle contains the processed code ready for bundling
type ProcessedBundle struct {
	ProcessedCode string
	Packages      map[string]string // package name -> package definition
}

// NewPackageProcessor creates a new package processor
func NewPackageProcessor(bundle *FileBundle) *PackageProcessor {
	return &PackageProcessor{
		bundle: bundle,
	}
}

// ProcessBundle processes the bundle to extract packages and modify imports
func (pp *PackageProcessor) ProcessBundle() (*ProcessedBundle, error) {
	processed := &ProcessedBundle{
		Packages: make(map[string]string),
	}

	// First, extract packages from dependency files
	for filename, content := range pp.bundle.Files {
		if filename != pp.bundle.MainFile {
			packageDef, err := pp.extractPackageDefinition(content, filename)
			if err != nil {
				return nil, fmt.Errorf("failed to process file %s: %w", filename, err)
			}
			if packageDef != "" {
				basename := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
				processed.Packages[basename] = packageDef
			}
		}
	}

	// Then, process the main file
	mainContent, exists := pp.bundle.Files[pp.bundle.MainFile]
	if !exists {
		return nil, fmt.Errorf("main file not found in bundle")
	}

	processedMainContent, err := pp.processMainFile(mainContent, processed.Packages)
	if err != nil {
		return nil, fmt.Errorf("failed to process main file: %w", err)
	}

	processed.ProcessedCode = processedMainContent
	return processed, nil
}

// extractPackageDefinition extracts package definition from a file content
func (pp *PackageProcessor) extractPackageDefinition(content, filename string) (string, error) {
	l := lexer.New(content)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		return "", fmt.Errorf("parse errors in %s: %v", filename, p.Errors())
	}

	// Look for package statements and extract their content
	for _, stmt := range program.Statements {
		if exprStmt, ok := stmt.(*ast.ExpressionStatement); ok {
			if pkgExpr, ok := exprStmt.Expression.(*ast.Package); ok {
				return pkgExpr.String(), nil
			}
		}
	}

	// If no package found, return the entire content as a package
	// This handles files that don't have explicit package declarations
	basename := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	return fmt.Sprintf("package %s {\n%s\n}", basename, content), nil
}

// processMainFile processes the main file to replace imports with direct package definitions
func (pp *PackageProcessor) processMainFile(mainContent string, packages map[string]string) (string, error) {
	l := lexer.New(mainContent)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		return "", fmt.Errorf("parse errors in main file: %v", p.Errors())
	}

	var result strings.Builder

	// First, add all package definitions
	for _, packageDef := range packages {
		result.WriteString(packageDef)
		result.WriteString("\n\n")
	}

	// Then, process the main file content, skipping imports that we've already embedded
	for _, stmt := range program.Statements {
		// Check if this is an import statement
		if exprStmt, ok := stmt.(*ast.ExpressionStatement); ok {
			if importExpr, ok := exprStmt.Expression.(*ast.Import); ok {
				// Check if any of the imported modules are in our packages or are built-in
				var identifiersToKeep []*ast.Identifier
				for _, ident := range importExpr.Identifiers {
					// if the module is a custom module, we skip it
					if _, exists := packages[ident.Value]; exists {
						continue
					}
					// if the module is a built-in module, we keep it
					if _, exists := module.Mapper[ident.Value]; exists {
						identifiersToKeep = append(identifiersToKeep, ident)
						continue
					}
					// otherwise, we keep it and let the runtime handle it
					identifiersToKeep = append(identifiersToKeep, ident)
				}

				if len(identifiersToKeep) == 0 {
					continue // Skip this import statement entirely
				}

				// Rebuild the import statement with only the identifiers to keep
				identifiersMap := make(map[string]*ast.Identifier)
				for _, ident := range identifiersToKeep {
					identifiersMap[ident.Value] = ident
				}
				newImportExpr := &ast.Import{
					Token:       importExpr.Token,
					Identifiers: identifiersMap,
				}
				result.WriteString(newImportExpr.String())
				result.WriteString("\n")
				continue
			}
		}

		// Add the statement to the result
		result.WriteString(stmt.String())
		result.WriteString("\n")
	}

	return result.String(), nil
}