package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	// Checks if the identifier exists in the environment, returns its value if found.
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	// Checks if the identifier is a built-in function, returns the built-in function if found.
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	// Returns an error if the identifier is not found in the environment or built-ins.
	return newError("Line %d: Identifier not recognized: %s", node.Token.Line, node.Value)
}
