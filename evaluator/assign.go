package evaluator

import (
	"github.com/ekilie/vint-lang/ast"
	"github.com/ekilie/vint-lang/object"
)

func evalAssign(node *ast.Assign, env *object.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}

	obj := env.Set(node.Name.Value, val)
	return obj
}
