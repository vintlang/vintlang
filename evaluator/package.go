package evaluator

import (
	"github.com/ekilie/vint-lang/ast"
	"github.com/ekilie/vint-lang/object"
)

func evalPackage(node *ast.Package, env *object.Environment) object.Object {
	Package := &object.Package{
		Name:  node.Name,
		Env:   env,
		Scope: object.NewEnclosedEnvironment(env),
	}

	Eval(node.Block, Package.Scope)
	env.Set(node.Name.Value, Package)
	return Package
}
