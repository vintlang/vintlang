package module

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
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
	TermFunctions["layout"] = termLayout
	TermFunctions["grid"] = termGrid
	TermFunctions["tabs"] = termTabs
	TermFunctions["accordion"] = termAccordion
	TermFunctions["tree"] = termTree
	TermFunctions["chart"] = termChart
	TermFunctions["gauge"] = termGauge
	TermFunctions["heatmap"] = termHeatmap
	TermFunctions["calendar"] = termCalendar
	TermFunctions["timeline"] = termTimeline
	TermFunctions["kanban"] = termKanban
	TermFunctions["split"] = termSplit
	TermFunctions["modal"] = termModal
	TermFunctions["tooltip"] = termTooltip
	TermFunctions["badge"] = termBadge
	TermFunctions["avatar"] = termAvatar
	TermFunctions["card"] = termCard
	TermFunctions["list"] = termList
	TermFunctions["form"] = termForm
	TermFunctions["wizard"] = termWizard
	TermFunctions["dashboard"] = termDashboard
}

// termPrint prints a message with optional color
func termPrint(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termPrintln(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termClear(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) > 0 {
		return &object.Error{Message: "term.clear does not accept any arguments"}
	}

	fmt.Print("\033[H\033[2J")
	return &object.Null{}
}

// termSpinner creates a simple loading indicator
func termSpinner(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.spinner requires exactly one argument: the message"}
	}

	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "message must be a string"}
	}

	// Display a simple loading message
	fmt.Printf("⏳ %s\n", msg.Value)

	// Return null for now - in future we could return a function to stop the spinner
	return &object.Null{}
}

// termProgress creates a progress bar function
func termProgress(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.progress requires exactly one argument: the total value"}
	}

	total, ok := args[0].(*object.Integer)
	if !ok {
		return &object.Error{Message: "total must be an integer"}
	}

	// Return a simple progress function that can be called with current value
	return &object.String{Value: fmt.Sprintf("Progress initialized with total: %d", total.Value)}
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
func termTable(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termBox(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termStyle(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termCursor(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termBeep(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) > 0 {
		return &object.Error{Message: "term.beep does not accept any arguments"}
	}

	fmt.Print("\a")
	return &object.Null{}
}

// termMoveCursor moves the cursor to a specific position
func termMoveCursor(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termGetSize(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termInput(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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

// termMenu creates an interactive menu with numbered options
func termMenu(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.menu requires exactly one argument: an array of menu items"}
	}

	items, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "menu items must be an array"}
	}

	// Show menu items with numbers
	fmt.Println("Menu:")
	for i, item := range items.Elements {
		fmt.Printf("%d. %s\n", i+1, item.Inspect())
	}

	// Get user selection
	var choice int
	for {
		fmt.Print("Select option (1-" + fmt.Sprintf("%d", len(items.Elements)) + "): ")

		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			continue
		}

		if n, err := fmt.Sscanf(input, "%d", &choice); n == 1 && err == nil {
			if choice >= 1 && choice <= len(items.Elements) {
				return items.Elements[choice-1]
			}
		}
		fmt.Println("Invalid selection. Please try again.")
	}
}

// termAlert shows an alert message
func termAlert(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termBanner(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termCountdown(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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

// termSelect creates a select menu with simple numbered selection
func termSelect(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.select requires exactly one argument: an array of options"}
	}

	options, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "options must be an array"}
	}

	// Convert to string slice for easier handling
	var optionStrs []string
	for _, option := range options.Elements {
		optionStrs = append(optionStrs, option.Inspect())
	}

	// Show options
	for i, option := range optionStrs {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	// Get user selection
	var choice int
	for {
		fmt.Print("Select option (1-" + fmt.Sprintf("%d", len(optionStrs)) + "): ")

		// Read input using a more reliable method
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			continue
		}

		// Parse the input
		if n, err := fmt.Sscanf(input, "%d", &choice); n == 1 && err == nil {
			if choice >= 1 && choice <= len(options.Elements) {
				return options.Elements[choice-1]
			}
		}
		fmt.Println("Invalid selection. Please try again.")
	}
}

// termCheckbox creates a checkbox list with simple number-based selection
func termCheckbox(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.checkbox requires exactly one argument: an array of options"}
	}

	options, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "options must be an array"}
	}

	// Convert to string slice
	var optionStrs []string
	for _, option := range options.Elements {
		optionStrs = append(optionStrs, option.Inspect())
	}

	// Show options
	fmt.Println("Select multiple options (separate numbers with spaces, e.g., '1 3 4'):")
	for i, option := range optionStrs {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	// Get user selections
	fmt.Print("Enter your choices: ")
	var input string
	fmt.Scanln(&input)

	// Parse selections
	selected := make(map[int]bool)
	var choice int
	for _, part := range strings.Fields(input) {
		if n, err := fmt.Sscanf(part, "%d", &choice); n == 1 && err == nil {
			if choice >= 1 && choice <= len(options.Elements) {
				selected[choice-1] = true
			}
		}
	}

	// Build result array
	var result []object.VintObject
	for i, option := range options.Elements {
		if selected[i] {
			result = append(result, option)
		}
	}

	return &object.Array{Elements: result}
}

// termRadio creates a radio button list with numbered selection
func termRadio(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.radio requires exactly one argument: an array of options"}
	}

	options, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "options must be an array"}
	}

	// Convert to string slice
	var optionStrs []string
	for _, option := range options.Elements {
		optionStrs = append(optionStrs, option.Inspect())
	}

	// Show options
	for i, option := range optionStrs {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	// Get user selection
	var choice int
	for {
		fmt.Print("Select option (1-" + fmt.Sprintf("%d", len(optionStrs)) + "): ")

		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			continue
		}

		if n, err := fmt.Sscanf(input, "%d", &choice); n == 1 && err == nil {
			if choice >= 1 && choice <= len(options.Elements) {
				return options.Elements[choice-1]
			}
		}
		fmt.Println("Invalid selection. Please try again.")
	}
}

// termPassword gets password input with hidden characters
func termPassword(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.password requires exactly one argument: the prompt message"}
	}

	prompt, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "prompt must be a string"}
	}

	fmt.Print(prompt.Value)

	// For now, just use regular input. In a real implementation,
	// we would use golang.org/x/term for proper password input
	var input string
	fmt.Scanln(&input)

	return &object.String{Value: input}
}

// termConfirm asks for yes/no confirmation
func termConfirm(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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

// termLoading shows a loading message
func termLoading(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.loading requires exactly one argument: the message"}
	}

	msg, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "message must be a string"}
	}

	// Display loading message with spinner
	fmt.Printf("⏳ %s\n", msg.Value)

	return &object.Null{}
}

// termNotify shows a notification message
func termNotify(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termError(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termSuccess(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termInfo(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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
func termWarning(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
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

// termLayout creates a flexible layout system
func termLayout(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.layout requires exactly one argument: layout configuration"}
	}

	config, ok := args[0].(*object.Dict)
	if !ok {
		return &object.Error{Message: "layout configuration must be a dictionary"}
	}

	// Parse layout configuration
	layout := lipgloss.NewStyle()
	for _, pair := range config.Pairs {
		key := pair.Key.(*object.String).Value
		value := pair.Value.(*object.String).Value

		switch key {
		case "direction":
			switch value {
			case "horizontal":
				layout = layout.Width(80).Height(24)
			case "vertical":
				layout = layout.Width(24).Height(80)
			}
		case "padding":
			layout = layout.Padding(1, 2)
		case "border":
			layout = layout.BorderStyle(lipgloss.RoundedBorder())
		}
	}

	return &object.String{Value: layout.Render("")}
}

// termGrid creates a grid layout
func termGrid(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return &object.Error{Message: "term.grid requires exactly two arguments: items array and grid configuration"}
	}

	items, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "items must be an array"}
	}

	config, ok := args[1].(*object.Dict)
	if !ok {
		return &object.Error{Message: "grid configuration must be a dictionary"}
	}

	// Parse grid configuration
	columns := 3 // default
	for _, pair := range config.Pairs {
		key := pair.Key.(*object.String).Value
		if key == "columns" {
			columns = int(pair.Value.(*object.Integer).Value)
		}
	}

	// Create grid layout
	var grid []string
	for i := 0; i < len(items.Elements); i += columns {
		var row []string
		for j := 0; j < columns && i+j < len(items.Elements); j++ {
			row = append(row, items.Elements[i+j].Inspect())
		}
		grid = append(grid, strings.Join(row, " | "))
	}

	return &object.String{Value: strings.Join(grid, "\n")}
}

// termTabs creates a tabbed interface
func termTabs(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.tabs requires exactly one argument: tabs configuration"}
	}

	tabs, ok := args[0].(*object.Dict)
	if !ok {
		return &object.Error{Message: "tabs configuration must be a dictionary"}
	}

	// Create tab style
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240"))

	// Create tabs
	var tabNames []string
	for _, pair := range tabs.Pairs {
		tabNames = append(tabNames, pair.Key.(*object.String).Value)
	}

	return &object.String{Value: style.Render(strings.Join(tabNames, " | "))}
}

// termAccordion creates a collapsible accordion
func termAccordion(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.accordion requires exactly one argument: sections configuration"}
	}

	sections, ok := args[0].(*object.Dict)
	if !ok {
		return &object.Error{Message: "sections configuration must be a dictionary"}
	}

	// Create accordion style
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240"))

	// Create sections
	var content []string
	for _, pair := range sections.Pairs {
		title := pair.Key.(*object.String).Value
		content = append(content, fmt.Sprintf("▼ %s", title))
		content = append(content, pair.Value.Inspect())
	}

	return &object.String{Value: style.Render(strings.Join(content, "\n"))}
}

// termTree creates a tree view
func termTree(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.tree requires exactly one argument: tree configuration"}
	}

	tree, ok := args[0].(*object.Dict)
	if !ok {
		return &object.Error{Message: "tree configuration must be a dictionary"}
	}

	// Create tree style
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240"))

	// Create tree structure
	var content []string
	for _, pair := range tree.Pairs {
		content = append(content, fmt.Sprintf("├─ %s", pair.Key.(*object.String).Value))
		if subTree, ok := pair.Value.(*object.Dict); ok {
			for _, subPair := range subTree.Pairs {
				content = append(content, fmt.Sprintf("│  └─ %s", subPair.Key.(*object.String).Value))
			}
		}
	}

	return &object.String{Value: style.Render(strings.Join(content, "\n"))}
}

// termChart creates a simple chart
func termChart(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.chart requires exactly one argument: data array"}
	}

	data, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "data must be an array"}
	}

	// Create chart
	var chart []string
	maxValue := 0
	for _, item := range data.Elements {
		if num, ok := item.(*object.Integer); ok {
			if int(num.Value) > maxValue {
				maxValue = int(num.Value)
			}
		}
	}

	for _, item := range data.Elements {
		if num, ok := item.(*object.Integer); ok {
			bar := strings.Repeat("█", int(float64(num.Value)/float64(maxValue)*20))
			chart = append(chart, bar)
		}
	}

	return &object.String{Value: strings.Join(chart, "\n")}
}

// termGauge creates a gauge/progress indicator
func termGauge(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.gauge requires exactly one argument: value (0-100)"}
	}

	value, ok := args[0].(*object.Integer)
	if !ok {
		return &object.Error{Message: "value must be an integer"}
	}

	// Create gauge
	width := 20
	filled := int(float64(value.Value) / 100.0 * float64(width))
	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)

	return &object.String{Value: fmt.Sprintf("[%s] %d%%", bar, value.Value)}
}

// termHeatmap creates a heatmap visualization
func termHeatmap(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.heatmap requires exactly one argument: data array"}
	}

	data, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "data must be an array"}
	}

	// Create heatmap
	var heatmap []string
	maxValue := 0
	for _, item := range data.Elements {
		if num, ok := item.(*object.Integer); ok {
			if int(num.Value) > maxValue {
				maxValue = int(num.Value)
			}
		}
	}

	colors := []string{"░", "▒", "▓", "█"}
	for _, item := range data.Elements {
		if num, ok := item.(*object.Integer); ok {
			colorIndex := int(float64(num.Value) / float64(maxValue) * float64(len(colors)-1))
			heatmap = append(heatmap, colors[colorIndex])
		}
	}

	return &object.String{Value: strings.Join(heatmap, "")}
}

// termCalendar creates a calendar view
func termCalendar(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.calendar requires exactly one argument: month configuration"}
	}

	config, ok := args[0].(*object.Dict)
	if !ok {
		return &object.Error{Message: "month configuration must be a dictionary"}
	}

	// Create calendar
	now := time.Now()
	year := now.Year()
	month := now.Month()

	for _, pair := range config.Pairs {
		key := pair.Key.(*object.String).Value
		if key == "year" {
			year = int(pair.Value.(*object.Integer).Value)
		} else if key == "month" {
			month = time.Month(pair.Value.(*object.Integer).Value)
		}
	}

	// Generate calendar
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	lastDay := firstDay.AddDate(0, 1, -1)

	calendar := fmt.Sprintf("%s %d\n", month, year)
	calendar += "Su Mo Tu We Th Fr Sa\n"

	// Add leading spaces
	for i := 0; i < int(firstDay.Weekday()); i++ {
		calendar += "   "
	}

	// Add days
	for day := 1; day <= lastDay.Day(); day++ {
		calendar += fmt.Sprintf("%2d ", day)
		if (int(firstDay.Weekday())+day)%7 == 0 {
			calendar += "\n"
		}
	}

	return &object.String{Value: calendar}
}

// termTimeline creates a timeline view
func termTimeline(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.timeline requires exactly one argument: events array"}
	}

	events, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "events must be an array"}
	}

	// Create timeline
	var timeline []string
	for i, event := range events.Elements {
		if eventDict, ok := event.(*object.Dict); ok {
			var title, time string
			for _, pair := range eventDict.Pairs {
				key := pair.Key.(*object.String).Value
				if key == "title" {
					title = pair.Value.(*object.String).Value
				} else if key == "time" {
					time = pair.Value.(*object.String).Value
				}
			}
			timeline = append(timeline, fmt.Sprintf("%s | %s", time, title))
			if i < len(events.Elements)-1 {
				timeline = append(timeline, "    |")
			}
		}
	}

	return &object.String{Value: strings.Join(timeline, "\n")}
}

// termKanban creates a kanban board
func termKanban(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.kanban requires exactly one argument: columns configuration"}
	}

	columns, ok := args[0].(*object.Dict)
	if !ok {
		return &object.Error{Message: "columns configuration must be a dictionary"}
	}

	// Create kanban board
	var board []string
	for _, pair := range columns.Pairs {
		title := pair.Key.(*object.String).Value
		items := pair.Value.(*object.Array)

		column := fmt.Sprintf("=== %s ===\n", title)
		for _, item := range items.Elements {
			column += fmt.Sprintf("• %s\n", item.Inspect())
		}
		board = append(board, column)
	}

	return &object.String{Value: strings.Join(board, "\n")}
}

// termSplit creates a split view
func termSplit(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return &object.Error{Message: "term.split requires exactly two arguments: left content and right content"}
	}

	left, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "left content must be a string"}
	}

	right, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "right content must be a string"}
	}

	// Create split view
	width := 80
	leftWidth := width / 2
	rightWidth := width - leftWidth

	leftStyle := lipgloss.NewStyle().Width(leftWidth)
	rightStyle := lipgloss.NewStyle().Width(rightWidth)

	return &object.String{Value: leftStyle.Render(left.Value) + rightStyle.Render(right.Value)}
}

// termModal creates a modal dialog
func termModal(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.modal requires exactly one argument: modal configuration"}
	}

	config, ok := args[0].(*object.Dict)
	if !ok {
		return &object.Error{Message: "modal configuration must be a dictionary"}
	}

	// Create modal style
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2)

	// Create modal content
	var title, content string
	for _, pair := range config.Pairs {
		key := pair.Key.(*object.String).Value
		if key == "title" {
			title = pair.Value.(*object.String).Value
		} else if key == "content" {
			content = pair.Value.(*object.String).Value
		}
	}

	modal := fmt.Sprintf("%s\n%s", title, content)
	return &object.String{Value: style.Render(modal)}
}

// termTooltip creates a tooltip
func termTooltip(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 2 {
		return &object.Error{Message: "term.tooltip requires exactly two arguments: text and tooltip message"}
	}

	text, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "text must be a string"}
	}

	tooltip, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "tooltip message must be a string"}
	}

	// Create tooltip style
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))

	return &object.String{Value: text.Value + style.Render(" [?] "+tooltip.Value)}
}

// termBadge creates a badge
func termBadge(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.badge requires exactly one argument: badge text"}
	}

	text, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "badge text must be a string"}
	}

	// Create badge style
	style := lipgloss.NewStyle().
		Background(lipgloss.Color("240")).
		Foreground(lipgloss.Color("255")).
		Padding(0, 1)

	return &object.String{Value: style.Render(text.Value)}
}

// termAvatar creates an avatar
func termAvatar(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.avatar requires exactly one argument: avatar text"}
	}

	text, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "avatar text must be a string"}
	}

	// Create avatar style
	style := lipgloss.NewStyle().
		Background(lipgloss.Color("240")).
		Foreground(lipgloss.Color("255")).
		Padding(0, 1)

	return &object.String{Value: style.Render(text.Value[:1])}
}

// termCard creates a card component
func termCard(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.card requires exactly one argument: card configuration"}
	}

	config, ok := args[0].(*object.Dict)
	if !ok {
		return &object.Error{Message: "card configuration must be a dictionary"}
	}

	// Create card style
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2)

	// Create card content
	var title, content string
	for _, pair := range config.Pairs {
		key := pair.Key.(*object.String).Value
		if key == "title" {
			title = pair.Value.(*object.String).Value
		} else if key == "content" {
			content = pair.Value.(*object.String).Value
		}
	}

	card := fmt.Sprintf("%s\n%s", title, content)
	return &object.String{Value: style.Render(card)}
}

// termList creates a list component
func termList(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.list requires exactly one argument: items array"}
	}

	items, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "items must be an array"}
	}

	// Create list style
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240"))

	// Create list items
	var list []string
	for _, item := range items.Elements {
		list = append(list, fmt.Sprintf("• %s", item.Inspect()))
	}

	return &object.String{Value: style.Render(strings.Join(list, "\n"))}
}

// termForm creates a form component
func termForm(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.form requires exactly one argument: form configuration"}
	}

	config, ok := args[0].(*object.Dict)
	if !ok {
		return &object.Error{Message: "form configuration must be a dictionary"}
	}

	// Create form style
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2)

	// Create form fields
	var form []string
	for _, pair := range config.Pairs {
		field := pair.Key.(*object.String).Value
		form = append(form, fmt.Sprintf("%s: [          ]", field))
	}

	return &object.String{Value: style.Render(strings.Join(form, "\n"))}
}

// termWizard creates a wizard/step-by-step form
func termWizard(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.wizard requires exactly one argument: steps array"}
	}

	steps, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "steps must be an array"}
	}

	// Create wizard style
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2)

	// Create wizard steps
	var wizard []string
	for i, step := range steps.Elements {
		wizard = append(wizard, fmt.Sprintf("Step %d: %s", i+1, step.Inspect()))
	}

	return &object.String{Value: style.Render(strings.Join(wizard, "\n"))}
}

// termDashboard creates a dashboard layout
func termDashboard(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) != 1 {
		return &object.Error{Message: "term.dashboard requires exactly one argument: widgets configuration"}
	}

	widgets, ok := args[0].(*object.Dict)
	if !ok {
		return &object.Error{Message: "widgets configuration must be a dictionary"}
	}

	// Create dashboard style
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2)

	// Create dashboard widgets
	var dashboard []string
	for _, pair := range widgets.Pairs {
		title := pair.Key.(*object.String).Value
		content := pair.Value.Inspect()
		dashboard = append(dashboard, fmt.Sprintf("=== %s ===\n%s", title, content))
	}

	return &object.String{Value: style.Render(strings.Join(dashboard, "\n"))}
}
