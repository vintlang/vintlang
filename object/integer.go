package object

import "fmt"

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (i *Integer) Method(method string, args []Object) Object {
	switch method {
	case "abs":
		return i.abs(args)
	case "is_even":
		return i.isEven(args)
	case "is_odd":
		return i.isOdd(args)
	case "to_string":
		return i.toString(args)
	case "sign":
		return i.sign(args)
	default:
		return newError("Sorry, the method '%s' is not supported for Integer.", method)
	}
}

func (i *Integer) abs(args []Object) Object {
	if len(args) != 0 {
		return newError("abs() expects 0 arguments, got %d", len(args))
	}
	v := i.Value
	if v < 0 {
		v = -v
	}
	return &Integer{Value: v}
}

func (i *Integer) isEven(args []Object) Object {
	if len(args) != 0 {
		return newError("is_even() expects 0 arguments, got %d", len(args))
	}
	return &Boolean{Value: i.Value%2 == 0}
}

func (i *Integer) isOdd(args []Object) Object {
	if len(args) != 0 {
		return newError("is_odd() expects 0 arguments, got %d", len(args))
	}
	return &Boolean{Value: i.Value%2 != 0}
}

func (i *Integer) toString(args []Object) Object {
	if len(args) != 0 {
		return newError("to_string() expects 0 arguments, got %d", len(args))
	}
	return &String{Value: i.Inspect()}
}

func (i *Integer) sign(args []Object) Object {
	if len(args) != 0 {
		return newError("sign() expects 0 arguments, got %d", len(args))
	}
	if i.Value > 0 {
		return &Integer{Value: 1}
	} else if i.Value < 0 {
		return &Integer{Value: -1}
	}
	return &Integer{Value: 0}
}
