package builtins

import (
	"os"

	"github.com/vintlang/vintlang/object"
)

func init() {
	registerIOBuiltins()
}

func registerIOBuiltins() {
	RegisterBuiltin("open", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) != 1 {
				return newError("open() requires 1 argument, you provided %d", len(args))
			}
			str, ok := args[0].(*object.String)
			if !ok {
				return newError("argument to open() must be a string, got %s", args[0].Type())
			}
			filename := str.Value

			file, err := os.ReadFile(filename)
			if err != nil {
				return newError("Failed to read file '%s': %s", filename, err.Error())
			}
			return &object.File{Filename: filename, Content: string(file)}
		},
	})

	RegisterBuiltin("write", &object.Builtin{
		Fn: func(args ...object.VintObject) object.VintObject {
			if len(args) < 2 || len(args) > 3 {
				return newError("write() requires 2 or 3 arguments (filename, content, [mode]), got %d", len(args))
			}

			filenameObj, ok := args[0].(*object.String)
			if !ok {
				return newError("first argument to write() must be a string, got %s", args[0].Type())
			}
			filename := filenameObj.Value

			contentObj, ok := args[1].(*object.String)
			if !ok {
				return newError("second argument to write() must be a string, got %s", args[1].Type())
			}
			content := contentObj.Value

			// Default mode: overwrite
			mode := "w"
			if len(args) == 3 {
				modeObj, ok := args[2].(*object.String)
				if !ok {
					return newError("third argument to write() must be a string, got %s", args[2].Type())
				}
				mode = modeObj.Value
			}

			var flag int
			switch mode {
			case "w":
				flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
			case "a":
				flag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
			default:
				return newError("write() mode must be 'w' (overwrite) or 'a' (append), got '%s'", mode)
			}

			f, err := os.OpenFile(filename, flag, 0644)
			if err != nil {
				return newError("Failed to open file '%s' for writing: %s", filename, err.Error())
			}
			defer f.Close()

			n, err := f.WriteString(content)
			if err != nil {
				return newError("Failed to write to file '%s': %s", filename, err.Error())
			}

			return &object.Integer{Value: int64(n)}
		},
	})
}
