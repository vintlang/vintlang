package evaluator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/module"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
)

var searchPaths []string

// Error messages
const (
	ErrModuleNotFound     = "Module '%s' not found. Searched in:\n%s"
	ErrImportFailed       = "Failed to import module '%s': %s"
	ErrFileNotFound       = "File '%s' not found in any of the search paths"
	ErrFileReadFailed     = "Failed to read file '%s': %s"
	ErrSyntaxError        = "Syntax error in file '%s':\n%s"
	ErrRuntimeError       = "Runtime error in file '%s': %s"
	ErrIdentifierNotFound = "Identifier '%s' not found in module '%s'"
	ErrInvalidModule      = "Invalid module name '%s'. Module names must be valid identifiers"
	ErrCircularImport     = "Circular import detected: '%s' is already being imported"
)

// Track imported modules to detect circular imports
var importedModules = make(map[string]bool)

func evalImport(node *ast.Import, env *object.Environment) object.Object {
	// Reset imported modules for new import chain
	importedModules = make(map[string]bool)

	for alias, modName := range node.Identifiers {
		// Validate module name
		if !isValidModuleName(modName.Value) {
			return newError(ErrInvalidModule, modName.Value)
		}

		// Check for circular imports
		if importedModules[modName.Value] {
			return newError(ErrCircularImport, modName.Value)
		}
		importedModules[modName.Value] = true

		if mod, exists := module.Mapper[modName.Value]; exists {
			env.Set(alias, mod)
		} else {
			result := evalImportFile(alias, modName, env)
			if isError(result) {
				return result
			}
		}
	}
	return NULL
}

func evalImportFile(name string, ident *ast.Identifier, env *object.Environment) object.Object {
	// Add default search paths
	addSearchPath("")
	addSearchPath("./modules")
	addSearchPath("./vintLang/modules")

	filename := findFile(name)
	if filename == "" {
		// Format search paths for better readability
		formattedPaths := formatSearchPaths()
		return newError(ErrModuleNotFound, name, formattedPaths)
	}

	scope, err := evaluateFile(filename, env)
	if err != nil {
		return newError(ErrImportFailed, name, err.Inspect())
	}

	return importFile(name, ident, env, scope)
}

func addSearchPath(path string) {
	// Convert to absolute path if relative
	if !filepath.IsAbs(path) {
		absPath, err := filepath.Abs(path)
		if err == nil {
			path = absPath
		}
	}

	// Only add if not already in search paths
	for _, existingPath := range searchPaths {
		if existingPath == path {
			return
		}
	}
	searchPaths = append(searchPaths, path)
}

func findFile(name string) string {
	// Try different file extensions
	extensions := []string{".vint", ".VINT", ".Vint"}
	basename := name

	for _, ext := range extensions {
		filename := basename + ext
		for _, path := range searchPaths {
			file := filepath.Join(path, filename)
			if fileExists(file) {
				return file
			}
		}
	}
	return ""
}

func fileExists(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func evaluateFile(file string, env *object.Environment) (*object.Environment, object.Object) {
	source, err := os.ReadFile(file)
	if err != nil {
		return nil, newError(ErrFileReadFailed, file, err.Error())
	}

	l := lexer.New(string(source))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		// Format syntax errors for better readability
		formattedErrors := formatSyntaxErrors(p.Errors())
		return nil, newError(ErrSyntaxError, file, formattedErrors)
	}

	scope := object.NewEnvironment()
	result := Eval(program, scope)

	if isError(result) {
		return nil, newError(ErrRuntimeError, file, result.Inspect())
	}

	return scope, nil
}

func importFile(name string, ident *ast.Identifier, env *object.Environment, scope *object.Environment) object.Object {
	value, found := scope.Get(ident.Value)
	if !found {
		return newError(ErrIdentifierNotFound, ident.Value, name)
	}
	env.Set(name, value)
	return NULL
}

// Helper functions for better error formatting

func formatSearchPaths() string {
	var paths []string
	for i, path := range searchPaths {
		paths = append(paths, fmt.Sprintf("  %d. %s", i+1, path))
	}
	return strings.Join(paths, "\n")
}

func formatSyntaxErrors(errors []string) string {
	var formatted []string
	for i, err := range errors {
		formatted = append(formatted, fmt.Sprintf("  %d. %s", i+1, err))
	}
	return strings.Join(formatted, "\n")
}

func isValidModuleName(name string) bool {
	// Check if the module name is a valid identifier
	if len(name) == 0 {
		return false
	}

	// First character must be a letter or underscore
	first := name[0]
	if !((first >= 'a' && first <= 'z') || (first >= 'A' && first <= 'Z') || first == '_') {
		return false
	}

	// Rest of the characters must be letters, numbers, or underscores
	for i := 1; i < len(name); i++ {
		c := name[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
			return false
		}
	}

	return true
}
