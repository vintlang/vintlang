package evaluator

import (
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

func evalSwitchStatement(se *ast.SwitchExpression, env *object.Environment) object.VintObject {
	obj := Eval(se.Value, env)
	if isError(obj) {
		return obj
	}

	for _, opt := range se.Choices {
		if opt.Default {
			continue
		}

		// Handle variable binding cases (case x if condition)
		if opt.Variable != nil {
			// Create new scope for the variable
			caseEnv := object.NewEnclosedEnvironment(env)
			caseEnv.Define(opt.Variable.Value, obj)

			// Evaluate guard condition in the new scope
			if opt.Guard != nil {
				guardResult := Eval(opt.Guard, caseEnv)
				if isError(guardResult) {
					return guardResult
				}
				if isTruthy(guardResult) {
					result := evalBlockStatement(opt.Block, caseEnv)
					if isError(result) {
						return result
					}
					return result
				}
			} else {
				// No guard, always match
				result := evalBlockStatement(opt.Block, caseEnv)
				if isError(result) {
					return result
				}
				return result
			}
		} else {
			// Handle regular value-based cases
			for _, val := range opt.Expr {
				out := Eval(val, env)
				if isError(out) {
					return out
				}
				if obj.Type() == out.Type() && obj.Inspect() == out.Inspect() {
					// Check guard condition if present
					if opt.Guard != nil {
						guardResult := Eval(opt.Guard, env)
						if isError(guardResult) {
							return guardResult
						}
						if !isTruthy(guardResult) {
							continue // Guard failed, try next case
						}
					}
					result := evalBlockStatement(opt.Block, env)
					if isError(result) {
						return result
					}
					return result
				}
			}
		}
	}

	// Handle default cases
	for _, opt := range se.Choices {
		if opt.Default {
			result := evalBlockStatement(opt.Block, env)
			if isError(result) {
				return result
			}
			return result
		}
	}

	return NULL
}
