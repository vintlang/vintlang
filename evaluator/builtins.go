package evaluator

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/toolkit"
)

func init() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
}

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
						return newError("Operation cannot be performed on null")
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
	"args": {
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
	},
	
	// Channel operations
	"send": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("send() takes exactly 2 arguments (channel, value), got %d", len(args))
			}
			
			ch, ok := args[0].(*object.Channel)
			if !ok {
				return newError("first argument to send() must be a channel, got %T", args[0])
			}
			
			if err := ch.Send(args[1]); err != nil {
				return newError("send error: %s", err.Error())
			}
			
			return NULL
		},
	},
	
	"receive": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("receive() takes exactly 1 argument (channel), got %d", len(args))
			}
			
			ch, ok := args[0].(*object.Channel)
			if !ok {
				return newError("argument to receive() must be a channel, got %T", args[0])
			}
			
			value, ok := ch.Receive()
			if !ok {
				return NULL // Channel closed
			}
			
			return value
		},
	},
	
	"close": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("close() takes exactly 1 argument (channel), got %d", len(args))
			}
			
			ch, ok := args[0].(*object.Channel)
			if !ok {
				return newError("argument to close() must be a channel, got %T", args[0])
			}
			
			ch.Close()
			return NULL
		},
	},

	// Math functions
	"abs": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("abs() takes exactly 1 argument, got %d", len(args))
			}
			
			switch arg := args[0].(type) {
			case *object.Integer:
				value := arg.Value
				if value < 0 {
					value = -value
				}
				return &object.Integer{Value: value}
			case *object.Float:
				return &object.Float{Value: math.Abs(arg.Value)}
			default:
				return newError("argument to abs() must be a number, got %s", arg.Type())
			}
		},
	},
	
	"min": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return newError("min() requires at least 1 argument")
			}
			
			var minVal object.Object
			var isFloat bool
			
			for i, arg := range args {
				switch val := arg.(type) {
				case *object.Integer:
					if i == 0 {
						minVal = val
					} else {
						if isFloat {
							currentFloat := float64(val.Value)
							if currentFloat < minVal.(*object.Float).Value {
								minVal = &object.Float{Value: currentFloat}
							}
						} else {
							if val.Value < minVal.(*object.Integer).Value {
								minVal = val
							}
						}
					}
				case *object.Float:
					if i == 0 || !isFloat {
						if i == 0 {
							minVal = val
							isFloat = true
						} else {
							// Convert previous min to float if it was integer
							prevInt := minVal.(*object.Integer).Value
							if val.Value < float64(prevInt) {
								minVal = val
							} else {
								minVal = &object.Float{Value: float64(prevInt)}
							}
							isFloat = true
						}
					} else {
						if val.Value < minVal.(*object.Float).Value {
							minVal = val
						}
					}
				default:
					return newError("all arguments to min() must be numbers, got %s at position %d", val.Type(), i)
				}
			}
			
			return minVal
		},
	},
	
	"max": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return newError("max() requires at least 1 argument")
			}
			
			var maxVal object.Object
			var isFloat bool
			
			for i, arg := range args {
				switch val := arg.(type) {
				case *object.Integer:
					if i == 0 {
						maxVal = val
					} else {
						if isFloat {
							currentFloat := float64(val.Value)
							if currentFloat > maxVal.(*object.Float).Value {
								maxVal = &object.Float{Value: currentFloat}
							}
						} else {
							if val.Value > maxVal.(*object.Integer).Value {
								maxVal = val
							}
						}
					}
				case *object.Float:
					if i == 0 || !isFloat {
						if i == 0 {
							maxVal = val
							isFloat = true
						} else {
							// Convert previous max to float if it was integer
							prevInt := maxVal.(*object.Integer).Value
							if val.Value > float64(prevInt) {
								maxVal = val
							} else {
								maxVal = &object.Float{Value: float64(prevInt)}
							}
							isFloat = true
						}
					} else {
						if val.Value > maxVal.(*object.Float).Value {
							maxVal = val
						}
					}
				default:
					return newError("all arguments to max() must be numbers, got %s at position %d", val.Type(), i)
				}
			}
			
			return maxVal
		},
	},
	
	"round": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("round() takes exactly 1 argument, got %d", len(args))
			}
			
			switch arg := args[0].(type) {
			case *object.Integer:
				return arg // Already an integer
			case *object.Float:
				return &object.Integer{Value: int64(math.Round(arg.Value))}
			default:
				return newError("argument to round() must be a number, got %s", arg.Type())
			}
		},
	},
	
	"floor": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("floor() takes exactly 1 argument, got %d", len(args))
			}
			
			switch arg := args[0].(type) {
			case *object.Integer:
				return arg // Already an integer
			case *object.Float:
				return &object.Integer{Value: int64(math.Floor(arg.Value))}
			default:
				return newError("argument to floor() must be a number, got %s", arg.Type())
			}
		},
	},
	
	"ceil": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("ceil() takes exactly 1 argument, got %d", len(args))
			}
			
			switch arg := args[0].(type) {
			case *object.Integer:
				return arg // Already an integer
			case *object.Float:
				return &object.Integer{Value: int64(math.Ceil(arg.Value))}
			default:
				return newError("argument to ceil() must be a number, got %s", arg.Type())
			}
		},
	},
	
	"sqrt": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("sqrt() takes exactly 1 argument, got %d", len(args))
			}
			
			switch arg := args[0].(type) {
			case *object.Integer:
				if arg.Value < 0 {
					return newError("sqrt() of negative number is not supported")
				}
				return &object.Float{Value: math.Sqrt(float64(arg.Value))}
			case *object.Float:
				if arg.Value < 0 {
					return newError("sqrt() of negative number is not supported")
				}
				return &object.Float{Value: math.Sqrt(arg.Value)}
			default:
				return newError("argument to sqrt() must be a number, got %s", arg.Type())
			}
		},
	},

	// String functions
	"upper": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("upper() takes exactly 1 argument, got %d", len(args))
			}
			
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to upper() must be a string, got %s", args[0].Type())
			}
			
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.ToUpper(str)}
		},
	},
	
	"lower": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("lower() takes exactly 1 argument, got %d", len(args))
			}
			
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to lower() must be a string, got %s", args[0].Type())
			}
			
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.ToLower(str)}
		},
	},
	
	"trim": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("trim() takes exactly 1 argument, got %d", len(args))
			}
			
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to trim() must be a string, got %s", args[0].Type())
			}
			
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.TrimSpace(str)}
		},
	},
	
	"contains": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("contains() takes exactly 2 arguments, got %d", len(args))
			}
			
			// Check if both arguments are strings for string.contains
			if args[0].Type() == object.STRING_OBJ && args[1].Type() == object.STRING_OBJ {
				str := args[0].(*object.String).Value
				substr := args[1].(*object.String).Value
				return &object.Boolean{Value: strings.Contains(str, substr)}
			}
			
			// Check if first argument is array for array.contains
			if args[0].Type() == object.ARRAY_OBJ {
				arr := args[0].(*object.Array)
				for _, element := range arr.Elements {
					if element.Inspect() == args[1].Inspect() {
						return TRUE
					}
				}
				return FALSE
			}
			
			return newError("contains() expects (string, string) or (array, element), got (%s, %s)", 
				args[0].Type(), args[1].Type())
		},
	},
	
	"startsWith": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("startsWith() takes exactly 2 arguments, got %d", len(args))
			}
			
			if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
				return newError("both arguments to startsWith() must be strings, got (%s, %s)", 
					args[0].Type(), args[1].Type())
			}
			
			str := args[0].(*object.String).Value
			prefix := args[1].(*object.String).Value
			return &object.Boolean{Value: strings.HasPrefix(str, prefix)}
		},
	},
	
	"endsWith": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("endsWith() takes exactly 2 arguments, got %d", len(args))
			}
			
			if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
				return newError("both arguments to endsWith() must be strings, got (%s, %s)", 
					args[0].Type(), args[1].Type())
			}
			
			str := args[0].(*object.String).Value
			suffix := args[1].(*object.String).Value
			return &object.Boolean{Value: strings.HasSuffix(str, suffix)}
		},
	},

	// Array functions
	"reverse": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("reverse() takes exactly 1 argument, got %d", len(args))
			}
			
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to reverse() must be an array, got %s", args[0].Type())
			}
			
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			reversed := make([]object.Object, length)
			
			for i, element := range arr.Elements {
				reversed[length-1-i] = element
			}
			
			return &object.Array{Elements: reversed}
		},
	},
	
	"indexOf": {
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
	},

	// Type checking functions
	"isInt": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("isInt() takes exactly 1 argument, got %d", len(args))
			}
			
			return &object.Boolean{Value: args[0].Type() == object.INTEGER_OBJ}
		},
	},
	
	"isFloat": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("isFloat() takes exactly 1 argument, got %d", len(args))
			}
			
			return &object.Boolean{Value: args[0].Type() == object.FLOAT_OBJ}
		},
	},
	
	"isString": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("isString() takes exactly 1 argument, got %d", len(args))
			}
			
			return &object.Boolean{Value: args[0].Type() == object.STRING_OBJ}
		},
	},
	
	"isBool": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("isBool() takes exactly 1 argument, got %d", len(args))
			}
			
			return &object.Boolean{Value: args[0].Type() == object.BOOLEAN_OBJ}
		},
	},
	
	"isArray": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("isArray() takes exactly 1 argument, got %d", len(args))
			}
			
			return &object.Boolean{Value: args[0].Type() == object.ARRAY_OBJ}
		},
	},
	
	"isDict": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("isDict() takes exactly 1 argument, got %d", len(args))
			}
			
			return &object.Boolean{Value: args[0].Type() == object.DICT_OBJ}
		},
	},
	
	"isNull": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("isNull() takes exactly 1 argument, got %d", len(args))
			}
			
			return &object.Boolean{Value: args[0].Type() == object.NULL_OBJ}
		},
	},

	// Additional Array functions
	"sort": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("sort() takes exactly 1 argument, got %d", len(args))
			}
			
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to sort() must be an array, got %s", args[0].Type())
			}
			
			arr := args[0].(*object.Array)
			
			// Create a copy of the array to avoid modifying the original
			sortedElements := make([]object.Object, len(arr.Elements))
			copy(sortedElements, arr.Elements)
			
			// Sort based on the type of elements
			sort.Slice(sortedElements, func(i, j int) bool {
				a, b := sortedElements[i], sortedElements[j]
				
				// Handle different types
				switch aVal := a.(type) {
				case *object.Integer:
					if bVal, ok := b.(*object.Integer); ok {
						return aVal.Value < bVal.Value
					} else if bVal, ok := b.(*object.Float); ok {
						return float64(aVal.Value) < bVal.Value
					}
				case *object.Float:
					if bVal, ok := b.(*object.Float); ok {
						return aVal.Value < bVal.Value
					} else if bVal, ok := b.(*object.Integer); ok {
						return aVal.Value < float64(bVal.Value)
					}
				case *object.String:
					if bVal, ok := b.(*object.String); ok {
						return aVal.Value < bVal.Value
					}
				}
				
				// Fallback to string comparison
				return a.Inspect() < b.Inspect()
			})
			
			return &object.Array{Elements: sortedElements}
		},
	},

	// Random number functions
	"rand": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return newError("rand() takes no arguments, got %d", len(args))
			}
			
			return &object.Float{Value: rand.Float64()}
		},
	},
	
	"randInt": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return newError("randInt() takes 1 or 2 arguments, got %d", len(args))
			}
			
			if len(args) == 1 {
				// randInt(max) - returns 0 to max-1
				maxVal, err := getIntValue(args[0])
				if err != nil {
					return newError("argument to randInt() must be an integer")
				}
				if maxVal <= 0 {
					return newError("argument to randInt() must be positive")
				}
				return &object.Integer{Value: rand.Int63n(maxVal)}
			} else {
				// randInt(min, max) - returns min to max-1
				minVal, err := getIntValue(args[0])
				if err != nil {
					return newError("first argument to randInt() must be an integer")
				}
				maxVal, err := getIntValue(args[1])
				if err != nil {
					return newError("second argument to randInt() must be an integer")
				}
				if maxVal <= minVal {
					return newError("max must be greater than min in randInt()")
				}
				return &object.Integer{Value: minVal + rand.Int63n(maxVal-minVal)}
			}
		},
	},

	// String conversion functions for numbers
	"parseFloat": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("parseFloat() takes exactly 1 argument, got %d", len(args))
			}
			
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to parseFloat() must be a string, got %s", args[0].Type())
			}
			
			str := args[0].(*object.String).Value
			val, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return newError("cannot parse '%s' as float: %s", str, err.Error())
			}
			
			return &object.Float{Value: val}
		},
	},
	
	"parseInt": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("parseInt() takes exactly 1 argument, got %d", len(args))
			}
			
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to parseInt() must be a string, got %s", args[0].Type())
			}
			
			str := args[0].(*object.String).Value
			val, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return newError("cannot parse '%s' as integer: %s", str, err.Error())
			}
			
			return &object.Integer{Value: val}
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

func getFloatValue(obj object.Object) (float64, error) {
	switch obj := obj.(type) {
	case *object.Float:
		return obj.Value, nil
	case *object.Integer:
		return float64(obj.Value), nil
	default:
		return 0, fmt.Errorf("expected Float or Integer, got %s", obj.Type())
	}
}

func getNumericValue(obj object.Object) (float64, bool, error) {
	switch obj := obj.(type) {
	case *object.Integer:
		return float64(obj.Value), false, nil
	case *object.Float:
		return obj.Value, true, nil
	default:
		return 0, false, fmt.Errorf("expected numeric type, got %s", obj.Type())
	}
}
