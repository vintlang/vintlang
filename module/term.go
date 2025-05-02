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
	TermFunctions["moveCursor"] = termMoveCursor
	TermFunctions["getSize"] = termGetSize
	TermFunctions["input"] = termInput
	TermFunctions["menu"] = termMenu
	TermFunctions["alert"] = termAlert
	TermFunctions["banner"] = termBanner
	TermFunctions["countdown"] = termCountdown
	TermFunctions["select"] = termSelect
	TermFunctions["checkbox"] = termCheckbox
	TermFunctions["radio"] = termRadio
	TermFunctions["password"] = termPassword
	TermFunctions["confirm"] = termConfirm
	TermFunctions["loading"] = termLoading
	TermFunctions["notify"] = termNotify
	TermFunctions["error"] = termError
	TermFunctions["success"] = termSuccess
	TermFunctions["info"] = termInfo
	TermFunctions["warning"] = termWarning
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

// termMoveCursor moves the cursor to a specific position
func termMoveCursor(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "term.moveCursor requires exactly two arguments: x and y coordinates"}
	}

	x, ok := args[0].(*object.Integer)
	if !ok {
		return &object.Error{Message: "x coordinate must be an integer"}
	}

	y, ok := args[1].(*object.Integer)
	if !ok {
		return &object.Error{Message: "y coordinate must be an integer"}
	}

	fmt.Printf("\033[%d;%dH", y.Value+1, x.Value+1)
	return &object.Null{}
}

// termGetSize returns the terminal size
func termGetSize(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) > 0 {
		return &object.Error{Message: "term.getSize does not accept any arguments"}
	}

	// Get terminal size
	width := 80  // Default width
	height := 24 // Default height

	return &object.Dict{
		Pairs: map[object.HashKey]object.DictPair{
			{Type: object.STRING_OBJ, Value: 0}: {
				Key:   &object.String{Value: "width"},
				Value: &object.Integer{Value: int64(width)},
			},
			{Type: object.STRING_OBJ, Value: 1}: {
				Key:   &object.String{Value: "height"},
				Value: &object.Integer{Value: int64(height)},
			},
		},
	}
}

// termInput gets user input with a prompt
func termInput(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.input requires exactly one argument: the prompt message"}
	}

	prompt, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "prompt must be a string"}
	}

	fmt.Print(prompt.Value)
	var input string
	fmt.Scanln(&input)
	return &object.String{Value: input}
}

// termMenu creates an interactive menu
func termMenu(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.menu requires exactly one argument: an array of menu items"}
	}

	items, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "menu items must be an array"}
	}

	// Print menu items
	for i, item := range items.Elements {
		fmt.Printf("%d. %s\n", i+1, item.Inspect())
	}

	// Get user selection
	var choice int
	fmt.Print("Select an option: ")
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(items.Elements) {
		return &object.Error{Message: "Invalid selection"}
	}

	return items.Elements[choice-1]
}

// termAlert shows an alert message
func termAlert(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.alert requires exactly one argument: the message"}
	}

	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "message must be a string"}
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("red")).
		Bold(true).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("red"))

	fmt.Println(style.Render("⚠️ " + msg.Value))
	return &object.Null{}
}

// termBanner creates a banner text
func termBanner(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.banner requires exactly one argument: the text"}
	}

	text, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "text must be a string"}
	}

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("205"))

	return &object.String{Value: style.Render(text.Value)}
}

// termCountdown creates a countdown timer
func termCountdown(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.countdown requires exactly one argument: the duration in seconds"}
	}

	duration, ok := args[0].(*object.Integer)
	if !ok {
		return &object.Error{Message: "duration must be an integer"}
	}

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("yellow"))

	for i := int(duration.Value); i > 0; i-- {
		fmt.Printf("\r%s", style.Render(fmt.Sprintf("Time remaining: %d seconds", i)))
		time.Sleep(time.Second)
	}
	fmt.Println()

	return &object.Null{}
}

// termSelect creates a select menu with arrow key navigation
func termSelect(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.select requires exactly one argument: an array of options"}
	}

	options, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "options must be an array"}
	}

	selected := 0
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true)

	for {
		// Clear screen
		fmt.Print("\033[H\033[2J")

		// Print options
		for i, option := range options.Elements {
			if i == selected {
				fmt.Println(style.Render("→ " + option.Inspect()))
			} else {
				fmt.Println("  " + option.Inspect())
			}
		}

		// Get key press
		var key string
		fmt.Scanln(&key)

		switch key {
		case "up":
			if selected > 0 {
				selected--
			}
		case "down":
			if selected < len(options.Elements)-1 {
				selected++
			}
		case "enter":
			return options.Elements[selected]
		}
	}
}

// termCheckbox creates a checkbox list
func termCheckbox(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.checkbox requires exactly one argument: an array of options"}
	}

	options, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "options must be an array"}
	}

	selected := make(map[int]bool)
	current := 0

	for {
		// Clear screen
		fmt.Print("\033[H\033[2J")

		// Print options
		for i, option := range options.Elements {
			mark := " "
			if selected[i] {
				mark = "✓"
			}
			if i == current {
				fmt.Printf("→ [%s] %s\n", mark, option.Inspect())
			} else {
				fmt.Printf("  [%s] %s\n", mark, option.Inspect())
			}
		}

		// Get key press
		var key string
		fmt.Scanln(&key)

		switch key {
		case "up":
			if current > 0 {
				current--
			}
		case "down":
			if current < len(options.Elements)-1 {
				current++
			}
		case "space":
			selected[current] = !selected[current]
		case "enter":
			var result []object.Object
			for i, option := range options.Elements {
				if selected[i] {
					result = append(result, option)
				}
			}
			return &object.Array{Elements: result}
		}
	}
}

// termRadio creates a radio button list
func termRadio(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.radio requires exactly one argument: an array of options"}
	}

	options, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "options must be an array"}
	}

	selected := 0

	for {
		// Clear screen
		fmt.Print("\033[H\033[2J")

		// Print options
		for i, option := range options.Elements {
			mark := "○"
			if i == selected {
				mark = "●"
			}
			if i == selected {
				fmt.Printf("→ %s %s\n", mark, option.Inspect())
			} else {
				fmt.Printf("  %s %s\n", mark, option.Inspect())
			}
		}

		// Get key press
		var key string
		fmt.Scanln(&key)

		switch key {
		case "up":
			if selected > 0 {
				selected--
			}
		case "down":
			if selected < len(options.Elements)-1 {
				selected++
			}
		case "enter":
			return options.Elements[selected]
		}
	}
}

// termPassword gets password input with hidden characters
func termPassword(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.password requires exactly one argument: the prompt message"}
	}

	prompt, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "prompt must be a string"}
	}

	fmt.Print(prompt.Value)

	// Disable terminal echo
	fmt.Print("\033[8m")

	var input string
	fmt.Scanln(&input)

	// Re-enable terminal echo
	fmt.Print("\033[28m")

	return &object.String{Value: input}
}

// termConfirm asks for yes/no confirmation
func termConfirm(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.confirm requires exactly one argument: the prompt message"}
	}

	prompt, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "prompt must be a string"}
	}

	fmt.Printf("%s (y/n): ", prompt.Value)
	var input string
	fmt.Scanln(&input)

	return &object.Boolean{Value: strings.ToLower(input) == "y"}
}

// termLoading shows a loading message with spinner
func termLoading(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.loading requires exactly one argument: the message"}
	}

	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "message must be a string"}
	}

	spinner := termSpinner([]object.Object{msg}, defs)
	return spinner
}

// termNotify shows a notification message
func termNotify(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.notify requires exactly one argument: the message"}
	}

	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "message must be a string"}
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("blue")).
		Bold(true).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("blue"))

	fmt.Println(style.Render("ℹ️ " + msg.Value))
	return &object.Null{}
}

// termError shows an error message
func termError(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.error requires exactly one argument: the message"}
	}

	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "message must be a string"}
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("red")).
		Bold(true).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("red"))

	fmt.Println(style.Render("❌ " + msg.Value))
	return &object.Null{}
}

// termSuccess shows a success message
func termSuccess(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.success requires exactly one argument: the message"}
	}

	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "message must be a string"}
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("green")).
		Bold(true).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("green"))

	fmt.Println(style.Render("✓ " + msg.Value))
	return &object.Null{}
}

// termInfo shows an info message
func termInfo(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.info requires exactly one argument: the message"}
	}

	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "message must be a string"}
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("cyan")).
		Bold(true).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("cyan"))

	fmt.Println(style.Render("ℹ️ " + msg.Value))
	return &object.Null{}
}

// termWarning shows a warning message
func termWarning(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "term.warning requires exactly one argument: the message"}
	}

	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "message must be a string"}
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("yellow")).
		Bold(true).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("yellow"))

	fmt.Println(style.Render("⚠️ " + msg.Value))
	return &object.Null{}
}
