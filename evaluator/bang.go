package evaluator

import "github.com/vintlang/vintlang/object"

func evalBangOperatorExpression(right object.VintObject) object.VintObject {
	switch obj := right.(type) {
	case *object.Boolean:
		if obj.Value {
			return FALSE
		}
		return TRUE
	case *object.Null:
		return TRUE
	default:
		return FALSE
	}
}
