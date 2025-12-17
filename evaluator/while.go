package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalWhileExpression(we *ast.WhileExpression, env *object.Environment) object.VintObject {
	var result object.VintObject
	for {
		condition := Eval(we.Condition, env)
		if isError(condition) {
			return condition
		}
		if !isTruthy(condition) {
			break
		}
		result = Eval(we.Consequence, env)
		if isError(result) {
			return result
		}
		if result != nil {
			switch result.Type() {
			case object.BREAK_OBJ:
				return NULL
			case object.CONTINUE_OBJ:
				continue // Continue to next iteration (re-evaluate condition)
			case object.RETURN_VALUE_OBJ:
				return result
			}
		}
	}
	return result
}
