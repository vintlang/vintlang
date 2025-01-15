package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalAssign(node *ast.Assign, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}

	obj := env.Set(node.Name.Value, val)
	return obj
}
