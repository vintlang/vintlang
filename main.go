package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/NuruProgramming/Nuru/repl"
	"github.com/NuruProgramming/Nuru/styles"
	"github.com/charmbracelet/lipgloss"
)

var (
	Title = styles.TitleStyle.
		Render(`
█░░ █░█ █▀▀ █░█ ▄▀█   █▄█ ▄▀█   █▄░█ █░█ █▀█ █░█
█▄▄ █▄█ █▄█ █▀█ █▀█   ░█░ █▀█   █░▀█ █▄█ █▀▄ █▄█`)
	Version = styles.VersionStyle.Render("v0.5.17")
	Author  = styles.AuthorStyle.Render("by vint Org")
	NewLogo = lipgloss.JoinVertical(lipgloss.Center, Title, lipgloss.JoinHorizontal(lipgloss.Center, Author, " | ", Version))
	Help    = styles.HelpStyle.Italic(false).Render(fmt.Sprintf(`💡 Namna ya kutumia vint:
	%s: Kuanza programu ya vint
	%s: Kuendesha faili la vint
	%s: Kusoma nyaraka za vint
	%s: Kufahamu toleo la vint
`,
		styles.HelpStyle.Bold(true).Render("vint"),
		styles.HelpStyle.Bold(true).Render("vint jinaLaFile.nr"),
		styles.HelpStyle.Bold(true).Render("vint --nyaraka"),
		styles.HelpStyle.Bold(true).Render("vint --toleo")))
)

func main() {

	args := os.Args
	if len(args) < 2 {

		help := styles.HelpStyle.Render("💡 Tumia exit() au toka() kuondoka")
		fmt.Println(lipgloss.JoinVertical(lipgloss.Left, NewLogo, "\n", help))
		repl.Start()
		return
	}

	if len(args) == 2 {
		switch args[1] {
		case "msaada", "-msaada", "--msaada", "help", "-help", "--help", "-h":
			fmt.Println(Help)
		case "version", "-version", "--version", "-v", "v", "--toleo", "-toleo":
			fmt.Println(NewLogo)
		case "-docs", "--docs", "-nyaraka", "--nyaraka":
			repl.Docs()
		default:
			file := args[1]

			if strings.HasSuffix(file, ".vint") {
				contents, err := os.ReadFile(file)
				if err != nil {
					fmt.Println(styles.ErrorStyle.Render("Error: vint imeshindwa kusoma faili: ", args[1]))
					os.Exit(1)
				}

				repl.Read(string(contents))
			} else {
				fmt.Println(styles.ErrorStyle.Render("'"+file+"'", "sii faili sahihi. Tumia faili la '.vint' "))
				os.Exit(1)
			}
		}
	} else {
		fmt.Println(styles.ErrorStyle.Render("Error: Operesheni imeshindikana boss."))
		fmt.Println(Help)
		os.Exit(1)
	}
}
