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
	case "charAt":
		return s.charAt(args)
	case "substring":
		return s.substring(args)
	case "indexOf":
		return s.indexOf(args)
	case "lastIndexOf":
		return s.lastIndexOf(args)
	case "times":
		return s.times(args)
	case "padStart":
		return s.padStart(args)
	case "padEnd":
		return s.padEnd(args)
	case "slice":
		return s.slice(args)
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

// charAt returns the character at the specified index
func (s *String) charAt(args []Object) Object {
	if len(args) != 1 {
		return newError("charAt() expects exactly 1 argument, got %d", len(args))
	}
	
	index, ok := args[0].(*Integer)
	if !ok {
		return newError("charAt() index must be an integer, got %s", args[0].Type())
	}
	
	idx := int(index.Value)
	runes := []rune(s.Value)
	
	if idx < 0 || idx >= len(runes) {
		return &String{Value: ""} // Return empty string for out of bounds
	}
	
	return &String{Value: string(runes[idx])}
}

// substring extracts a substring between start and end indices
func (s *String) substring(args []Object) Object {
	if len(args) < 1 || len(args) > 2 {
		return newError("substring() expects 1 or 2 arguments, got %d", len(args))
	}
	
	start, ok := args[0].(*Integer)
	if !ok {
		return newError("substring() start index must be an integer, got %s", args[0].Type())
	}
	
	runes := []rune(s.Value)
	startIdx := int(start.Value)
	endIdx := len(runes)
	
	if len(args) == 2 {
		end, ok := args[1].(*Integer)
		if !ok {
			return newError("substring() end index must be an integer, got %s", args[1].Type())
		}
		endIdx = int(end.Value)
	}
	
	// Handle negative indices and bounds
	if startIdx < 0 {
		startIdx = 0
	}
	if endIdx < 0 {
		endIdx = 0
	}
	if startIdx > len(runes) {
		startIdx = len(runes)
	}
	if endIdx > len(runes) {
		endIdx = len(runes)
	}
	if startIdx > endIdx {
		startIdx, endIdx = endIdx, startIdx // Swap if start > end
	}
	
	return &String{Value: string(runes[startIdx:endIdx])}
}

// indexOf finds the index of the first occurrence of a substring
func (s *String) indexOf(args []Object) Object {
	if len(args) < 1 || len(args) > 2 {
		return newError("indexOf() expects 1 or 2 arguments, got %d", len(args))
	}
	
	searchStr, ok := args[0].(*String)
	if !ok {
		return newError("indexOf() search string must be a string, got %s", args[0].Type())
	}
	
	fromIndex := 0
	if len(args) == 2 {
		from, ok := args[1].(*Integer)
		if !ok {
			return newError("indexOf() from index must be an integer, got %s", args[1].Type())
		}
		fromIndex = int(from.Value)
	}
	
	if fromIndex < 0 {
		fromIndex = 0
	}
	
	if fromIndex >= len(s.Value) {
		return &Integer{Value: -1}
	}
	
	substr := s.Value[fromIndex:]
	index := strings.Index(substr, searchStr.Value)
	if index == -1 {
		return &Integer{Value: -1}
	}
	
	return &Integer{Value: int64(fromIndex + index)}
}

// lastIndexOf finds the index of the last occurrence of a substring
func (s *String) lastIndexOf(args []Object) Object {
	if len(args) < 1 || len(args) > 2 {
		return newError("lastIndexOf() expects 1 or 2 arguments, got %d", len(args))
	}
	
	searchStr, ok := args[0].(*String)
	if !ok {
		return newError("lastIndexOf() search string must be a string, got %s", args[0].Type())
	}
	
	fromIndex := len(s.Value)
	if len(args) == 2 {
		from, ok := args[1].(*Integer)
		if !ok {
			return newError("lastIndexOf() from index must be an integer, got %s", args[1].Type())
		}
		fromIndex = int(from.Value)
	}
	
	if fromIndex < 0 {
		return &Integer{Value: -1}
	}
	if fromIndex >= len(s.Value) {
		fromIndex = len(s.Value) - 1
	}
	
	substr := s.Value[:fromIndex+1]
	index := strings.LastIndex(substr, searchStr.Value)
	
	return &Integer{Value: int64(index)}
}

// times repeats the string n times
func (s *String) times(args []Object) Object {
	if len(args) != 1 {
		return newError("times() expects exactly 1 argument, got %d", len(args))
	}
	
	count, ok := args[0].(*Integer)
	if !ok {
		return newError("times() count must be an integer, got %s", args[0].Type())
	}
	
	n := int(count.Value)
	if n < 0 {
		return newError("times() count cannot be negative")
	}
	
	return &String{Value: strings.Repeat(s.Value, n)}
}

// padStart pads the string with a specified string until it reaches the target length
func (s *String) padStart(args []Object) Object {
	if len(args) < 1 || len(args) > 2 {
		return newError("padStart() expects 1 or 2 arguments, got %d", len(args))
	}
	
	targetLength, ok := args[0].(*Integer)
	if !ok {
		return newError("padStart() target length must be an integer, got %s", args[0].Type())
	}
	
	padString := " " // Default padding
	if len(args) == 2 {
		pad, ok := args[1].(*String)
		if !ok {
			return newError("padStart() pad string must be a string, got %s", args[1].Type())
		}
		padString = pad.Value
	}
	
	targetLen := int(targetLength.Value)
	currentLen := len([]rune(s.Value))
	
	if targetLen <= currentLen {
		return s // No padding needed
	}
	
	padLen := targetLen - currentLen
	if len(padString) == 0 {
		return s // Cannot pad with empty string
	}
	
	// Repeat pad string and truncate to exact length needed
	repeats := (padLen / len([]rune(padString))) + 1
	padding := strings.Repeat(padString, repeats)
	paddingRunes := []rune(padding)
	
	if len(paddingRunes) > padLen {
		paddingRunes = paddingRunes[:padLen]
	}
	
	return &String{Value: string(paddingRunes) + s.Value}
}

// padEnd pads the string with a specified string until it reaches the target length
func (s *String) padEnd(args []Object) Object {
	if len(args) < 1 || len(args) > 2 {
		return newError("padEnd() expects 1 or 2 arguments, got %d", len(args))
	}
	
	targetLength, ok := args[0].(*Integer)
	if !ok {
		return newError("padEnd() target length must be an integer, got %s", args[0].Type())
	}
	
	padString := " " // Default padding
	if len(args) == 2 {
		pad, ok := args[1].(*String)
		if !ok {
			return newError("padEnd() pad string must be a string, got %s", args[1].Type())
		}
		padString = pad.Value
	}
	
	targetLen := int(targetLength.Value)
	currentLen := len([]rune(s.Value))
	
	if targetLen <= currentLen {
		return s // No padding needed
	}
	
	padLen := targetLen - currentLen
	if len(padString) == 0 {
		return s // Cannot pad with empty string
	}
	
	// Repeat pad string and truncate to exact length needed
	repeats := (padLen / len([]rune(padString))) + 1
	padding := strings.Repeat(padString, repeats)
	paddingRunes := []rune(padding)
	
	if len(paddingRunes) > padLen {
		paddingRunes = paddingRunes[:padLen]
	}
	
	return &String{Value: s.Value + string(paddingRunes)}
}

// slice extracts a portion of the string between start and end indices
func (s *String) slice(args []Object) Object {
	if len(args) < 1 || len(args) > 2 {
		return newError("slice() expects 1 or 2 arguments, got %d", len(args))
	}
	
	start, ok := args[0].(*Integer)
	if !ok {
		return newError("slice() start index must be an integer, got %s", args[0].Type())
	}
	
	runes := []rune(s.Value)
	length := len(runes)
	startIdx := int(start.Value)
	endIdx := length
	
	if len(args) == 2 {
		end, ok := args[1].(*Integer)
		if !ok {
			return newError("slice() end index must be an integer, got %s", args[1].Type())
		}
		endIdx = int(end.Value)
	}
	
	// Handle negative indices
	if startIdx < 0 {
		startIdx = length + startIdx
	}
	if endIdx < 0 {
		endIdx = length + endIdx
	}
	
	// Bound check
	if startIdx < 0 {
		startIdx = 0
	}
	if endIdx > length {
		endIdx = length
	}
	if startIdx > endIdx {
		startIdx = endIdx
	}
	
	return &String{Value: string(runes[startIdx:endIdx])}
}

// format applies formatting to the string with provided arguments.
// func (s *String) format(args []Object) Object {
// 	value, err := formatStr(s.Value, args)
// 	if err != nil {
// 		return newError(err.Error())
// 	}
// 	return &String{Value: value}
// }
