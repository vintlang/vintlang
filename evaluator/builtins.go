package evaluator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/vintlang/vintlang/object"
)

func handlePrint(w io.Writer, args []object.Object, addNewline bool) object.Object {
	var arr []string
	for _, arg := range args {
		if arg == nil {
			return newError("Operation cannot be performed on nil")
		}
		arr = append(arr, arg.Inspect())
	}
	str := strings.Join(arr, " ")
	if addNewline {
		fmt.Fprintln(w, str)
	} else {
		fmt.Fprint(w, str)
	}
	return nil
}

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
			return handlePrint(os.Stdout, args, false)
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
	"println": {
		Fn: func(args ...object.Object) object.Object {
			return handlePrint(os.Stdout, args, true)
		},
	},
	"printErr": {
		Fn: func(args ...object.Object) object.Object {
			return handlePrint(os.Stderr, args, false)
		},
	},
	"printlnErr": {
		Fn: func(args ...object.Object) object.Object {
			return handlePrint(os.Stderr, args, true)
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

	"eq": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("Sorry, convert requires 2 arguments, you provided %d", len(args))
			}

			a := args[0]
			b := args[1]

			if a == b {
				return &object.Boolean{Value: true}
			}
			return &object.Boolean{Value: false}
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
	"string": {
		// Converts a given object to a string representation
		Fn: func(args ...object.Object) object.Object {
			// Check that exactly one argument is provided
			if len(args) != 1 {
				return newError("string() requires exactly 1 argument, you provided %d", len(args))
			}

			value := args[0]

			// Use the existing convertToString function to handle conversion
			return convertToString(value)
		},
	},
	"int": {
		// Converts a given object to an integer, if possible
		Fn: func(args ...object.Object) object.Object {
			// Check that exactly one argument is provided
			if len(args) != 1 {
				return newError("int() requires exactly 1 argument, you provided %d", len(args))
			}

			value := args[0]

			// Use the existing convertToInteger function to handle conversion
			return convertToInteger(value)
		},
	},

	"and": {
		Fn: func(args ...object.Object) object.Object {
			// Ensure that there are exactly 2 arguments
			if len(args) != 2 {
				return newError("and requires 2 arguments, you provided %d", len(args))
			}

			// Get the boolean value of the first argument
			bool1, err := getBooleanValue(args[0])
			if err != nil {
				// Return an error if the first argument is not a boolean
				return newError("First argument must be a boolean")
			}

			// Get the boolean value of the second argument
			bool2, err := getBooleanValue(args[1])
			if err != nil {
				// Return an error if the second argument is not a boolean
				return newError("Second argument must be a boolean")
			}

			// Perform the logical AND operation and return the result as a boolean
			return &object.Boolean{Value: bool1 && bool2}
		},
	},
	"or": {
		Fn: func(args ...object.Object) object.Object {
			// Ensure that there are exactly 2 arguments
			if len(args) != 2 {
				return newError("or requires 2 arguments, you provided %d", len(args))
			}

			// Get the boolean value of the first argument
			bool1, err := getBooleanValue(args[0])
			if err != nil {
				// Return an error if the first argument is not a boolean
				return newError("First argument must be a boolean")
			}

			// Get the boolean value of the second argument
			bool2, err := getBooleanValue(args[1])
			if err != nil {
				// Return an error if the second argument is not a boolean
				return newError("Second argument must be a boolean")
			}

			// Perform the logical OR operation and return the result as a boolean
			return &object.Boolean{Value: bool1 || bool2}
		},
	},
	"not": {
		Fn: func(args ...object.Object) object.Object {
			// Ensure that there is exactly 1 argument
			if len(args) != 1 {
				return newError("not requires 1 argument, you provided %d", len(args))
			}

			// Get the boolean value of the argument
			boolVal, err := getBooleanValue(args[0])
			if err != nil {
				// Return an error if the argument is not a boolean
				return newError("Argument must be a boolean")
			}

			// Perform the logical NOT operation and return the result as a boolean
			return &object.Boolean{Value: !boolVal}
		},
	},
	"len": {
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
	},
	"append": {
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
	},
	"pop": {
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
	},
	"keys": {
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
	},
	"values": {
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
	},
	"sleep": {
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
	},
	"exit": {
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
	},
	"chr": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.INTEGER_OBJ {
				return newError("argument to `chr` must be an integer, got %s", args[0].Type())
			}
			code := args[0].(*object.Integer).Value
			return &object.String{Value: string(rune(code))}
		},
	},
	"ord": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `ord` must be a string, got %s", args[0].Type())
			}
			s := args[0].(*object.String).Value
			if len(s) != 1 {
				return newError("argument to `ord` must be a single character string")
			}
			return &object.Integer{Value: int64(s[0])}
		},
	},
	"has_key": {
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
