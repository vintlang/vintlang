package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalMethodExpression(node *ast.MethodExpression, env *object.Environment) object.VintObject {
	obj := Eval(node.Object, env)
	if isError(obj) {
		return obj
	}
	args := evalExpressions(node.Arguments, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	defs := make(map[string]object.VintObject)

	for k, v := range node.Defaults {
		defs[k] = Eval(v, env)
	}
	return applyMethod(obj, node.Method, args, defs, node.Token.Line)
}

func applyMethod(obj object.VintObject, method ast.Expression, args []object.VintObject, defs map[string]object.VintObject, l int) object.VintObject {
	switch obj := obj.(type) {
	case *object.String:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.File:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.Time:
		return obj.Method(method.(*ast.Identifier).Value, args, defs)
	case *object.Duration:
		return obj.Method(method.(*ast.Identifier).Value, args, defs)
	case *object.Integer:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.Float:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.Boolean:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.Array:
		switch method.(*ast.Identifier).Value {
		case "map":
			return maap(obj, args)
		case "filter":
			return filter(obj, args)
		case "sortBy":
			return sortBy(obj, args)
		default:
			return obj.Method(method.(*ast.Identifier).Value, args)
		}
	case *object.Dict:
		switch method.(*ast.Identifier).Value {
		case "has_key":
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			key, ok := args[0].(object.Hashable)
			if !ok {
				return newError("argument to `has_key` must be hashable, got %s", args[0].Type())
			}
			if _, ok := obj.Pairs[key.HashKey()]; ok {
				return TRUE
			}
			return FALSE
		case "filter":
			return dictFilter(obj, args)
		case "map":
			return dictMap(obj, args)
		case "reduce":
			return dictReduce(obj, args)
		case "forEach":
			return dictForEach(obj, args)
		case "find":
			return dictFind(obj, args)
		case "some":
			return dictSome(obj, args)
		case "every":
			return dictEvery(obj, args)
		default:
			return obj.Method(method.(*ast.Identifier).Value, args)
		}
	case *object.Module:
		if fn, ok := obj.Functions[method.(*ast.Identifier).Value]; ok {
			return fn(args, defs)
		}
	case *object.Instance:
		if val, ok := obj.Package.Scope.Get(method.(*ast.Identifier).Value); ok {
			fn, ok := val.(*object.Function)
			if !ok {
				return newError("'%s' is not a function", method.(*ast.Identifier).Value)
			}
			fn.Env.Define("@", obj)
			ret := applyFunction(fn, args, l)
			fn.Env.Del("@")
			return ret
		}
	case *object.Package:
		// Here We Use GetPublic to enforce private member access control
		if val, ok := obj.GetPublic(method.(*ast.Identifier).Value); ok {
			fn, ok := val.(*object.Function)
			if !ok {
				return newError("'%s' is not a function", method.(*ast.Identifier).Value)
			}
			fn.Env.Define("@", obj)
			ret := applyFunction(fn, args, l)
			fn.Env.Del("@")
			return ret
		}
	case *object.StructInstance:
		methodName := method.(*ast.Identifier).Value
		return callStructMethod(obj, methodName, args, defs, l)
	}
	return newError("Sorry, %s does not have a function '%s()'", obj.Inspect(), method.(*ast.Identifier).Value)
}

// ///////////////////////////////////////////////////////////////
// //////// Some methods here because of loop dependency ////////
// /////////////////////////////////////////////////////////////
func maap(a *object.Array, args []object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.FUNCTION_OBJ {
		return newError("Sorry, the argument is not valid")
	}

	fn, ok := args[0].(*object.Function)
	if !ok {
		return newError("Sorry, the argument is not valid")
	}
	newArr := object.Array{Elements: []object.VintObject{}}
	for _, obj := range a.Elements {
		env := object.NewEnclosedEnvironment(fn.Env)
		env.Define(fn.Parameters[0].Value, obj)
		r := Eval(fn.Body, env)
		if o, ok := r.(*object.ReturnValue); ok {
			r = o.Value
		}
		newArr.Elements = append(newArr.Elements, r)
	}
	return &newArr
}

func filter(a *object.Array, args []object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.FUNCTION_OBJ {
		return newError("Sorry, the argument is not valid")
	}

	fn, ok := args[0].(*object.Function)
	if !ok {
		return newError("Sorry, the argument is not valid")
	}
	newArr := object.Array{Elements: []object.VintObject{}}
	for _, obj := range a.Elements {
		env := object.NewEnclosedEnvironment(fn.Env)
		env.Define(fn.Parameters[0].Value, obj)
		cond := Eval(fn.Body, env)
		if cond.Inspect() == "true" {
			newArr.Elements = append(newArr.Elements, obj)
		}
	}
	return &newArr
}

func sortBy(a *object.Array, args []object.VintObject) object.VintObject {
	if len(args) != 1 || args[0].Type() != object.FUNCTION_OBJ {
		return newError("Sorry, the argument is not valid")
	}

	fn, ok := args[0].(*object.Function)
	if !ok {
		return newError("Sorry, the argument is not valid")
	}

	if len(a.Elements) <= 1 {
		return a
	}

	// Use bubble sort for simplicity with custom comparison
	elements := make([]object.VintObject, len(a.Elements))
	copy(elements, a.Elements)

	for i := 0; i < len(elements)-1; i++ {
		for j := 0; j < len(elements)-i-1; j++ {
			// Call comparison function with two elements
			env := object.NewEnclosedEnvironment(fn.Env)
			env.Define(fn.Parameters[0].Value, elements[j])
			if len(fn.Parameters) > 1 {
				env.Define(fn.Parameters[1].Value, elements[j+1])
			}

			result1 := Eval(fn.Body, env)
			if o, ok := result1.(*object.ReturnValue); ok {
				result1 = o.Value
			}

			env2 := object.NewEnclosedEnvironment(fn.Env)
			env2.Define(fn.Parameters[0].Value, elements[j+1])
			if len(fn.Parameters) > 1 {
				env2.Define(fn.Parameters[1].Value, elements[j])
			}

			result2 := Eval(fn.Body, env2)
			if o, ok := result2.(*object.ReturnValue); ok {
				result2 = o.Value
			}

			// Compare results and swap if needed
			shouldSwap := false
			if intResult1, ok1 := result1.(*object.Integer); ok1 {
				if intResult2, ok2 := result2.(*object.Integer); ok2 {
					shouldSwap = intResult1.Value > intResult2.Value
				}
			} else if floatResult1, ok1 := result1.(*object.Float); ok1 {
				if floatResult2, ok2 := result2.(*object.Float); ok2 {
					shouldSwap = floatResult1.Value > floatResult2.Value
				}
			} else if strResult1, ok1 := result1.(*object.String); ok1 {
				if strResult2, ok2 := result2.(*object.String); ok2 {
					shouldSwap = strResult1.Value > strResult2.Value
				}
			}

			if shouldSwap {
				elements[j], elements[j+1] = elements[j+1], elements[j]
			}
		}
	}

	a.Elements = elements
	return a
}

// evalDictCallback evaluates a dict callback function with the given arguments,
// binding them to the function's parameter names.
func evalDictCallback(fn *object.Function, args []object.VintObject) object.VintObject {
	env := object.NewEnclosedEnvironment(fn.Env)
	for i, param := range fn.Parameters {
		if i < len(args) {
			env.Define(param.Value, args[i])
		}
	}
	result := Eval(fn.Body, env)
	if rv, ok := result.(*object.ReturnValue); ok {
		return rv.Value
	}
	return result
}

func dictFilter(d *object.Dict, args []object.VintObject) object.VintObject {
	if len(args) != 1 {
		return newError("filter() expects 1 argument (function), got %d", len(args))
	}
	fn, ok := args[0].(*object.Function)
	if !ok {
		return newError("filter() argument must be a function")
	}

	newPairs := make(map[object.HashKey]object.DictPair)
	for _, pair := range d.Pairs {
		result := evalDictCallback(fn, []object.VintObject{pair.Key, pair.Value})
		if isError(result) {
			return result
		}
		if isTruthy(result) {
			newPairs[pair.Key.(object.Hashable).HashKey()] = pair
		}
	}
	return &object.Dict{Pairs: newPairs}
}

func dictMap(d *object.Dict, args []object.VintObject) object.VintObject {
	if len(args) != 1 {
		return newError("map() expects 1 argument (function), got %d", len(args))
	}
	fn, ok := args[0].(*object.Function)
	if !ok {
		return newError("map() argument must be a function")
	}

	newPairs := make(map[object.HashKey]object.DictPair)
	for _, pair := range d.Pairs {
		result := evalDictCallback(fn, []object.VintObject{pair.Key, pair.Value})
		if isError(result) {
			return result
		}
		newPair := object.DictPair{Key: pair.Key, Value: result}
		newPairs[pair.Key.(object.Hashable).HashKey()] = newPair
	}
	return &object.Dict{Pairs: newPairs}
}

func dictReduce(d *object.Dict, args []object.VintObject) object.VintObject {
	if len(args) < 1 || len(args) > 2 {
		return newError("reduce() expects 1 or 2 arguments (function, initial), got %d", len(args))
	}
	fn, ok := args[0].(*object.Function)
	if !ok {
		return newError("reduce() first argument must be a function")
	}

	var accumulator object.VintObject
	if len(args) == 2 {
		accumulator = args[1]
	} else {
		if len(d.Pairs) == 0 {
			return newError("Cannot reduce empty dictionary without initial value")
		}
		for _, pair := range d.Pairs {
			accumulator = pair.Value
			break
		}
	}

	for _, pair := range d.Pairs {
		result := evalDictCallback(fn, []object.VintObject{accumulator, pair.Key, pair.Value})
		if isError(result) {
			return result
		}
		accumulator = result
	}
	return accumulator
}

func dictForEach(d *object.Dict, args []object.VintObject) object.VintObject {
	if len(args) != 1 {
		return newError("forEach() expects 1 argument (function), got %d", len(args))
	}
	fn, ok := args[0].(*object.Function)
	if !ok {
		return newError("forEach() argument must be a function")
	}

	for _, pair := range d.Pairs {
		result := evalDictCallback(fn, []object.VintObject{pair.Key, pair.Value})
		if isError(result) {
			return result
		}
	}
	return NULL
}

func dictFind(d *object.Dict, args []object.VintObject) object.VintObject {
	if len(args) != 1 {
		return newError("find() expects 1 argument (function), got %d", len(args))
	}
	fn, ok := args[0].(*object.Function)
	if !ok {
		return newError("find() argument must be a function")
	}

	for _, pair := range d.Pairs {
		result := evalDictCallback(fn, []object.VintObject{pair.Key, pair.Value})
		if isError(result) {
			return result
		}
		if isTruthy(result) {
			return &object.Array{Elements: []object.VintObject{pair.Key, pair.Value}}
		}
	}
	return NULL
}

func dictSome(d *object.Dict, args []object.VintObject) object.VintObject {
	if len(args) != 1 {
		return newError("some() expects 1 argument (function), got %d", len(args))
	}
	fn, ok := args[0].(*object.Function)
	if !ok {
		return newError("some() argument must be a function")
	}

	for _, pair := range d.Pairs {
		result := evalDictCallback(fn, []object.VintObject{pair.Key, pair.Value})
		if isError(result) {
			return result
		}
		if isTruthy(result) {
			return TRUE
		}
	}
	return FALSE
}

func dictEvery(d *object.Dict, args []object.VintObject) object.VintObject {
	if len(args) != 1 {
		return newError("every() expects 1 argument (function), got %d", len(args))
	}
	fn, ok := args[0].(*object.Function)
	if !ok {
		return newError("every() argument must be a function")
	}

	for _, pair := range d.Pairs {
		result := evalDictCallback(fn, []object.VintObject{pair.Key, pair.Value})
		if isError(result) {
			return result
		}
		if !isTruthy(result) {
			return FALSE
		}
	}
	return TRUE
}
