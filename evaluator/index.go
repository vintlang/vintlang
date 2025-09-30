package evaluator

import "github.com/vintlang/vintlang/object"

// evalIndexExpression handles indexing operations for arrays and dictionaries.
func evalIndexExpression(left, index object.VintObject, line int) object.VintObject {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.ARRAY_OBJ && index.Type() != object.INTEGER_OBJ:
		return newError("Line %d: Please use a number, not: %s", line, index.Type())
	case left.Type() == object.DICT_OBJ:
		return evalDictIndexExpression(left, index, line)
	default:
		return newError("Line %d: This operation is not possible for: %s", line, left.Type())
	}
}

// evalSliceExpression handles slicing operations for arrays.
func evalSliceExpression(left, start, end object.VintObject, line int) object.VintObject {
	if left.Type() != object.ARRAY_OBJ {
		return newError("Line %d: Slicing is only supported for arrays, not: %s", line, left.Type())
	}

	arrayObject := left.(*object.Array)
	arrayLen := len(arrayObject.Elements)

	// Default values for start and end
	startIdx := 0
	endIdx := arrayLen

	// Parse start index
	if start != nil {
		if start.Type() != object.INTEGER_OBJ {
			return newError("Line %d: Slice start index must be an integer, not: %s", line, start.Type())
		}
		startIdx = int(start.(*object.Integer).Value)
		if startIdx < 0 {
			startIdx = arrayLen + startIdx // Handle negative indices
		}
		if startIdx < 0 {
			startIdx = 0
		}
		if startIdx > arrayLen {
			startIdx = arrayLen
		}
	}

	// Parse end index
	if end != nil {
		if end.Type() != object.INTEGER_OBJ {
			return newError("Line %d: Slice end index must be an integer, not: %s", line, end.Type())
		}
		endIdx = int(end.(*object.Integer).Value)
		if endIdx < 0 {
			endIdx = arrayLen + endIdx // Handle negative indices
		}
		if endIdx < 0 {
			endIdx = 0
		}
		if endIdx > arrayLen {
			endIdx = arrayLen
		}
	}

	// Ensure start <= end
	if startIdx > endIdx {
		startIdx = endIdx
	}

	// Create the sliced array
	slicedElements := make([]object.VintObject, endIdx-startIdx)
	copy(slicedElements, arrayObject.Elements[startIdx:endIdx])

	return &object.Array{Elements: slicedElements}
}

// evalArrayIndexExpression evaluates an array index expression.
func evalArrayIndexExpression(array, index object.VintObject) object.VintObject {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return arrayObject.Elements[idx]
}

// evalDictIndexExpression evaluates a dictionary index expression.
func evalDictIndexExpression(dict, index object.VintObject, line int) object.VintObject {
	dictObject := dict.(*object.Dict)

	// Ensure the index can be used as a key in the dictionary (Hashable).
	key, ok := index.(object.Hashable)
	if !ok {
		return newError("Line %d: Sorry, %s cannot be used as a key", line, index.Type())
	}

	// Look up the key in the dictionary and return the value.
	pair, ok := dictObject.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}

	return pair.Value
}
