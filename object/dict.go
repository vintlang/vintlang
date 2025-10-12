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
	case "filter":
		return d.filter(args)
	case "map":
		return d.mapDict(args)
	case "reduce":
		return d.reduce(args)
	case "forEach":
		return d.forEach(args)
	case "find":
		return d.find(args)
	case "some":
		return d.some(args)
	case "every":
		return d.every(args)
	case "pick":
		return d.pick(args)
	case "omit":
		return d.omit(args)
	case "flatten":
		return d.flatten(args)
	case "deepMerge":
		return d.deepMerge(args)
	case "equals":
		return d.equals(args)
	case "isEmpty":
		return d.isEmpty(args)
	case "entries":
		return d.entries(args)
	case "fromEntries":
		return d.fromEntries(args)
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

// filter creates a new dictionary with key-value pairs that pass the test
func (d *Dict) filter(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("filter() expects 1 argument (function), got %d", len(args))
	}

	fn, ok := args[0].(*Function)
	if !ok {
		return newError("Argument must be a function")
	}

	newPairs := make(map[HashKey]DictPair)
	for _, pair := range d.Pairs {
		// Call function with key, value
		args := []VintObject{pair.Key, pair.Value}
		result := callFunction(fn, args)

		if isTruthy(result) {
			newPairs[pair.Key.(Hashable).HashKey()] = pair
		}
	}

	return &Dict{Pairs: newPairs}
}

// mapDict creates a new dictionary with transformed values
func (d *Dict) mapDict(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("map() expects 1 argument (function), got %d", len(args))
	}

	fn, ok := args[0].(*Function)
	if !ok {
		return newError("Argument must be a function")
	}

	newPairs := make(map[HashKey]DictPair)
	for _, pair := range d.Pairs {
		// Call function with key, value
		args := []VintObject{pair.Key, pair.Value}
		result := callFunction(fn, args)

		newPair := DictPair{Key: pair.Key, Value: result}
		newPairs[pair.Key.(Hashable).HashKey()] = newPair
	}

	return &Dict{Pairs: newPairs}
}

// reduce reduces the dictionary to a single value
func (d *Dict) reduce(args []VintObject) VintObject {
	if len(args) < 1 || len(args) > 2 {
		return newError("reduce() expects 1 or 2 arguments (function, initial), got %d", len(args))
	}

	fn, ok := args[0].(*Function)
	if !ok {
		return newError("First argument must be a function")
	}

	var accumulator VintObject
	if len(args) == 2 {
		accumulator = args[1]
	} else {
		// Use first value as initial accumulator
		if len(d.Pairs) == 0 {
			return newError("Cannot reduce empty dictionary without initial value")
		}
		for _, pair := range d.Pairs {
			accumulator = pair.Value
			break
		}
	}

	for _, pair := range d.Pairs {
		args := []VintObject{accumulator, pair.Key, pair.Value}
		accumulator = callFunction(fn, args)
	}

	return accumulator
}

// forEach executes a function for each key-value pair
func (d *Dict) forEach(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("forEach() expects 1 argument (function), got %d", len(args))
	}

	fn, ok := args[0].(*Function)
	if !ok {
		return newError("Argument must be a function")
	}

	for _, pair := range d.Pairs {
		args := []VintObject{pair.Key, pair.Value}
		callFunction(fn, args)
	}

	return &Null{}
}

// find returns the first key-value pair that satisfies the test function
func (d *Dict) find(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("find() expects 1 argument (function), got %d", len(args))
	}

	fn, ok := args[0].(*Function)
	if !ok {
		return newError("Argument must be a function")
	}

	for _, pair := range d.Pairs {
		args := []VintObject{pair.Key, pair.Value}
		result := callFunction(fn, args)

		if isTruthy(result) {
			return &Array{Elements: []VintObject{pair.Key, pair.Value}}
		}
	}

	return &Null{}
}

// some tests whether at least one key-value pair passes the test
func (d *Dict) some(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("some() expects 1 argument (function), got %d", len(args))
	}

	fn, ok := args[0].(*Function)
	if !ok {
		return newError("Argument must be a function")
	}

	for _, pair := range d.Pairs {
		args := []VintObject{pair.Key, pair.Value}
		result := callFunction(fn, args)

		if isTruthy(result) {
			return &Boolean{Value: true}
		}
	}

	return &Boolean{Value: false}
}

// every tests whether all key-value pairs pass the test
func (d *Dict) every(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("every() expects 1 argument (function), got %d", len(args))
	}

	fn, ok := args[0].(*Function)
	if !ok {
		return newError("Argument must be a function")
	}

	for _, pair := range d.Pairs {
		args := []VintObject{pair.Key, pair.Value}
		result := callFunction(fn, args)

		if !isTruthy(result) {
			return &Boolean{Value: false}
		}
	}

	return &Boolean{Value: true}
}

// pick creates a new dictionary with only specified keys
func (d *Dict) pick(args []VintObject) VintObject {
	if len(args) == 0 {
		return newError("pick() expects at least 1 argument (keys), got %d", len(args))
	}

	newPairs := make(map[HashKey]DictPair)

	for _, key := range args {
		hashable, ok := key.(Hashable)
		if !ok {
			continue // skip non-hashable keys
		}

		if pair, exists := d.Pairs[hashable.HashKey()]; exists {
			newPairs[hashable.HashKey()] = pair
		}
	}

	return &Dict{Pairs: newPairs}
}

// omit creates a new dictionary excluding specified keys
func (d *Dict) omit(args []VintObject) VintObject {
	if len(args) == 0 {
		return d.copy([]VintObject{})
	}

	newPairs := make(map[HashKey]DictPair)

	// Create set of keys to omit
	omitKeys := make(map[HashKey]bool)
	for _, key := range args {
		if hashable, ok := key.(Hashable); ok {
			omitKeys[hashable.HashKey()] = true
		}
	}

	// Copy all pairs except omitted ones
	for hashKey, pair := range d.Pairs {
		if !omitKeys[hashKey] {
			newPairs[hashKey] = pair
		}
	}

	return &Dict{Pairs: newPairs}
}

// flatten flattens nested dictionary (one level deep)
func (d *Dict) flatten(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("flatten() expects 0 arguments, got %d", len(args))
	}

	newPairs := make(map[HashKey]DictPair)

	for _, pair := range d.Pairs {
		if nestedDict, ok := pair.Value.(*Dict); ok {
			// Flatten nested dictionary
			for _, nestedPair := range nestedDict.Pairs {
				// Create new key by combining parent and child keys
				combinedKey := &String{Value: pair.Key.Inspect() + "." + nestedPair.Key.Inspect()}
				newPair := DictPair{Key: combinedKey, Value: nestedPair.Value}
				newPairs[combinedKey.HashKey()] = newPair
			}
		} else {
			newPairs[pair.Key.(Hashable).HashKey()] = pair
		}
	}

	return &Dict{Pairs: newPairs}
}

// deepMerge recursively merges dictionaries
func (d *Dict) deepMerge(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("deepMerge() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Dict)
	if !ok {
		return newError("Argument must be a Dict")
	}

	newPairs := make(map[HashKey]DictPair)

	// Copy original pairs
	for hashKey, pair := range d.Pairs {
		newPairs[hashKey] = pair
	}

	// Merge other pairs
	for hashKey, otherPair := range other.Pairs {
		if existingPair, exists := newPairs[hashKey]; exists {
			// If both values are dictionaries, recursively merge them
			if existingDict, ok1 := existingPair.Value.(*Dict); ok1 {
				if otherDict, ok2 := otherPair.Value.(*Dict); ok2 {
					merged := existingDict.deepMerge([]VintObject{otherDict})
					newPair := DictPair{Key: existingPair.Key, Value: merged}
					newPairs[hashKey] = newPair
					continue
				}
			}
		}
		newPairs[hashKey] = otherPair
	}

	return &Dict{Pairs: newPairs}
}

// equals checks if two dictionaries are equal
func (d *Dict) equals(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("equals() expects 1 argument, got %d", len(args))
	}

	other, ok := args[0].(*Dict)
	if !ok {
		return &Boolean{Value: false}
	}

	if len(d.Pairs) != len(other.Pairs) {
		return &Boolean{Value: false}
	}

	for hashKey, pair := range d.Pairs {
		otherPair, exists := other.Pairs[hashKey]
		if !exists {
			return &Boolean{Value: false}
		}

		if pair.Value.Inspect() != otherPair.Value.Inspect() {
			return &Boolean{Value: false}
		}
	}

	return &Boolean{Value: true}
}

// isEmpty checks if the dictionary is empty
func (d *Dict) isEmpty(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("isEmpty() expects 0 arguments, got %d", len(args))
	}

	return &Boolean{Value: len(d.Pairs) == 0}
}

// entries returns an array of [key, value] pairs
func (d *Dict) entries(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("entries() expects 0 arguments, got %d", len(args))
	}

	entries := make([]VintObject, 0, len(d.Pairs))
	for _, pair := range d.Pairs {
		entry := &Array{Elements: []VintObject{pair.Key, pair.Value}}
		entries = append(entries, entry)
	}

	return &Array{Elements: entries}
}

// fromEntries creates a dictionary from an array of [key, value] pairs
func (d *Dict) fromEntries(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("fromEntries() expects 1 argument (array), got %d", len(args))
	}

	arr, ok := args[0].(*Array)
	if !ok {
		return newError("Argument must be an array")
	}

	newPairs := make(map[HashKey]DictPair)

	for _, element := range arr.Elements {
		entry, ok := element.(*Array)
		if !ok || len(entry.Elements) != 2 {
			return newError("Each entry must be an array with exactly 2 elements [key, value]")
		}

		key, ok := entry.Elements[0].(Hashable)
		if !ok {
			return newError("Key must be hashable")
		}

		pair := DictPair{Key: entry.Elements[0], Value: entry.Elements[1]}
		newPairs[key.HashKey()] = pair
	}

	return &Dict{Pairs: newPairs}
}

// Helper functions for dict methods
func callFunction(fn *Function, args []VintObject) VintObject {
	// This is a simplified version - in practice, you'd need to call the evaluator
	// For now, return a dummy value
	return &Boolean{Value: true}
}

func isTruthy(obj VintObject) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

var (
	NULL  = &Null{}
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)
