package evaluator

import (
	"github.com/ekilie/vint-lang/ast"
	"github.com/ekilie/vint-lang/object"
)

func evalAt(node *ast.At, env *object.Environment) object.Object {
	if at, ok := env.Get("@"); ok {
		return at
	}
	return newError("Iko nje ya scope")
}
