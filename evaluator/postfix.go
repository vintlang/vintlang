package evaluator

import (
	"github.com/ekilie/vint-lang/ast"
	"github.com/ekilie/vint-lang/object"
)

func evalPostfixExpression(env *object.Environment, operator string, node *ast.PostfixExpression) object.Object {
	val, ok := env.Get(node.Token.Literal)
	if !ok {
		return newError("Use a NUMBER or DECIMAL IDENTIFIER, not %s", node.Token.Type)
	}
	switch operator {
	case "++":
		switch arg := val.(type) {
		case *object.Integer:
			v := arg.Value + 1
			return env.Set(node.Token.Literal, &object.Integer{Value: v})
		case *object.Float:
			v := arg.Value + 1
			return env.Set(node.Token.Literal, &object.Float{Value: v})
		default:
			return newError("Line %d: %s is not a numeric identifier. Use '++' with a number or decimal identifier.\nExample:\tlet i = 2; i++", node.Token.Line, node.Token.Literal)
		}
	case "--":
		switch arg := val.(type) {
		case *object.Integer:
			v := arg.Value - 1
			return env.Set(node.Token.Literal, &object.Integer{Value: v})
		case *object.Float:
			v := arg.Value - 1
			return env.Set(node.Token.Literal, &object.Float{Value: v})
		default:
			return newError("Line %d: %s is not a numeric identifier. Use '--' with a number or decimal identifier.\nExample:\tlet i = 2; i--", node.Token.Line, node.Token.Literal)
		}
	default:
		return newError("Unknown operator: %s", operator)
	}
}
