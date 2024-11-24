package evaluator

import (
	"github.com/ekilie/vint-lang/ast"
	"github.com/ekilie/vint-lang/object"
)

// evalCall evaluates a function call expression by:
// 1. Evaluating the function to be called.
// 2. Evaluating the arguments for the function.
// 3. Applying the function with the evaluated arguments.
func evalCall(node *ast.CallExpression, env *object.Environment) object.Object {
	// Evaluate the function expression
	function := Eval(node.Function, env)

	// Check if an error occurred while evaluating the function
	if isError(function) {
		return function
	}

	var args []object.Object

	// Evaluate the arguments based on the function type
	switch fn := function.(type) {
	case *object.Function:
		// If it's a regular function, evaluate its arguments
		args = evalArgsExpressions(node, fn, env)
	case *object.Package:
		// If it's a package, look for the 'init' function inside the package
		obj, ok := fn.Scope.Get("init")
		if !ok {
			return newError("Package does not have 'init'") // Return an error if 'init' is not found in the package
		}
		// If 'init' is found, evaluate its arguments
		args = evalArgsExpressions(node, obj.(*object.Function), env)
	default:
		// If the function is of unknown type, evaluate the arguments in the default manner
		args = evalExpressions(node.Arguments, env)
	}

	// If there is exactly one argument and it's an error, return the error
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	// Apply the evaluated function with the arguments
	return applyFunction(function, args, node.Token.Line)
}

// evalArgsExpressions evaluates the arguments passed to the function call.
// It handles both positional arguments and keyword arguments (assigned with `=`).
func evalArgsExpressions(node *ast.CallExpression, fn *object.Function, env *object.Environment) []object.Object {
	// Initialize an array for positional arguments and a dictionary for keyword arguments
	argsList := &object.Array{}
	argsHash := &object.Dict{}
	argsHash.Pairs = make(map[object.HashKey]object.DictPair)

	// Iterate through the arguments in the function call expression
	for _, exprr := range node.Arguments {
		switch exp := exprr.(type) {
		case *ast.Assign:
			// If the argument is an assignment (i.e., a keyword argument)
			val := Eval(exp.Value, env)
			if isError(val) {
				return []object.Object{val} // Return error if evaluation fails
			}
			var keyHash object.HashKey
			key := &object.String{Value: exp.Name.Value}
			keyHash = key.HashKey()
			// Add the keyword argument to the dictionary
			pair := object.DictPair{Key: key, Value: val}
			argsHash.Pairs[keyHash] = pair
		default:
			// For regular arguments, evaluate the expression and add to the positional argument list
			evaluated := Eval(exp, env)
			if isError(evaluated) {
				return []object.Object{evaluated} // Return error if evaluation fails
			}
			argsList.Elements = append(argsList.Elements, evaluated)
		}
	}

	// Prepare the final list of arguments, ensuring they match the function's parameters
	var result []object.Object
	var params = map[string]bool{}
	for _, exp := range fn.Parameters {
		params[exp.Value] = true
		if len(argsList.Elements) > 0 {
			// Use the positional arguments first
			result = append(result, argsList.Elements[0])
			argsList.Elements = argsList.Elements[1:]
		} else {
			// If no more positional arguments, try to use keyword arguments
			keyParam := &object.String{Value: exp.Value}
			keyParamHash := keyParam.HashKey()
			if valParam, ok := argsHash.Pairs[keyParamHash]; ok {
				// If a keyword argument is found for the parameter, use it
				result = append(result, valParam.Value)
				delete(argsHash.Pairs, keyParamHash)
			} else {
				// If no value is found for the parameter, check if a default value is provided
				if _e, _ok := fn.Defaults[exp.Value]; _ok {
					evaluated := Eval(_e, env)
					if isError(evaluated) {
						return []object.Object{evaluated} // Return error if default evaluation fails
					}
					result = append(result, evaluated)
				} else {
					// If no default is provided and no value is found, return an error
					return []object.Object{&object.Error{Message: "Missing argument"}}
				}
			}
		}
	}

	// Check if any extra keyword arguments are provided that don't match function parameters
	for _, pair := range argsHash.Pairs {
		if _, ok := params[pair.Key.(*object.String).Value]; ok {
			return []object.Object{&object.Error{Message: "Multiple arguments for a single parameter"}} // Return error if multiple values are given for a parameter
		}
	}

	// Return the list of evaluated arguments
	return result
}
