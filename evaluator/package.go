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

	// Automatically runs the init function if it exists
	if initFunc, ok := Package.Scope.Get("init"); ok {
		if fn, ok := initFunc.(*object.Function); ok {
			// Call the init function
			// We can create a new environment for the function call if needed
			// but for an init function, it should probably run in the package's scope
			extendedEnv := object.NewEnclosedEnvironment(Package.Scope)
			Eval(fn.Body, extendedEnv)
		}
	}
	env.Define(node.Name.Value, Package)
	return Package
}
