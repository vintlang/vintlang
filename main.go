package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/ekilie/vint-lang/repl"
	"github.com/ekilie/vint-lang/styles"
	"github.com/ekilie/vint-lang/toolkit"
)
const VINT_VERSION = "0.1.2"
// Constants for styled output
var (
	// Title banner for the CLI
	Title = styles.TitleStyle.Render(`
                        ██╗   ██╗██╗███╗   ██╗████████╗
                        ██║   ██║██║████╗  ██║╚══██╔══╝
                        ██║   ██║██║██╔██╗ ██║   ██║   
                        ╚██╗ ██╔╝██║██║╚██╗██║   ██║   
                         ╚████╔╝ ██║██║ ╚████║   ██║   
                          ╚═══╝  ╚═╝╚═╝  ╚═══╝   ╚═╝ 
`)

	// CLI metadata
	Version = styles.VersionStyle.Render("v"+VINT_VERSION)
	Author  = styles.AuthorStyle.Render("by Tachera Sasi")

	// Combined logo with title, author, and version
	NewLogo = lipgloss.JoinVertical(lipgloss.Center, Title,
		lipgloss.JoinHorizontal(lipgloss.Center, Author, " | ", Version))

	// Help message for the CLI usage
	Help = styles.HelpStyle.Italic(false).Render(fmt.Sprintf(`💡 How to use vint:
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
	// Retrieve command-line arguments
	args := os.Args

	// No arguments provided: Start REPL mode
	if len(args) < 2 {
		help := styles.HelpStyle.Render("💡 Use exit() to exit")
		fmt.Println(lipgloss.JoinVertical(lipgloss.Left, NewLogo, "\n", help))
		repl.Start()
		return
	}

	// Single argument provided: Handle specific commands
	if len(args) >= 2 {//Greater or equal accounting for the cli-args and the package manager for vint 
		switch args[1] {
		case "help", "-help", "--help", "-h":
			fmt.Println(Help)
		case "version", "-version", "--version", "-v", "v":
			fmt.Println(NewLogo)
		case "--docs", "-docs":
			repl.Docs()
		case "get":
			toolkit.Get(os.Args[2])
		case ".":
			// Runs default main.vint file
			run("main.vint")
		default:
			// Attempts to run the specified file
			file := args[1]
			run(file)
		}
	} else {
		// Invalid usage: Displays error and help
		fmt.Println(styles.ErrorStyle.Render("Error: Operation failed."))
		fmt.Println(Help)
		os.Exit(1)
	}
}

// runs and executes the specified Vint file
func run(file string) {
	if len(os.Args) > 2 {
		// Appends all arguments after the first two directly to toolkit.CLI_ARGS
		toolkit.CLI_ARGS = append(toolkit.CLI_ARGS, os.Args[2:]...)
	}
	
	// Ensures the file has a .vint extension
	if strings.HasSuffix(file, ".vint") {
		contents, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(styles.ErrorStyle.Render("Error: vint Failed to read the file: ", file))
			os.Exit(1)
		}

		// Passes the file contents to the REPL for execution
		repl.Read(string(contents))
	} else {
		// Handles invalid file type
		fmt.Println(styles.ErrorStyle.Render("'"+file+"'", "is not a correct file type. Use '.vint'"))
		os.Exit(1)
	}
}
