package builtins

import "github.com/vintlang/vintlang/object"

func init() {
	registerDictBuiltins()
}

func registerDictBuiltins() {
	RegisterBuiltin("keys", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.DICT_OBJ {
				return newError("argument to `keys` must be a dictionary, got %s", args[0].Type())
			}
			dict := args[0].(*object.Dict)
			keys := make([]object.Object, 0, len(dict.Pairs))
			for _, pair := range dict.Pairs {
				keys = append(keys, pair.Key)
			}
			return &object.Array{Elements: keys}
		},
	})

	RegisterBuiltin("values", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.DICT_OBJ {
				return newError("argument to `values` must be a dictionary, got %s", args[0].Type())
			}
			dict := args[0].(*object.Dict)
			values := make([]object.Object, 0, len(dict.Pairs))
			for _, pair := range dict.Pairs {
				values = append(values, pair.Value)
			}
			return &object.Array{Elements: values}
		},
	})

	RegisterBuiltin("has_key", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.DICT_OBJ {
				return newError("first argument to `has_key` must be a dictionary, got %s", args[0].Type())
			}
			dict := args[0].(*object.Dict)
			key, ok := args[1].(object.Hashable)
			if !ok {
				return newError("second argument to `has_key` must be hashable, got %s", args[1].Type())
			}
			if _, ok := dict.Pairs[key.HashKey()]; ok {
				return TRUE
			}
			return FALSE
		},
	})
}