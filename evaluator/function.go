package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalFunction(node *ast.FunctionLiteral, env *object.Environment) object.VintObject {
	function := &object.Function{
		Name:       node.Name,
		Parameters: node.Parameters,
		Defaults:   node.Defaults,
		Body:       node.Body,
		Env:        env,
	}

	return function
}

func evalTypedFunction(node *ast.TypedFunctionLiteral, env *object.Environment) object.VintObject {
	params := make([]*ast.Identifier, len(node.Parameters))
	paramTypes := make([]ast.Type, len(node.Parameters))
	for i, tp := range node.Parameters {
		params[i] = tp.Identifier
		paramTypes[i] = tp.Type
	}

	defaults := make(map[string]ast.Expression)
	for _, tp := range node.Parameters {
		if tp.Default != nil {
			defaults[tp.Identifier.Value] = tp.Default
		}
	}

	function := &object.Function{
		Name:       "",
		Parameters: params,
		ParamTypes: paramTypes,
		ReturnType: node.ReturnType,
		Defaults:   defaults,
		Body:       node.Body,
		Env:        env,
	}

	return function
}
