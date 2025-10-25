package module

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vintlang/vintlang/object"
)

var FmtFunctions = map[string]object.ModuleFunction{}

func init() {
	// String formatting functions
	FmtFunctions["sprintf"] = sprintf
	FmtFunctions["printf"] = printf
	FmtFunctions["fprintf"] = fprintf
	FmtFunctions["errorf"] = errorf

	// Padding and alignment functions
	FmtFunctions["padLeft"] = padLeft
	FmtFunctions["padRight"] = padRight
	FmtFunctions["padCenter"] = padCenter

	// Number formatting functions
	FmtFunctions["formatInt"] = formatInt
	FmtFunctions["formatFloat"] = formatFloat
	FmtFunctions["formatHex"] = formatHex
	FmtFunctions["formatOct"] = formatOct
	FmtFunctions["formatBin"] = formatBin

	// Width and precision functions
	FmtFunctions["width"] = width
	FmtFunctions["precision"] = precision

	// Utility functions
	FmtFunctions["repeat"] = repeat
	FmtFunctions["truncate"] = truncate
}

// Helper function to convert VintObject to interface{}
func VintObjectToInterface(obj object.VintObject) interface{} {
	if obj == nil {
		return nil
	}

	switch o := obj.(type) {
	case *object.String:
		return o.Value
	case *object.Integer:
		return o.Value
	case *object.Float:
		return o.Value
	case *object.Boolean:
		return o.Value
	case *object.Null:
		return nil
	default:
		return obj.Inspect()
	}
}

// sprintf formats a string using Go's fmt.Sprintf
func sprintf(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"fmt", "sprintf",
			"at least 1 string argument (format string)",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.sprintf("Hello, %s!", "World") -> returns "Hello, World!"`,
		)
	}

	format := args[0].(*object.String).Value
	var formatArgs []interface{}
	for _, arg := range args[1:] {
		formatArgs = append(formatArgs, VintObjectToInterface(arg))
	}

	result := fmt.Sprintf(format, formatArgs...)
	return &object.String{Value: result}
}

// printf prints formatted text to stdout
func printf(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"fmt", "printf",
			"at least 1 string argument (format string)",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.printf("Hello, %s!\n", "World") -> prints "Hello, World!"`,
		)
	}

	format := args[0].(*object.String).Value
	var formatArgs []interface{}
	for _, arg := range args[1:] {
		formatArgs = append(formatArgs, VintObjectToInterface(arg))
	}

	fmt.Printf(format, formatArgs...)
	return nil
}

// fprintf prints formatted text to a file or writer
func fprintf(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 || args[1].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"fmt", "fprintf",
			"file object and format string",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.fprintf(file, "Hello, %s!\n", "World") -> writes to file`,
		)
	}

	// For now, we'll write to stdout if first arg is not a file
	// In a full implementation, you'd handle file objects properly
	format := args[1].(*object.String).Value
	var formatArgs []interface{}
	for _, arg := range args[2:] {
		formatArgs = append(formatArgs, VintObjectToInterface(arg))
	}

	// Write to stdout for now (could be enhanced to write to actual files)
	fmt.Printf(format, formatArgs...)
	return nil
}

// errorf creates a formatted error
func errorf(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || args[0].Type() != object.STRING_OBJ {
		return ErrorMessage(
			"fmt", "errorf",
			"at least 1 string argument (format string)",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.errorf("error: %s", "something went wrong") -> returns error object`,
		)
	}

	format := args[0].(*object.String).Value
	var formatArgs []interface{}
	for _, arg := range args[1:] {
		formatArgs = append(formatArgs, VintObjectToInterface(arg))
	}

	errorMsg := fmt.Sprintf(format, formatArgs...)
	return &object.Error{Message: errorMsg}
}

// padLeft pads a string to the left with spaces or a specified character
func padLeft(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"fmt", "padLeft",
			"string and integer (width) arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.padLeft("hello", 10) -> "     hello"`,
		)
	}

	str := args[0].(*object.String).Value
	width := int(args[1].(*object.Integer).Value)
	padChar := " "

	if len(args) > 2 && args[2].Type() == object.STRING_OBJ {
		padChar = args[2].(*object.String).Value
		if len(padChar) == 0 {
			padChar = " "
		}
	}

	if len(str) >= width {
		return &object.String{Value: str}
	}

	padding := strings.Repeat(padChar, width-len(str))
	return &object.String{Value: padding + str}
}

// padRight pads a string to the right with spaces or a specified character
func padRight(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"fmt", "padRight",
			"string and integer (width) arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.padRight("hello", 10) -> "hello     "`,
		)
	}

	str := args[0].(*object.String).Value
	width := int(args[1].(*object.Integer).Value)
	padChar := " "

	if len(args) > 2 && args[2].Type() == object.STRING_OBJ {
		padChar = args[2].(*object.String).Value
		if len(padChar) == 0 {
			padChar = " "
		}
	}

	if len(str) >= width {
		return &object.String{Value: str}
	}

	padding := strings.Repeat(padChar, width-len(str))
	return &object.String{Value: str + padding}
}

// padCenter centers a string with padding
func padCenter(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"fmt", "padCenter",
			"string and integer (width) arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.padCenter("hello", 11) -> "   hello   "`,
		)
	}

	str := args[0].(*object.String).Value
	width := int(args[1].(*object.Integer).Value)
	padChar := " "

	if len(args) > 2 && args[2].Type() == object.STRING_OBJ {
		padChar = args[2].(*object.String).Value
		if len(padChar) == 0 {
			padChar = " "
		}
	}

	if len(str) >= width {
		return &object.String{Value: str}
	}

	totalPadding := width - len(str)
	leftPadding := totalPadding / 2
	rightPadding := totalPadding - leftPadding

	left := strings.Repeat(padChar, leftPadding)
	right := strings.Repeat(padChar, rightPadding)

	return &object.String{Value: left + str + right}
}

// formatInt formats an integer with specified base and width
func formatInt(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || args[0].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"fmt", "formatInt",
			"integer argument",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.formatInt(42, 10, 5) -> "   42" (base 10, width 5)`,
		)
	}

	num := args[0].(*object.Integer).Value
	base := 10
	width := 0

	if len(args) > 1 && args[1].Type() == object.INTEGER_OBJ {
		base = int(args[1].(*object.Integer).Value)
	}
	if len(args) > 2 && args[2].Type() == object.INTEGER_OBJ {
		width = int(args[2].(*object.Integer).Value)
	}

	formatted := strconv.FormatInt(num, base)
	if width > len(formatted) {
		padding := strings.Repeat(" ", width-len(formatted))
		formatted = padding + formatted
	}

	return &object.String{Value: formatted}
}

// formatFloat formats a float with specified precision
func formatFloat(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || args[0].Type() != object.FLOAT_OBJ {
		return ErrorMessage(
			"fmt", "formatFloat",
			"float argument",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.formatFloat(3.14159, 2) -> "3.14"`,
		)
	}

	num := args[0].(*object.Float).Value
	precision := 2

	if len(args) > 1 && args[1].Type() == object.INTEGER_OBJ {
		precision = int(args[1].(*object.Integer).Value)
	}

	formatted := strconv.FormatFloat(num, 'f', precision, 64)
	return &object.String{Value: formatted}
}

// formatHex formats an integer as hexadecimal
func formatHex(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || args[0].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"fmt", "formatHex",
			"integer argument",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.formatHex(255) -> "ff"`,
		)
	}

	num := args[0].(*object.Integer).Value
	uppercase := false

	if len(args) > 1 && args[1].Type() == object.BOOLEAN_OBJ {
		uppercase = args[1].(*object.Boolean).Value
	}

	formatted := strconv.FormatInt(num, 16)
	if uppercase {
		formatted = strings.ToUpper(formatted)
	}

	return &object.String{Value: formatted}
}

// formatOct formats an integer as octal
func formatOct(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || args[0].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"fmt", "formatOct",
			"integer argument",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.formatOct(64) -> "100"`,
		)
	}

	num := args[0].(*object.Integer).Value
	formatted := strconv.FormatInt(num, 8)
	return &object.String{Value: formatted}
}

// formatBin formats an integer as binary
func formatBin(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 || args[0].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"fmt", "formatBin",
			"integer argument",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.formatBin(15) -> "1111"`,
		)
	}

	num := args[0].(*object.Integer).Value
	formatted := strconv.FormatInt(num, 2)
	return &object.String{Value: formatted}
}

// width formats a string to a specific width
func width(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"fmt", "width",
			"string and integer (width) arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.width("hello", 10) -> "hello     " (right-padded)`,
		)
	}

	str := args[0].(*object.String).Value
	w := int(args[1].(*object.Integer).Value)

	if len(str) >= w {
		return &object.String{Value: str[:w]} // truncate if too long
	}

	padding := strings.Repeat(" ", w-len(str))
	return &object.String{Value: str + padding}
}

// precision limits a float to a certain number of decimal places (returns formatted string)
func precision(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 || args[0].Type() != object.FLOAT_OBJ || args[1].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"fmt", "precision",
			"float and integer (precision) arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.precision(3.14159, 2) -> "3.14"`,
		)
	}

	num := args[0].(*object.Float).Value
	prec := int(args[1].(*object.Integer).Value)

	formatted := strconv.FormatFloat(num, 'f', prec, 64)
	return &object.String{Value: formatted}
}

// repeat repeats a string n times
func repeat(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 2 || args[0].Type() != object.STRING_OBJ || args[1].Type() != object.INTEGER_OBJ {
		return ErrorMessage(
			"fmt", "repeat",
			"string and integer (count) arguments",
			fmt.Sprintf("%d arguments", len(args)),
			`fmt.repeat("hi", 3) -> "hihihi"`,
		)
	}

	str := args[0].(*object.String).Value
	count := int(args[1].(*object.Integer).Value)

	if count < 0 {
		count = 0
	}

	result := strings.Repeat(str, count)
	return &object.String{Value: result}
}
