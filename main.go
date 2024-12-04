package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/ekilie/vint-lang/repl"
	"github.com/ekilie/vint-lang/styles"
)

var (
	Title = styles.TitleStyle.
		Render(`
                        ██╗   ██╗██╗███╗   ██╗████████╗
                        ██║   ██║██║████╗  ██║╚══██╔══╝
                        ██║   ██║██║██╔██╗ ██║   ██║   
                        ╚██╗ ██╔╝██║██║╚██╗██║   ██║   
                         ╚████╔╝ ██║██║ ╚████║   ██║   
                          ╚═══╝  ╚═╝╚═╝  ╚═══╝   ╚═╝ 
`)
	Version = styles.VersionStyle.Render("v0.1.0")
	Author  = styles.AuthorStyle.Render("by ekilie")
	NewLogo = lipgloss.JoinVertical(lipgloss.Center, Title, lipgloss.JoinHorizontal(lipgloss.Center, Author, " | ", Version))
	Help    = styles.HelpStyle.Italic(false).Render(fmt.Sprintf(`💡 How to use vint:
    %s: Start the vint program
    %s: Run a vint file
    %s: Read vint documentation
    %s: Know vint version
`,
		styles.HelpStyle.Bold(true).Render("vint"),
		styles.HelpStyle.Bold(true).Render("vint filename.vint"),
		styles.HelpStyle.Bold(true).Render("vint --docs"),
		styles.HelpStyle.Bold(true).Render("vint --version")))
)

func main() {

	args := os.Args
	if len(args) < 2 {

		help := styles.HelpStyle.Render("💡 Use exit() to exit")
		fmt.Println(lipgloss.JoinVertical(lipgloss.Left, NewLogo, "\n", help))
		repl.Start()
		return
	}

	if len(args) == 2 {
		switch args[1] {
		case "help", "-help", "--help", "-h":
			fmt.Println(Help)
		case "version", "-version", "--version", "-v", "v":
			fmt.Println(NewLogo)
		case "--docs", "-docs":
			repl.Docs()
		case ".":
			run("main.vint") //running the main.vint file if: 'vint .'
		default:
			file := args[1]

			//Running the vint file
			run(file)
		}
	} else {
		fmt.Println(styles.ErrorStyle.Render("Error: Operation failed."))
		fmt.Println(Help)
		os.Exit(1)
	}
}

func run(file string) {

	if strings.HasSuffix(file, ".vint") {
		contents, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(styles.ErrorStyle.Render("Error: vint Failed to read the file: ", file))
			os.Exit(1)
		}

		repl.Read(string(contents))

	} else {
		fmt.Println(styles.ErrorStyle.Render("'"+file+"'", "is not a correct file type. Use '.vint'"))
		os.Exit(1)
	}

}
