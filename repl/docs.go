package repl

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
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

	tableOfContentStyle = lipgloss.NewStyle().Margin(1, 2).BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#aa6f5a")).
				Foreground(lipgloss.Color("#aa6f5a")).
				Padding(2)
)

type Item struct {
	title, desc, filename string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.desc }
func (i Item) FilterValue() string { return i.title }

type playground struct {
	id           string
	output       viewport.Model
	code         string
	editor       textarea.Model
	docs         viewport.Model
	ready        bool
	filename     string
	content      []byte
	mybutton     string
	fileSelected bool
	toc          list.Model
	windowWidth  int
	windowHeight int
	docRenderer  *glamour.TermRenderer
}

func (pg playground) Init() tea.Cmd {
	return textarea.Blink
}

func (pg playground) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		edCmd  tea.Cmd
		opCmd  tea.Cmd
		docCmd tea.Cmd
		tocCmd tea.Cmd
	)

	pg.editor, edCmd = pg.editor.Update(msg)
	pg.output, opCmd = pg.output.Update(msg)
	if !pg.fileSelected {
		pg.toc, tocCmd = pg.toc.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			fmt.Println(pg.editor.Value())
			return pg, tea.Quit
		case tea.KeyEnter:
			i, ok := pg.toc.SelectedItem().(Item)
			if ok {
				pg.filename = i.filename
				content, err := res.ReadFile(pg.filename)
				if err != nil {
					pg.docs.SetContent(styles.ErrorStyle.Render("Documentation file not found: ") + pg.filename)
					pg.fileSelected = true
					pg.editor.Focus()
				} else {
					pg.content = content
					str, err := pg.docRenderer.Render(string(pg.content))
					if err != nil {
						pg.docs.SetContent(styles.ErrorStyle.Render("Error rendering documentation: ") + err.Error())
						pg.fileSelected = true
						pg.editor.Focus()
					} else {
						pg.docs.SetContent(str + "\n\n\n\n\n\n")
						pg.fileSelected = true
						pg.editor.Focus()
					}
				}
			}
		case tea.KeyCtrlR:
			if strings.Contains(pg.editor.Value(), "input") {
				pg.output.SetContent(styles.HelpStyle.Italic(false).Render("Sorry, you cannot use input() function in this playground."))
			} else {
				code := strings.ReplaceAll(pg.editor.Value(), "print", "_print")
				pg.code = code
				env := object.NewEnvironment()
				l := lexer.New(pg.code)
				p := parser.New(l)
				program := p.ParseProgram()
				if len(p.Errors()) != 0 {
					pg.output.Style = styles.ErrorStyle.PaddingLeft(3)
					pg.output.SetContent(strings.Join(p.Errors(), "\n"))
				} else {
					evaluated := evaluator.Eval(program, env)
					if evaluated != nil {
						if evaluated.Type() != object.NULL_OBJ {
							pg.output.Style = styles.ReplStyle.PaddingLeft(3)
							content := evaluated.Inspect()
							l := strings.Split(content, "\n")
							if len(l) > 15 {
								content = strings.Join(l[len(l)-16:], "\n")
							}
							pg.output.SetContent(content)
						}
					}
				}
			}
		case tea.KeyEsc:
			if pg.fileSelected {
				pg.fileSelected = false
				pg.editor.Blur()
			}
		}

	case tea.MouseMsg:
		if zone.Get(pg.id + "docs").InBounds(msg) {
			pg.docs, docCmd = pg.docs.Update(msg)
		}
		switch msg.Type {
		case tea.MouseLeft:
			if zone.Get(pg.id + "run").InBounds(msg) {
				if strings.Contains(pg.editor.Value(), "input") {
					pg.output.SetContent(styles.HelpStyle.Italic(false).Render("Sorry, you cannot use input() function in this playground."))
				} else {
					code := strings.ReplaceAll(pg.editor.Value(), "print", "_print")
					pg.code = code
					env := object.NewEnvironment()
					l := lexer.New(pg.code)
					p := parser.New(l)
					program := p.ParseProgram()
					if len(p.Errors()) != 0 {
						pg.output.Style = styles.ErrorStyle.PaddingLeft(3)
						pg.output.SetContent(strings.Join(p.Errors(), "\n"))
					} else {
						evaluated := evaluator.Eval(program, env)
						if evaluated != nil {
							if evaluated.Type() != object.NULL_OBJ {
								pg.output.Style = styles.ReplStyle.PaddingLeft(3)
								content := evaluated.Inspect()
								l := strings.Split(content, "\n")
								if len(l) > 15 {
									content = strings.Join(l[len(l)-16:], "\n")
								}
								pg.output.SetContent(content)
							}
						}
					}
				}
			}
		}
	case tea.WindowSizeMsg:
		if !pg.ready {
			pg.editor = textarea.New()
			pg.editor.Placeholder = "Write vint code here..."
			pg.editor.Prompt = "┃ "
			pg.editor.SetWidth(msg.Width / 2)
			pg.editor.SetHeight((2 * msg.Height / 3) - 4)
			pg.editor.CharLimit = 0
			pg.editor.FocusedStyle.CursorLine = lipgloss.NewStyle()
			pg.editor.FocusedStyle.Base = lipgloss.NewStyle().PaddingTop(2).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("238"))
			pg.editor.ShowLineNumbers = true

			pg.output = viewport.New(msg.Width/2, msg.Height/3-4)
			pg.output.Style = lipgloss.NewStyle().PaddingLeft(3)
			var output string
			output = "Your code output will be displayed here..." + strings.Repeat(" ", msg.Width-6)
			pg.output.SetContent(output)

			pg.docs = viewport.New(msg.Width/2, msg.Height)
			pg.docs.KeyMap = viewport.KeyMap{
				Up: key.NewBinding(
					key.WithKeys("up"),
				),
				Down: key.NewBinding(
					key.WithKeys("down"),
				),
			}
			pg.docs.Style = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62")).
				Padding(2)

			renderer, err := glamour.NewTermRenderer(
				glamour.WithAutoStyle(),
				glamour.WithWordWrap(msg.Width/2-4),
			)
			if err != nil {
				panic(err)
			}

			pg.docRenderer = renderer
			pg.toc.SetSize(msg.Width, msg.Height-8)
			pg.windowWidth = msg.Width
			pg.windowHeight = msg.Height
			pg.mybutton = activeButtonStyle.Width(msg.Width / 2).Height(1).Align(lipgloss.Center).Render("Run (CTRL + R)")
			pg.ready = true

		} else {
			pg.editor.SetHeight((2 * msg.Height / 3) - 4)
			pg.editor.SetWidth(msg.Width / 2)
			pg.output.Height = msg.Height/3 - 4
			pg.output.Width = msg.Width / 2

			renderer, err := glamour.NewTermRenderer(
				glamour.WithAutoStyle(),
				glamour.WithWordWrap(msg.Width/2-4),
			)
			if err != nil {
				panic(err)
			}

			pg.docRenderer = renderer
			str, err := pg.docRenderer.Render(string(pg.content))
			if err != nil {
				panic(err)
			}
			pg.docs.Height = msg.Height
			pg.docs.Width = msg.Width / 2

			pg.docs.SetContent(str + "\n\n\n\n\n\n")
			pg.mybutton = activeButtonStyle.Width(msg.Width / 2).Height(1).Align(lipgloss.Center).Render("Run (CTRL + R)")
			pg.toc.SetSize(msg.Width, msg.Height-8)
			pg.windowWidth = msg.Width
			pg.windowHeight = msg.Height
		}
	}

	return pg, tea.Batch(edCmd, opCmd, docCmd, tocCmd)
}

func (pg playground) View() string {
	if !pg.ready {
		return "\n  Preparing documentation..."
	}
	var docs string
	if !pg.fileSelected {
		docs = zone.Mark(pg.id+"toc", tableOfContentStyle.Width(pg.windowWidth/2-4).Height(pg.windowHeight-8).Render(pg.toc.View()))
	} else {
		docs = zone.Mark(pg.id+"docs", pg.docs.View())
	}
	button := zone.Mark(pg.id+"run", pg.mybutton)
	return zone.Scan(lipgloss.JoinHorizontal(lipgloss.Center, docs, lipgloss.JoinVertical(lipgloss.Left, pg.editor.View(), button, pg.output.View())))
}
