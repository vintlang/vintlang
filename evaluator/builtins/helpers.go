package builtins

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/vintlang/vintlang/object"
)

// Helper functions used by builtin functions

func handlePrint(w io.Writer, args []object.Object, addNewline bool) object.Object {
	var arr []string
	for _, arg := range args {
		if arg == nil {
			return newError("Operation cannot be performed on nil")
		}
		arr = append(arr, arg.Inspect())
	}
	str := strings.Join(arr, " ")
	if addNewline {
		fmt.Fprintln(w, str)
	} else {
		fmt.Fprint(w, str)
	}
	return nil
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func getIntValue(obj object.Object) (int64, error) {
	switch obj := obj.(type) {
	case *object.Integer:
		return obj.Value, nil
	default:
		return 0, fmt.Errorf("expected Integer, got %s", obj.Type())
	}
}

func getBooleanValue(obj object.Object) (bool, error) {
	switch obj := obj.(type) {
	case *object.Boolean:
		return obj.Value, nil
	default:
		return false, fmt.Errorf("expected Boolean, got %s", obj.Type())
	}
}

func convertToInteger(value object.Object) object.Object {
	switch value := value.(type) {
	case *object.Integer:
		return value
	case *object.Float:
		return &object.Integer{Value: int64(value.Value)}
	case *object.String:
		val, err := strconv.ParseInt(value.Value, 10, 64)
		if err != nil {
			return newError("Cannot convert '%s' to integer", value.Value)
		}
		return &object.Integer{Value: val}
	case *object.Boolean:
		if value.Value {
			return &object.Integer{Value: 1}
		}
		return &object.Integer{Value: 0}
	default:
		return newError("Cannot convert %s to integer", value.Type())
	}
}

func convertToFloat(value object.Object) object.Object {
	switch value := value.(type) {
	case *object.Float:
		return value
	case *object.Integer:
		return &object.Float{Value: float64(value.Value)}
	case *object.String:
		val, err := strconv.ParseFloat(value.Value, 64)
		if err != nil {
			return newError("Cannot convert '%s' to float", value.Value)
		}
		return &object.Float{Value: val}
	case *object.Boolean:
		if value.Value {
			return &object.Float{Value: 1.0}
		}
		return &object.Float{Value: 0.0}
	default:
		return newError("Cannot convert %s to float", value.Type())
	}
}

func convertToString(value object.Object) object.Object {
	return &object.String{Value: value.Inspect()}
}

func convertToBoolean(value object.Object) object.Object {
	switch value := value.(type) {
	case *object.Boolean:
		return value
	case *object.Integer:
		return &object.Boolean{Value: value.Value != 0}
	case *object.Float:
		return &object.Boolean{Value: value.Value != 0.0}
	case *object.String:
		val, err := strconv.ParseBool(value.Value)
		if err != nil {
			return newError("Cannot convert '%s' to boolean", value.Value)
		}
		return &object.Boolean{Value: val}
	default:
		return newError("Cannot convert %s to boolean", value.Type())
	}
}

// Common constants
var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)
