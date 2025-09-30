package object

import (
	"fmt"
	"math"
)

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (i *Integer) Method(method string, args []VintObject) VintObject {
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
	case "pow":
		return i.pow(args)
	case "sqrt":
		return i.sqrt(args)
	case "gcd":
		return i.gcd(args)
	case "lcm":
		return i.lcm(args)
	case "factorial":
		return i.factorial(args)
	default:
		return newError("Sorry, the method '%s' is not supported for Integer.", method)
	}
}

func (i *Integer) abs(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("abs() expects 0 arguments, got %d", len(args))
	}
	v := i.Value
	if v < 0 {
		v = -v
	}
	return &Integer{Value: v}
}

func (i *Integer) isEven(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("is_even() expects 0 arguments, got %d", len(args))
	}
	return &Boolean{Value: i.Value%2 == 0}
}

func (i *Integer) isOdd(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("is_odd() expects 0 arguments, got %d", len(args))
	}
	return &Boolean{Value: i.Value%2 != 0}
}

func (i *Integer) toString(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("to_string() expects 0 arguments, got %d", len(args))
	}
	return &String{Value: i.Inspect()}
}

func (i *Integer) sign(args []VintObject) VintObject {
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

func (i *Integer) pow(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("pow() expects 1 argument, got %d", len(args))
	}

	exponent, ok := args[0].(*Integer)
	if !ok {
		return newError("Exponent must be an integer")
	}

	if exponent.Value < 0 {
		return newError("Negative exponents not supported for integers")
	}

	result := int64(1)
	base := i.Value
	exp := exponent.Value

	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}

	return &Integer{Value: result}
}

func (i *Integer) sqrt(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("sqrt() expects 0 arguments, got %d", len(args))
	}

	if i.Value < 0 {
		return newError("Cannot calculate square root of negative number")
	}

	return &Float{Value: math.Sqrt(float64(i.Value))}
}

func (i *Integer) gcd(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("gcd() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Integer)
	if !ok {
		return newError("Argument must be an integer")
	}

	a, b := i.Value, other.Value
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for b != 0 {
		a, b = b, a%b
	}

	return &Integer{Value: a}
}

func (i *Integer) lcm(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("lcm() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Integer)
	if !ok {
		return newError("Argument must be an integer")
	}

	// LCM(a, b) = |a * b| / GCD(a, b)
	gcdResult := i.gcd(args)
	gcdValue := gcdResult.(*Integer).Value

	a, b := i.Value, other.Value
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	return &Integer{Value: (a * b) / gcdValue}
}

func (i *Integer) factorial(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("factorial() expects 0 arguments, got %d", len(args))
	}

	if i.Value < 0 {
		return newError("Factorial is not defined for negative numbers")
	}

	if i.Value > 20 {
		return newError("Factorial of numbers greater than 20 may cause overflow")
	}

	result := int64(1)
	for n := int64(2); n <= i.Value; n++ {
		result *= n
	}

	return &Integer{Value: result}
}
