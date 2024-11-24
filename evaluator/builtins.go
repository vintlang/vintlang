package evaluator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ekilie/vint-lang/object"
)

var builtins = map[string]*object.Builtin{
	"input": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) > 1 {
				return newError("Sorry, this function accepts 0 or 1 argument, you provided %d", len(args))
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
	},
	"print": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				fmt.Println("")
			} else {
				var arr []string
				for _, arg := range args {
					if arg == nil {
						return newError("Operation cannot be performed on nil")
					}
					arr = append(arr, arg.Inspect())
				}
				str := strings.Join(arr, " ")
				fmt.Println(str)
			}
			return nil
		},
	},
	"write": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return &object.String{Value: "\n"}
			} else {
				var arr []string
				for _, arg := range args {
					if arg == nil {
						return newError("Operation cannot be performed on nil")
					}
					arr = append(arr, arg.Inspect())
				}
				str := strings.Join(arr, " ")
				return &object.String{Value: str}
			}
		},
	},
	"type": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Sorry, this function requires 1 argument, you provided %d", len(args))
			}

			return &object.String{Value: string(args[0].Type())}
		},
	},
	"open": {
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
	},
	"range": {
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
	},

	"convert": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("Sorry, convert requires 2 arguments, you provided %d", len(args))
			}

			value := args[0]
			targetType := args[1]

			if targetType.Type() != object.STRING_OBJ {
				return newError("Target type must be a string")
			}

			targetTypeStr := targetType.(*object.String).Value

			switch targetTypeStr {
			case "INTEGER":
				return convertToInteger(value)
			case "FLOAT":
				return convertToFloat(value)
			case "STRING":
				return convertToString(value)
			case "BOOLEAN":
				return convertToBoolean(value)
			default:
				return newError("Unknown type: %s", targetTypeStr)
			}
		},
	},
}

func getIntValue(obj object.Object) (int64, error) {
	switch obj := obj.(type) {
	case *object.Integer:
		return obj.Value, nil
	default:
		return 0, fmt.Errorf("expected Integer, got %s", obj.Type())
	}
}
