package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

// evalEnumStatement evaluates enum declarations
func evalEnumStatement(node *ast.EnumStatement, env *object.Environment) object.VintObject {
	// Create the enum object
	enum := &object.Enum{
		Name:    node.Name.Value,
		Members: make(map[string]object.VintObject),
	}

	// Evaluate each member's value
	for memberName, valueExpr := range node.Values {
		value := Eval(valueExpr, env)
		if isError(value) {
			return value
		}

		// Store the member value
		enum.Members[memberName] = value
	}

	// Define the enum in the environment as a constant (immutable)
	return env.DefineConst(node.Name.Value, enum)
}
