package object

import (
	"bytes"
	"strings"
)

type Array struct {
	Elements []Object
	offset   int
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	if len(ao.Elements) != 0 {
		for _, e := range ao.Elements {
			if e.Inspect() != "" {
				elements = append(elements, e.Inspect())
			}
		}
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (ao *Array) Next() (Object, Object) {
	idx := ao.offset
	if len(ao.Elements) > idx {
		ao.offset = idx + 1
		return &Integer{Value: int64(idx)}, ao.Elements[idx]
	}
	return nil, nil
}

func (ao *Array) Reset() {
	ao.offset = 0
}

func (a *Array) Method(method string, args []Object) Object {
	switch method {
	case "length":
		return a.len(args)
	case "push":
		return a.push(args)
	case "last":
		return a.last()
	case "join":
		return a.join(args)
	case "filter":
		return a.filter(args)
	case "find":
		return a.find(args)
	default:
		return newError("Sorry, the method '%s' is not supported for this object.", method)
	}
}

func (a *Array) len(args []Object) Object {
	if len(args) != 0 {
		return newError("Error: Expected 0 arguments, but got %d.", len(args))
	}
	return &Integer{Value: int64(len(a.Elements))}
}

func (a *Array) last() Object {
	length := len(a.Elements)
	if length > 0 {
		return a.Elements[length-1]
	}
	return &Null{}
}

func (a *Array) push(args []Object) Object {
	a.Elements = append(a.Elements, args...)
	return a
}

func (a *Array) join(args []Object) Object {
	if len(args) > 1 {
		return newError("Error: Expected at most 1 argument, but got %d.", len(args))
	}
	if len(a.Elements) > 0 {
		glue := ""
		if len(args) == 1 {
			glue = args[0].(*String).Value
		}
		length := len(a.Elements)
		newElements := make([]string, length)
		for k, v := range a.Elements {
			newElements[k] = v.Inspect()
		}
		return &String{Value: strings.Join(newElements, glue)}
	} else {
		return &String{Value: ""}
	}
}

func (a *Array) filter(args []Object) Object {
	if len(args) != 1 {
		return newError("Error: Expected exactly 1 argument, but got %d.", len(args))
	}

	dummy := []Object{}
	filteredArr := Array{Elements: dummy}
	for _, obj := range a.Elements {
		if obj.Inspect() == args[0].Inspect() && obj.Type() == args[0].Type() {
			filteredArr.Elements = append(filteredArr.Elements, obj)
		}
	}
	return &filteredArr
}

func (a *Array) find(args []Object) Object {
	if len(args) != 1 {
		return newError("Error: Expected exactly 1 argument, but got %d.", len(args))
	}

	for _, obj := range a.Elements {
		if obj.Inspect() == args[0].Inspect() && obj.Type() == args[0].Type() {
			return obj
		}
	}
	return &Null{}
}
