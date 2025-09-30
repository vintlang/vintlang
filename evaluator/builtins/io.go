package builtins

import (
	"os"
	"strings"

	"github.com/vintlang/vintlang/object"
)

func init() {
	registerIOBuiltins()
}

func registerIOBuiltins() {
	RegisterBuiltin("open", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Sorry, this function requires 1 argument, you provided %d", len(args))
			}
			filename := args[0].(*object.String).Value

			file, err := os.ReadFile(filename)
			if err != nil {
				return &object.Error{Message: "Failed to read file or file does not exist"}
			}
			return &object.File{Filename: filename, Content: string(file)}
		},
	})

	RegisterBuiltin("write", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return &object.String{Value: "\n"}
			} else {
				var arr []string
				for _, arg := range args {
					if arg == nil {
						return newError("Operation cannot be performed on null")
					}
					arr = append(arr, arg.Inspect())
				}
				str := strings.Join(arr, " ")
				return &object.String{Value: str}
			}
		},
	})
}