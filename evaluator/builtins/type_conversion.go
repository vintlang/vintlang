package builtins

import (
	"strconv"

	"github.com/vintlang/vintlang/object"
)

func init() {
	registerTypeConversionBuiltins()
}

func registerTypeConversionBuiltins() {
	RegisterBuiltin("convert", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("Sorry, convert requires 2 arguments, you provided %d", len(args))
			}

			value := args[0]
			targetType := args[1]

			if targetType.Type() != object.STRING_OBJ {
				return newError("Target type must be a string")
			}

			targetTypeStr := targetType.(*object.String).Value

			switch targetTypeStr {
			case "INTEGER":
				return convertToInteger(value)
			case "FLOAT":
				return convertToFloat(value)
			case "STRING":
				return convertToString(value)
			case "BOOLEAN":
				return convertToBoolean(value)
			default:
				return newError("Unknown type: %s", targetTypeStr)
			}
		},
	})

	RegisterBuiltin("string", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("string() requires exactly 1 argument, you provided %d", len(args))
			}

			value := args[0]
			return convertToString(value)
		},
	})

	RegisterBuiltin("int", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("int() requires exactly 1 argument, you provided %d", len(args))
			}

			value := args[0]
			return convertToInteger(value)
		},
	})

	RegisterBuiltin("parseInt", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("parseInt() takes exactly 1 argument, got %d", len(args))
			}

			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to parseInt() must be a string, got %s", args[0].Type())
			}

			str := args[0].(*object.String).Value
			val, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return newError("cannot parse '%s' as integer: %s", str, err.Error())
			}

			return &object.Integer{Value: val}
		},
	})

	RegisterBuiltin("parseFloat", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("parseFloat() takes exactly 1 argument, got %d", len(args))
			}

			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to parseFloat() must be a string, got %s", args[0].Type())
			}

			str := args[0].(*object.String).Value
			val, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return newError("cannot parse '%s' as float: %s", str, err.Error())
			}

			return &object.Float{Value: val}
		},
	})
}