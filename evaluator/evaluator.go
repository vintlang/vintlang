package evaluator

import (
	"fmt"
	"os"

	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
)

var (
	NULL     = &object.Null{}
	TRUE     = &object.Boolean{Value: true}
	FALSE    = &object.Boolean{Value: false}
	BREAK    = &object.Break{}
	CONTINUE = &object.Continue{}

	deferredCalls []*object.DeferredCall
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right, node.Token.Line)

	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) && right != nil {
			return right
		}
		return evalInfixExpression(node.Operator, left, right, node.Token.Line)
	case *ast.PostfixExpression:
		return evalPostfixExpression(env, node.Operator, node)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.IfExpression:
		return evalIfExpression(node, env)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		if f, ok := val.(*object.Function); ok {
			f.Name = node.Name.Value
		}
		return env.Define(node.Name.Value, val)

	case *ast.ConstStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		return env.DefineConst(node.Name.Value, val)

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.FunctionLiteral:
		return evalFunction(node, env)

	case *ast.MethodExpression:
		return evalMethodExpression(node, env)

	case *ast.Import:
		return evalImport(node, env)

	case *ast.CallExpression:
		return evalCall(node, env)

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.At:
		return evalAt(node, env)
	case *ast.ArrayLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}
	case *ast.RangeExpression:
		return evalRangeExpression(node, env)
	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index, node.Token.Line)
	case *ast.SliceExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		var start, end object.Object
		if node.Start != nil {
			start = Eval(node.Start, env)
			if isError(start) {
				return start
			}
		}
		if node.End != nil {
			end = Eval(node.End, env)
			if isError(end) {
				return end
			}
		}
		return evalSliceExpression(left, start, end, node.Token.Line)
	case *ast.DictLiteral:
		return evalDictLiteral(node, env)
	case *ast.WhileExpression:
		return evalWhileExpression(node, env)
	case *ast.Break:
		return evalBreak(node)
	case *ast.Continue:
		return evalContinue(node)
	case *ast.SwitchExpression:
		return evalSwitchStatement(node, env)
	case *ast.Null:
		return NULL
	// case *ast.For:
	// 	return evalForExpression(node, env)
	case *ast.ForIn:
		return evalForInExpression(node, env, node.Token.Line)
	case *ast.Package:
		return evalPackage(node, env)
	case *ast.PropertyExpression:
		return evalPropertyExpression(node, env)
	case *ast.PropertyAssignment:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		return evalPropertyAssignment(node.Name, val, env)
	case *ast.Assign:
		return evalAssign(node, env)
	case *ast.AssignEqual:
		return evalAssignEqual(node, env)

	case *ast.AssignmentExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		value := Eval(node.Value, env)
		if isError(value) {
			return value
		}

		// This is an easy way to assign operators like +=, -= etc
		// for index expressions (arrays and dicts) where applicable
		op := node.Token.Literal
		if len(op) >= 2 {
			op = op[:len(op)-1]
			value = evalInfixExpression(op, left, value, node.Token.Line)
			if isError(value) {
				return value
			}
		}

		if ident, ok := node.Left.(*ast.Identifier); ok {
			newVal, ok := env.Assign(ident.Value, value)
			if !ok {
				return newError("assignment to undeclared variable '%s'", ident.Value)
			}
			return newVal
		} else if ie, ok := node.Left.(*ast.IndexExpression); ok {
			obj := Eval(ie.Left, env)
			if isError(obj) {
				return obj
			}

			if array, ok := obj.(*object.Array); ok {
				index := Eval(ie.Index, env)
				if isError(index) {
					return index
				}
				if idx, ok := index.(*object.Integer); ok {
					if int(idx.Value) >= len(array.Elements) {
						return newError("Index exceeds the number of elements")
					}
					array.Elements[idx.Value] = value
				} else {
					return newError("Cannot perform this operation with %#v", index)
				}
			} else if hash, ok := obj.(*object.Dict); ok {
				key := Eval(ie.Index, env)
				if isError(key) {
					return key
				}
				if hashKey, ok := key.(object.Hashable); ok {
					hashed := hashKey.HashKey()
					hash.Pairs[hashed] = object.DictPair{Key: key, Value: value}
				} else {
					return newError("Cannot perform this operation with %T", key)
				}
			} else {
				return newError("%T does not support this operation", obj)
			}
		} else {
			return newError("Use an identifier instead of %T", node.Left)
		}
	case *ast.IncludeStatement:
		return evalIncludeStatement(node, env)

	case *ast.TodoStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		msg := val.Inspect()
		fmt.Printf("\n\u001b[1;33m[TODO]\u001b[0m: %s\n\n", msg)
		return NULL
	case *ast.WarnStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		fmt.Printf("\n\u001b[1;33m[WARN]\u001b[0m: %s\n\n", val.Inspect())
		return NULL
	case *ast.ErrorStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		return newError(val.Inspect())
	case *ast.DeferStatement:
		call, ok := node.Call.(*ast.CallExpression)
		if !ok {
			return newError("defer statement must be followed by a function call")
		}

		fn := Eval(call.Function, env)
		if isError(fn) {
			return fn
		}

		args := evalExpressions(call.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		deferredCall := &object.DeferredCall{Fn: fn, Args: args}
		deferredCalls = append(deferredCalls, deferredCall)

		return NULL
	case *ast.InfoStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		fmt.Printf("\n\u001b[1;36m[INFO]\u001b[0m: %s\n\n", val.Inspect())
		return NULL
	case *ast.DebugStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		fmt.Printf("\n\u001b[1;35m[DEBUG]\u001b[0m: %s\n\n", val.Inspect())
		return NULL
	case *ast.NoteStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		fmt.Printf("\n\u001b[1;34m[NOTE]\u001b[0m: %s\n\n", val.Inspect())
		return NULL
	case *ast.SuccessStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		fmt.Printf("\n\u001b[1;32m[SUCCESS]\u001b[0m: %s\n\n", val.Inspect())
		return NULL
	case *ast.RepeatStatement:
		countObj := Eval(node.Count, env)
		if isError(countObj) {
			return countObj
		}
		count, ok := countObj.(*object.Integer)
		if !ok {
			return newError("repeat expects an integer count, got %s", countObj.Type())
		}
		var result object.Object = NULL
		varName := node.VarName
		if varName == "" {
			varName = "i"
		}
		for i := int64(0); i < count.Value; i++ {
			loopEnv := object.NewEnclosedEnvironment(env)
			loopEnv.Define(varName, &object.Integer{Value: i})
			res := Eval(node.Block, loopEnv)
			if isError(res) {
				return res
			}
			if res != nil {
				switch res.Type() {
				case object.BREAK_OBJ:
					return NULL
				case object.CONTINUE_OBJ:
					continue
				case object.RETURN_VALUE_OBJ:
					return res
				}
			}
			result = res
		}
		return result
	
	// Async/Concurrency constructs
	case *ast.AsyncFunctionLiteral:
		return &object.AsyncFunction{
			Parameters: node.Parameters,
			Body:       node.Body,
			Env:        env,
		}

	case *ast.AwaitExpression:
		promise := Eval(node.Value, env)
		if isError(promise) {
			return promise
		}
		
		promiseObj, ok := promise.(*object.Promise)
		if !ok {
			return newError("await can only be used with promises, got %T", promise)
		}
		
		// Block until promise resolves using channel-based waiting
		promiseObj.Wait()
		
		if promiseObj.Error != nil {
			return promiseObj.Error
		}
		return promiseObj.Value

	case *ast.GoStatement:
		// Execute the expression concurrently
		go func() {
			Eval(node.Expression, env)
		}()
		return NULL

	case *ast.ChannelExpression:
		if node.Buffer != nil {
			bufferSize := Eval(node.Buffer, env)
			if isError(bufferSize) {
				return bufferSize
			}
			
			size, ok := bufferSize.(*object.Integer)
			if !ok {
				return newError("channel buffer size must be an integer, got %T", bufferSize)
			}
			
			return object.NewBufferedChannel(int(size.Value))
		}
		
		return object.NewChannel()
	}

	return nil
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func isTruthy(obj object.Object) bool {
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

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}

	return false
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}

		result = append(result, evaluated)
	}

	return result
}

func applyFunction(fn object.Object, args []object.Object, line int) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		prevDefersCount := len(deferredCalls)
		defer func() {
			if len(deferredCalls) > prevDefersCount {
				callsToExecute := deferredCalls[prevDefersCount:]
				for i := len(callsToExecute) - 1; i >= 0; i-- {
					dc := callsToExecute[i]
					applyFunction(dc.Fn, dc.Args, 0)
				}
				deferredCalls = deferredCalls[:prevDefersCount]
			}
		}()

		if fn.Name != "" {
			fn.Env.Define(fn.Name, fn)
		}
		extendedEnv := extendedFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.AsyncFunction:
		// Execute async function and return a promise
		return fn.Execute(args, Eval)
	case *object.Builtin:
		if result := fn.Fn(args...); result != nil {
			return result
		}
		return NULL
	case *object.Package:
		obj := &object.Instance{
			Package: fn,
			Env:     object.NewEnclosedEnvironment(fn.Env),
		}
		obj.Env.Define("@", obj)
		node, ok := fn.Scope.Get("init")
		if !ok {
			return newError("Line %d: The package does not have an 'init' function", line)
		}
		node.(*object.Function).Env.Define("@", obj)
		applyFunction(node, args, fn.Name.Token.Line)
		node.(*object.Function).Env.Del("@")
		return obj
	default:
		return newError("not a function: %s", fn.Type())
	}
}

func extendedFunctionEnv(
	fn *object.Function,
	args []object.Object,
) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for i, param := range fn.Parameters {
		if i < len(args) {
			env.Define(param.Value, args[i])
		} else if _, ok := fn.Defaults[param.Value]; ok {
			env.Define(param.Value, Eval(fn.Defaults[param.Value], env))
		}
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}

func evalBreak(node *ast.Break) object.Object {
	return BREAK
}

func evalContinue(node *ast.Continue) object.Object {
	return CONTINUE
}

// func evalForExpression(fe *ast.For, env *object.Environment) object.Object {
// 	obj, ok := env.Get(fe.Identifier)
// 	defer func() { // stay safe and not reassign an existing variable
// 		if ok {
// 			env.Set(fe.Identifier, obj)
// 		}
// 	}()
// 	val := Eval(fe.StarterValue, env)
// 	if isError(val) {
// 		return val
// 	}

// 	env.Set(fe.StarterName.Value, val)

// 	// err := Eval(fe.Starter, env)
// 	// if isError(err) {
// 	// 	return err
// 	// }
// 	for {
// 		evaluated := Eval(fe.Condition, env)
// 		if isError(evaluated) {
// 			return evaluated
// 		}
// 		if !isTruthy(evaluated) {
// 			break
// 		}
// 		res := Eval(fe.Block, env)
// 		if isError(res) {
// 			return res
// 		}
// 		if res.Type() == object.BREAK_OBJ {
// 			break
// 		}
// 		if res.Type() == object.CONTINUE_OBJ {
// 			err := Eval(fe.Closer, env)
// 			if isError(err) {
// 				return err
// 			}
// 			continue
// 		}
// 		if res.Type() == object.RETURN_VALUE_OBJ {
// 			return res
// 		}
// 		err := Eval(fe.Closer, env)
// 		if isError(err) {
// 			return err
// 		}
// 	}
// 	return NULL
// }

func loopIterable(
	next func() (object.Object, object.Object),
	env *object.Environment,
	fi *ast.ForIn,
) object.Object {
	var ret object.Object
	k, v := next()
	for k != nil {
		loopEnv := object.NewEnclosedEnvironment(env)
		loopEnv.Define(fi.Key, k)
		if fi.Value != "" {
			loopEnv.Define(fi.Value, v)
		}
		ret = Eval(fi.Block, loopEnv)
		if ret != nil {
			if ret.Type() == object.BREAK_OBJ {
				return NULL
			}
			if ret.Type() == object.CONTINUE_OBJ {
				k, v = next()
				continue
			}
			if ret.Type() == object.RETURN_VALUE_OBJ {
				return ret
			}
		}
		k, v = next()
	}
	return NULL
}

func evalIncludeStatement(node *ast.IncludeStatement, env *object.Environment) object.Object {
	pathObj := Eval(node.Path, env)
	if isError(pathObj) {
		return pathObj
	}
	path, ok := pathObj.(*object.String)
	if !ok {
		return newError("include path must be a string, got %s", pathObj.Type())
	}

	// Read file content
	content, err := os.ReadFile(path.Value)
	if err != nil {
		return newError("could not include file '%s': %s", path.Value, err)
	}

	// Create a new lexer, parser and evaluate the program
	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return newError("errors parsing included file '%s': %s", path.Value, p.Errors()[0])
	}

	return Eval(program, env)
}

func evalRangeExpression(node *ast.RangeExpression, env *object.Environment) object.Object {
	start := Eval(node.Start, env)
	if isError(start) {
		return start
	}

	end := Eval(node.End, env)
	if isError(end) {
		return end
	}

	startInt, ok := start.(*object.Integer)
	if !ok {
		return newError("range start must be an integer, got %T", start)
	}

	endInt, ok := end.(*object.Integer)
	if !ok {
		return newError("range end must be an integer, got %T", end)
	}

	return &object.Range{
		Start:   startInt.Value,
		End:     endInt.Value,
		Current: startInt.Value,
	}
}
