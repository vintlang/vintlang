package evaluator

import (
	"github.com/ekilie/vint-lang/ast"
	"github.com/ekilie/vint-lang/object"
)

func evalFunction(node *ast.FunctionLiteral, env *object.Environment) object.Object {
	function := &object.Function{
		Name:       node.Name,
		Parameters: node.Parameters,
		Defaults:   node.Defaults,
		Body:       node.Body,
		Env:        env,
	}

	return function
}
