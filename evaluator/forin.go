package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalForInExpression(fie *ast.ForIn, env *object.Environment, line int) object.Object {
	// Evaluates the iterable expression
	iterable := Eval(fie.Iterable, env)

	// Check if the iterable object supports iteration
	switch i := iterable.(type) {
	case object.Iterable:
		defer func() {
			i.Reset() // Reset iterable after iteration
		}()
		return loopIterable(i.Next, env, fie, line) // Start looping through the iterable
	default:
		// Returns an error if the iterable object does not support iteration
		return newError("Line %d: for..in loop requires an iterable object, but got %s", line, i.Type())
	}
}

func loopIterable(
	next func() (object.Object, object.Object),
	env *object.Environment,
	fi *ast.ForIn,
	line int,
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
		if isError(ret) {
			return ret
		}
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