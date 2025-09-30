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

			bool2, err := getBooleanValue(args[1])
			if err != nil {
				return newError("Second argument must be a boolean")
			}

			return &object.Boolean{Value: bool1 && bool2}
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

			bool2, err := getBooleanValue(args[1])
			if err != nil {
				return newError("Second argument must be a boolean")
			}

			return &object.Boolean{Value: bool1 || bool2}
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

			if a == b {
				return &object.Boolean{Value: true}
			}
			return &object.Boolean{Value: false}
		},
	})
}
