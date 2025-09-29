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
	"github.com/vintlang/vintlang/evaluator"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
	"github.com/vintlang/vintlang/styles"
)

const PROMPT = ">>> "

// go:embed docs
var res embed.FS

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
		item{title: "CSV", desc: "📈 Handle CSV data with ease", filename: "csv.md"},
		item{title: "Dictionaries", desc: "📚 Unlock the knowledge of dictionaries in vint", filename: "dictionaries.md"},
		item{title: "Files", desc: "💾 Handle files effortlessly in vint", filename: "files.md"},
		item{title: "For", desc: "🔄 Loop like a pro with 'for' in vint", filename: "for.md"},
		item{title: "Function", desc: "🔧 Create powerful functions in vint", filename: "function.md"},
		item{title: "Identifiers", desc: "🔖 Give your variables their own identity in vint", filename: "identifiers.md"},
		item{title: "If Statements", desc: "🔮 Control the flow with 'if' statements in vint", filename: "ifStatements.md"},
		item{title: "JSON", desc: "📄 Master the art of JSON in vint", filename: "json.md"},
		item{title: "Keywords", desc: "🔑 Learn the secret language of vint's keywords", filename: "keywords.md"},
		item{title: "MySQL", desc: "🗄️ Interact with MySQL databases", filename: "mysql.md"},
		item{title: "PostgreSQL", desc: "🐘 Work with PostgreSQL databases", filename: "postgres.md"},
		item{title: "Net", desc: "🌐 Explore the world of networking in vint", filename: "net.md"},
		item{title: "Null", desc: "🌌 Embrace the void with Null in vint", filename: "null.md"},
		item{title: "Numbers", desc: "🔢 Discover the magic of numbers in vint", filename: "numbers.md"},
		item{title: "Operators", desc: "🧙 Perform spells with vint's operators", filename: "operators.md"},
		item{title: "Packages", desc: "📦 Harness the power of packages in vint", filename: "packages.md"},
		item{title: "Path", desc: "🛤️ Manipulate file paths like a pro", filename: "path.md"},
		item{title: "Random", desc: "🎲 Generate random numbers and more", filename: "random.md"},
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
