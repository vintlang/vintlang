package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalWhileExpression(we *ast.WhileExpression, env *object.Environment) object.VintObject {
	condition := Eval(we.Condition, env)
	var evaluated object.VintObject
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		evaluated = Eval(we.Consequence, env)
		if isError(evaluated) {
			return evaluated
		}
		if evaluated != nil && evaluated.Type() == object.BREAK_OBJ {
			return evaluated
		}
		evaluated = evalWhileExpression(we, env)
	}
	return evaluated
}
