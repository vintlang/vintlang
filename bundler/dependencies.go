package bundler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/module"
	"github.com/vintlang/vintlang/parser"
)

// FileBundle represents a collection of bundled files
type FileBundle struct {
	MainFile     string
	Files        map[string]string // filename -> content
	IncludeFiles map[string]bool   // filename -> true if included via include statement
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
			Files:        make(map[string]string),
			IncludeFiles: make(map[string]bool),
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

// processFile processes a single file and its imports/includes
func (da *DependencyAnalyzer) processFile(filename string) error {
	return da.processFileWithType(filename, false)
}

// processFileWithType processes a single file and tracks if it's an include
func (da *DependencyAnalyzer) processFileWithType(filename string, isInclude bool) error {
	// Skip if already processed
	if da.processed[filename] {
		return nil
	}
	
	// Read the file
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file '%s': %w", filename, err)
	}
	
	// Parse the file to find imports and includes
	imports, includes, err := da.extractImportsAndIncludes(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse file '%s': %w", filename, err)
	}
	
	// Add to bundle
	da.bundle.Files[filename] = string(content)
	if isInclude {
		da.bundle.IncludeFiles[filename] = true
	}
	da.processed[filename] = true
	
	// Process each import (module-based)
	for _, importName := range imports {
		// Skip built-in modules
		if da.isBuiltinModule(importName) {
			continue
		}
		
		// Find the imported file
		importedFile := da.findModuleFile(importName)
		if importedFile == "" {
			// This is not necessarily an error - the file might be a built-in module
			// or might be optional. We'll let the runtime handle missing files.
			continue
		}
		
		// Recursively process the imported file (not an include)
		err = da.processFileWithType(importedFile, false)
		if err != nil {
			return err
		}
	}
	
	// Process each include (file path-based)
	for _, includePath := range includes {
		// Find the included file
		includedFile := da.findIncludeFile(includePath)
		if includedFile == "" {
			// Include files should exist, so this is more likely to be an error
			continue
		}
		
		// Recursively process the included file (marked as include)
		err = da.processFileWithType(includedFile, true)
		if err != nil {
			return err
		}
	}
	
	return nil
}

// extractImportsAndIncludes parses a vint file and returns separate lists of imports and includes
func (da *DependencyAnalyzer) extractImportsAndIncludes(content string) ([]string, []string, error) {
	l := lexer.New(content)
	p := parser.New(l)
	program := p.ParseProgram()
	
	if len(p.Errors()) > 0 {
		return nil, nil, fmt.Errorf("parse errors: %v", p.Errors())
	}
	
	var imports []string
	var includes []string
	
	// Walk the AST to find import and include statements
	for _, stmt := range program.Statements {
		da.findImportsAndIncludesInNode(stmt, &imports, &includes)
	}
	
	return imports, includes, nil
}

// findImportsAndIncludesInNode recursively finds import and include statements in AST nodes
func (da *DependencyAnalyzer) findImportsAndIncludesInNode(node ast.Node, imports *[]string, includes *[]string) {
	switch n := node.(type) {
	case *ast.Import:
		// Extract module names from the import
		for _, ident := range n.Identifiers {
			*imports = append(*imports, ident.Value)
		}
	case *ast.IncludeStatement:
		// Extract file path from include statement
		if stringLit, ok := n.Path.(*ast.StringLiteral); ok {
			*includes = append(*includes, stringLit.Value)
		}
	case *ast.ExpressionStatement:
		// Import statements might be wrapped in expression statements
		da.findImportsAndIncludesInNode(n.Expression, imports, includes)
	case *ast.Program:
		for _, stmt := range n.Statements {
			da.findImportsAndIncludesInNode(stmt, imports, includes)
		}
	case *ast.BlockStatement:
		for _, stmt := range n.Statements {
			da.findImportsAndIncludesInNode(stmt, imports, includes)
		}
	// Add more cases as needed for other node types that might contain imports/includes
	}
}

// isBuiltinModule checks if a module is a built-in module
func (da *DependencyAnalyzer) isBuiltinModule(moduleName string) bool {
	_, exists := module.Mapper[moduleName]
	return exists
}

// findFile finds a file in the search paths
func (da *DependencyAnalyzer) findFile(name string) string {
	// Check if this looks like a file path (for include statements)
	if da.isFilePath(name) {
		return da.findIncludeFile(name)
	}
	
	// Handle as module name (for import statements)
	return da.findModuleFile(name)
}

// isFilePath checks if a name looks like a file path rather than a module name
func (da *DependencyAnalyzer) isFilePath(name string) bool {
	// If it contains path separators or file extensions, treat as file path
	return strings.Contains(name, "/") || strings.Contains(name, "\\") || strings.Contains(name, ".")
}

// findIncludeFile finds a file by its path for include statements
func (da *DependencyAnalyzer) findIncludeFile(path string) string {
	// For include statements, the path might be relative to the main file
	if !filepath.IsAbs(path) {
		// Try relative to main file directory first
		mainDir := filepath.Dir(da.bundle.MainFile)
		candidatePath := filepath.Join(mainDir, path)
		if da.fileExists(candidatePath) {
			absFile, err := filepath.Abs(candidatePath)
			if err == nil {
				return absFile
			}
			return candidatePath
		}
		
		// Try relative to each search path
		for _, searchPath := range da.searchPaths {
			candidatePath := filepath.Join(searchPath, path)
			if da.fileExists(candidatePath) {
				absFile, err := filepath.Abs(candidatePath)
				if err == nil {
					return absFile
				}
				return candidatePath
			}
		}
	} else {
		// Absolute path, check if it exists
		if da.fileExists(path) {
			return path
		}
	}
	
	return ""
}

// findModuleFile finds a file by module name for import statements
func (da *DependencyAnalyzer) findModuleFile(name string) string {
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