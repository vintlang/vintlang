package object

import (
	"time"
)

type Duration struct {
	Value time.Duration
}

func (d *Duration) Type() ObjectType { return DURATION_OBJ }
func (d *Duration) Inspect() string  { return d.Value.String() }

func (d *Duration) Method(method string, args []VintObject, defs map[string]VintObject) VintObject {
	switch method {
	case "hours":
		return d.hours(args, defs)
	case "minutes":
		return d.minutes(args, defs)
	case "seconds":
		return d.seconds(args, defs)
	case "milliseconds":
		return d.milliseconds(args, defs)
	case "nanoseconds":
		return d.nanoseconds(args, defs)
	case "string":
		return d.string(args, defs)
	case "add":
		return d.add(args, defs)
	case "subtract":
		return d.subtract(args, defs)
	case "multiply":
		return d.multiply(args, defs)
	case "divide":
		return d.divide(args, defs)
	}
	return nil
}

func (d *Duration) hours(args []VintObject, defs map[string]VintObject) VintObject {
	if len(defs) != 0 {
		return newError("hours() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("hours() expects 0 arguments, got %d", len(args))
	}

	return &Float{Value: d.Value.Hours()}
}

func (d *Duration) minutes(args []VintObject, defs map[string]VintObject) VintObject {
	if len(defs) != 0 {
		return newError("minutes() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("minutes() expects 0 arguments, got %d", len(args))
	}

	return &Float{Value: d.Value.Minutes()}
}

func (d *Duration) seconds(args []VintObject, defs map[string]VintObject) VintObject {
	if len(defs) != 0 {
		return newError("seconds() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("seconds() expects 0 arguments, got %d", len(args))
	}

	return &Float{Value: d.Value.Seconds()}
}

func (d *Duration) milliseconds(args []VintObject, defs map[string]VintObject) VintObject {
	if len(defs) != 0 {
		return newError("milliseconds() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("milliseconds() expects 0 arguments, got %d", len(args))
	}

	return &Integer{Value: d.Value.Nanoseconds() / 1e6}
}

func (d *Duration) nanoseconds(args []VintObject, defs map[string]VintObject) VintObject {
	if len(defs) != 0 {
		return newError("nanoseconds() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("nanoseconds() expects 0 arguments, got %d", len(args))
	}

	return &Integer{Value: d.Value.Nanoseconds()}
}

func (d *Duration) string(args []VintObject, defs map[string]VintObject) VintObject {
	if len(defs) != 0 {
		return newError("string() does not accept keyword arguments")
	}
	if len(args) != 0 {
		return newError("string() expects 0 arguments, got %d", len(args))
	}

	return &String{Value: d.Value.String()}
}

func (d *Duration) add(args []VintObject, defs map[string]VintObject) VintObject {
	if len(defs) != 0 {
		return newError("add() does not accept keyword arguments")
	}
	if len(args) != 1 {
		return newError("add() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Duration)
	if !ok {
		return newError("add() argument must be a duration")
	}

	return &Duration{Value: d.Value + other.Value}
}

func (d *Duration) subtract(args []VintObject, defs map[string]VintObject) VintObject {
	if len(defs) != 0 {
		return newError("subtract() does not accept keyword arguments")
	}
	if len(args) != 1 {
		return newError("subtract() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Duration)
	if !ok {
		return newError("subtract() argument must be a duration")
	}

	return &Duration{Value: d.Value - other.Value}
}

func (d *Duration) multiply(args []VintObject, defs map[string]VintObject) VintObject {
	if len(defs) != 0 {
		return newError("multiply() does not accept keyword arguments")
	}
	if len(args) != 1 {
		return newError("multiply() expects 1 argument, got %d", len(args))
	}

	switch arg := args[0].(type) {
	case *Integer:
		return &Duration{Value: time.Duration(int64(arg.Value)) * d.Value}
	case *Float:
		return &Duration{Value: time.Duration(arg.Value * float64(d.Value))}
	default:
		return newError("multiply() argument must be a number")
	}
}

func (d *Duration) divide(args []VintObject, defs map[string]VintObject) VintObject {
	if len(defs) != 0 {
		return newError("divide() does not accept keyword arguments")
	}
	if len(args) != 1 {
		return newError("divide() expects 1 argument, got %d", len(args))
	}

	switch arg := args[0].(type) {
	case *Integer:
		if arg.Value == 0 {
			return newError("division by zero")
		}
		return &Duration{Value: d.Value / time.Duration(arg.Value)}
	case *Float:
		if arg.Value == 0.0 {
			return newError("division by zero")
		}
		return &Duration{Value: time.Duration(float64(d.Value) / arg.Value)}
	case *Duration:
		if arg.Value == 0 {
			return newError("division by zero")
		}
		return &Float{Value: float64(d.Value) / float64(arg.Value)}
	default:
		return newError("divide() argument must be a number or duration")
	}
}
