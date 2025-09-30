package evaluator

import (
	builtinpkg "github.com/vintlang/vintlang/evaluator/builtins"
	"github.com/vintlang/vintlang/object"
)

// GetBuiltinFunction returns a builtin function by name
func GetBuiltinFunction(name string) (*object.Builtin, bool) {
	return builtinpkg.GetBuiltin(name)
}

// GetAllBuiltinFunctions returns all registered builtin functions
func GetAllBuiltinFunctions() map[string]*object.Builtin {
	return builtinpkg.GetAllBuiltins()
}

// We initialize builtins variable to maintain compatibility
// This will be populated once all the builtin packages are imported
func init() {
	// The builtins will be automatically registered by importing the subpackages
}

// Compatibility map - this gets populated by the registry
var builtins = builtinpkg.BuiltinRegistry
