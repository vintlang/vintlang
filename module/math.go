package module

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/vintlang/vintlang/object"
)

var MathFunctions = map[string]object.ModuleFunction{
	"PI":        pi,
	"e":         e,
	"phi":       phi,
	"ln10":      ln10,
	"ln2":       ln2,
	"log10e":    log10e,
	"log2e":     log2e,
	"log2":      log2,
	"sqrt1_2":   sqrt1_2,
	"sqrt2":     sqrt2,
	"sqrt3":     sqrt3,
	"sqrt5":     sqrt5,
	"EPSILON":   epsilon,
	"abs":       abs,
	"sign":      sign,
	"ceil":      ceil,
	"floor":     floor,
	"sqrt":      sqrt,
	"cbrt":      cbrt,
	"root":      root,
	"hypot":     hypot,
	"random":    random,
	"factorial": factorial,
	"round":     round,
	"max":       max,
	"min":       min,
	"exp":       exp,
	"expm1":     expm1,
	// "log":       log,
	"log10": log10,
	"log1p": log1p,
	"cos":   cos,
	"sin":   sin,
	"tan":   tan,
	"acos":  acos,
	"asin":  asin,
	"atan":  atan,
	"cosh":  cosh,
	"sinh":  sinh,
	"tanh":  tanh,
	"acosh": acosh,
	"asinh": asinh,
	"atanh": atanh,
	"atan2": atan2,
	// Statistics functions
	"mean":     mean,
	"stddev":   stddev,
	"variance": variance,
	"median":   median,
	// Complex number functions
	"complex": complexNum,
	// Big integer functions
	"bigint": bigint,
	// Linear algebra operations
	"dot":       dot,
	"cross":     cross,
	"magnitude": magnitude,
	// Numerical methods
	"gcd":       gcd,
	"lcm":       lcm,
	"clamp":     clamp,
	"lerp":      lerp,
}

var Constants = map[string]object.VintObject{
	"PI":      &object.Float{Value: math.Pi},
	"e":       &object.Float{Value: math.E},
	"phi":     &object.Float{Value: (1 + math.Sqrt(5)) / 2},
	"ln10":    &object.Float{Value: math.Log10E},
	"ln2":     &object.Float{Value: math.Ln2},
	"log10e":  &object.Float{Value: math.Log10E},
	"log2e":   &object.Float{Value: math.Log2E},
	"sqrt1_2": &object.Float{Value: 1 / math.Sqrt2},
	"sqrt2":   &object.Float{Value: math.Sqrt2},
	"sqrt3":   &object.Float{Value: math.Sqrt(3)},
	"sqrt5":   &object.Float{Value: math.Sqrt(5)},
	"EPSILON": &object.Float{Value: 2.220446049250313e-16},
}

func pi(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: math.Pi}
}

func e(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: math.E}
}

func phi(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: (1 + math.Sqrt(5)) / 2}
}

func ln10(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: math.Log10E}
}

func ln2(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: math.Ln2}
}

func log10e(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: math.Log10E}
}

func log2e(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: math.Log2E}
}

func sqrt1_2(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: 1 / math.Sqrt2}
}

func sqrt2(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: math.Sqrt2}
}

func sqrt3(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: math.Sqrt(3)}
}

func sqrt5(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: math.Sqrt(5)}
}

func epsilon(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return &object.Float{Value: 2.220446049250313e-16}
}
func abs(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{
			Message: "\033[1;31m -> math.abs()\033[0m:\n" +
				"  This function does not accept keyword arguments.\n" +
				"  Usage: math.abs(-5) -> 5\n",
		}
	}
	if len(args) != 1 {
		return ErrorMessage(
			"math", "abs",
			"1 numeric argument (number to get absolute value)",
			fmt.Sprintf("%d arguments", len(args)),
			"math.abs(-5) -> 5",
		)
	}

	// Check if it's a complex number (dict with real and imag keys)
	if dict, ok := args[0].(*object.Dict); ok {
		realKey := &object.String{Value: "real"}
		imagKey := &object.String{Value: "imag"}
		
		realPair, hasReal := dict.Pairs[realKey.HashKey()]
		imagPair, hasImag := dict.Pairs[imagKey.HashKey()]
		
		if hasReal && hasImag {
			// It's a complex number, calculate magnitude
			real := extractFloatValue(realPair.Value)
			imag := extractFloatValue(imagPair.Value)
			magnitude := math.Sqrt(real*real + imag*imag)
			return &object.Float{Value: magnitude}
		}
	}

	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return ErrorMessage(
			"math", "abs",
			"numeric argument (integer, float, or complex number)",
			string(args[0].Type()),
			"math.abs(-5) -> 5 or math.abs(complex(3, 4)) -> 5",
		)
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		if arg.Value < 0 {
			return &object.Integer{Value: -arg.Value}
		}
		return arg
	case *object.Float:
		if arg.Value < 0 {
			return &object.Float{Value: -arg.Value}
		}
		return arg
	default:
		return &object.Error{Message: fmt.Sprintf("math.abs() internal error: unexpected argument type %s", args[0].Type())}
	}
}

func sign(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return ErrorMessage(
			"math", "sign",
			"no definitions allowed",
			fmt.Sprintf("%d definitions provided", len(defs)),
			"math.sign(5)",
		)
	}
	if len(args) != 1 {
		return ErrorMessage(
			"math", "sign",
			"1 number argument",
			fmt.Sprintf("%d arguments", len(args)),
			"math.sign(5) or math.sign(-3.5)",
		)
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		if arg.Value == 0 {
			return &object.Integer{Value: 0}
		} else if arg.Value > 0 {
			return &object.Integer{Value: 1}
		} else {
			return &object.Integer{Value: -1}
		}
	case *object.Float:
		if arg.Value == 0 {
			return &object.Integer{Value: 0}
		} else if arg.Value > 0 {
			return &object.Integer{Value: 1}
		} else {
			return &object.Integer{Value: -1}
		}
	default:
		return ErrorMessage(
			"math", "sign",
			"number argument (integer or float)",
			string(args[0].Type()),
			"math.sign(5) or math.sign(-3.5)",
		)
	}
}

func ceil(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return ErrorMessage(
			"math", "ceil",
			"no definitions allowed",
			fmt.Sprintf("%d definitions provided", len(defs)),
			"math.ceil(4.3)",
		)
	}
	if len(args) != 1 {
		return ErrorMessage(
			"math", "ceil",
			"1 number argument",
			fmt.Sprintf("%d arguments", len(args)),
			"math.ceil(4.3)",
		)
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return ErrorMessage(
			"math", "ceil",
			"number argument (integer or float)",
			string(args[0].Type()),
			"math.ceil(4.3)",
		)
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Integer{Value: arg.Value}
	case *object.Float:
		return &object.Integer{Value: int64(math.Ceil(arg.Value))}
	default:
		return &object.Error{Message: "The argument must be a number."}
	}
}

func floor(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This operation does not allow definitions."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This operation requires exactly one argument."}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a number."}
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Integer{Value: arg.Value}
	case *object.Float:
		return &object.Integer{Value: int64(math.Floor(arg.Value))}
	default:
		return &object.Error{Message: "The argument must be a number."}
	}
}

func sqrt(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This operation does not allow definitions."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This operation requires exactly one argument."}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a number."}
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Sqrt(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Sqrt(arg.Value)}
	default:
		return &object.Error{Message: "The argument must be a number."}
	}
}

func cbrt(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This operation does not allow definitions."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This operation requires exactly one argument."}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a number."}
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Cbrt(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Cbrt(arg.Value)}
	default:
		return &object.Error{Message: "The argument must be a number."}
	}
}

func root(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This operation does not allow definitions."}
	}
	if len(args) != 2 {
		return &object.Error{Message: "This operation requires exactly two arguments."}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The first argument must be a number."}
	}
	if args[1].Type() != object.INTEGER_OBJ {
		return &object.Error{Message: "The second argument must be a number."}
	}
	base, ok := args[0].(*object.Float)
	if !ok {
		base = &object.Float{Value: float64(args[0].(*object.Integer).Value)}
	}
	exp := args[1].(*object.Integer).Value

	if exp == 0 {
		return &object.Float{Value: 1.0}
	} else if exp < 0 {
		return &object.Error{Message: "The second argument must be a non-negative integer"}
	}

	x := 1.0
	for i := 0; i < 10; i++ {
		x = x - (math.Pow(x, float64(exp))-base.Value)/(float64(exp)*math.Pow(x, float64(exp-1)))
	}

	return &object.Float{Value: x}
}

func hypot(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This operation does not allow definitions."}
	}
	if len(args) < 2 {
		return &object.Error{Message: "This operation requires at least two arguments."}
	}
	var sumOfSquares float64
	for _, arg := range args {
		if arg.Type() != object.INTEGER_OBJ && arg.Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "Arguments must be numbers."}
		}
		switch num := arg.(type) {
		case *object.Integer:
			sumOfSquares += float64(num.Value) * float64(num.Value)
		case *object.Float:
			sumOfSquares += num.Value * num.Value
		}
	}
	return &object.Float{Value: math.Sqrt(sumOfSquares)}
}

func factorial(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This operation does not allow definitions."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This operation requires exactly one argument."}
	}
	if args[0].Type() != object.INTEGER_OBJ {
		return &object.Error{Message: "The argument must be a number."}
	}
	n := args[0].(*object.Integer).Value
	if n < 0 {
		return &object.Error{Message: "The argument must be a non-negative integer"}
	}
	result := int64(1)
	for i := int64(2); i <= n; i++ {
		result *= i
	}
	return &object.Integer{Value: result}
}
func round(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}

	num := args[0].(*object.Float).Value
	return &object.Integer{Value: int64(num + 0.5)}
}

func max(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}

	arg, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "The argument must be an array."}
	}

	if len(arg.Elements) == 0 {
		return &object.Error{Message: "The array cannot be empty."}
	}

	var maxNum float64

	for _, element := range arg.Elements {
		if element.Type() != object.INTEGER_OBJ && element.Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "All elements in the array must be numbers."}
		}

		switch num := element.(type) {
		case *object.Integer:
			if float64(num.Value) > maxNum {
				maxNum = float64(num.Value)
			}
		case *object.Float:
			if num.Value > maxNum {
				maxNum = num.Value
			}
		default:
			return &object.Error{Message: "All elements in the array must be numbers."}
		}
	}

	return &object.Float{Value: maxNum}
}

func min(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}

	arg, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "The argument must be an array."}
	}

	if len(arg.Elements) == 0 {
		return &object.Error{Message: "The array cannot be empty."}
	}

	minNum := math.MaxFloat64

	for _, element := range arg.Elements {
		if element.Type() != object.INTEGER_OBJ && element.Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "All elements in the array must be numbers."}
		}

		switch num := element.(type) {
		case *object.Integer:
			if float64(num.Value) < minNum {
				minNum = float64(num.Value)
			}
		case *object.Float:
			if num.Value < minNum {
				minNum = num.Value
			}
		default:
			return &object.Error{Message: "All elements in the array must be numbers."}
		}
	}

	return &object.Float{Value: minNum}
}

func exp(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Exp(num)}
}

func expm1(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Expm1(num)}
}

// func log(args []object.Object, defs map[string]object.Object) object.Object {
// 	if len(defs) != 0 {
// 		return &object.Error{Message: "This function does not accept keyword arguments."}
// 	}
// 	if len(args) != 1 {
// 		return &object.Error{Message: "This function requires exactly one argument."}
// 	}
// 	if args[0].Type() != object.FLOAT_OBJ {
// 		return &object.Error{Message: "The argument must be a float."}
// 	}
// 	num := args[0].(*object.Float).Value
// 	return &object.Float{Value: math.Log(num)}
// }

func log10(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Log10(num)}
}

func log2(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a number."}
	}

	arg := extractFloatValue(args[0])

	if arg <= 0 {
		return &object.Error{Message: "The argument must be greater than 0."}
	}

	return &object.Float{Value: math.Log2(arg)}
}

func extractFloatValue(obj object.VintObject) float64 {
	switch obj := obj.(type) {
	case *object.Integer:
		return float64(obj.Value)
	case *object.Float:
		return obj.Value
	default:
		return 0
	}
}

func log1p(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Log1p(num)}
}

func cos(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Cos(num)}
}

func sin(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Sin(num)}
}

func tan(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Tan(num)}
}

func acos(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Acos(num)}
}

func asin(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Asin(num)}
}

func atan(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Atan(num)}
}

func cosh(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Cosh(num)}
}

func sinh(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Sinh(num)}
}

func tanh(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Tanh(num)}
}

func acosh(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Acosh(num)}
}

func asinh(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Asinh(num)}
}

func atan2(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 2 {
		return &object.Error{Message: "This function requires exactly two arguments."}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Arguments must be numbers."}
	}
	if args[1].Type() != object.INTEGER_OBJ && args[1].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Arguments must be numbers."}
	}

	y := extractFloatValue(args[0])
	x := extractFloatValue(args[1])

	return &object.Float{Value: math.Atan2(y, x)}
}

func atanh(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument."}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "The argument must be a float."}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Atanh(num)}
}

func random(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}

	if len(args) != 0 {
		return &object.Error{Message: "This function takes no arguments."}
	}

	rand.Seed(time.Now().UnixNano())
	value := rand.Float64()

	return &object.Float{Value: value}
}

// Statistics functions

func mean(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument (array of numbers)."}
	}

	arr, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "The argument must be an array."}
	}

	if len(arr.Elements) == 0 {
		return &object.Error{Message: "Cannot calculate mean of empty array."}
	}

	var sum float64
	for _, element := range arr.Elements {
		if element.Type() != object.INTEGER_OBJ && element.Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "All elements in the array must be numbers."}
		}
		sum += extractFloatValue(element)
	}

	return &object.Float{Value: sum / float64(len(arr.Elements))}
}

func variance(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument (array of numbers)."}
	}

	arr, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "The argument must be an array."}
	}

	if len(arr.Elements) == 0 {
		return &object.Error{Message: "Cannot calculate variance of empty array."}
	}

	// Calculate mean first
	var sum float64
	for _, element := range arr.Elements {
		if element.Type() != object.INTEGER_OBJ && element.Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "All elements in the array must be numbers."}
		}
		sum += extractFloatValue(element)
	}
	meanVal := sum / float64(len(arr.Elements))

	// Calculate variance
	var varianceSum float64
	for _, element := range arr.Elements {
		diff := extractFloatValue(element) - meanVal
		varianceSum += diff * diff
	}

	return &object.Float{Value: varianceSum / float64(len(arr.Elements))}
}

func stddev(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	varianceResult := variance(args, defs)
	if varianceResult.Type() == object.ERROR_OBJ {
		return varianceResult
	}

	varianceVal := varianceResult.(*object.Float).Value
	return &object.Float{Value: math.Sqrt(varianceVal)}
}

func median(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument (array of numbers)."}
	}

	arr, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "The argument must be an array."}
	}

	if len(arr.Elements) == 0 {
		return &object.Error{Message: "Cannot calculate median of empty array."}
	}

	// Extract and validate numbers
	numbers := make([]float64, len(arr.Elements))
	for i, element := range arr.Elements {
		if element.Type() != object.INTEGER_OBJ && element.Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "All elements in the array must be numbers."}
		}
		numbers[i] = extractFloatValue(element)
	}

	// Sort the numbers
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	// Calculate median
	n := len(numbers)
	if n%2 == 0 {
		return &object.Float{Value: (numbers[n/2-1] + numbers[n/2]) / 2}
	}
	return &object.Float{Value: numbers[n/2]}
}

// Complex number support

func complexNum(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 2 {
		return &object.Error{Message: "This function requires exactly two arguments (real, imaginary)."}
	}

	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Both arguments must be numbers."}
	}
	if args[1].Type() != object.INTEGER_OBJ && args[1].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Both arguments must be numbers."}
	}

	real := extractFloatValue(args[0])
	imag := extractFloatValue(args[1])

	// Return as a dict with real and imag properties
	dict := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	
	realKey := &object.String{Value: "real"}
	imagKey := &object.String{Value: "imag"}
	
	dict.Pairs[realKey.HashKey()] = object.DictPair{
		Key:   realKey,
		Value: &object.Float{Value: real},
	}
	dict.Pairs[imagKey.HashKey()] = object.DictPair{
		Key:   imagKey,
		Value: &object.Float{Value: imag},
	}

	return dict
}

// Big integer support

func bigint(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument (string or integer)."}
	}

	var value string
	switch arg := args[0].(type) {
	case *object.String:
		value = arg.Value
	case *object.Integer:
		value = fmt.Sprintf("%d", arg.Value)
	default:
		return &object.Error{Message: "Argument must be a string or integer."}
	}

	// Return as a dict with value and type properties
	dict := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	
	valueKey := &object.String{Value: "value"}
	typeKey := &object.String{Value: "type"}
	
	dict.Pairs[valueKey.HashKey()] = object.DictPair{
		Key:   valueKey,
		Value: &object.String{Value: value},
	}
	dict.Pairs[typeKey.HashKey()] = object.DictPair{
		Key:   typeKey,
		Value: &object.String{Value: "bigint"},
	}

	return dict
}

// Linear algebra operations

func dot(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 2 {
		return &object.Error{Message: "This function requires exactly two arguments (two arrays)."}
	}

	arr1, ok1 := args[0].(*object.Array)
	arr2, ok2 := args[1].(*object.Array)
	
	if !ok1 || !ok2 {
		return &object.Error{Message: "Both arguments must be arrays."}
	}

	if len(arr1.Elements) != len(arr2.Elements) {
		return &object.Error{Message: "Arrays must have the same length."}
	}

	if len(arr1.Elements) == 0 {
		return &object.Error{Message: "Arrays cannot be empty."}
	}

	var result float64
	for i := 0; i < len(arr1.Elements); i++ {
		if arr1.Elements[i].Type() != object.INTEGER_OBJ && arr1.Elements[i].Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "All elements must be numbers."}
		}
		if arr2.Elements[i].Type() != object.INTEGER_OBJ && arr2.Elements[i].Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "All elements must be numbers."}
		}
		result += extractFloatValue(arr1.Elements[i]) * extractFloatValue(arr2.Elements[i])
	}

	return &object.Float{Value: result}
}

func cross(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 2 {
		return &object.Error{Message: "This function requires exactly two arguments (two 3D vectors)."}
	}

	arr1, ok1 := args[0].(*object.Array)
	arr2, ok2 := args[1].(*object.Array)
	
	if !ok1 || !ok2 {
		return &object.Error{Message: "Both arguments must be arrays."}
	}

	if len(arr1.Elements) != 3 || len(arr2.Elements) != 3 {
		return &object.Error{Message: "Both vectors must be 3D (length 3)."}
	}

	// Extract components
	var a1, a2, a3, b1, b2, b3 float64
	for i := 0; i < 3; i++ {
		if arr1.Elements[i].Type() != object.INTEGER_OBJ && arr1.Elements[i].Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "All elements must be numbers."}
		}
		if arr2.Elements[i].Type() != object.INTEGER_OBJ && arr2.Elements[i].Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "All elements must be numbers."}
		}
	}
	a1, a2, a3 = extractFloatValue(arr1.Elements[0]), extractFloatValue(arr1.Elements[1]), extractFloatValue(arr1.Elements[2])
	b1, b2, b3 = extractFloatValue(arr2.Elements[0]), extractFloatValue(arr2.Elements[1]), extractFloatValue(arr2.Elements[2])

	// Calculate cross product
	result := &object.Array{
		Elements: []object.VintObject{
			&object.Float{Value: a2*b3 - a3*b2},
			&object.Float{Value: a3*b1 - a1*b3},
			&object.Float{Value: a1*b2 - a2*b1},
		},
	}

	return result
}

func magnitude(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "This function requires exactly one argument (array)."}
	}

	arr, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "The argument must be an array."}
	}

	if len(arr.Elements) == 0 {
		return &object.Error{Message: "Array cannot be empty."}
	}

	var sumOfSquares float64
	for _, element := range arr.Elements {
		if element.Type() != object.INTEGER_OBJ && element.Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "All elements must be numbers."}
		}
		val := extractFloatValue(element)
		sumOfSquares += val * val
	}

	return &object.Float{Value: math.Sqrt(sumOfSquares)}
}

// Numerical methods

func gcd(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 2 {
		return &object.Error{Message: "This function requires exactly two arguments (two integers)."}
	}

	if args[0].Type() != object.INTEGER_OBJ || args[1].Type() != object.INTEGER_OBJ {
		return &object.Error{Message: "Both arguments must be integers."}
	}

	a := args[0].(*object.Integer).Value
	b := args[1].(*object.Integer).Value

	// Make positive
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}

	return &object.Integer{Value: a}
}

func lcm(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 2 {
		return &object.Error{Message: "This function requires exactly two arguments (two integers)."}
	}

	gcdResult := gcd(args, defs)
	if gcdResult.Type() == object.ERROR_OBJ {
		return gcdResult
	}

	a := args[0].(*object.Integer).Value
	b := args[1].(*object.Integer).Value
	gcdVal := gcdResult.(*object.Integer).Value

	if gcdVal == 0 {
		return &object.Integer{Value: 0}
	}

	// Make positive
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	return &object.Integer{Value: (a * b) / gcdVal}
}

func clamp(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 3 {
		return &object.Error{Message: "This function requires exactly three arguments (value, min, max)."}
	}

	for i := 0; i < 3; i++ {
		if args[i].Type() != object.INTEGER_OBJ && args[i].Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "All arguments must be numbers."}
		}
	}

	value := extractFloatValue(args[0])
	minVal := extractFloatValue(args[1])
	maxVal := extractFloatValue(args[2])

	if minVal > maxVal {
		return &object.Error{Message: "Min value must be less than or equal to max value."}
	}

	if value < minVal {
		value = minVal
	} else if value > maxVal {
		value = maxVal
	}

	return &object.Float{Value: value}
}

func lerp(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(defs) != 0 {
		return &object.Error{Message: "This function does not accept keyword arguments."}
	}
	if len(args) != 3 {
		return &object.Error{Message: "This function requires exactly three arguments (start, end, t)."}
	}

	for i := 0; i < 3; i++ {
		if args[i].Type() != object.INTEGER_OBJ && args[i].Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "All arguments must be numbers."}
		}
	}

	start := extractFloatValue(args[0])
	end := extractFloatValue(args[1])
	t := extractFloatValue(args[2])

	return &object.Float{Value: start + (end-start)*t}
}
