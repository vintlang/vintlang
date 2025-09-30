package builtins

import "github.com/vintlang/vintlang/object"

func init() {
	registerArrayBuiltins()
}

func registerArrayBuiltins() {
	RegisterBuiltin("range", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 3 {
				return newError("Sorry, range requires 1 to 3 arguments, you provided %d", len(args))
			}

			var start, end, step int64
			var err error

			switch len(args) {
			case 1:
				end, err = getIntValue(args[0])
				if err != nil {
					return newError("Argument must be an integer")
				}
				start, step = 0, 1
			case 2:
				start, err = getIntValue(args[0])
				if err != nil {
					return newError("First argument must be an integer")
				}
				end, err = getIntValue(args[1])
				if err != nil {
					return newError("Second argument must be an integer")
				}
				step = 1
			case 3:
				start, err = getIntValue(args[0])
				if err != nil {
					return newError("First argument must be an integer")
				}
				end, err = getIntValue(args[1])
				if err != nil {
					return newError("Second argument must be an integer")
				}
				step, err = getIntValue(args[2])
				if err != nil {
					return newError("Third argument must be an integer")
				}
				if step == 0 {
					return newError("Step cannot be zero")
				}
			}

			elements := []object.Object{}
			for i := start; (step > 0 && i < end) || (step < 0 && i > end); i += step {
				elements = append(elements, &object.Integer{Value: i})
			}

			return &object.Array{Elements: elements}
		},
	})

	RegisterBuiltin("append", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d, want>=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `append` must be an array, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			newElements := make([]object.Object, len(arr.Elements), len(arr.Elements)+len(args)-1)
			copy(newElements, arr.Elements)
			newElements = append(newElements, args[1:]...)
			return &object.Array{Elements: newElements}
		},
	})

	RegisterBuiltin("pop", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `pop` must be an array, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length == 0 {
				return newError("cannot pop from an empty array")
			}
			popped := arr.Elements[length-1]
			arr.Elements = arr.Elements[:length-1]
			return popped
		},
	})

	RegisterBuiltin("indexOf", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("indexOf() takes exactly 2 arguments, got %d", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to indexOf() must be an array, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			for i, element := range arr.Elements {
				if element.Inspect() == args[1].Inspect() {
					return &object.Integer{Value: int64(i)}
				}
			}

			return &object.Integer{Value: -1} // Not found
		},
	})

	RegisterBuiltin("unique", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("unique() takes exactly 1 argument, got %d", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to unique() must be an array, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) == 0 {
				// Return empty array if input is empty
				return &object.Array{Elements: []object.Object{}}
			}

			seen := make(map[string]bool)
			unique := []object.Object{}

			for _, element := range arr.Elements {
				// Use Inspect() method to get string representation for comparison
				key := element.Inspect()
				if !seen[key] {
					seen[key] = true
					unique = append(unique, element)
				}
			}

			return &object.Array{Elements: unique}
		},
	})
}