package evaluator

import "github.com/vintlang/vintlang/object"

func evalBangOperatorExpression(right object.VintObject) object.VintObject {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}
