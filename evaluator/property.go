package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalPropertyExpression(node *ast.PropertyExpression, env *object.Environment) object.Object {
	left := Eval(node.Object, env)
	if isError(left) {
		return left
	}
	switch left.(type) {
	case *object.Instance:
		obj := left.(*object.Instance)
		prop := node.Property.(*ast.Identifier).Value
		if val, ok := obj.Env.Get(prop); ok {
			return val
		}
	case *object.Package:
		obj := left.(*object.Package)
		prop := node.Property.(*ast.Identifier).Value
		if val, ok := obj.Env.Get(prop); ok {
			return val
		}
		// case *object.Module:
		// 	mod := left.(*object.Module)
		// 	prop := node.Property.(*ast.Identifier).Value
		// 	if val, ok := mod.Properties[prop]; ok {
		// 		return val()
		// 	}
	}
	return newError("Value %s is not valid for %s", node.Property.(*ast.Identifier).Value, left.Inspect())
}

func evalPropertyAssignment(name *ast.PropertyExpression, val object.Object, env *object.Environment) object.Object {
	left := Eval(name.Object, env)
	if isError(left) {
		return left
	}
	switch left.(type) {
	case *object.Instance:
		obj := left.(*object.Instance)
		prop := name.Property.(*ast.Identifier).Value
		if _, ok := obj.Env.Get(prop); ok {
			obj.Env.Set(prop, val)
			return NULL
		}
		obj.Env.Set(prop, val)
		return NULL
	case *object.Package:
		obj := left.(*object.Package)
		prop := name.Property.(*ast.Identifier).Value
		if _, ok := obj.Env.Get(prop); ok {
			obj.Env.Set(prop, val)
			return NULL
		}
		obj.Env.Set(prop, val)
		return NULL
	default:
		return newError("Failed to set in package %s", left.Type())
	}
}
