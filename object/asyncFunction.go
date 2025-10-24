package object

import (
	"fmt"
	"strings"

	"github.com/vintlang/vintlang/ast"
)

// AsyncFunction represents an async function object
type AsyncFunction struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (af *AsyncFunction) Type() VintObjectType {
	return ASYNC_FUNC_OBJ
}

func (af *AsyncFunction) Inspect() string {
	var out strings.Builder

	params := []string{}
	for _, p := range af.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("async func")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(af.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// Execute executes the async function and returns a promise
func (af *AsyncFunction) Execute(args []VintObject, eval func(ast.Node, *Environment) VintObject) *Promise {
	promise := NewPromise()

	// Execute the function asynchronously
	go func() {
		defer func() {
			if r := recover(); r != nil {
				promise.Reject(&Error{Message: fmt.Sprintf("async function panic: %v", r)})
			}
		}()

		// Create new environment for the function
		extEnv := NewEnclosedEnvironment(af.Env)

		// Bind parameters
		for paramIdx, param := range af.Parameters {
			if paramIdx < len(args) {
				extEnv.Define(param.Value, args[paramIdx])
			}
		}

		// Execute the function body
		result := eval(af.Body, extEnv)

		// Handle return values and errors
		if returnValue, ok := result.(*ReturnValue); ok {
			promise.Resolve(returnValue.Value)
		} else if errorObj, ok := result.(*Error); ok {
			promise.Reject(errorObj)
		} else {
			promise.Resolve(result)
		}
	}()

	return promise
}
