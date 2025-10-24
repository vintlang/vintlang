package object

import (
	"hash/fnv"
	"math"
	"strconv"
)

type Float struct {
	Value float64
}

func (f *Float) Inspect() string      { return strconv.FormatFloat(f.Value, 'f', -1, 64) }
func (f *Float) Type() VintObjectType { return FLOAT_OBJ }

func (f *Float) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(f.Inspect()))
	return HashKey{Type: f.Type(), Value: h.Sum64()}
}

func (f *Float) Method(method string, args []VintObject) VintObject {
	switch method {
	case "abs":
		return f.abs(args)
	case "ceil":
		return f.ceil(args)
	case "floor":
		return f.floor(args)
	case "round":
		return f.round(args)
	case "sqrt":
		return f.sqrt(args)
	case "pow":
		return f.pow(args)
	case "is_nan":
		return f.isNaN(args)
	case "is_infinite":
		return f.isInfinite(args)
	case "to_string":
		return f.toString(args)
	case "clamp":
		return f.clamp(args)
	case "toPrecision":
		return f.toPrecision(args)
	case "toFixed":
		return f.toFixed(args)
	case "sign":
		return f.sign(args)
	case "truncate":
		return f.truncate(args)
	case "mod":
		return f.mod(args)
	case "degrees":
		return f.degrees(args)
	case "radians":
		return f.radians(args)
	case "sin":
		return f.sin(args)
	case "cos":
		return f.cos(args)
	case "tan":
		return f.tan(args)
	case "log":
		return f.log(args)
	case "exp":
		return f.exp(args)
	default:
		return newError("Method '%s' is not supported for Float objects", method)
	}
}

func (f *Float) abs(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("abs() expects 0 arguments, got %d", len(args))
	}
	return &Float{Value: math.Abs(f.Value)}
}

func (f *Float) ceil(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("ceil() expects 0 arguments, got %d", len(args))
	}
	return &Integer{Value: int64(math.Ceil(f.Value))}
}

func (f *Float) floor(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("floor() expects 0 arguments, got %d", len(args))
	}
	return &Integer{Value: int64(math.Floor(f.Value))}
}

func (f *Float) round(args []VintObject) VintObject {
	if len(args) > 1 {
		return newError("round() expects 0 or 1 arguments, got %d", len(args))
	}

	if len(args) == 0 {
		return &Float{Value: math.Round(f.Value)}
	}

	digits, ok := args[0].(*Integer)
	if !ok {
		return newError("Digits argument must be an integer")
	}

	multiplier := math.Pow(10, float64(digits.Value))
	return &Float{Value: math.Round(f.Value*multiplier) / multiplier}
}

func (f *Float) sqrt(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("sqrt() expects 0 arguments, got %d", len(args))
	}

	if f.Value < 0 {
		return newError("Cannot calculate square root of negative number")
	}

	return &Float{Value: math.Sqrt(f.Value)}
}

func (f *Float) pow(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("pow() expects 1 argument, got %d", len(args))
	}

	var exponent float64
	switch exp := args[0].(type) {
	case *Integer:
		exponent = float64(exp.Value)
	case *Float:
		exponent = exp.Value
	default:
		return newError("Exponent must be a number")
	}

	return &Float{Value: math.Pow(f.Value, exponent)}
}

func (f *Float) isNaN(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("is_nan() expects 0 arguments, got %d", len(args))
	}
	return &Boolean{Value: math.IsNaN(f.Value)}
}

func (f *Float) isInfinite(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("is_infinite() expects 0 arguments, got %d", len(args))
	}
	return &Boolean{Value: math.IsInf(f.Value, 0)}
}

func (f *Float) toString(args []VintObject) VintObject {
	if len(args) > 1 {
		return newError("to_string() expects 0 or 1 arguments, got %d", len(args))
	}

	if len(args) == 0 {
		return &String{Value: f.Inspect()}
	}

	precision, ok := args[0].(*Integer)
	if !ok {
		return newError("Precision argument must be an integer")
	}

	return &String{Value: strconv.FormatFloat(f.Value, 'f', int(precision.Value), 64)}
}

func (f *Float) clamp(args []VintObject) VintObject {
	if len(args) != 2 {
		return newError("clamp() expects 2 arguments, got %d", len(args))
	}

	var min, max float64

	switch minVal := args[0].(type) {
	case *Integer:
		min = float64(minVal.Value)
	case *Float:
		min = minVal.Value
	default:
		return newError("Min value must be a number")
	}

	switch maxVal := args[1].(type) {
	case *Integer:
		max = float64(maxVal.Value)
	case *Float:
		max = maxVal.Value
	default:
		return newError("Max value must be a number")
	}

	if min > max {
		return newError("Min value cannot be greater than max value")
	}

	value := f.Value
	if value < min {
		value = min
	} else if value > max {
		value = max
	}

	return &Float{Value: value}
}

// toPrecision formats the float to specified precision
func (f *Float) toPrecision(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("toPrecision() expects 1 argument, got %d", len(args))
	}

	precision, ok := args[0].(*Integer)
	if !ok {
		return newError("Precision must be an integer")
	}

	if precision.Value < 1 || precision.Value > 21 {
		return newError("Precision must be between 1 and 21")
	}

	result := strconv.FormatFloat(f.Value, 'g', int(precision.Value), 64)
	return &String{Value: result}
}

// toFixed formats the float to fixed decimal places
func (f *Float) toFixed(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("toFixed() expects 1 argument, got %d", len(args))
	}

	places, ok := args[0].(*Integer)
	if !ok {
		return newError("Decimal places must be an integer")
	}

	if places.Value < 0 || places.Value > 20 {
		return newError("Decimal places must be between 0 and 20")
	}

	result := strconv.FormatFloat(f.Value, 'f', int(places.Value), 64)
	return &String{Value: result}
}

// sign returns the sign of the float
func (f *Float) sign(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("sign() expects 0 arguments, got %d", len(args))
	}

	if math.IsNaN(f.Value) {
		return &Float{Value: math.NaN()}
	}

	if f.Value > 0 {
		return &Float{Value: 1.0}
	} else if f.Value < 0 {
		return &Float{Value: -1.0}
	}
	return &Float{Value: 0.0}
}

// truncate removes the fractional part
func (f *Float) truncate(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("truncate() expects 0 arguments, got %d", len(args))
	}

	return &Float{Value: math.Trunc(f.Value)}
}

// mod calculates the floating-point remainder
func (f *Float) mod(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("mod() expects 1 argument, got %d", len(args))
	}

	var divisor float64
	switch arg := args[0].(type) {
	case *Float:
		divisor = arg.Value
	case *Integer:
		divisor = float64(arg.Value)
	default:
		return newError("Divisor must be a number")
	}

	if divisor == 0 {
		return newError("Division by zero")
	}

	return &Float{Value: math.Mod(f.Value, divisor)}
}

// degrees converts radians to degrees
func (f *Float) degrees(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("degrees() expects 0 arguments, got %d", len(args))
	}

	return &Float{Value: f.Value * 180.0 / math.Pi}
}

// radians converts degrees to radians
func (f *Float) radians(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("radians() expects 0 arguments, got %d", len(args))
	}

	return &Float{Value: f.Value * math.Pi / 180.0}
}

// sin calculates the sine
func (f *Float) sin(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("sin() expects 0 arguments, got %d", len(args))
	}

	return &Float{Value: math.Sin(f.Value)}
}

// cos calculates the cosine
func (f *Float) cos(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("cos() expects 0 arguments, got %d", len(args))
	}

	return &Float{Value: math.Cos(f.Value)}
}

// tan calculates the tangent
func (f *Float) tan(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("tan() expects 0 arguments, got %d", len(args))
	}

	return &Float{Value: math.Tan(f.Value)}
}

// log calculates the natural logarithm
func (f *Float) log(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("log() expects 0 arguments, got %d", len(args))
	}

	if f.Value <= 0 {
		return newError("Cannot calculate log of non-positive number")
	}

	return &Float{Value: math.Log(f.Value)}
}

// exp calculates e raised to the power of the float
func (f *Float) exp(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("exp() expects 0 arguments, got %d", len(args))
	}

	return &Float{Value: math.Exp(f.Value)}
}
