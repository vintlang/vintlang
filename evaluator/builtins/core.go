package builtins

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/vintlang/vintlang/object"
)

func init() {
	registerCoreBuiltins()
}

func registerCoreBuiltins() {
	RegisterBuiltin("input", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 1 {
				return newError("Function '%s' accepts 0 or 1 argument, got %d", "input", len(args))
			}

			if len(args) > 0 && args[0].Type() != object.STRING_OBJ {
				return newError(fmt.Sprintf(`Please use quotes: "%s"`, args[0].Inspect()))
			}
			if len(args) == 1 {
				prompt := args[0].(*object.String).Value
				fmt.Fprint(os.Stdout, prompt)
			}

			buffer := bufio.NewReader(os.Stdin)

			line, _, err := buffer.ReadLine()
			if err != nil && err != io.EOF {
				return newError("Failed to read input")
			}

			return &object.String{Value: string(line)}
		},
	})

	RegisterBuiltin("print", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return handlePrint(os.Stdout, args, false)
		},
	})

	RegisterBuiltin("println", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return handlePrint(os.Stdout, args, true)
		},
	})

	RegisterBuiltin("printErr", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return handlePrint(os.Stderr, args, false)
		},
	})

	RegisterBuiltin("printlnErr", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return handlePrint(os.Stderr, args, true)
		},
	})

	RegisterBuiltin("type", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Function 'type' requires exactly 1 argument, got %d", len(args))
			}

			return &object.String{Value: string(args[0].Type())}
		},
	})

	RegisterBuiltin("len", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.Dict:
				return &object.Integer{Value: int64(len(arg.Pairs))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	})
}