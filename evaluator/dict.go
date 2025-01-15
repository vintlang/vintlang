package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalDictLiteral(node *ast.DictLiteral, env *object.Environment) object.Object {
	// Create a map to store key-value pairs for the dictionary
	pairs := make(map[object.HashKey]object.DictPair)

	// Iterate over the pairs in the dictionary literal node
	for keyNode, valueNode := range node.Pairs {
		// Evaluate the key and check for errors
		key := Eval(keyNode, env)
		if isError(key) {
			return key // Return the error if key evaluation fails
		}

		// Ensure the key is hashable, as it will be used for dictionary lookup
		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("Line %d: Hashing failed: %s", node.Token.Line, key.Type()) // Return an error if key is not hashable
		}

		// Evaluate the value and check for errors
		value := Eval(valueNode, env)
		if isError(value) {
			return value // Return the error if value evaluation fails
		}

		// Add the key-value pair to the dictionary's map
		hashed := hashKey.HashKey()
		pairs[hashed] = object.DictPair{Key: key, Value: value}
	}

	// Return the constructed dictionary object
	return &object.Dict{Pairs: pairs}
}
