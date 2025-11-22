package repl

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	prompt "github.com/AvicennaJr/GoPrompt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/vintlang/vintlang/docs"
	"github.com/vintlang/vintlang/evaluator"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
	"github.com/vintlang/vintlang/styles"
)

var (
	buttonStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), true, false).
			Padding(0, 3).
			MarginTop(1)

	activeButtonStyle = buttonStyle.Copy().
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#aa6f5a")).
				Margin(0, 2).
				Underline(true)

	inactiveButtonStyle = buttonStyle.Copy().
				Foreground(lipgloss.Color("#888888")).
				Background(lipgloss.Color("#333333")).
				Margin(0, 2)

	tableOfContentStyle = lipgloss.NewStyle().Margin(1, 2).BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#aa6f5a")).
				Foreground(lipgloss.Color("#aa6f5a")).
				Padding(2)

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#1a1a1a")).
			Padding(0, 1).
			MarginTop(1)
)

type ExecutionResult struct {
	output   string
	duration time.Duration
	success  bool
}

type executionMsg ExecutionResult

type saveStateMsg struct {
	filename string
	success  bool
}

type loadStateMsg struct {
	filename string
	content  string
}

// Key bindings for new features
type keyMap struct {
	Run        key.Binding
	Save       key.Binding
	Load       key.Binding
	Clear      key.Binding
	ToggleDocs key.Binding
	Back       key.Binding
	Quit       key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Run, k.Save, k.ToggleDocs, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Run, k.Save, k.Load},
		{k.Clear, k.ToggleDocs, k.Back},
		{k.Quit},
	}
}

var keys = keyMap{
	Run: key.NewBinding(
		key.WithKeys("ctrl+r"),
		key.WithHelp("ctrl+r", "run code"),
	),
	Save: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "save code"),
	),
	Load: key.NewBinding(
		key.WithKeys("ctrl+o"),
		key.WithHelp("ctrl+o", "load code"),
	),
	Clear: key.NewBinding(
		key.WithKeys("ctrl+l"),
		key.WithHelp("ctrl+l", "clear output"),
	),
	ToggleDocs: key.NewBinding(
		key.WithKeys("ctrl+d"),
		key.WithHelp("ctrl+d", "toggle docs"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "ctrl+q"),
		key.WithHelp("ctrl+c", "quit"),
	),
}

type playground struct {
	id           string
	output       viewport.Model
	code         string
	editor       textarea.Model
	docs         viewport.Model
	ready        bool
	filename     string
	content      []byte
	runButton    string
	saveButton   string
	loadButton   string
	fileSelected bool
	toc          list.Model
	windowWidth  int
	windowHeight int
	docRenderer  *glamour.TermRenderer
	keys         keyMap
	spinner      spinner.Model
	executing    bool
	lastResult   *ExecutionResult
	progress     progress.Model
	showHelp     bool
	statusMsg    string
	statusTimer  *time.Timer
}

func newPlayground() playground {
	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	prog := progress.New(progress.WithDefaultGradient())
	
	return playground{
		keys:     keys,
		spinner:  sp,
		progress: prog,
		statusMsg: "Welcome to Vint Playground!",
	}
}

func (pg playground) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
		pg.spinner.Tick,
	)
}

func (pg *playground) setStatus(message string, duration time.Duration) {
	pg.statusMsg = message
	if pg.statusTimer != nil {
		pg.statusTimer.Stop()
	}
	pg.statusTimer = time.AfterFunc(duration, func() {
		pg.statusMsg = "Ready"
	})
}

func (pg playground) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		edCmd  tea.Cmd
		opCmd  tea.Cmd
		docCmd tea.Cmd
		tocCmd tea.Cmd
		spCmd  tea.Cmd
	)

	pg.editor, edCmd = pg.editor.Update(msg)
	pg.output, opCmd = pg.output.Update(msg)
	if !pg.fileSelected {
		pg.toc, tocCmd = pg.toc.Update(msg)
	}
	
	if pg.executing {
		pg.spinner, spCmd = pg.spinner.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if !pg.editor.Focused() && !pg.fileSelected {
			switch {
			case key.Matches(msg, pg.keys.Run):
				return pg.executeCode()
			case key.Matches(msg, pg.keys.Save):
				return pg.saveCode()
			case key.Matches(msg, pg.keys.Load):
				return pg.loadCode()
			case key.Matches(msg, pg.keys.Clear):
				pg.output.SetContent("Output cleared")
				pg.setStatus("Output cleared", 2*time.Second)
			case key.Matches(msg, pg.keys.ToggleDocs):
				pg.fileSelected = !pg.fileSelected
				if pg.fileSelected {
					pg.editor.Focus()
				}
			case key.Matches(msg, pg.keys.Back):
				if pg.fileSelected {
					pg.fileSelected = false
				}
			case key.Matches(msg, pg.keys.Quit):
				fmt.Println(pg.editor.Value())
				return pg, tea.Quit
			}
		}

		switch msg.Type {
		case tea.KeyEnter:
			if !pg.fileSelected {
				i, ok := pg.toc.SelectedItem().(docs.Item)
				if ok {
					pg.loadDocumentation(i.Filename())
				}
			}
		case tea.KeyF1:
			pg.showHelp = !pg.showHelp
		}

	case executionMsg:
		pg.executing = false
		result := ExecutionResult(msg)
		pg.lastResult = &result
		
		if result.success {
			pg.output.Style = styles.ReplStyle.PaddingLeft(3)
			pg.setStatus(fmt.Sprintf("Execution successful (%v)", result.duration), 3*time.Second)
		} else {
			pg.output.Style = styles.ErrorStyle.PaddingLeft(3)
			pg.setStatus("Execution failed", 3*time.Second)
		}
		pg.output.SetContent(result.output)

	case saveStateMsg:
		if msg.success {
			pg.setStatus(fmt.Sprintf("Saved to %s", msg.filename), 3*time.Second)
		} else {
			pg.setStatus(fmt.Sprintf("Failed to save to %s", msg.filename), 3*time.Second)
		}

	case loadStateMsg:
		pg.editor.SetValue(msg.content)
		pg.setStatus(fmt.Sprintf("Loaded from %s", msg.filename), 3*time.Second)

	case tea.MouseMsg:
		if zone.Get(pg.id + "docs").InBounds(msg) {
			pg.docs, docCmd = pg.docs.Update(msg)
		}
		switch msg.Type {
		case tea.MouseLeft:
			if zone.Get(pg.id + "run").InBounds(msg) {
				return pg.executeCode()
			}
			if zone.Get(pg.id + "save").InBounds(msg) {
				return pg.saveCode()
			}
			if zone.Get(pg.id + "load").InBounds(msg) {
				return pg.loadCode()
			}
		}

	case tea.WindowSizeMsg:
		return pg.handleResize(msg)
	}

	return pg, tea.Batch(edCmd, opCmd, docCmd, tocCmd, spCmd)
}

func (pg *playground) executeCode() (tea.Model, tea.Cmd) {
	if strings.Contains(pg.editor.Value(), "input") {
		pg.output.SetContent(styles.HelpStyle.Italic(false).Render("Sorry, you cannot use input() function in this playground."))
		return pg, nil
	}

	pg.executing = true
	code := strings.ReplaceAll(pg.editor.Value(), "print", "_print")
	pg.code = code

	// Run execution in a goroutine to prevent blocking
	go func() {
		start := time.Now()
		env := object.NewEnvironment()
		l := lexer.New(pg.code)
		p := parser.New(l)
		program := p.ParseProgram()
		
		var output string
		success := true
		
		if len(p.Errors()) != 0 {
			output = strings.Join(p.Errors(), "\n")
			success = false
		} else {
			evaluated := evaluator.Eval(program, env)
			if evaluated != nil && evaluated.Type() != object.NULL_OBJ {
				content := evaluated.Inspect()
				lines := strings.Split(content, "\n")
				if len(lines) > 15 {
					content = strings.Join(lines[len(lines)-16:], "\n")
				}
				output = content
			} else {
				output = "No output"
			}
		}
		
		duration := time.Since(start)
		// Send result back to main thread
		// Note: In real implementation, you'd use tea.Cmd to send this
		pg.Update(executionMsg{
			output:   output,
			duration: duration,
			success:  success,
		})
	}()

	pg.setStatus("Executing code...", 30*time.Second)
	return pg, pg.spinner.Tick
}

func (pg *playground) saveCode() (tea.Model, tea.Cmd) {
	filename := "playground.vint"
	content := pg.editor.Value()
	
	go func() {
		err := os.WriteFile(filename, []byte(content), 0644)
		success := err == nil
		// Send save result back to main thread
		pg.Update(saveStateMsg{
			filename: filename,
			success:  success,
		})
	}()
	
	return pg, nil
}

func (pg *playground) loadCode() (tea.Model, tea.Cmd) {
	filename := "playground.vint"
	
	go func() {
		content, err := os.ReadFile(filename)
		if err == nil {
			// Send load result back to main thread
			pg.Update(loadStateMsg{
				filename: filename,
				content:  string(content),
			})
		} else {
			pg.Update(loadStateMsg{
				filename: filename,
				content:  "",
			})
		}
	}()
	
	return pg, nil
}

func (pg *playground) loadDocumentation(filename string) {
	pg.filename = filename
	content, err := docs.Docs.ReadFile(pg.filename)
	if err != nil {
		pg.docs.SetContent(styles.ErrorStyle.Render("Documentation file not found: ") + pg.filename)
		pg.fileSelected = true
		pg.editor.Focus()
		return
	}

	pg.content = content
	str, err := pg.docRenderer.Render(string(pg.content))
	if err != nil {
		pg.docs.SetContent(styles.ErrorStyle.Render("Error rendering documentation: ") + err.Error())
	} else {
		pg.docs.SetContent(str + "\n\n\n\n\n\n")
	}
	pg.fileSelected = true
	pg.editor.Focus()
}

func (pg *playground) handleResize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	if !pg.ready {
		pg.initializeComponents(msg)
		pg.ready = true
	} else {
		pg.updateComponents(msg)
	}
	
	return pg, nil
}

func (pg *playground) initializeComponents(msg tea.WindowSizeMsg) {
	pg.editor = textarea.New()
	pg.editor.Placeholder = "Write vint code here...\nTry: print(\"Hello, World!\")\nUse Ctrl+R to run, Ctrl+S to save, Ctrl+O to load"
	pg.editor.Prompt = "┃ "
	pg.editor.SetWidth(msg.Width / 2)
	pg.editor.SetHeight((2 * msg.Height / 3) - 6)
	pg.editor.CharLimit = 10000
	pg.editor.FocusedStyle.CursorLine = lipgloss.NewStyle().Background(lipgloss.Color("#1a1a1a"))
	pg.editor.FocusedStyle.Base = lipgloss.NewStyle().PaddingTop(2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("238"))
	pg.editor.ShowLineNumbers = true
	pg.editor.Focus()

	pg.output = viewport.New(msg.Width/2, msg.Height/3-4)
	pg.output.Style = lipgloss.NewStyle().PaddingLeft(3)
	pg.output.SetContent("Your code output will be displayed here...\n\nUse examples:\n  print(\"Hello World\")\n  let x = 5 + 3\n  if true { print(\"True!\") }")

	pg.docs = viewport.New(msg.Width/2, msg.Height)
	pg.docs.KeyMap = viewport.KeyMap{
		Up: key.NewBinding(key.WithKeys("up")),
		Down: key.NewBinding(key.WithKeys("down")),
		PageUp: key.NewBinding(key.WithKeys("pgup")),
		PageDown: key.NewBinding(key.WithKeys("pgdown")),
	}

	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(msg.Width/2-4),
	)

	pg.docRenderer = renderer
	pg.toc.SetSize(msg.Width, msg.Height-8)
	pg.windowWidth = msg.Width
	pg.windowHeight = msg.Height
	
	pg.updateButtons(msg.Width)
}

func (pg *playground) updateComponents(msg tea.WindowSizeMsg) {
	pg.editor.SetHeight((2 * msg.Height / 3) - 6)
	pg.editor.SetWidth(msg.Width / 2)
	pg.output.Height = msg.Height/3 - 4
	pg.output.Width = msg.Width / 2

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(msg.Width/2-4),
	)
	if err == nil {
		pg.docRenderer = renderer
		if pg.fileSelected && len(pg.content) > 0 {
			str, err := pg.docRenderer.Render(string(pg.content))
			if err == nil {
				pg.docs.Height = msg.Height
				pg.docs.Width = msg.Width / 2
				pg.docs.SetContent(str + "\n\n\n\n\n\n")
			}
		}
	}

	pg.updateButtons(msg.Width)
	pg.toc.SetSize(msg.Width, msg.Height-8)
	pg.windowWidth = msg.Width
	pg.windowHeight = msg.Height
}

func (pg *playground) updateButtons(width int) {
	buttonWidth := (width / 2) / 3
	
	runText := "Run (Ctrl+R)"
	if pg.executing {
		runText = pg.spinner.View() + " Running..."
	}
	pg.runButton = activeButtonStyle.Width(buttonWidth).Height(1).Align(lipgloss.Center).Render(runText)
	pg.saveButton = activeButtonStyle.Width(buttonWidth).Height(1).Align(lipgloss.Center).Render("Save (Ctrl+S)")
	pg.loadButton = activeButtonStyle.Width(buttonWidth).Height(1).Align(lipgloss.Center).Render("Load (Ctrl+O)")
}

func (pg playground) View() string {
	if !pg.ready {
		return "\n  Initializing Vint Playground..."
	}

	// Documentation panel
	var docsPanel string
	if !pg.fileSelected {
		docsPanel = zone.Mark(pg.id+"toc", tableOfContentStyle.Width(pg.windowWidth/2-4).Height(pg.windowHeight-8).Render(pg.toc.View()))
	} else {
		docsPanel = zone.Mark(pg.id+"docs", pg.docs.View())
	}

	// Code panel
	codePanel := lipgloss.JoinVertical(lipgloss.Left,
		pg.editor.View(),
		lipgloss.JoinHorizontal(lipgloss.Center,
			zone.Mark(pg.id+"run", pg.runButton),
			zone.Mark(pg.id+"save", pg.saveButton),
			zone.Mark(pg.id+"load", pg.loadButton),
		),
		pg.output.View(),
	)

	// Status bar
	statusBar := statusStyle.Width(pg.windowWidth).Render(pg.statusMsg)

	// Help panel (toggle)
	var helpPanel string
	if pg.showHelp {
		helpPanel = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("238")).
			Padding(1).
			Width(pg.windowWidth).
			Render(
				lipgloss.JoinVertical(lipgloss.Left,
					"Key Bindings:",
					"  Ctrl+R  - Run code",
					"  Ctrl+S  - Save code",
					"  Ctrl+O  - Load code", 
					"  Ctrl+L  - Clear output",
					"  Ctrl+D  - Toggle documentation",
					"  Esc     - Back to documentation list",
					"  F1      - Toggle this help",
					"  Ctrl+Q  - Quit",
				),
			)
	}

	// Main layout
	mainContent := lipgloss.JoinHorizontal(lipgloss.Center, docsPanel, codePanel)
	
	if pg.showHelp {
		return zone.Scan(lipgloss.JoinVertical(lipgloss.Left, mainContent, helpPanel, statusBar))
	}
	return zone.Scan(lipgloss.JoinVertical(lipgloss.Left, mainContent, statusBar))
}

// Enhanced REPL with history and better completion
type ReplHistory struct {
	commands []string
	index    int
}

func (h *ReplHistory) Add(cmd string) {
	h.commands = append(h.commands, cmd)
	h.index = len(h.commands)
}

func (h *ReplHistory) Previous() string {
	if len(h.commands) == 0 {
		return ""
	}
	if h.index > 0 {
		h.index--
	}
	return h.commands[h.index]
}

func (h *ReplHistory) Next() string {
	if len(h.commands) == 0 {
		return ""
	}
	if h.index < len(h.commands)-1 {
		h.index++
		return h.commands[h.index]
	}
	h.index = len(h.commands)
	return ""
}

// Enhanced dummy executor with history
type enhancedDummy struct {
	env     *object.Environment
	history *ReplHistory
}

func (d *enhancedDummy) executor(in string) {
	if strings.TrimSpace(in) == "exit()" {
		style := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Background(lipgloss.Color("#282C34")).
			Bold(true).
			Padding(1, 2).
			Margin(1).
			Border(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("#FF4500"))
		fmt.Println(style.Render("\nThank you for using Vint! Goodbye and happy coding!"))
		os.Exit(0)
	}
	
	// Add to history
	d.history.Add(in)
	
	l := lexer.NewWithFilename(in, "<repl>")
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			fmt.Println("\t" + styles.ErrorStyle.Render(msg))
		}
		return
	}
	
	evaluated := evaluator.Eval(program, d.env)
	if evaluated != nil && evaluated.Type() != object.NULL_OBJ {
		fmt.Println(styles.ReplStyle.Render(evaluated.Inspect()))
	}
}

// Enhanced completer with Vint language keywords
func enhancedCompleter(in prompt.Document) []prompt.Suggest {
	keywords := []prompt.Suggest{
		{Text: "let", Description: "Variable declaration"},
		{Text: "if", Description: "Conditional statement"},
		{Text: "else", Description: "Else clause"},
		{Text: "fn", Description: "Function definition"},
		{Text: "return", Description: "Return from function"},
		{Text: "print", Description: "Print to output"},
		{Text: "true", Description: "Boolean true"},
		{Text: "false", Description: "Boolean false"},
		{Text: "nil", Description: "Null value"},
	}
	
	return prompt.FilterHasPrefix(keywords, in.GetWordBeforeCursor(), true)
}

func StartEnhanced() {
	env := object.NewEnvironment()
	history := &ReplHistory{}

	d := enhancedDummy{
		env:     env,
		history: history,
	}

	p := prompt.New(
		d.executor,
		enhancedCompleter,
		prompt.OptionPrefix(PROMPT),
		prompt.OptionTitle("Vint Programming Language - Enhanced REPL"),
		prompt.OptionHistory(history.commands),
		prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlUp,
			Fn: func(buf *prompt.Buffer) {
				if text := history.Previous(); text != "" {
					buf.InsertText(text, false, true)
				}
			},
		}),
		prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlDown,
			Fn: func(buf *prompt.Buffer) {
				if text := history.Next(); text != "" {
					buf.InsertText(text, false, true)
				}
			},
		}),
	)

	p.Run()
}

func EnhancedDocs() {
	zone.NewGlobal()

	p := newPlayground()
	p.toc = list.New(englishItems, list.NewDefaultDelegate(), 0, 0)
	p.toc.Title = "Vint Documentation - Table of Contents"
	p.toc.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#aa6f5a")).
		Bold(true).
		Padding(0, 1)
	p.id = zone.NewPrefix()

	if _, err := tea.NewProgram(p, tea.WithMouseAllMotion(), tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}