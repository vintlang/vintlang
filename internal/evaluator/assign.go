package evaluator

import (
	"github.com/vintlang/vintlang/internal/ast"
	"github.com/vintlang/vintlang/internal/object"
)

func evalAssign(node *ast.Assign, env *object.Environment) object.VintObject {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}

	// Check declared type
	if declaredType, ok := env.GetDeclaredType(node.Name.Value); ok {
		if !compatible(declaredType, val) {
			return newTypeError(node.Token.Line,
				"cannot assign %s to variable '%s' of type %s",
				val.Type(), node.Name.Value, declaredType.String())
		}
	}

	newVal, ok := env.Assign(node.Name.Value, val)
	if !ok {
		return newError("Line %d: Assignment to undeclared variable '%s'. Use 'let' to declare the variable first", node.Token.Line, node.Name.Value)
	}
	return newVal
}
