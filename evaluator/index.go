package evaluator

import "github.com/vintlang/vintlang/object"

// evalIndexExpression handles indexing operations for arrays and dictionaries.
func evalIndexExpression(left, index object.Object, line int) object.Object {
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

// evalArrayIndexExpression evaluates an array index expression.
func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return arrayObject.Elements[idx]
}

// evalDictIndexExpression evaluates a dictionary index expression.
func evalDictIndexExpression(dict, index object.Object, line int) object.Object {
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
