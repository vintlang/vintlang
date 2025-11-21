package repl

import (
	"embed"
	"fmt"
	"log"
	"os"
	"strings"

	prompt "github.com/AvicennaJr/GoPrompt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/vintlang/vintlang/docs"
	"github.com/vintlang/vintlang/evaluator"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
	"github.com/vintlang/vintlang/styles"
)

const PROMPT = ">>> "

var res embed.FS = docs.Docs

func Read(contents string) {
	ReadWithFilename(contents, "<input>")
}

func ReadWithFilename(contents string, filename string) {
	env := object.NewEnvironment()

	l := lexer.NewWithFilename(contents, filename)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println(styles.ErrorStyle.Italic(false).Render("These errors occured:"))

		for _, msg := range p.Errors() {
			fmt.Println("\t" + styles.ErrorStyle.Render(msg))
		}
		return // Don't evaluate if there are parser errors
	}
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() != object.NULL_OBJ {
			fmt.Println(styles.ReplStyle.Render(evaluated.Inspect()))
		}
	}
}

func Start() {
	env := object.NewEnvironment()

	var d dummy
	d.env = env
	p := prompt.New(
		d.executor,
		completer,
		prompt.OptionPrefix(PROMPT),
		prompt.OptionTitle("Vint Programming Language"),
	)

	p.Run()
}

type dummy struct {
	env *object.Environment
}

func (d *dummy) executor(in string) {
	if strings.TrimSpace(in) == "exit()" {
		style := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")). // Gold text
			Background(lipgloss.Color("#282C34")). // Dark background
			Bold(true).
			Padding(1, 2).
			Margin(1).
			Border(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("#FF4500")) // Bright orange border

		message := style.Render("\nThank you for using Vint! Goodbye and happy coding!")
		fmt.Println(message)
		os.Exit(0)
	}
	l := lexer.NewWithFilename(in, "<repl>")
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			fmt.Println("\t" + styles.ErrorStyle.Render(msg))
		}
	}
	env := d.env
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() != object.NULL_OBJ {
			fmt.Println(styles.ReplStyle.Render(evaluated.Inspect()))
		}
	}
}

func completer(in prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func Docs() {
	zone.NewGlobal()

	var p playground

	p.toc = list.New(englishItems, list.NewDefaultDelegate(), 0, 0)
	p.toc.Title = "Table of Contents"
	p.id = zone.NewPrefix()

	if _, err := tea.NewProgram(p, tea.WithMouseAllMotion()).Run(); err != nil {
		log.Fatal(err)
	}
}

var englishItems []list.Item

func init() {
	for _, docItem := range docs.GetDocsItem() {
		englishItems = append(englishItems, docItem)
	}
}
