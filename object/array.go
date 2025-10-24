package object

import (
	"bytes"
	"math"
	"sort"
	"strings"
)

type Array struct {
	Elements []VintObject
	offset   int
}

func (ao *Array) Type() VintObjectType { return ARRAY_OBJ }
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

func (ao *Array) Next() (VintObject, VintObject) {
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

func (a *Array) Method(method string, args []VintObject) VintObject {
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
	case "slice":
		return a.slice(args)
	case "concat":
		return a.concat(args)
	case "includes":
		return a.includes(args)
	case "every":
		return a.every(args)
	case "some":
		return a.some(args)
	case "reduce":
		return a.reduce(args)
	case "flatten":
		return a.flatten(args)
	case "unique":
		return a.unique(args)
	case "fill":
		return a.fill(args)
	case "lastIndexOf":
		return a.lastIndexOf(args)
	case "sum":
		return a.sum(args)
	case "average":
		return a.average(args)
	case "mean":
		return a.average(args) // alias for average
	case "min":
		return a.min(args)
	case "max":
		return a.max(args)
	case "median":
		return a.median(args)
	case "mode":
		return a.mode(args)
	case "variance":
		return a.variance(args)
	case "standardDeviation":
		return a.standardDeviation(args)
	case "product":
		return a.product(args)
	case "sortBy":
		return a.sortBy(args)
	case "sortDesc":
		return a.sortDesc(args)
	case "sortAsc":
		return a.sortAsc(args)
	default:
		return newError("Sorry, the method '%s' is not supported for this object.", method)
	}
}

func (a *Array) len(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("Expected 0 arguments, but got %d.", len(args))
	}
	return &Integer{Value: int64(len(a.Elements))}
}

func (a *Array) last() VintObject {
	length := len(a.Elements)
	if length > 0 {
		return a.Elements[length-1]
	}
	return &Null{}
}

func (a *Array) push(args []VintObject) VintObject {
	a.Elements = append(a.Elements, args...)
	return a
}

func (a *Array) join(args []VintObject) VintObject {
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

func (a *Array) filter(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("Expected exactly 1 argument, but got %d.", len(args))
	}

	dummy := []VintObject{}
	filteredArr := Array{Elements: dummy}
	for _, obj := range a.Elements {
		if obj.Inspect() == args[0].Inspect() && obj.Type() == args[0].Type() {
			filteredArr.Elements = append(filteredArr.Elements, obj)
		}
	}
	return &filteredArr
}

func (a *Array) find(args []VintObject) VintObject {
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

func (a *Array) pop() VintObject {
	if len(a.Elements) == 0 {
		return &Null{}
	}
	last := a.Elements[len(a.Elements)-1]
	a.Elements = a.Elements[:len(a.Elements)-1]
	return last
}

func (a *Array) shift() VintObject {
	if len(a.Elements) == 0 {
		return &Null{}
	}
	first := a.Elements[0]
	a.Elements = a.Elements[1:]
	return first
}

func (a *Array) unshift(args []VintObject) VintObject {
	if len(args) == 0 {
		return newError("unshift() expects at least 1 argument, got 0")
	}
	a.Elements = append(args, a.Elements...)
	return a
}

func (a *Array) reverse() VintObject {
	length := len(a.Elements)
	for i := 0; i < length/2; i++ {
		a.Elements[i], a.Elements[length-1-i] = a.Elements[length-1-i], a.Elements[i]
	}
	return a
}

func (a *Array) sort() VintObject {
	// Only sorts arrays of Integers, Floats, or Strings for simplicity
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
				return newError("sort() only supports arrays of integers, floats, or strings")
			}
			ints[i] = int(intEl.Value)
		}
		sort.Ints(ints)
		for i, v := range ints {
			a.Elements[i] = &Integer{Value: int64(v)}
		}
	case FLOAT_OBJ:
		floats := make([]float64, len(a.Elements))
		for i, el := range a.Elements {
			floatEl, ok := el.(*Float)
			if !ok {
				return newError("sort() only supports arrays of integers, floats, or strings")
			}
			floats[i] = floatEl.Value
		}
		sort.Float64s(floats)
		for i, v := range floats {
			a.Elements[i] = &Float{Value: v}
		}
	case STRING_OBJ:
		strs := make([]string, len(a.Elements))
		for i, el := range a.Elements {
			strEl, ok := el.(*String)
			if !ok {
				return newError("sort() only supports arrays of integers, floats, or strings")
			}
			strs[i] = strEl.Value
		}
		sort.Strings(strs)
		for i, v := range strs {
			a.Elements[i] = &String{Value: v}
		}
	default:
		return newError("sort() only supports arrays of integers, floats, or strings")
	}
	return a
}

func (a *Array) mapMethod(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("map() expects exactly 1 argument, got %d", len(args))
	}
	fn, ok := args[0].(*Function)
	if !ok {
		return newError("map() expects a function as its argument")
	}
	mapped := make([]VintObject, len(a.Elements))
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

// slice extracts a portion of the array between start and end indices
func (a *Array) slice(args []VintObject) VintObject {
	if len(args) < 1 || len(args) > 2 {
		return newError("slice() expects 1 or 2 arguments, got %d", len(args))
	}

	start, ok := args[0].(*Integer)
	if !ok {
		return newError("slice() start index must be an integer, got %s", args[0].Type())
	}

	startIdx := int(start.Value)
	endIdx := len(a.Elements)

	if len(args) == 2 {
		end, ok := args[1].(*Integer)
		if !ok {
			return newError("slice() end index must be an integer, got %s", args[1].Type())
		}
		endIdx = int(end.Value)
	}

	// Handle negative indices
	length := len(a.Elements)
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

	return &Array{Elements: a.Elements[startIdx:endIdx]}
}

// concat concatenates multiple arrays together
func (a *Array) concat(args []VintObject) VintObject {
	newElements := make([]VintObject, len(a.Elements))
	copy(newElements, a.Elements)

	for _, arg := range args {
		arr, ok := arg.(*Array)
		if !ok {
			return newError("concat() arguments must be arrays, got %s", arg.Type())
		}
		newElements = append(newElements, arr.Elements...)
	}

	return &Array{Elements: newElements}
}

// includes checks if the array contains a specific element
func (a *Array) includes(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("includes() expects exactly 1 argument, got %d", len(args))
	}

	target := args[0]
	for _, element := range a.Elements {
		if element.Inspect() == target.Inspect() && element.Type() == target.Type() {
			return &Boolean{Value: true}
		}
	}
	return &Boolean{Value: false}
}

// every checks if all elements satisfy a condition function
func (a *Array) every(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("every() expects exactly 1 argument, got %d", len(args))
	}

	fn, ok := args[0].(*Function)
	if !ok {
		return newError("every() expects a function as its argument")
	}

	for _, el := range a.Elements {
		callObj, found := fn.Env.Get("__call__")
		if !found {
			return newError("every() function does not have a __call__ method")
		}
		builtin, ok := callObj.(*Builtin)
		if !ok {
			return newError("every() function's __call__ is not a builtin function")
		}
		result := builtin.Fn(el)

		// Check if result is truthy
		if boolResult, ok := result.(*Boolean); ok && !boolResult.Value {
			return &Boolean{Value: false}
		}
	}

	return &Boolean{Value: true}
}

// some checks if any element satisfies a condition function
func (a *Array) some(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("some() expects exactly 1 argument, got %d", len(args))
	}

	fn, ok := args[0].(*Function)
	if !ok {
		return newError("some() expects a function as its argument")
	}

	for _, el := range a.Elements {
		callObj, found := fn.Env.Get("__call__")
		if !found {
			return newError("some() function does not have a __call__ method")
		}
		builtin, ok := callObj.(*Builtin)
		if !ok {
			return newError("some() function's __call__ is not a builtin function")
		}
		result := builtin.Fn(el)

		// Check if result is truthy
		if boolResult, ok := result.(*Boolean); ok && boolResult.Value {
			return &Boolean{Value: true}
		}
	}

	return &Boolean{Value: false}
}

// reduce reduces the array to a single value using an accumulator function
func (a *Array) reduce(args []VintObject) VintObject {
	if len(args) < 1 || len(args) > 2 {
		return newError("reduce() expects 1 or 2 arguments, got %d", len(args))
	}

	fn, ok := args[0].(*Function)
	if !ok {
		return newError("reduce() first argument must be a function")
	}

	if len(a.Elements) == 0 {
		if len(args) == 2 {
			return args[1] // Return initial value if array is empty
		}
		return newError("reduce() of empty array with no initial value")
	}

	var accumulator VintObject
	startIdx := 0

	if len(args) == 2 {
		accumulator = args[1]
	} else {
		accumulator = a.Elements[0]
		startIdx = 1
	}

	for i := startIdx; i < len(a.Elements); i++ {
		callObj, found := fn.Env.Get("__call__")
		if !found {
			return newError("reduce() function does not have a __call__ method")
		}
		builtin, ok := callObj.(*Builtin)
		if !ok {
			return newError("reduce() function's __call__ is not a builtin function")
		}
		accumulator = builtin.Fn(accumulator, a.Elements[i], &Integer{Value: int64(i)})
	}

	return accumulator
}

// flatten flattens nested arrays into a single array
func (a *Array) flatten(args []VintObject) VintObject {
	if len(args) > 1 {
		return newError("flatten() expects at most 1 argument, got %d", len(args))
	}

	depth := 1
	if len(args) == 1 {
		depthArg, ok := args[0].(*Integer)
		if !ok {
			return newError("flatten() depth must be an integer, got %s", args[0].Type())
		}
		depth = int(depthArg.Value)
		if depth < 0 {
			depth = -1 // Infinite depth
		}
	}

	flattened := a.flattenHelper(a.Elements, depth)
	return &Array{Elements: flattened}
}

// flattenHelper recursively flattens arrays
func (a *Array) flattenHelper(elements []VintObject, depth int) []VintObject {
	if depth == 0 {
		return elements
	}

	var result []VintObject
	for _, element := range elements {
		if arr, ok := element.(*Array); ok {
			if depth == -1 {
				result = append(result, a.flattenHelper(arr.Elements, -1)...)
			} else {
				result = append(result, a.flattenHelper(arr.Elements, depth-1)...)
			}
		} else {
			result = append(result, element)
		}
	}
	return result
}

// unique removes duplicate elements from the array
func (a *Array) unique(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("unique() expects 0 arguments, got %d", len(args))
	}

	seen := make(map[string]bool)
	var unique []VintObject

	for _, element := range a.Elements {
		key := string(element.Type()) + ":" + element.Inspect()
		if !seen[key] {
			seen[key] = true
			unique = append(unique, element)
		}
	}

	return &Array{Elements: unique}
}

// fill fills the array with a specified value
func (a *Array) fill(args []VintObject) VintObject {
	if len(args) < 1 || len(args) > 3 {
		return newError("fill() expects 1 to 3 arguments, got %d", len(args))
	}

	value := args[0]
	start := 0
	end := len(a.Elements)

	if len(args) >= 2 {
		startArg, ok := args[1].(*Integer)
		if !ok {
			return newError("fill() start index must be an integer, got %s", args[1].Type())
		}
		start = int(startArg.Value)
	}

	if len(args) == 3 {
		endArg, ok := args[2].(*Integer)
		if !ok {
			return newError("fill() end index must be an integer, got %s", args[2].Type())
		}
		end = int(endArg.Value)
	}

	// Handle negative indices
	length := len(a.Elements)
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}

	// Bound check
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if start > end {
		start = end
	}

	// Fill the array
	for i := start; i < end; i++ {
		a.Elements[i] = value
	}

	return a
}

// lastIndexOf finds the last index of an element in the array
func (a *Array) lastIndexOf(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("lastIndexOf() expects exactly 1 argument, got %d", len(args))
	}

	target := args[0]
	for i := len(a.Elements) - 1; i >= 0; i-- {
		element := a.Elements[i]
		if element.Inspect() == target.Inspect() && element.Type() == target.Type() {
			return &Integer{Value: int64(i)}
		}
	}

	return &Integer{Value: -1} // Not found
}

// Mathematical array methods

// sum calculates the sum of all numeric elements in the array
func (a *Array) sum(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("sum() expects 0 arguments, got %d", len(args))
	}

	if len(a.Elements) == 0 {
		return &Integer{Value: 0} // Sum of empty array is 0
	}

	var sum float64
	for _, element := range a.Elements {
		switch elem := element.(type) {
		case *Integer:
			sum += float64(elem.Value)
		case *Float:
			sum += elem.Value
		default:
			return newError("sum() can only be applied to arrays containing numbers")
		}
	}

	// Return integer if the sum is a whole number, otherwise return float
	if sum == float64(int64(sum)) {
		return &Integer{Value: int64(sum)}
	}
	return &Float{Value: sum}
}

// average calculates the average (mean) of all numeric elements in the array
func (a *Array) average(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("average() expects 0 arguments, got %d", len(args))
	}

	if len(a.Elements) == 0 {
		return newError("average() cannot be calculated for an empty array")
	}

	var sum float64
	for _, element := range a.Elements {
		switch elem := element.(type) {
		case *Integer:
			sum += float64(elem.Value)
		case *Float:
			sum += elem.Value
		default:
			return newError("average() can only be applied to arrays containing numbers")
		}
	}

	avg := sum / float64(len(a.Elements))
	return &Float{Value: avg}
}

// min finds the minimum value in the array
func (a *Array) min(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("min() expects 0 arguments, got %d", len(args))
	}

	if len(a.Elements) == 0 {
		return newError("min() cannot be calculated for an empty array")
	}

	var min float64
	initialized := false

	for _, element := range a.Elements {
		var value float64
		switch elem := element.(type) {
		case *Integer:
			value = float64(elem.Value)
		case *Float:
			value = elem.Value
		default:
			return newError("min() can only be applied to arrays containing numbers")
		}

		if !initialized || value < min {
			min = value
			initialized = true
		}
	}

	// Return in the same type as the minimum element if possible
	for _, element := range a.Elements {
		switch elem := element.(type) {
		case *Integer:
			if float64(elem.Value) == min {
				return elem
			}
		case *Float:
			if elem.Value == min {
				return elem
			}
		}
	}

	return &Float{Value: min}
}

// max finds the maximum value in the array
func (a *Array) max(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("max() expects 0 arguments, got %d", len(args))
	}

	if len(a.Elements) == 0 {
		return newError("max() cannot be calculated for an empty array")
	}

	var max float64
	initialized := false

	for _, element := range a.Elements {
		var value float64
		switch elem := element.(type) {
		case *Integer:
			value = float64(elem.Value)
		case *Float:
			value = elem.Value
		default:
			return newError("max() can only be applied to arrays containing numbers")
		}

		if !initialized || value > max {
			max = value
			initialized = true
		}
	}

	// Return in the same type as the maximum element if possible
	for _, element := range a.Elements {
		switch elem := element.(type) {
		case *Integer:
			if float64(elem.Value) == max {
				return elem
			}
		case *Float:
			if elem.Value == max {
				return elem
			}
		}
	}

	return &Float{Value: max}
}

// median calculates the median value of numeric elements in the array
func (a *Array) median(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("median() expects 0 arguments, got %d", len(args))
	}

	if len(a.Elements) == 0 {
		return newError("median() cannot be calculated for an empty array")
	}

	// Convert all elements to float64 and sort them
	values := make([]float64, len(a.Elements))
	for i, element := range a.Elements {
		switch elem := element.(type) {
		case *Integer:
			values[i] = float64(elem.Value)
		case *Float:
			values[i] = elem.Value
		default:
			return newError("median() can only be applied to arrays containing numbers")
		}
	}

	sort.Float64s(values)

	n := len(values)
	if n%2 == 1 {
		// Odd number of elements
		median := values[n/2]
		return &Float{Value: median}
	} else {
		// Even number of elements
		median := (values[n/2-1] + values[n/2]) / 2
		return &Float{Value: median}
	}
}

// mode finds the most frequently occurring value(s) in the array
func (a *Array) mode(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("mode() expects 0 arguments, got %d", len(args))
	}

	if len(a.Elements) == 0 {
		return newError("mode() cannot be calculated for an empty array")
	}

	// Count frequencies
	freq := make(map[string]int)
	elementMap := make(map[string]VintObject)

	for _, element := range a.Elements {
		key := string(element.Type()) + ":" + element.Inspect()
		freq[key]++
		elementMap[key] = element
	}

	// Find maximum frequency
	maxFreq := 0
	for _, count := range freq {
		if count > maxFreq {
			maxFreq = count
		}
	}

	// Collect all elements with maximum frequency
	var modes []VintObject
	for key, count := range freq {
		if count == maxFreq {
			modes = append(modes, elementMap[key])
		}
	}

	return &Array{Elements: modes}
}

// variance calculates the variance of numeric elements in the array
func (a *Array) variance(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("variance() expects 0 arguments, got %d", len(args))
	}

	if len(a.Elements) == 0 {
		return newError("variance() cannot be calculated for an empty array")
	}

	// Calculate mean first
	var sum float64
	for _, element := range a.Elements {
		switch elem := element.(type) {
		case *Integer:
			sum += float64(elem.Value)
		case *Float:
			sum += elem.Value
		default:
			return newError("variance() can only be applied to arrays containing numbers")
		}
	}

	mean := sum / float64(len(a.Elements))

	// Calculate variance
	var variance float64
	for _, element := range a.Elements {
		var value float64
		switch elem := element.(type) {
		case *Integer:
			value = float64(elem.Value)
		case *Float:
			value = elem.Value
		}
		variance += math.Pow(value-mean, 2)
	}

	variance /= float64(len(a.Elements))
	return &Float{Value: variance}
}

// standardDeviation calculates the standard deviation of numeric elements in the array
func (a *Array) standardDeviation(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("standardDeviation() expects 0 arguments, got %d", len(args))
	}

	varianceResult := a.variance(args)
	if varianceResult.Type() == ERROR_OBJ {
		return varianceResult
	}

	variance := varianceResult.(*Float).Value
	stdDev := math.Sqrt(variance)
	return &Float{Value: stdDev}
}

// product calculates the product of all numeric elements in the array
func (a *Array) product(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("product() expects 0 arguments, got %d", len(args))
	}

	if len(a.Elements) == 0 {
		return &Integer{Value: 1} // Mathematical convention: empty product is 1
	}

	var product float64 = 1
	for _, element := range a.Elements {
		switch elem := element.(type) {
		case *Integer:
			product *= float64(elem.Value)
		case *Float:
			product *= elem.Value
		default:
			return newError("product() can only be applied to arrays containing numbers")
		}
	}

	// Return integer if the product is a whole number, otherwise return float
	if product == float64(int64(product)) {
		return &Integer{Value: int64(product)}
	}
	return &Float{Value: product}
}

// Enhanced sorting methods

// sortBy sorts the array using a custom comparison function
func (a *Array) sortBy(args []VintObject) VintObject {
	// This method is handled by the evaluator for proper function calling
	return newError("sortBy() should be handled by the evaluator")
}

// sortDesc sorts the array in descending order
func (a *Array) sortDesc(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("sortDesc() expects 0 arguments, got %d", len(args))
	}

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
				return newError("sortDesc() only supports arrays of integers or strings")
			}
			ints[i] = int(intEl.Value)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(ints)))
		for i, v := range ints {
			a.Elements[i] = &Integer{Value: int64(v)}
		}
	case STRING_OBJ:
		strs := make([]string, len(a.Elements))
		for i, el := range a.Elements {
			strEl, ok := el.(*String)
			if !ok {
				return newError("sortDesc() only supports arrays of integers or strings")
			}
			strs[i] = strEl.Value
		}
		sort.Sort(sort.Reverse(sort.StringSlice(strs)))
		for i, v := range strs {
			a.Elements[i] = &String{Value: v}
		}
	case FLOAT_OBJ:
		floats := make([]float64, len(a.Elements))
		for i, el := range a.Elements {
			floatEl, ok := el.(*Float)
			if !ok {
				return newError("sortDesc() only supports arrays of integers, floats, or strings")
			}
			floats[i] = floatEl.Value
		}
		sort.Sort(sort.Reverse(sort.Float64Slice(floats)))
		for i, v := range floats {
			a.Elements[i] = &Float{Value: v}
		}
	default:
		return newError("sortDesc() only supports arrays of integers, floats, or strings")
	}
	return a
}

// sortAsc sorts the array in ascending order (alias for existing sort method)
func (a *Array) sortAsc(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("sortAsc() expects 0 arguments, got %d", len(args))
	}
	return a.sort()
}
