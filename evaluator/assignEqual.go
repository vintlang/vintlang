package evaluator

import (
	"strings"

	"github.com/ekilie/vint-lang/ast"
	"github.com/ekilie/vint-lang/object"
)

func evalAssignEqual(node *ast.AssignEqual, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}

	value := Eval(node.Value, env)
	if isError(value) {
		return value
	}

	switch node.Token.Literal {
	case "+=":
		switch arg := left.(type) {
		case *object.Integer:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value + val.Value
				return env.Set(node.Left.Token.Literal, &object.Integer{Value: v})
			case *object.Float:
				v := float64(arg.Value) + val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '+=' to add %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.Float:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value + float64(val.Value)
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			case *object.Float:
				v := arg.Value + val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '+=' to add %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.String:
			switch val := value.(type) {
			case *object.String:
				v := arg.Value + val.Value
				return env.Set(node.Left.Token.Literal, &object.String{Value: v})
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
				return env.Set(node.Left.Token.Literal, &object.Integer{Value: v})
			case *object.Float:
				v := float64(arg.Value) - val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '-=' to subtract %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.Float:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value - float64(val.Value)
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			case *object.Float:
				v := arg.Value - val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
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
				return env.Set(node.Left.Token.Literal, &object.Integer{Value: v})
			case *object.Float:
				v := float64(arg.Value) * val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			case *object.String:
				v := strings.Repeat(val.Value, int(arg.Value))
				return env.Set(node.Left.Token.Literal, &object.String{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '*=' to multiply %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.Float:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value * float64(val.Value)
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			case *object.Float:
				v := arg.Value * val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '*=' to multiply %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.String:
			switch val := value.(type) {
			case *object.Integer:
				v := strings.Repeat(arg.Value, int(val.Value))
				return env.Set(node.Left.Token.Literal, &object.String{Value: v})
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
				return env.Set(node.Left.Token.Literal, &object.Integer{Value: v})
			case *object.Float:
				v := float64(arg.Value) / val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			default:
				// Check for invalid operation between different types
				return newError("Line %d: Cannot use '/=' to divide %v and %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.Float:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value / float64(val.Value)
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			case *object.Float:
				v := arg.Value / val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
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
