package object

import (
	// "fmt"
	"hash/fnv"
	"strconv"
	// "strconv"
	"strings"
)

// String represents a string object in the system.
type String struct {
	Value  string
	offset int
}

// Inspect returns the string representation of the String object.
func (s *String) Inspect() string { return s.Value }

// Type returns the object type for String.
func (s *String) Type() ObjectType { return STRING_OBJ }

// HashKey generates a hash key for the String object.
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// Next implements an iterator for the String, returning the next character and its index.
func (s *String) Next() (Object, Object) {
	if s.offset < len(s.Value) {
		char := string(s.Value[s.offset])
		index := &Integer{Value: int64(s.offset)}
		s.offset++
		return index, &String{Value: char}
	}
	return nil, nil
}

// Reset resets the iterator offset to the beginning of the string.
func (s *String) Reset() {
	s.offset = 0
}

// Method dynamically dispatches string-related methods.
func (s *String) Method(method string, args []Object) Object {
	switch method {
	case "len":
		return s.len(args)
	case "upper":
		return s.upper(args)
	case "lower":
		return s.lower(args)
	case "split":
		return s.split(args)
	case "trim":
		return s.trim(args)
	case "contains":
		return s.contains(args)
	case "replace":
		return s.replace(args)
	case "reverse":
		return s.reverse(args)
	// case "format":
	// 	return s.format(args)
	default:
		return newError("Method '%s' is not supported on strings", method)
	}
}

// len returns the length of the string.
func (s *String) len(args []Object) Object {
	if len(args) != 0 {
		return newError("len() expects 0 arguments, got %d", len(args))
	}
	return &Integer{Value: int64(len(s.Value))}
}

// toInt converts the string to an integer
func (s *String) toInt(args []Object) Object {
	// Ensure no arguments are provided
	if len(args) != 0 {
		return newError("toInt() expects 0 arguments, got %d", len(args))
	}

	// Try to convert the string value to an integer
	numVal, err := strconv.Atoi(s.Value)
	if err != nil {
		// Return an error if conversion fails
		return newError("Failed to convert '%s' to an integer", s.Value)
	}

	// Return the integer value as an Integer object
	return &Integer{Value: int64(numVal)}
}


// upper converts the string to uppercase.
func (s *String) upper(args []Object) Object {
	if len(args) != 0 {
		return newError("upper() expects 0 arguments, got %d", len(args))
	}
	return &String{Value: strings.ToUpper(s.Value)}
}

// lower converts the string to lowercase.
func (s *String) lower(args []Object) Object {
	if len(args) != 0 {
		return newError("lower() expects 0 arguments, got %d", len(args))
	}
	return &String{Value: strings.ToLower(s.Value)}
}

// split splits the string by a given delimiter.
func (s *String) split(args []Object) Object {
	sep := ""
	if len(args) == 1 {
		arg, ok := args[0].(*String)
		if !ok {
			return newError("split() expects a STRING argument, got %s", args[0].Type())
		}
		sep = arg.Value
	} else if len(args) > 1 {
		return newError("split() expects at most 1 argument, got %d", len(args))
	}

	parts := strings.Split(s.Value, sep)
	elements := make([]Object, len(parts))
	for i, part := range parts {
		elements[i] = &String{Value: part}
	}
	return &Array{Elements: elements}
}

// trim removes leading and trailing whitespace or specified characters.
func (s *String) trim(args []Object) Object {
	if len(args) > 1 {
		return newError("trim() expects at most 1 argument, got %d", len(args))
	}

	chars := " \t\n\r"
	if len(args) == 1 {
		arg, ok := args[0].(*String)
		if !ok {
			return newError("trim() expects a STRING argument, got %s", args[0].Type())
		}
		chars = arg.Value
	}

	return &String{Value: strings.Trim(s.Value, chars)}
}

// contains checks if the string contains a given substring.
func (s *String) contains(args []Object) Object {
	if len(args) != 1 {
		return newError("contains() expects 1 argument, got %d", len(args))
	}

	arg, ok := args[0].(*String)
	if !ok {
		return newError("contains() expects a STRING argument, got %s", args[0].Type())
	}

	return &Boolean{Value: strings.Contains(s.Value, arg.Value)}
}

// replace replaces occurrences of a substring with another substring.
func (s *String) replace(args []Object) Object {
	if len(args) != 2 {
		return newError("replace() expects 2 arguments, got %d", len(args))
	}

	old, ok1 := args[0].(*String)
	new, ok2 := args[1].(*String)
	if !ok1 || !ok2 {
		return newError("replace() expects STRING arguments, got %s and %s", args[0].Type(), args[1].Type())
	}

	return &String{Value: strings.ReplaceAll(s.Value, old.Value, new.Value)}
}

// reverse reverses the string.
func (s *String) reverse(args []Object) Object {
	if len(args) != 0 {
		return newError("reverse() expects 0 arguments, got %d", len(args))
	}

	runes := []rune(s.Value)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return &String{Value: string(runes)}
}

// format applies formatting to the string with provided arguments.
// func (s *String) format(args []Object) Object {
// 	value, err := formatStr(s.Value, args)
// 	if err != nil {
// 		return newError(err.Error())
// 	}
// 	return &String{Value: value}
// }
