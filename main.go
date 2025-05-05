package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/vintlang/vintlang/bundler"
	"github.com/vintlang/vintlang/repl"
	"github.com/vintlang/vintlang/styles"
	"github.com/vintlang/vintlang/toolkit"
)

const VINT_VERSION = "0.1.8"

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
	Version = styles.VersionStyle.Render("v" + VINT_VERSION)
	Author  = styles.AuthorStyle.Render("Tachera W")

	NewLogo = lipgloss.JoinVertical(lipgloss.Center, Title,
		lipgloss.JoinHorizontal(lipgloss.Center, Author, " | ", Version))

	// Help message for the CLI usage
	Help = styles.HelpStyle.Italic(false).Render(fmt.Sprintf(`💡 How to use vint:
    %s: Start the vint program
    %s: Run a vint file
    %s: Build a vint file into binary
    %s: Initialize a new vint project
    %s: Install a vint package
    %s: Run tests in current directory
    %s: Format vint code
    %s: Show vint version
    %s: Show this help message
`,
		styles.HelpStyle.Bold(true).Render("vint"),
		styles.HelpStyle.Bold(true).Render("vint filename.vint"),
		styles.HelpStyle.Bold(true).Render("vint build filename.vint"),
		styles.HelpStyle.Bold(true).Render("vint init"),
		styles.HelpStyle.Bold(true).Render("vint get package"),
		styles.HelpStyle.Bold(true).Render("vint test"),
		styles.HelpStyle.Bold(true).Render("vint fmt filename.vint"),
		styles.HelpStyle.Bold(true).Render("vint version"),
		styles.HelpStyle.Bold(true).Render("vint help")))
)

func main() {
	versionMsg := lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Center, "VintLang", " : ", Version))

	args := os.Args

	if len(args) < 2 {
		help := styles.HelpStyle.Render("💡 Use exit() to exit")
		fmt.Println(lipgloss.JoinVertical(lipgloss.Left, NewLogo, "\n", help))
		repl.Start()
		return
	}

	if len(args) >= 2 {
		switch args[1] {
		case "help", "-help", "--help", "-h":
			fmt.Println(Help)
		case "version", "-version", "--version", "-v", "v":
			fmt.Println(versionMsg)
		case "build", "-build", "--build", "-b":
			if len(args) < 3 {
				fmt.Println(styles.ErrorStyle.Render("Error: Please specify a Vint file to build"))
				os.Exit(1)
			}
			if err := bundler.Bundle(args[2]); err != nil {
				fmt.Println(styles.ErrorStyle.Render(fmt.Sprintf("Build failed: %v", err)))
				os.Exit(1)
			}
			fmt.Println(styles.HelpStyle.Render("Build successful!"))
		case "get":
			if len(args) < 3 {
				fmt.Println(styles.ErrorStyle.Render("Error: Please specify a package to install"))
				os.Exit(1)
			}
			toolkit.Get(args[2])
		case "init":
			toolkit.Init(args)
		case "test", "-test", "--test", "-t":
			runTests()
		case "fmt", "-fmt", "--fmt", "-f":
			if len(args) < 3 {
				fmt.Println(styles.ErrorStyle.Render("Error: Please specify a Vint file to format"))
				os.Exit(1)
			}
			formatFile(args[2])
		case ".":
			run("main.vint")
		default:
			file := args[1]
			run(file)
		}
	} else {
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

// runTests executes all tests in the current directory
func runTests() {
	fmt.Println(styles.HelpStyle.Render("Running tests..."))
	cmd := exec.Command("go", "test", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(styles.ErrorStyle.Render("Tests failed"))
		os.Exit(1)
	}
	fmt.Println(styles.HelpStyle.Render("All tests passed!"))
}

// formatFile formats a Vint source file
func formatFile(file string) {
	if !strings.HasSuffix(file, ".vint") {
		fmt.Println(styles.ErrorStyle.Render("Error: Can only format .vint files"))
		os.Exit(1)
	}

	_, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error: Failed to read file:", file))
		os.Exit(1)
	}

	// TODO: Implement actual formatting logic
	// For now, just show a placeholder message
	fmt.Println(styles.HelpStyle.Render("Formatting", file))
	fmt.Println(styles.HelpStyle.Render("Note: Formatting is not yet implemented"))
}
