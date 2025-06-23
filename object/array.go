package object

import (
	"bytes"
	"sort"
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
	case "pop":
		return a.pop()
	case "shift":
		return a.shift()
	case "unshift":
		return a.unshift(args)
	case "reverse":
		return a.reverse()
	case "sort":
		return a.sort()
	case "map":
		return a.mapMethod(args)
	default:
		return newError("Sorry, the method '%s' is not supported for this object.", method)
	}
}

func (a *Array) len(args []Object) Object {
	if len(args) != 0 {
		return newError("Expected 0 arguments, but got %d.", len(args))
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
		return newError("Expected at most 1 argument, but got %d.", len(args))
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
		return newError("Expected exactly 1 argument, but got %d.", len(args))
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
		return newError("Expected exactly 1 argument, but got %d.", len(args))
	}

	for _, obj := range a.Elements {
		if obj.Inspect() == args[0].Inspect() && obj.Type() == args[0].Type() {
			return obj
		}
	}
	return &Null{}
}

func (a *Array) pop() Object {
	if len(a.Elements) == 0 {
		return &Null{}
	}
	last := a.Elements[len(a.Elements)-1]
	a.Elements = a.Elements[:len(a.Elements)-1]
	return last
}

func (a *Array) shift() Object {
	if len(a.Elements) == 0 {
		return &Null{}
	}
	first := a.Elements[0]
	a.Elements = a.Elements[1:]
	return first
}

func (a *Array) unshift(args []Object) Object {
	if len(args) == 0 {
		return newError("unshift() expects at least 1 argument, got 0")
	}
	a.Elements = append(args, a.Elements...)
	return a
}

func (a *Array) reverse() Object {
	length := len(a.Elements)
	for i := 0; i < length/2; i++ {
		a.Elements[i], a.Elements[length-1-i] = a.Elements[length-1-i], a.Elements[i]
	}
	return a
}

func (a *Array) sort() Object {
	// Only sorts arrays of Integers or Strings for simplicity
	if len(a.Elements) == 0 {
		return a
	}
	firstType := a.Elements[0].Type()
	switch firstType {
	case INTEGER_OBJ:
		ints := make([]int, len(a.Elements))
		for i, el := range a.Elements {
			intEl, ok := el.(*Integer)
			if !ok {
				return newError("sort() only supports arrays of integers or strings")
			}
			ints[i] = int(intEl.Value)
		}
		sort.Ints(ints)
		for i, v := range ints {
			a.Elements[i] = &Integer{Value: int64(v)}
		}
	case STRING_OBJ:
		strs := make([]string, len(a.Elements))
		for i, el := range a.Elements {
			strEl, ok := el.(*String)
			if !ok {
				return newError("sort() only supports arrays of integers or strings")
			}
			strs[i] = strEl.Value
		}
		sort.Strings(strs)
		for i, v := range strs {
			a.Elements[i] = &String{Value: v}
		}
	default:
		return newError("sort() only supports arrays of integers or strings")
	}
	return a
}

func (a *Array) mapMethod(args []Object) Object {
	if len(args) != 1 {
		return newError("map() expects exactly 1 argument, got %d", len(args))
	}
	fn, ok := args[0].(*Function)
	if !ok {
		return newError("map() expects a function as its argument")
	}
	mapped := make([]Object, len(a.Elements))
	for i, el := range a.Elements {
		// For simplicity, only pass the element as argument
		callObj, found := fn.Env.Get("__call__")
		if !found {
			return newError("map() function does not have a __call__ method")
		}
		builtin, ok := callObj.(*Builtin)
		if !ok {
			return newError("map() function's __call__ is not a builtin function")
		}
		result := builtin.Fn(el)
		mapped[i] = result
	}
	return &Array{Elements: mapped}
}
