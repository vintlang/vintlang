package builtins

import "github.com/vintlang/vintlang/object"

// BuiltinRegistry holds all registered builtin functions
var BuiltinRegistry = make(map[string]*object.Builtin)

// RegisterBuiltin registers a builtin function
func RegisterBuiltin(name string, builtin *object.Builtin) {
	BuiltinRegistry[name] = builtin
}

// GetAllBuiltins returns the complete map of builtin functions
func GetAllBuiltins() map[string]*object.Builtin {
	return BuiltinRegistry
}

// GetBuiltin returns a specific builtin function by name
func GetBuiltin(name string) (*object.Builtin, bool) {
	builtin, exists := BuiltinRegistry[name]
	return builtin, exists
}
