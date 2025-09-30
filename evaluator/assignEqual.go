package evaluator

import (
	"strings"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalAssignEqual(node *ast.AssignEqual, env *object.Environment) object.VintObject {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}

	value := Eval(node.Value, env)
	if isError(value) {
		return value
	}

	assign := func(val object.VintObject) object.VintObject {
		newVal, ok := env.Assign(node.Left.Token.Literal, val)
		if !ok {
			return newError("assignment to undeclared variable '%s'", node.Left.Token.Literal)
		}
		return newVal
	}

	switch node.Token.Literal {
	case "+=":
		switch arg := left.(type) {
		case *object.Integer:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value + val.Value
				return assign(&object.Integer{Value: v})
			case *object.Float:
				v := float64(arg.Value) + val.Value
				return assign(&object.Float{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '+=' to add %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.Float:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value + float64(val.Value)
				return assign(&object.Float{Value: v})
			case *object.Float:
				v := arg.Value + val.Value
				return assign(&object.Float{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '+=' to add %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.String:
			switch val := value.(type) {
			case *object.String:
				v := arg.Value + val.Value
				return assign(&object.String{Value: v})
			default:
				// Check for invalid operation for non-strings
				return newError("Line %d: Cannot use '+=' with %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		default:
			// Check for invalid operation on unsupported types
			return newError("Line %d: Cannot use '+=' with %v", node.Token.Line, arg.Type())
		}
	case "-=":
		switch arg := left.(type) {
		case *object.Integer:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value - val.Value
				return assign(&object.Integer{Value: v})
			case *object.Float:
				v := float64(arg.Value) - val.Value
				return assign(&object.Float{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '-=' to subtract %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.Float:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value - float64(val.Value)
				return assign(&object.Float{Value: v})
			case *object.Float:
				v := arg.Value - val.Value
				return assign(&object.Float{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '-=' to subtract %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		default:
			// Check for invalid operation on unsupported types
			return newError("Line %d: Cannot use '-=' with %v", node.Token.Line, arg.Type())
		}
	case "*=":
		switch arg := left.(type) {
		case *object.Integer:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value * val.Value
				return assign(&object.Integer{Value: v})
			case *object.Float:
				v := float64(arg.Value) * val.Value
				return assign(&object.Float{Value: v})
			case *object.String:
				v := strings.Repeat(val.Value, int(arg.Value))
				return assign(&object.String{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '*=' to multiply %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.Float:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value * float64(val.Value)
				return assign(&object.Float{Value: v})
			case *object.Float:
				v := arg.Value * val.Value
				return assign(&object.Float{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '*=' to multiply %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.String:
			switch val := value.(type) {
			case *object.Integer:
				v := strings.Repeat(arg.Value, int(val.Value))
				return assign(&object.String{Value: v})
			default:
				// Check for invalid operation for non-integer multiplications
				return newError("Line %d: Cannot use '*=' with %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		default:
			// Check for invalid operation on unsupported types
			return newError("Line %d: Cannot use '*=' with %v", node.Token.Line, arg.Type())
		}
	case "/=":
		switch arg := left.(type) {
		case *object.Integer:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value / val.Value
				return assign(&object.Integer{Value: v})
			case *object.Float:
				v := float64(arg.Value) / val.Value
				return assign(&object.Float{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '/=' to divide %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.Float:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value / float64(val.Value)
				return assign(&object.Float{Value: v})
			case *object.Float:
				v := arg.Value / val.Value
				return assign(&object.Float{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '/=' to divide %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		default:
			// Check for invalid operation on unsupported types
			return newError("Line %d: Cannot use '/=' with %v", node.Token.Line, arg.Type())
		}
	default:
		// Check for an unknown operation
		return newError("Line %d: Unknown operation %s", node.Token.Line, node.Token.Literal)
	}
}
