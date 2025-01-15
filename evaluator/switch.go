package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalSwitchStatement(se *ast.SwitchExpression, env *object.Environment) object.Object {
	obj := Eval(se.Value, env)
	for _, opt := range se.Choices {

		if opt.Default {
			continue
		}
		for _, val := range opt.Expr {
			out := Eval(val, env)
			if obj.Type() == out.Type() && obj.Inspect() == out.Inspect() {
				blockOut := evalBlockStatement(opt.Block, env)
				return blockOut
			}
		}
	}
	for _, opt := range se.Choices {
		if opt.Default {
			out := evalBlockStatement(opt.Block, env)
			return out
		}
	}
	return nil
}
