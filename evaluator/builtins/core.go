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

	// String utility functions
	RegisterBuiltin("format", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return newError("wrong number of arguments. got=%d, want=1+", len(args))
			}

			formatStr, ok := args[0].(*object.String)
			if !ok {
				return newError("first argument to `format` must be STRING, got %T", args[0])
			}

			formatArgs := make([]interface{}, len(args)-1)
			for i, arg := range args[1:] {
				formatArgs[i] = arg.Inspect()
			}

			result := fmt.Sprintf(formatStr.Value, formatArgs...)
			return &object.String{Value: result}
		},
	})

	RegisterBuiltin("startsWith", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			str, ok := args[0].(*object.String)
			if !ok {
				return newError("first argument to `startsWith` must be STRING, got %T", args[0])
			}

			prefix, ok := args[1].(*object.String)
			if !ok {
				return newError("second argument to `startsWith` must be STRING, got %T", args[1])
			}

			result := len(str.Value) >= len(prefix.Value) && str.Value[:len(prefix.Value)] == prefix.Value
			if result {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("endsWith", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			str, ok := args[0].(*object.String)
			if !ok {
				return newError("first argument to `endsWith` must be STRING, got %T", args[0])
			}

			suffix, ok := args[1].(*object.String)
			if !ok {
				return newError("second argument to `endsWith` must be STRING, got %T", args[1])
			}

			result := len(str.Value) >= len(suffix.Value) && str.Value[len(str.Value)-len(suffix.Value):] == suffix.Value
			if result {
				return TRUE
			}
			return FALSE
		},
	})

	RegisterBuiltin("chr", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			num, ok := args[0].(*object.Integer)
			if !ok {
				return newError("argument to `chr` must be INTEGER, got %T", args[0])
			}

			if num.Value < 0 || num.Value > 127 {
				return newError("chr() arg not in range(128)")
			}

			return &object.String{Value: string(rune(num.Value))}
		},
	})

	RegisterBuiltin("ord", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			str, ok := args[0].(*object.String)
			if !ok {
				return newError("argument to `ord` must be STRING, got %T", args[0])
			}

			if len(str.Value) != 1 {
				return newError("ord() expected a character, but string of length %d found", len(str.Value))
			}

			return &object.Integer{Value: int64(str.Value[0])}
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
