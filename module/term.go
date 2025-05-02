package module

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/vintlang/vintlang/ast"
	"github.com/vintlang/vintlang/object"
)

var TermFunctions = map[string]object.ModuleFunction{}

func init() {
	TermFunctions["print"] = termPrint
	TermFunctions["println"] = termPrintln
	TermFunctions["clear"] = termClear
	TermFunctions["spinner"] = termSpinner
	TermFunctions["progress"] = termProgress
	TermFunctions["table"] = termTable
	TermFunctions["box"] = termBox
	TermFunctions["style"] = termStyle
	TermFunctions["cursor"] = termCursor
	TermFunctions["beep"] = termBeep
}

// termPrint prints a message with optional color
func termPrint(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return &object.Error{Message: "term.print requires 1-2 arguments: message and optional color"}
	}

	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "message must be a string"}
	}

	style := lipgloss.NewStyle()
	if len(args) == 2 {
		color, ok := args[1].(*object.String)
		if !ok {
			return &object.Error{Message: "color must be a string"}
		}
		style = style.Foreground(lipgloss.Color(color.Value))
	}

	fmt.Print(style.Render(msg.Value))
	return &object.Null{}
}

// termPrintln prints a message with optional color and adds a newline
func termPrintln(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return &object.Error{Message: "term.println requires 1-2 arguments: message and optional color"}
	}

	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "message must be a string"}
	}

	style := lipgloss.NewStyle()
	if len(args) == 2 {
		color, ok := args[1].(*object.String)
		if !ok {
			return &object.Error{Message: "color must be a string"}
		}
		style = style.Foreground(lipgloss.Color(color.Value))
	}

	fmt.Println(style.Render(msg.Value))
	return &object.Null{}
}

// termClear clears the terminal screen
func termClear(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 0 {
		return &object.Error{Message: "term.clear does not accept any arguments"}
	}

	fmt.Print("\033[H\033[2J")
	return &object.Null{}
}

// termSpinner creates a loading spinner
func termSpinner(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.spinner requires exactly one argument: the message"}
	}

	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "message must be a string"}
	}

	spinner := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	i := 0

	// Create a channel to stop the spinner
	stop := make(chan bool)

	// Start the spinner in a goroutine
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				fmt.Printf("\r%s %s", spinner[i], msg.Value)
				i = (i + 1) % len(spinner)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Return a function to stop the spinner
	return &object.Function{
		Parameters: []*ast.Identifier{},
		Body: &ast.BlockStatement{
			Statements: []ast.Statement{
				&ast.ExpressionStatement{
					Expression: &ast.CallExpression{
						Function:  &ast.Identifier{Value: "stop"},
						Arguments: []ast.Expression{},
					},
				},
			},
		},
		Env: object.NewEnclosedEnvironment(nil),
	}
}

// termProgress creates a progress bar
func termProgress(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.progress requires exactly one argument: the total value"}
	}

	total, ok := args[0].(*object.Integer)
	if !ok {
		return &object.Error{Message: "total must be an integer"}
	}

	width := 30
	bar := make([]rune, width)
	for i := range bar {
		bar[i] = ' '
	}

	return &object.Function{
		Parameters: []*ast.Identifier{
			{Value: "value"},
		},
		Body: &ast.BlockStatement{
			Statements: []ast.Statement{
				&ast.ExpressionStatement{
					Expression: &ast.CallExpression{
						Function: &ast.Identifier{Value: "updateProgress"},
						Arguments: []ast.Expression{
							&ast.Identifier{Value: "value"},
							&ast.IntegerLiteral{Value: total.Value},
							&ast.IntegerLiteral{Value: int64(width)},
						},
					},
				},
			},
		},
		Env: object.NewEnclosedEnvironment(nil),
	}
}

// updateProgress updates the progress bar
func updateProgress(current, total, width int64) {
	percentage := float64(current) / float64(total)
	filled := int(float64(width) * percentage)

	bar := make([]rune, width)
	for i := range bar {
		if i < filled {
			bar[i] = '█'
		} else {
			bar[i] = '░'
		}
	}

	fmt.Printf("\r[%s] %d%%", string(bar), int(percentage*100))
}

// termTable creates a formatted table
func termTable(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.table requires exactly one argument: an array of rows"}
	}

	rows, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "rows must be an array"}
	}

	// Convert rows to strings
	var tableRows []string
	for _, row := range rows.Elements {
		if rowArray, ok := row.(*object.Array); ok {
			var cells []string
			for _, cell := range rowArray.Elements {
				cells = append(cells, cell.Inspect())
			}
			tableRows = append(tableRows, strings.Join(cells, " | "))
		}
	}

	// Create table style
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240"))

	// Join rows with newlines
	table := strings.Join(tableRows, "\n")
	return &object.String{Value: style.Render(table)}
}

// termBox creates a boxed text
func termBox(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.box requires exactly one argument: the text"}
	}

	text, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "text must be a string"}
	}

	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2)

	return &object.String{Value: style.Render(text.Value)}
}

// termStyle creates a styled text
func termStyle(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return &object.Error{Message: "term.style requires 1-2 arguments: text and optional style options"}
	}

	text, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "text must be a string"}
	}

	style := lipgloss.NewStyle()
	if len(args) == 2 {
		options, ok := args[1].(*object.Dict)
		if !ok {
			return &object.Error{Message: "style options must be a dictionary"}
		}

		for _, pair := range options.Pairs {
			key := pair.Key.(*object.String).Value
			value := pair.Value.(*object.String).Value

			switch key {
			case "color":
				style = style.Foreground(lipgloss.Color(value))
			case "background":
				style = style.Background(lipgloss.Color(value))
			case "bold":
				style = style.Bold(true)
			case "italic":
				style = style.Italic(true)
			case "underline":
				style = style.Underline(true)
			}
		}
	}

	return &object.String{Value: style.Render(text.Value)}
}

// termCursor controls cursor visibility
func termCursor(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.cursor requires exactly one argument: visibility (true/false)"}
	}

	visible, ok := args[0].(*object.Boolean)
	if !ok {
		return &object.Error{Message: "visibility must be a boolean"}
	}

	if visible.Value {
		fmt.Print("\033[?25h") // Show cursor
	} else {
		fmt.Print("\033[?25l") // Hide cursor
	}

	return &object.Null{}
}

// termBeep plays a terminal beep
func termBeep(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 0 {
		return &object.Error{Message: "term.beep does not accept any arguments"}
	}

	fmt.Print("\a")
	return &object.Null{}
}
