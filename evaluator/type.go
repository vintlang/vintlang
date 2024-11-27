package evaluator

import (
	"fmt"
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
		// Parse the string to an integer
		i, err := strconv.ParseInt(obj.Value, 10, 64)
		if err != nil {
			return newError("Cannot convert '%s' to INTEGER", obj.Value)
		}
		return &object.Integer{Value: i}
	case *object.Boolean:
		// Convert boolean to integer: true -> 1, false -> 0
		if obj.Value {
			return &object.Integer{Value: 1}
		}
		return &object.Integer{Value: 0}
	default:
		// Return error if type is not convertible to integer
		return newError("Cannot convert %s to INTEGER", obj.Type())
	}
}

// Converts an object to a float, if possible
func convertToFloat(obj object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.Float:
		return obj
	case *object.Integer:
		// Convert integer to float
		return &object.Float{Value: float64(obj.Value)}
	case *object.String:
		// Parse the string to a float
		f, err := strconv.ParseFloat(obj.Value, 64)
		if err != nil {
			return newError("Cannot convert '%s' to FLOAT", obj.Value)
		}
		return &object.Float{Value: f}
	case *object.Boolean:
		// Convert boolean to float: true -> 1.0, false -> 0.0
		if obj.Value {
			return &object.Float{Value: 1.0}
		}
		return &object.Float{Value: 0.0}
	default:
		// Return error if type is not convertible to float
		return newError("Cannot convert %s to FLOAT", obj.Type())
	}
}

// Converts an object to a string
func convertToString(obj object.Object) object.Object {
	// Simply return the string representation of the object
	return &object.String{Value: obj.Inspect()}
}

// Converts an object to a boolean
func convertToBoolean(obj object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.Boolean:
		// Return the boolean object as is
		return obj
	case *object.Integer:
		// Convert integer to boolean: non-zero -> true, zero -> false
		return &object.Boolean{Value: obj.Value != 0}
	case *object.Float:
		// Convert float to boolean: non-zero -> true, zero -> false
		return &object.Boolean{Value: obj.Value != 0}
	case *object.String:
		// Convert string to boolean: empty string -> false, non-empty -> true
		return &object.Boolean{Value: len(obj.Value) > 0}
	case *object.Null:
		// Null is considered as false
		return &object.Boolean{Value: false}
	default:
		// Default to true for any other type
		return &object.Boolean{Value: true}
	}
}

// Helper function to extract the boolean value from an object
// Returns an error if the object is not a boolean
func getBooleanValue(obj object.Object) (bool, error) {
	switch obj := obj.(type) {
	case *object.Boolean:
		return obj.Value, nil
	case *object.Integer:
		return obj.Value != 0, nil
	case *object.Float:
		return obj.Value != 0, nil
	case *object.String:
		return len(obj.Value) > 0, nil
	default:
		return false, fmt.Errorf("expected Boolean, Integer, Float, or String, got %s", obj.Type())
	}
}
