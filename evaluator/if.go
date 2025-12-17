package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.VintObject {
	condition := Eval(ie.Condition, env)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		result := Eval(ie.Consequence, env)
		if isError(result) {
			return result
		}
		return result
	} else if ie.Alternative != nil {
		result := Eval(ie.Alternative, env)
		if isError(result) {
			return result
		}
		return result
	} else {
		return NULL
	}
}
