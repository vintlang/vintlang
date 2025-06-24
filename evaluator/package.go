package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalPackage(node *ast.Package, env *object.Environment) object.Object {
	Package := &object.Package{
		Name:  node.Name,
		Env:   env,
		Scope: object.NewEnclosedEnvironment(env),
	}

	Eval(node.Block, Package.Scope)
	env.Define(node.Name.Value, Package)
	return Package
}
