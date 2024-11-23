package object

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
)

// String represents a string object in the system
type String struct {
	Value  string
	offset int
}

func (s *String) Inspect() string  { return s.Value }
func (s *String) Type() ObjectType { return STRING_OBJ }

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

func (s *String) Next() (Object, Object) {
	offset := s.offset
	if len(s.Value) > offset {
		s.offset = offset + 1
		return &Integer{Value: int64(offset)}, &String{Value: string(s.Value[offset])}
	}
	return nil, nil
}

func (s *String) Reset() {
	s.offset = 0
}

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
	case "format":
		return s.format(args)
	default:
		return newError("Method '%s' is not supported on strings", method)
	}
}

func (s *String) len(args []Object) Object {
	if len(args) != 0 {
		return newError("Expected 0 arguments, but got %d", len(args))
	}
	return &Integer{Value: int64(len(s.Value))}
}

func (s *String) upper(args []Object) Object {
	if len(args) != 0 {
		return newError("Expected 0 arguments, but got %d", len(args))
	}
	return &String{Value: strings.ToUpper(s.Value)}
}

func (s *String) lower(args []Object) Object {
	if len(args) != 0 {
		return newError("Expected 0 arguments, but got %d", len(args))
	}
	return &String{Value: strings.ToLower(s.Value)}
}

func (s *String) split(args []Object) Object {
	if len(args) > 1 {
		return newError("Expected 1 or 0 arguments, but got %d", len(args))
	}
	sep := " "
	if len(args) == 1 {
		strArg, ok := args[0].(*String)
		if !ok {
			return newError("Expected argument of type STRING, but got %s", args[0].Type())
		}
		sep = strArg.Value
	}
	parts := strings.Split(s.Value, sep)
	elements := make([]Object, len(parts))
	for i, v := range parts {
		elements[i] = &String{Value: v}
	}
	return &Array{Elements: elements}
}

func (s *String) format(args []Object) Object {
	value, err := formatStr(s.Value, args)
	if err != nil {
		return newError(err.Error())
	}
	return &String{Value: value}
}

func formatStr(format string, options []Object) (string, error) {
	var str, val strings.Builder
	checkVal := false
	escapeChar := false
	optsLen := len(options)

	type optM struct {
		used bool
		obj  Object
	}

	optionsMap := make(map[int]optM, optsLen)
	for i, opt := range options {
		optionsMap[i] = optM{used: false, obj: opt}
	}

	for _, ch := range format {
		if !escapeChar && ch == '\\' {
			escapeChar = true
			continue
		}

		if ch == '{' && !escapeChar {
			checkVal = true
			continue
		}

		if escapeChar {
			if ch != '{' && ch != '}' {
				str.WriteRune('\\')
			}
			escapeChar = false
		}

		if checkVal && ch == '}' {
			index, err := strconv.Atoi(strings.TrimSpace(val.String()))
			if err != nil {
				return "", fmt.Errorf("invalid placeholder: `%s` is not a number", val.String())
			}
			if index >= optsLen {
				return "", fmt.Errorf("placeholder index %d exceeds available arguments (%d)", index, optsLen)
			}

			opt := optionsMap[index]
			str.WriteString(opt.obj.Inspect())
			optionsMap[index] = optM{used: true, obj: opt.obj}

			checkVal = false
			val.Reset()
			continue
		}

		if checkVal {
			val.WriteRune(ch)
			continue
		}

		str.WriteRune(ch)
	}

	if checkVal {
		return "", fmt.Errorf("unmatched '{' detected: `%s`", val.String())
	}

	for i, opt := range optionsMap {
		if !opt.used {
			return "", fmt.Errorf("argument at index %d (%s) was provided but not used", i, opt.obj.Inspect())
		}
	}

	return str.String(), nil
}
