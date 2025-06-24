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

	newVal, ok := env.Assign(node.Name.Value, val)
	if !ok {
		return newError("assignment to undeclared variable '%s'", node.Name.Value)
	}
	return newVal
}
