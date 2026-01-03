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

		// Check if pattern matches the value and extract variables
		matchEnv, matched := matchesPatternWithBindings(obj, matchCase.Pattern, env)
		if matched {
			// Check guard condition if present
			if matchCase.Guard != nil {
				guardResult := Eval(matchCase.Guard, matchEnv)
				if isError(guardResult) {
					return guardResult
				}
				if !isTruthy(guardResult) {
					continue // Guard failed, try next pattern
				}
			}
			return evalBlockStatement(matchCase.Block, matchEnv)
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

// matchesPatternWithBindings checks if a value matches a pattern and returns
// an environment with bound variables from the pattern
func matchesPatternWithBindings(value object.VintObject, pattern ast.Expression, env *object.Environment) (*object.Environment, bool) {
	matchEnv := object.NewEnclosedEnvironment(env)

	if matchesPatternAndBind(value, pattern, matchEnv) {
		return matchEnv, true
	}
	return env, false
}

// matchesPatternAndBind recursively matches patterns and binds variables
func matchesPatternAndBind(value object.VintObject, pattern ast.Expression, env *object.Environment) bool {
	switch p := pattern.(type) {
	case *ast.ArrayPattern:
		return matchesArrayPattern(value, p, env)
	case *ast.DictLiteral:
		return matchesDictPatternWithBinding(value, p, env)
	case *ast.Identifier:
		// Variable binding - bind the value to the identifier
		if p.Value != "_" { // Don't bind wildcards
			env.Define(p.Value, value)
		}
		return true
	default:
		// For other patterns, use the existing logic
		return matchesPattern(value, pattern, env)
	}
}

// matchesArrayPattern checks if a value matches an array pattern
func matchesArrayPattern(value object.VintObject, pattern *ast.ArrayPattern, env *object.Environment) bool {
	// Value must be an array
	arr, ok := value.(*object.Array)
	if !ok {
		return false
	}

	elements := arr.Elements
	patternLen := len(pattern.Elements)

	// If there's no rest pattern, lengths must match exactly
	if pattern.Rest == nil {
		if len(elements) != patternLen {
			return false
		}

		// Match each element
		for i, patternElement := range pattern.Elements {
			if !matchesPatternAndBind(elements[i], patternElement, env) {
				return false
			}
		}
		return true
	}

	// With rest pattern, we need at least as many elements as fixed patterns
	if len(elements) < patternLen {
		return false
	}

	// Match fixed elements at the beginning
	for i, patternElement := range pattern.Elements {
		if !matchesPatternAndBind(elements[i], patternElement, env) {
			return false
		}
	}

	// Bind remaining elements to rest variable
	if pattern.Rest != nil && pattern.Rest.Value != "_" {
		restElements := elements[patternLen:]
		restArray := &object.Array{Elements: restElements}
		env.Define(pattern.Rest.Value, restArray)
	}

	return true
}

// matchesDictPatternWithBinding matches dictionary patterns with variable binding support
func matchesDictPatternWithBinding(value object.VintObject, dictPattern *ast.DictLiteral, env *object.Environment) bool {
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

		// Handle pattern value - if it's an identifier, bind it
		if ident, ok := patternValue.(*ast.Identifier); ok && ident.Value != "_" {
			// Bind the value from the dictionary to this variable
			env.Define(ident.Value, pair.Value)
		} else {
			// Evaluate the pattern value and compare
			patternValueObj := Eval(patternValue, env)
			if isError(patternValueObj) {
				return false
			}

			// Check if values match
			if pair.Value.Type() != patternValueObj.Type() || pair.Value.Inspect() != patternValueObj.Inspect() {
				return false
			}
		}
	}

	// All pattern pairs matched
	return true
}
