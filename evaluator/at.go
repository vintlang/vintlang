package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

// evalAt handles the evaluation of the "@" expression.
// It checks if the "@" symbol is available in the current environment and returns its value.
// If not found, it returns an error indicating that the scope is invalid.
func evalAt(node *ast.At, env *object.Environment) object.Object {
	// Checks if "@" exists in the environment
	if at, ok := env.Get("@"); ok {
		return at // Return the value associated with "@"
	}
	// Returns an error if "@" is not found in the scope
	return newError("Out of scope")
}
