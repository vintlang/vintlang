package evaluator

import (
	"strconv"

	"github.com/ekilie/vint-lang/object"
)

// Converts an object to an integer, if possible
func convertToInteger(obj object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.Integer:
		return obj
	case *object.Float:
		return &object.Integer{Value: int64(obj.Value)}
	case *object.String:
		i, err := strconv.ParseInt(obj.Value, 10, 64)
		if err != nil {
			return newError("Cannot convert '%s' to INTEGER", obj.Value)
		}
		return &object.Integer{Value: i}
	case *object.Boolean:
		if obj.Value {
			return &object.Integer{Value: 1}
		}
		return &object.Integer{Value: 0}
	default:
		return newError("Cannot convert %s to INTEGER", obj.Type())
	}
}

// Converts an object to a float, if possible
func convertToFloat(obj object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.Float:
		return obj
	case *object.Integer:
		return &object.Float{Value: float64(obj.Value)}
	case *object.String:
		f, err := strconv.ParseFloat(obj.Value, 64)
		if err != nil {
			return newError("Cannot convert '%s' to FLOAT", obj.Value)
		}
		return &object.Float{Value: f}
	case *object.Boolean:
		if obj.Value {
			return &object.Float{Value: 1.0}
		}
		return &object.Float{Value: 0.0}
	default:
		return newError("Cannot convert %s to FLOAT", obj.Type())
	}
}

// Converts an object to a string
func convertToString(obj object.Object) object.Object {
	return &object.String{Value: obj.Inspect()}
}

// Converts an object to a boolean
func convertToBoolean(obj object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.Boolean:
		return obj
	case *object.Integer:
		return &object.Boolean{Value: obj.Value != 0}
	case *object.Float:
		return &object.Boolean{Value: obj.Value != 0}
	case *object.String:
		return &object.Boolean{Value: len(obj.Value) > 0}
	case *object.Null:
		return &object.Boolean{Value: false}
	default:
		return &object.Boolean{Value: true}
	}
}
