package evaluator

import (
	"github.com/ekilie/vint-lang/ast"
	"github.com/ekilie/vint-lang/object"
)

func evalPackage(node *ast.Package, env *object.Environment) object.Object {
	pakeji := &object.Package{
		Name:  node.Name,
		Env:   env,
		Scope: object.NewEnclosedEnvironment(env),
	}

	Eval(node.Block, pakeji.Scope)
	env.Set(node.Name.Value, pakeji)
	return pakeji
}
