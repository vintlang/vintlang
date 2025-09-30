package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalMatchExpression(me *ast.MatchExpression, env *object.Environment) object.VintObject {
	obj := Eval(me.Value, env)
	if isError(obj) {
		return obj
	}

	for _, matchCase := range me.Cases {
		if matchCase.Pattern == nil {
			continue
		}

		// Check for wildcard pattern "_"
		if ident, ok := matchCase.Pattern.(*ast.Identifier); ok && ident.Value == "_" {
			// This is the default/wildcard case - execute if no other cases match
			// We'll handle this at the end
			continue
		}

		// Check if pattern matches the value
		if matchesPattern(obj, matchCase.Pattern, env) {
			return evalBlockStatement(matchCase.Block, env)
		}
	}

	// If no specific pattern matched, look for wildcard pattern
	for _, matchCase := range me.Cases {
		if ident, ok := matchCase.Pattern.(*ast.Identifier); ok && ident.Value == "_" {
			return evalBlockStatement(matchCase.Block, env)
		}
	}

	return NULL
}

func matchesPattern(value object.VintObject, pattern ast.Expression, env *object.Environment) bool {
	// Handle dict pattern matching
	if dictPattern, ok := pattern.(*ast.DictLiteral); ok {
		return matchesDictPattern(value, dictPattern, env)
	}

	// For other patterns, evaluate the pattern and compare
	patternObj := Eval(pattern, env)
	if isError(patternObj) {
		return false
	}

	return value.Type() == patternObj.Type() && value.Inspect() == patternObj.Inspect()
}

func matchesDictPattern(value object.VintObject, dictPattern *ast.DictLiteral, env *object.Environment) bool {
	// Value must be a dict
	dict, ok := value.(*object.Dict)
	if !ok {
		return false
	}

	// For each key-value pair in the pattern, check if it exists in the dict
	for patternKey, patternValue := range dictPattern.Pairs {
		// Evaluate the pattern key
		keyObj := Eval(patternKey, env)
		if isError(keyObj) {
			return false
		}

		// Check if the key exists in the dict
		hashable, ok := keyObj.(object.Hashable)
		if !ok {
			return false
		}

		pair, exists := dict.Pairs[hashable.HashKey()]
		if !exists {
			return false
		}

		// Evaluate the pattern value
		patternValueObj := Eval(patternValue, env)
		if isError(patternValueObj) {
			return false
		}

		// Check if values match
		if pair.Value.Type() != patternValueObj.Type() || pair.Value.Inspect() != patternValueObj.Inspect() {
			return false
		}
	}

	// All pattern pairs matched
	return true
}
