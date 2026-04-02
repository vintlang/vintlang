package builtins

import (
	"math"

	"github.com/vintlang/vintlang/object"
)

func init() {
	registerUtilityBuiltins()
}

func registerUtilityBuiltins() {
	// copy() - shallow copy of arrays and dicts
	RegisterBuiltin("copy", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("copy() requires exactly 1 argument, got %d", len(args))
			}

			switch obj := args[0].(type) {
			case *object.Array:
				newElements := make([]object.VintObject, len(obj.Elements))
				copy(newElements, obj.Elements)
				return &object.Array{Elements: newElements}
			case *object.Dict:
				newPairs := make(map[object.HashKey]object.DictPair, len(obj.Pairs))
				for k, v := range obj.Pairs {
					newPairs[k] = object.DictPair{Key: v.Key, Value: v.Value}
				}
				return &object.Dict{Pairs: newPairs}
			case *object.String:
				return &object.String{Value: obj.Value}
			default:
				return newError("copy() does not support type %s", args[0].Type())
			}
		},
	})

	// clone() - deep copy of arrays and dicts
	RegisterBuiltin("clone", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("clone() requires exactly 1 argument, got %d", len(args))
			}
			return deepCopy(args[0])
		},
	})

	// pow(base, exp) - exponentiation
	RegisterBuiltin("pow", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 2 {
				return newError("pow() requires exactly 2 arguments, got %d", len(args))
			}

			var base, exp float64

			switch b := args[0].(type) {
			case *object.Integer:
				base = float64(b.Value)
			case *object.Float:
				base = b.Value
			default:
				return newError("first argument to pow() must be a number, got %s", args[0].Type())
			}

			switch e := args[1].(type) {
			case *object.Integer:
				exp = float64(e.Value)
			case *object.Float:
				exp = e.Value
			default:
				return newError("second argument to pow() must be a number, got %s", args[1].Type())
			}

			result := math.Pow(base, exp)

			// Return integer if both args were integers and result is whole
			_, baseIsInt := args[0].(*object.Integer)
			_, expIsInt := args[1].(*object.Integer)
			if baseIsInt && expIsInt && exp >= 0 && result == math.Trunc(result) && !math.IsInf(result, 0) {
				return &object.Integer{Value: int64(result)}
			}

			return &object.Float{Value: result}
		},
	})

	// Type predicate builtins
	RegisterBuiltin("is_null", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("is_null() requires exactly 1 argument, got %d", len(args))
			}
			if args[0].Type() == object.NULL_OBJ {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("is_int", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("is_int() requires exactly 1 argument, got %d", len(args))
			}
			if args[0].Type() == object.INTEGER_OBJ {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("is_float", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("is_float() requires exactly 1 argument, got %d", len(args))
			}
			if args[0].Type() == object.FLOAT_OBJ {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("is_string", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("is_string() requires exactly 1 argument, got %d", len(args))
			}
			if args[0].Type() == object.STRING_OBJ {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("is_bool", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("is_bool() requires exactly 1 argument, got %d", len(args))
			}
			if args[0].Type() == object.BOOLEAN_OBJ {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("is_array", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("is_array() requires exactly 1 argument, got %d", len(args))
			}
			if args[0].Type() == object.ARRAY_OBJ {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("is_dict", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("is_dict() requires exactly 1 argument, got %d", len(args))
			}
			if args[0].Type() == object.DICT_OBJ {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("is_function", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("is_function() requires exactly 1 argument, got %d", len(args))
			}
			t := args[0].Type()
			if t == object.FUNCTION_OBJ || t == object.BUILTIN_OBJ || t == object.ASYNC_FUNC_OBJ {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("is_error", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("is_error() requires exactly 1 argument, got %d", len(args))
			}
			if args[0].Type() == object.ERROR_OBJ || args[0].Type() == object.CUSTOM_ERROR_OBJ {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("is_number", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("is_number() requires exactly 1 argument, got %d", len(args))
			}
			t := args[0].Type()
			if t == object.INTEGER_OBJ || t == object.FLOAT_OBJ {
				return TRUE
			}
			return FALSE
		},
	})
}

func deepCopy(obj object.VintObject) object.VintObject {
	switch o := obj.(type) {
	case *object.Array:
		newElements := make([]object.VintObject, len(o.Elements))
		for i, el := range o.Elements {
			newElements[i] = deepCopy(el)
		}
		return &object.Array{Elements: newElements}
	case *object.Dict:
		newPairs := make(map[object.HashKey]object.DictPair, len(o.Pairs))
		for k, v := range o.Pairs {
			newPairs[k] = object.DictPair{Key: deepCopy(v.Key), Value: deepCopy(v.Value)}
		}
		return &object.Dict{Pairs: newPairs}
	case *object.String:
		return &object.String{Value: o.Value}
	case *object.Integer:
		return &object.Integer{Value: o.Value}
	case *object.Float:
		return &object.Float{Value: o.Value}
	case *object.Boolean:
		return &object.Boolean{Value: o.Value}
	case *object.Null:
		return NULL
	default:
		return obj
	}
}
