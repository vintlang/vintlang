package object

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type DictPair struct {
	Key   VintObject
	Value VintObject
}

type Dict struct {
	Pairs  map[HashKey]DictPair
	offset int
}

func (d *Dict) Type() ObjectType { return DICT_OBJ }
func (d *Dict) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}

	for _, pair := range d.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

func (d *Dict) Next() (VintObject, VintObject) {
	idx := 0
	dict := make(map[string]DictPair)
	var keys []string
	for _, v := range d.Pairs {
		dict[v.Key.Inspect()] = v
		keys = append(keys, v.Key.Inspect())
	}

	sort.Strings(keys)

	for _, k := range keys {
		if d.offset == idx {
			d.offset += 1
			return dict[k].Key, dict[k].Value
		}
		idx += 1
	}
	return nil, nil
}

func (d *Dict) Reset() {
	d.offset = 0
}

func (d *Dict) Method(method string, args []VintObject) VintObject {
	switch method {
	case "keys":
		return d.keys(args)
	case "values":
		return d.values(args)
	case "size":
		return d.size(args)
	case "has":
		return d.has(args)
	case "get":
		return d.get(args)
	case "set":
		return d.set(args)
	case "remove":
		return d.remove(args)
	case "clear":
		return d.clear(args)
	case "merge":
		return d.merge(args)
	case "copy":
		return d.copy(args)
	default:
		return newError("Method '%s' is not supported for Dict objects", method)
	}
}

func (d *Dict) keys(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("keys() expects 0 arguments, got %d", len(args))
	}

	keys := make([]VintObject, 0, len(d.Pairs))
	for _, pair := range d.Pairs {
		keys = append(keys, pair.Key)
	}

	return &Array{Elements: keys}
}

func (d *Dict) values(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("values() expects 0 arguments, got %d", len(args))
	}

	values := make([]VintObject, 0, len(d.Pairs))
	for _, pair := range d.Pairs {
		values = append(values, pair.Value)
	}

	return &Array{Elements: values}
}

func (d *Dict) size(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("size() expects 0 arguments, got %d", len(args))
	}

	return &Integer{Value: int64(len(d.Pairs))}
}

func (d *Dict) has(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("has() expects 1 argument, got %d", len(args))
	}

	key, ok := args[0].(Hashable)
	if !ok {
		return newError("Key must be hashable")
	}

	_, exists := d.Pairs[key.HashKey()]
	return &Boolean{Value: exists}
}

func (d *Dict) get(args []VintObject) VintObject {
	if len(args) < 1 || len(args) > 2 {
		return newError("get() expects 1 or 2 arguments, got %d", len(args))
	}

	key, ok := args[0].(Hashable)
	if !ok {
		return newError("Key must be hashable")
	}

	pair, exists := d.Pairs[key.HashKey()]
	if exists {
		return pair.Value
	}

	if len(args) == 2 {
		return args[1] // default value
	}

	return &Null{}
}

func (d *Dict) set(args []VintObject) VintObject {
	if len(args) != 2 {
		return newError("set() expects 2 arguments, got %d", len(args))
	}

	key, ok := args[0].(Hashable)
	if !ok {
		return newError("Key must be hashable")
	}

	hashKey := key.HashKey()
	d.Pairs[hashKey] = DictPair{Key: args[0], Value: args[1]}

	return d
}

func (d *Dict) remove(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("remove() expects 1 argument, got %d", len(args))
	}

	key, ok := args[0].(Hashable)
	if !ok {
		return newError("Key must be hashable")
	}

	hashKey := key.HashKey()
	if _, exists := d.Pairs[hashKey]; exists {
		delete(d.Pairs, hashKey)
		return &Boolean{Value: true}
	}

	return &Boolean{Value: false}
}

func (d *Dict) clear(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("clear() expects 0 arguments, got %d", len(args))
	}

	d.Pairs = make(map[HashKey]DictPair)
	return d
}

func (d *Dict) merge(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("merge() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Dict)
	if !ok {
		return newError("Argument must be a Dict")
	}

	for hashKey, pair := range other.Pairs {
		d.Pairs[hashKey] = pair
	}

	return d
}

func (d *Dict) copy(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("copy() expects 0 arguments, got %d", len(args))
	}

	newPairs := make(map[HashKey]DictPair)
	for hashKey, pair := range d.Pairs {
		newPairs[hashKey] = pair
	}

	return &Dict{Pairs: newPairs}
}
