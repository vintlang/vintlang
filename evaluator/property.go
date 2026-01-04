package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalPropertyExpression(node *ast.PropertyExpression, env *object.Environment) object.VintObject {
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

		// Use GetPublic to ensure only public members are accessible
		if val, ok := obj.GetPublic(prop); ok {
			return val
		}

		// If GetPublic failed, check if it's because the member is private
		if obj.IsPrivate(prop) {
			return newError("cannot access private property '%s' from package", prop)
		}
		// case *object.Module:
		// 	mod := left.(*object.Module)
		// 	prop := node.Property.(*ast.Identifier).Value
		// 	if val, ok := mod.Properties[prop]; ok {
		// 		return val()
		// 	}
	case *object.Enum:
		enum := left.(*object.Enum)
		prop := node.Property.(*ast.Identifier).Value
		member, ok := enum.GetMember(prop)
		if !ok {
			return newError("Enum '%s' has no member '%s'", enum.Name, prop)
		}
		return member
	}
	return newError("Value %s is not valid for %s", node.Property.(*ast.Identifier).Value, left.Inspect())
}

func evalPropertyAssignment(name *ast.PropertyExpression, val object.VintObject, env *object.Environment) object.VintObject {
	left := Eval(name.Object, env)
	if isError(left) {
		return left
	}
	switch left.(type) {
	case *object.Instance:
		obj := left.(*object.Instance)
		prop := name.Property.(*ast.Identifier).Value
		obj.Env.SetScoped(prop, val)
		return NULL
	case *object.Package:
		obj := left.(*object.Package)
		prop := name.Property.(*ast.Identifier).Value

		// Check if trying to assign to a private member from outside
		if obj.IsPrivate(prop) {
			return newError("Cannot assign to private member '%s' from outside package '%s'", prop, obj.Name.Value)
		}

		obj.Scope.SetScoped(prop, val)
		return NULL
	default:
		return newError("Failed to set in package %s", left.Type())
	}
}
