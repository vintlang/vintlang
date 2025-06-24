package evaluator

import "github.com/vintlang/vintlang/object"

func evalMinusPrefixOperatorExpression(right object.Object, line int) object.Object {
	switch obj := right.(type) {

	case *object.Integer:
		return &object.Integer{Value: -obj.Value}

	case *object.Float:
		return &object.Float{Value: -obj.Value}

	default:
		return newError("Line %d: Unknown operation: -%s", line, right.Type())
	}
}
func evalPlusPrefixOperatorExpression(right object.Object, line int) object.Object {
	switch obj := right.(type) {

	case *object.Integer:
		return &object.Integer{Value: obj.Value}

	case *object.Float:
		return &object.Float{Value: obj.Value}

	default:
		return newError("Line %d: Unknown operation: +%s", line, right.Type())
	}
}

func evalPrefixExpression(operator string, right object.Object, line int) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right, line)
	case "+":
		return evalPlusPrefixOperatorExpression(right, line)
	case "*":
		// derefences a pointer to get the value
		if p, ok := right.(*object.Pointer); ok {
			return p.Ref
		}
		return newError("Line %d: cannot dereference non-pointer", line)
	case "&":
		// Creates a new Pointer object that points to the value
		return &object.Pointer{Ref: right}
	default:
		return newError("Line %d: Unknown operation: %s%s", line, operator, right.Type())
	}
}
