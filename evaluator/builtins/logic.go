package builtins

import "github.com/vintlang/vintlang/object"

func init() {
	registerLogicBuiltins()
}

func registerLogicBuiltins() {
	RegisterBuiltin("and", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 2 {
				return newError("and requires 2 arguments, you provided %d", len(args))
			}

			bool1, err := getBooleanValue(args[0])
			if err != nil {
				return newError("First argument must be a boolean")
			}

			// Short-circuit: if first is false, return false immediately
			if !bool1 {
				return FALSE
			}

			bool2, err := getBooleanValue(args[1])
			if err != nil {
				return newError("Second argument must be a boolean")
			}

			if bool2 {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("or", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 2 {
				return newError("or requires 2 arguments, you provided %d", len(args))
			}

			bool1, err := getBooleanValue(args[0])
			if err != nil {
				return newError("First argument must be a boolean")
			}

			// Short-circuit: if first is true, return true immediately
			if bool1 {
				return TRUE
			}

			bool2, err := getBooleanValue(args[1])
			if err != nil {
				return newError("Second argument must be a boolean")
			}

			if bool2 {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("not", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("not requires 1 argument, you provided %d", len(args))
			}

			boolVal, err := getBooleanValue(args[0])
			if err != nil {
				return newError("Argument must be a boolean")
			}

			return &object.Boolean{Value: !boolVal}
		},
	})

	RegisterBuiltin("xor", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 2 {
				return newError("xor requires 2 arguments, you provided %d", len(args))
			}

			bool1, err := getBooleanValue(args[0])
			if err != nil {
				return newError("First argument must be a boolean")
			}

			bool2, err := getBooleanValue(args[1])
			if err != nil {
				return newError("Second argument must be a boolean")
			}

			return &object.Boolean{Value: (bool1 && !bool2) || (!bool1 && bool2)}
		},
	})

	RegisterBuiltin("nand", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 2 {
				return newError("nand requires 2 arguments, you provided %d", len(args))
			}

			bool1, err := getBooleanValue(args[0])
			if err != nil {
				return newError("First argument must be a boolean")
			}

			bool2, err := getBooleanValue(args[1])
			if err != nil {
				return newError("Second argument must be a boolean")
			}

			return &object.Boolean{Value: !(bool1 && bool2)}
		},
	})

	RegisterBuiltin("nor", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 2 {
				return newError("nor requires 2 arguments, you provided %d", len(args))
			}

			bool1, err := getBooleanValue(args[0])
			if err != nil {
				return newError("First argument must be a boolean")
			}

			bool2, err := getBooleanValue(args[1])
			if err != nil {
				return newError("Second argument must be a boolean")
			}

			return &object.Boolean{Value: !(bool1 || bool2)}
		},
	})

	RegisterBuiltin("eq", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 2 {
				return newError("Sorry, eq requires 2 arguments, you provided %d", len(args))
			}

			a := args[0]
			b := args[1]

			// Compare by type first, then by value
			if a.Type() != b.Type() {
				return FALSE
			}

			switch aVal := a.(type) {
			case *object.Integer:
				bVal := b.(*object.Integer)
				if aVal.Value == bVal.Value {
					return TRUE
				}
				return FALSE
			case *object.Float:
				bVal := b.(*object.Float)
				if aVal.Value == bVal.Value {
					return TRUE
				}
				return FALSE
			case *object.String:
				bVal := b.(*object.String)
				if aVal.Value == bVal.Value {
					return TRUE
				}
				return FALSE
			case *object.Boolean:
				bVal := b.(*object.Boolean)
				if aVal.Value == bVal.Value {
					return TRUE
				}
				return FALSE
			case *object.Null:
				return TRUE
			default:
				// Fallback: compare string representations
				if a.Inspect() == b.Inspect() {
					return TRUE
				}
				return FALSE
			}
		},
	})
}
