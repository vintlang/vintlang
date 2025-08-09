package bundler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/module"
	"github.com/vintlang/vintlang/parser"
)

// FileBundle represents a collection of bundled files
type FileBundle struct {
	MainFile string
	Files    map[string]string // filename -> content
}

// DependencyAnalyzer analyzes and collects all dependent files for bundling
type DependencyAnalyzer struct {
	processed   map[string]bool
	searchPaths []string
	bundle      *FileBundle
}

// NewDependencyAnalyzer creates a new dependency analyzer
func NewDependencyAnalyzer() *DependencyAnalyzer {
	return &DependencyAnalyzer{
		processed:   make(map[string]bool),
		searchPaths: []string{},
		bundle: &FileBundle{
			Files: make(map[string]string),
		},
	}
}

// AnalyzeDependencies analyzes a main file and returns all its dependencies
func (da *DependencyAnalyzer) AnalyzeDependencies(mainFile string) (*FileBundle, error) {
	// Set up search paths
	da.setupSearchPaths(mainFile)
	
	// Resolve absolute path for main file
	absMainFile, err := filepath.Abs(mainFile)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve absolute path for main file: %w", err)
	}
	
	da.bundle.MainFile = absMainFile
	
	// Process the main file and its dependencies recursively
	err = da.processFile(absMainFile)
	if err != nil {
		return nil, fmt.Errorf("failed to process dependencies: %w", err)
	}
	
	return da.bundle, nil
}

// setupSearchPaths sets up the search paths for finding imported files
func (da *DependencyAnalyzer) setupSearchPaths(mainFile string) {
	// Add the directory containing the main file
	mainDir := filepath.Dir(mainFile)
	da.addSearchPath(mainDir)
	
	// Add current directory
	da.addSearchPath(".")
	
	// Add modules directory if it exists
	modulesPath := "./modules"
	if da.dirExists(modulesPath) {
		da.addSearchPath(modulesPath)
	}
	
	// Add modules directory relative to main file
	relativeModulesPath := filepath.Join(mainDir, "modules")
	if da.dirExists(relativeModulesPath) {
		da.addSearchPath(relativeModulesPath)
	}
}

// addSearchPath adds a path to the search paths
func (da *DependencyAnalyzer) addSearchPath(path string) {
	absPath, err := filepath.Abs(path)
	if err == nil {
		path = absPath
	}
	
	// Only add if not already in search paths
	for _, existingPath := range da.searchPaths {
		if existingPath == path {
			return
		}
	}
	da.searchPaths = append(da.searchPaths, path)
}

// processFile processes a single file and its imports
func (da *DependencyAnalyzer) processFile(filename string) error {
	// Skip if already processed
	if da.processed[filename] {
		return nil
	}
	
	// Read the file
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file '%s': %w", filename, err)
	}
	
	// Parse the file to find imports
	imports, err := da.extractImports(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse file '%s': %w", filename, err)
	}
	
	// Add to bundle
	da.bundle.Files[filename] = string(content)
	da.processed[filename] = true
	
	// Process each import
	for _, importName := range imports {
		// Skip built-in modules
		if da.isBuiltinModule(importName) {
			continue
		}
		
		// Find the imported file
		importedFile := da.findFile(importName)
		if importedFile == "" {
			// This is not necessarily an error - the file might be a built-in module
			// or might be optional. We'll let the runtime handle missing files.
			continue
		}
		
		// Recursively process the imported file
		err = da.processFile(importedFile)
		if err != nil {
			return err
		}
	}
	
	return nil
}

// extractImports parses a vint file and returns the list of imported modules
func (da *DependencyAnalyzer) extractImports(content string) ([]string, error) {
	l := lexer.New(content)
	p := parser.New(l)
	program := p.ParseProgram()
	
	if len(p.Errors()) > 0 {
		return nil, fmt.Errorf("parse errors: %v", p.Errors())
	}
	
	var imports []string
	
	// Walk the AST to find import statements
	for _, stmt := range program.Statements {
		da.findImportsInNode(stmt, &imports)
	}
	
	return imports, nil
}

// findImportsInNode recursively finds import statements in AST nodes
func (da *DependencyAnalyzer) findImportsInNode(node ast.Node, imports *[]string) {
	switch n := node.(type) {
	case *ast.Import:
		// Extract module names from the import
		for _, ident := range n.Identifiers {
			*imports = append(*imports, ident.Value)
		}
	case *ast.ExpressionStatement:
		// Import statements might be wrapped in expression statements
		da.findImportsInNode(n.Expression, imports)
	case *ast.Program:
		for _, stmt := range n.Statements {
			da.findImportsInNode(stmt, imports)
		}
	case *ast.BlockStatement:
		for _, stmt := range n.Statements {
			da.findImportsInNode(stmt, imports)
		}
	// Add more cases as needed for other node types that might contain imports
	}
}

// isBuiltinModule checks if a module is a built-in module
func (da *DependencyAnalyzer) isBuiltinModule(moduleName string) bool {
	_, exists := module.Mapper[moduleName]
	return exists
}

// findFile finds a file in the search paths
func (da *DependencyAnalyzer) findFile(name string) string {
	extensions := []string{".vint", ".VINT", ".Vint"}
	basename := name
	
	for _, ext := range extensions {
		filename := basename + ext
		for _, path := range da.searchPaths {
			file := filepath.Join(path, filename)
			if da.fileExists(file) {
				absFile, err := filepath.Abs(file)
				if err == nil {
					return absFile
				}
				return file
			}
		}
	}
	return ""
}

// fileExists checks if a file exists
func (da *DependencyAnalyzer) fileExists(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// dirExists checks if a directory exists
func (da *DependencyAnalyzer) dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}