package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalPostfixExpression(env *object.Environment, operator string, node *ast.PostfixExpression) object.VintObject {
	val, ok := env.Get(node.Token.Literal)
	if !ok {
		return newError("Use a NUMBER or DECIMAL IDENTIFIER, not %s", node.Token.Type)
	}

	assign := func(val object.VintObject) object.VintObject {
		newVal, ok := env.Assign(node.Token.Literal, val)
		if !ok {
			return newError("assignment to undeclared variable '%s'", node.Token.Literal)
		}
		return newVal
	}

	switch operator {
	case "++":
		switch arg := val.(type) {
		case *object.Integer:
			v := arg.Value + 1
			return assign(&object.Integer{Value: v})
		case *object.Float:
			v := arg.Value + 1
			return assign(&object.Float{Value: v})
		default:
			return newError("Line %d: %s is not a numeric identifier. Use '++' with a number or decimal identifier.\nExample:\tlet i = 2; i++", node.Token.Line, node.Token.Literal)
		}
	case "--":
		switch arg := val.(type) {
		case *object.Integer:
			v := arg.Value - 1
			return assign(&object.Integer{Value: v})
		case *object.Float:
			v := arg.Value - 1
			return assign(&object.Float{Value: v})
		default:
			return newError("Line %d: %s is not a numeric identifier. Use '--' with a number or decimal identifier.\nExample:\tlet i = 2; i--", node.Token.Line, node.Token.Literal)
		}
	default:
		return newError("Unknown operator: %s", operator)
	}
}
