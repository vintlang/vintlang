package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalPackage(node *ast.Package, env *object.Environment) object.VintObject {
	Package := &object.Package{
		Name:  node.Name,
		Env:   env,
		Scope: object.NewEnclosedEnvironment(env),
	}

	Package.Scope.Define("@", Package)

	Eval(node.Block, Package.Scope)

	// Automatically run the init function if it exists
	if initFunc, ok := Package.Scope.Get("init"); ok {
		if fn, ok := initFunc.(*object.Function); ok {
			applyFunction(fn, []object.VintObject{}, 0)
		}
	}
	env.Define(node.Name.Value, Package)
	return Package
}
