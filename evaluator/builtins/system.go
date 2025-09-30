package builtins

import (
	"os"
	"time"

	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/toolkit"
)

func init() {
	registerSystemBuiltins()
}

func registerSystemBuiltins() {
	RegisterBuiltin("exit", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.INTEGER_OBJ {
				return newError("argument to `exit` must be an integer, got %s", args[0].Type())
			}
			code := args[0].(*object.Integer).Value
			os.Exit(int(code))
			return nil
		},
	})

	RegisterBuiltin("sleep", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.INTEGER_OBJ {
				return newError("argument to `sleep` must be an integer, got %s", args[0].Type())
			}
			ms := args[0].(*object.Integer).Value
			time.Sleep(time.Duration(ms) * time.Millisecond)
			return nil
		},
	})

	RegisterBuiltin("args", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 0 {
				return newError("args does not accept any arguments")
			}

			cliArgs := &object.Array{}
			for _, arg := range toolkit.GetCliArgs() {
				cliArgs.Elements = append(cliArgs.Elements, &object.String{Value: arg})
			}
			return cliArgs
		},
	})
}
