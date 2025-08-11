package object

import (
	"hash/fnv"
	"math"
	"strconv"
)

type Float struct {
	Value float64
}

func (f *Float) Inspect() string  { return strconv.FormatFloat(f.Value, 'f', -1, 64) }
func (f *Float) Type() ObjectType { return FLOAT_OBJ }

func (f *Float) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(f.Inspect()))
	return HashKey{Type: f.Type(), Value: h.Sum64()}
}

func (f *Float) Method(method string, args []Object) Object {
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
	default:
		return newError("Method '%s' is not supported for Float objects", method)
	}
}

func (f *Float) abs(args []Object) Object {
	if len(args) != 0 {
		return newError("abs() expects 0 arguments, got %d", len(args))
	}
	return &Float{Value: math.Abs(f.Value)}
}

func (f *Float) ceil(args []Object) Object {
	if len(args) != 0 {
		return newError("ceil() expects 0 arguments, got %d", len(args))
	}
	return &Integer{Value: int64(math.Ceil(f.Value))}
}

func (f *Float) floor(args []Object) Object {
	if len(args) != 0 {
		return newError("floor() expects 0 arguments, got %d", len(args))
	}
	return &Integer{Value: int64(math.Floor(f.Value))}
}

func (f *Float) round(args []Object) Object {
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

func (f *Float) sqrt(args []Object) Object {
	if len(args) != 0 {
		return newError("sqrt() expects 0 arguments, got %d", len(args))
	}
	
	if f.Value < 0 {
		return newError("Cannot calculate square root of negative number")
	}
	
	return &Float{Value: math.Sqrt(f.Value)}
}

func (f *Float) pow(args []Object) Object {
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

func (f *Float) isNaN(args []Object) Object {
	if len(args) != 0 {
		return newError("is_nan() expects 0 arguments, got %d", len(args))
	}
	return &Boolean{Value: math.IsNaN(f.Value)}
}

func (f *Float) isInfinite(args []Object) Object {
	if len(args) != 0 {
		return newError("is_infinite() expects 0 arguments, got %d", len(args))
	}
	return &Boolean{Value: math.IsInf(f.Value, 0)}
}

func (f *Float) toString(args []Object) Object {
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

func (f *Float) clamp(args []Object) Object {
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
