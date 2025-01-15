package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalForInExpression(fie *ast.ForIn, env *object.Environment, line int) object.Object {
	// Evaluates the iterable expression and checks for existing key and value identifiers in the environment
	iterable := Eval(fie.Iterable, env)
	existingKeyIdentifier, okk := env.Get(fie.Key) // Check if key identifier exists
	existingValueIdentifier, okv := env.Get(fie.Value) // Check if value identifier exists
	
	// Ensure the original key and value identifiers are restored after execution
	defer func() { 
		if okk {
			env.Set(fie.Key, existingKeyIdentifier) // Restore key identifier
		}
		if okv {
			env.Set(fie.Value, existingValueIdentifier) // Restore value identifier
		}
	}()
	
	// Check if the iterable object supports iteration
	switch i := iterable.(type) {
	case object.Iterable:
		defer func() {
			i.Reset() // Reset iterable after iteration
		}()
		return loopIterable(i.Next, env, fie) // Start looping through the iterable
	default:
		// Returns an error if the iterable object does not support iteration
		return newError("Line %d: Operation not supported on %s", line, i.Type())
	}
}
