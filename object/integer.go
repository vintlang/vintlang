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
	case "toBinary":
		return i.toBinary(args)
	case "toHex":
		return i.toHex(args)
	case "toOctal":
		return i.toOctal(args)
	case "isPrime":
		return i.isPrime(args)
	case "nthRoot":
		return i.nthRoot(args)
	case "mod":
		return i.mod(args)
	case "clamp":
		return i.clamp(args)
	case "inRange":
		return i.inRange(args)
	case "digits":
		return i.digits(args)
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

// toBinary converts the integer to binary representation
func (i *Integer) toBinary(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("toBinary() expects 0 arguments, got %d", len(args))
	}
	return &String{Value: fmt.Sprintf("%b", i.Value)}
}

// toHex converts the integer to hexadecimal representation
func (i *Integer) toHex(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("toHex() expects 0 arguments, got %d", len(args))
	}
	return &String{Value: fmt.Sprintf("%x", i.Value)}
}

// toOctal converts the integer to octal representation
func (i *Integer) toOctal(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("toOctal() expects 0 arguments, got %d", len(args))
	}
	return &String{Value: fmt.Sprintf("%o", i.Value)}
}

// isPrime checks if the integer is a prime number
func (i *Integer) isPrime(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("isPrime() expects 0 arguments, got %d", len(args))
	}

	n := i.Value
	if n < 2 {
		return &Boolean{Value: false}
	}
	if n == 2 {
		return &Boolean{Value: true}
	}
	if n%2 == 0 {
		return &Boolean{Value: false}
	}

	for j := int64(3); j*j <= n; j += 2 {
		if n%j == 0 {
			return &Boolean{Value: false}
		}
	}

	return &Boolean{Value: true}
}

// nthRoot calculates the nth root of the integer
func (i *Integer) nthRoot(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("nthRoot() expects 1 argument, got %d", len(args))
	}

	n, ok := args[0].(*Integer)
	if !ok {
		return newError("Root must be an integer")
	}

	if n.Value <= 0 {
		return newError("Root must be positive")
	}

	if i.Value < 0 && n.Value%2 == 0 {
		return newError("Cannot calculate even root of negative number")
	}

	result := math.Pow(float64(i.Value), 1.0/float64(n.Value))
	return &Float{Value: result}
}

// mod calculates the modulo (remainder) of the integer with another integer
func (i *Integer) mod(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("mod() expects 1 argument, got %d", len(args))
	}

	divisor, ok := args[0].(*Integer)
	if !ok {
		return newError("Divisor must be an integer")
	}

	if divisor.Value == 0 {
		return newError("Division by zero")
	}

	return &Integer{Value: i.Value % divisor.Value}
}

// clamp restricts the integer to be within specified bounds
func (i *Integer) clamp(args []VintObject) VintObject {
	if len(args) != 2 {
		return newError("clamp() expects 2 arguments (min, max), got %d", len(args))
	}

	min, ok1 := args[0].(*Integer)
	max, ok2 := args[1].(*Integer)
	if !ok1 || !ok2 {
		return newError("Both bounds must be integers")
	}

	if min.Value > max.Value {
		return newError("Minimum bound cannot be greater than maximum bound")
	}

	value := i.Value
	if value < min.Value {
		value = min.Value
	} else if value > max.Value {
		value = max.Value
	}

	return &Integer{Value: value}
}

// inRange checks if the integer is within the specified range (inclusive)
func (i *Integer) inRange(args []VintObject) VintObject {
	if len(args) != 2 {
		return newError("inRange() expects 2 arguments (min, max), got %d", len(args))
	}

	min, ok1 := args[0].(*Integer)
	max, ok2 := args[1].(*Integer)
	if !ok1 || !ok2 {
		return newError("Both bounds must be integers")
	}

	return &Boolean{Value: i.Value >= min.Value && i.Value <= max.Value}
}

// digits returns an array of individual digits
func (i *Integer) digits(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("digits() expects 0 arguments, got %d", len(args))
	}

	value := i.Value
	if value < 0 {
		value = -value
	}

	str := fmt.Sprintf("%d", value)
	digits := make([]VintObject, len(str))

	for i, char := range str {
		digit := int64(char - '0')
		digits[i] = &Integer{Value: digit}
	}

	return &Array{Elements: digits}
}
