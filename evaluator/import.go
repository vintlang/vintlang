package evaluator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ekilie/vint-lang/ast"
	"github.com/ekilie/vint-lang/lexer"
	"github.com/ekilie/vint-lang/module"
	"github.com/ekilie/vint-lang/object"
	"github.com/ekilie/vint-lang/parser"
)

var searchPaths []string

func evalImport(node *ast.Import, env *object.Environment) object.Object {
	for alias, modName := range node.Identifiers {
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
	addSearchPath("")
	addSearchPath("./modules")

	filename := findFile(name)
	if filename == "" {
		return newError("Module '%s' not found. Searched paths: %s", name, strings.Join(searchPaths, ", "))
	}

	scope, err := evaluateFile(filename, env)
	if err != nil {
		return newError("Error evaluating module '%s': %s", name, err)
	}

	return importFile(name, ident, env, scope)
}

func addSearchPath(path string) {
	searchPaths = append(searchPaths, path)
}

func findFile(name string) string {
	basename := fmt.Sprintf("%s.vint", name)
	for _, path := range searchPaths {
		file := filepath.Join(path, basename)
		if fileExists(file) {
			return file
		}
	}
	return ""
}

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func evaluateFile(file string, env *object.Environment) (*object.Environment, object.Object) {
	source, err := os.ReadFile(file)
	if err != nil {
		return nil, newError("Failed to open file '%s': %s", file, err.Error())
	}

	l := lexer.New(string(source))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		return nil, newError("Syntax errors in file '%s':\n%s", file, strings.Join(p.Errors(), "\n"))
	}

	scope := object.NewEnvironment()
	result := Eval(program, scope)

	if isError(result) {
		return nil, newError("Runtime error in file '%s': %s", file, result.Inspect())
	}

	return scope, nil
}

func importFile(name string, ident *ast.Identifier, env *object.Environment, scope *object.Environment) object.Object {
	value, found := scope.Get(ident.Value)
	if !found {
		return newError("Identifier '%s' not found in module '%s'", ident.Value, name)
	}
	env.Set(name, value)
	return NULL
}
