package module

import (
	"strings"

	"github.com/fatih/color"
	"github.com/vintlang/vintlang/object"
)

var (
	// Basic colors
	Cyan    = color.New(color.FgCyan)
	Green   = color.New(color.FgGreen)
	Red     = color.New(color.FgRed)
	Yellow  = color.New(color.FgYellow)
	Magenta = color.New(color.FgMagenta)
	White   = color.New(color.FgWhite)
	Blue    = color.New(color.FgBlue)

	// Text styles
	Bold      = color.New(color.Bold)
	Underline = color.New(color.Underline)

	// Custom styles
	HeaderStyle   = color.New(color.FgGreen, color.Bold)              // Titles & headers
	ErrorStyle    = color.New(color.FgRed, color.Bold, color.BgBlack) // Errors
	SuccessStyle  = color.New(color.FgHiGreen, color.Bold)            // Success messages
	WarningStyle  = color.New(color.FgYellow, color.Bold)             // Warnings
	InfoStyle     = color.New(color.FgCyan, color.Bold)               // Information
	DebugStyle    = color.New(color.FgMagenta, color.Bold)            // Debug messages
	InputPrompt   = color.New(color.FgBlue, color.Bold)               // Input prompts
	Highlight     = color.New(color.FgHiBlue, color.Bold)             // Highlights key info
	DimText       = color.New(color.FgWhite, color.Faint)             // Subtle text
	InvertedStyle = color.New(color.FgBlack, color.BgWhite)           // Inverted text
)

var StyledFunctions = map[string]object.ModuleFunction{}

func init() {
	// Basic colors
	StyledFunctions["red"] = red
	StyledFunctions["green"] = green
	StyledFunctions["yellow"] = yellow
	StyledFunctions["blue"] = blue

	StyledFunctions["header"] = header
	StyledFunctions["error"] = err
	StyledFunctions["success"] = success
	StyledFunctions["warning"] = warning
	StyledFunctions["info"] = info
	StyledFunctions["debug"] = debug
	StyledFunctions["inputPrompt"] = inputPrompt
	StyledFunctions["highlight"] = highlight
	StyledFunctions["dim"] = dim
	StyledFunctions["inverted"] = inverted
}

// red prints text in red color
func red(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, Red)
}

// green prints text in green color
func green(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, Green)
}

// yellow prints text in yellow color
func yellow(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, Yellow)
}

// blue prints text in blue color
func blue(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, Blue)
}

// header prints text using HeaderStyle
func header(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, HeaderStyle)
}

// err prints text using ErrorStyle
func err(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, ErrorStyle)
}

// success prints text using SuccessStyle
func success(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, SuccessStyle)
}

// warning prints text using WarningStyle
func warning(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, WarningStyle)
}

// info prints text using InfoStyle
func info(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, InfoStyle)
}

// debug prints text using DebugStyle
func debug(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, DebugStyle)
}

// inputPrompt prints text using InputPrompt style
func inputPrompt(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, InputPrompt)
}

// highlight prints text using Highlight style
func highlight(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, Highlight)
}

// dim prints text using DimText style
func dim(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, DimText)
}

// inverted prints text using InvertedStyle
func inverted(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	return printStyled(args, InvertedStyle)
}

// Utility function to handle styled printing
func printStyled(args []object.VintObject, style *color.Color) object.VintObject {
	if len(args) == 0 {
		style.Println("")
		return nil
	}
	var arr []string
	for _, arg := range args {
		if arg == nil {
			return &object.Error{Message: "Operation cannot be performed on nil"}
		}
		arr = append(arr, arg.Inspect())
	}
	str := strings.Join(arr, " ")
	style.Println(str)
	return nil
}
