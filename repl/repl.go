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
	"github.com/ekilie/vint-lang/evaluator"
	"github.com/ekilie/vint-lang/lexer"
	"github.com/ekilie/vint-lang/object"
	"github.com/ekilie/vint-lang/parser"
	"github.com/ekilie/vint-lang/styles"
	zone "github.com/lrstanley/bubblezone"
)

const PROMPT = ">>> "

//go:embed docs
var res embed.FS

func Read(contents string) {
	env := object.NewEnvironment()

	l := lexer.New(contents)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println(styles.ErrorStyle.Italic(false).Render("These errors occured:"))

		for _, msg := range p.Errors() {
			fmt.Println("\t" + styles.ErrorStyle.Render(msg))
		}
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
	l := lexer.New(in)
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

	languageChoice := []list.Item{
		languages{title: "English", desc: "Read documentation in English", dir: "en"},
	}

	var p playground

	p.languageCursor = list.New(languageChoice, list.NewDefaultDelegate(), 50, 8)
	p.languageCursor.Title = "Chagua Lugha"
	p.languageCursor.SetFilteringEnabled(false)
	p.languageCursor.SetShowStatusBar(false)
	p.languageCursor.SetShowPagination(false)
	p.languageCursor.SetShowHelp(false)
	p.toc = list.New(englishItems, list.NewDefaultDelegate(), 0, 0)
	p.toc.Title = "Table of Contents"
	p.id = zone.NewPrefix()

	if _, err := tea.NewProgram(p, tea.WithMouseAllMotion()).Run(); err != nil {
		log.Fatal(err)
	}
}

var (
	englishItems = []list.Item{
		item{title: "Arrays", desc: "🚀 Unleash the power of arrays in vint", filename: "arrays.md"},
		item{title: "Booleans", desc: "👍👎 Master the world of 'if' and 'else' with bools", filename: "bool.md"},
		item{title: "Builtins", desc: "💡 Reveal the secrets of builtin functions in vint", filename: "builtins.md"},
		item{title: "Comments", desc: "💬 Speak your mind with comments in vint", filename: "comments.md"},
		item{title: "Dictionaries", desc: "📚 Unlock the knowledge of dictionaries in vint", filename: "dictionaries.md"},
		item{title: "Files", desc: "💾 Handle files effortlessly in vint", filename: "files.md"},
		item{title: "For", desc: "🔄 Loop like a pro with 'for' in vint", filename: "for.md"},
		item{title: "Function", desc: "🔧 Create powerful functions in vint", filename: "function.md"},
		item{title: "Identifiers", desc: "🔖 Give your variables their own identity in vint", filename: "identifiers.md"},
		item{title: "If Statements", desc: "🔮 Control the flow with 'if' statements in vint", filename: "ifStatements.md"},
		item{title: "JSON", desc: "📄 Master the art of JSON in vint", filename: "json.md"},
		item{title: "Keywords", desc: "🔑 Learn the secret language of vint's keywords", filename: "keywords.md"},
		item{title: "Net", desc: "🌐 Explore the world of networking in vint", filename: "net.md"},
		item{title: "Null", desc: "🌌 Embrace the void with Null in vint", filename: "null.md"},
		item{title: "Numbers", desc: "🔢 Discover the magic of numbers in vint", filename: "numbers.md"},
		item{title: "Operators", desc: "🧙 Perform spells with vint's operators", filename: "operators.md"},
		item{title: "Packages", desc: "📦 Harness the power of packages in vint", filename: "packages.md"},
		item{title: "Strings", desc: "🎼 Compose stories with strings in vint", filename: "strings.md"},
		item{title: "Switch", desc: "🧭 Navigate complex scenarios with 'switch' in vint", filename: "switch.md"},
		item{title: "Time", desc: "⏰ Manage time with ease in vint", filename: "time.md"},
		item{title: "While", desc: "⌛ Learn the art of patience with 'while' loops in vint", filename: "while.md"},
	}
	kiswahiliItems = []list.Item{
		item{title: "Arrays", desc: "🚀 Unleash the power of arrays in vint", filename: "arrays.md"},
		item{title: "Booleans", desc: "👍👎 Master the world of 'if' and 'else' with bools", filename: "bool.md"},
		item{title: "Builtins", desc: "💡 Reveal the secrets of builtin functions in vint", filename: "builtins.md"},
		item{title: "Comments", desc: "💬 Speak your mind with comments in vint", filename: "comments.md"},
		item{title: "Dictionaries", desc: "📚 Unlock the knowledge of dictionaries in vint", filename: "dictionaries.md"},
		item{title: "Files", desc: "💾 Handle files effortlessly in vint", filename: "files.md"},
		item{title: "For", desc: "🔄 Loop like a pro with 'for' in vint", filename: "for.md"},
		item{title: "Function", desc: "🔧 Create powerful functions in vint", filename: "function.md"},
		item{title: "Identifiers", desc: "🔖 Give your variables their own identity in vint", filename: "identifiers.md"},
		item{title: "If Statements", desc: "🔮 Control the flow with 'if' statements in vint", filename: "ifStatements.md"},
		item{title: "JSON", desc: "📄 Master the art of JSON in vint", filename: "json.md"},
		item{title: "Keywords", desc: "🔑 Learn the secret language of vint's keywords", filename: "keywords.md"},
		item{title: "Net", desc: "🌐 Explore the world of networking in vint", filename: "net.md"},
		item{title: "Null", desc: "🌌 Embrace the void with Null in vint", filename: "null.md"},
		item{title: "Numbers", desc: "🔢 Discover the magic of numbers in vint", filename: "numbers.md"},
		item{title: "Operators", desc: "🧙 Perform spells with vint's operators", filename: "operators.md"},
		item{title: "Packages", desc: "📦 Harness the power of packages in vint", filename: "packages.md"},
		item{title: "Strings", desc: "🎼 Compose stories with strings in vint", filename: "strings.md"},
		item{title: "Switch", desc: "🧭 Navigate complex scenarios with 'switch' in vint", filename: "switch.md"},
		item{title: "Time", desc: "⏰ Manage time with ease in vint", filename: "time.md"},
		item{title: "While", desc: "⌛ Learn the art of patience with 'while' loops in vint", filename: "while.md"},
	}
)
