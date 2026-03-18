package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/vintlang/vintlang/bundler"
	"github.com/vintlang/vintlang/config"
	"github.com/vintlang/vintlang/evaluator"
	"github.com/vintlang/vintlang/lexer"
	"github.com/vintlang/vintlang/object"
	"github.com/vintlang/vintlang/parser"
	"github.com/vintlang/vintlang/repl"
	"github.com/vintlang/vintlang/styles"
	"github.com/vintlang/vintlang/token"
	"github.com/vintlang/vintlang/toolkit"
)

const VINT_VERSION = config.VINT_VERSION

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
    %s: Bundle a vint file into binary
    %s: Initialize a new vint project
    %s: Install a vint package
    %s: Run tests in current directory
    %s: Format vint code
    %s: Open interactive documentation
    %s: Trace pipeline stages to a txt file
    %s: Show vint version
    %s: Show this help message
`,
		styles.HelpStyle.Bold(true).Render("vint"),
		styles.HelpStyle.Bold(true).Render("vint filename.vint"),
		styles.HelpStyle.Bold(true).Render("vint bundler filename.vint"),
		styles.HelpStyle.Bold(true).Render("vint init"),
		styles.HelpStyle.Bold(true).Render("vint get package"),
		styles.HelpStyle.Bold(true).Render("vint test"),
		styles.HelpStyle.Bold(true).Render("vint fmt filename.vint"),
		styles.HelpStyle.Bold(true).Render("vint docs"),
		styles.HelpStyle.Bold(true).Render("vint --trace filename.vint"),
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
		case "play", "-play", "--play", "-p":
			repl.Playground()
		case "docs", "-docs", "--docs", "-d":
			repl.Docs()
		case "lines":
			toolkit.PrintCodebaseLines()
		case "version", "-version", "--version", "-v", "v":
			fmt.Println(versionMsg)
		case "bundle", "-bundle", "--bundle", "-b", "--b":
			if len(args) < 3 {
				fmt.Println(styles.ErrorStyle.Render("Error: Please specify a Vint file to bundle"))
				os.Exit(1)
			}
			if err := bundler.Bundle(args[2:]); err != nil {
				fmt.Println(styles.ErrorStyle.Render(fmt.Sprintf("Build failed: %v", err)))
				os.Exit(1)
			}
			fmt.Println(styles.HelpStyle.Render("Build successful!"))
		case "bundle-multi":
			if len(args) < 7 {
				fmt.Println(styles.ErrorStyle.Render("Error: usage: vint bundle-multi <file> \"\" <name> <dir> <targets> [quiet]"))
				os.Exit(1)
			}
			if err := bundler.BundleMulti(args[2:]); err != nil {
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
		case "new":
			toolkit.New(args)
		case "fmt", "-fmt", "--fmt", "-f":
			if len(args) < 3 {
				fmt.Println(styles.ErrorStyle.Render("Error: Please specify a Vint file to format"))
				os.Exit(1)
			}
			formatFile(args[2])
		case "trace", "-trace", "--trace":
			if len(args) < 3 {
				fmt.Println(styles.ErrorStyle.Render("Error: Please specify a Vint file to trace"))
				os.Exit(1)
			}
			runWithTrace(args[2])
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

		// Add the script's directory to the import search paths so that
		// multi-file projects with packages can be run from any working directory.
		scriptDir := filepath.Dir(file)
		if absDir, err := filepath.Abs(scriptDir); err == nil {
			evaluator.AddSearchPath(absDir)
		}

		// Passes the file contents to the REPL for execution
		repl.ReadWithFilename(string(contents), file)
	} else {
		// Handles invalid file type
		fmt.Println(styles.ErrorStyle.Render("'"+file+"'", "is not a correct file type. Use '.vint'"))
		os.Exit(1)
	}
}

// formatFile formats a Vint source file
func formatFile(file string) {
	if !strings.HasSuffix(file, ".vint") {
		fmt.Println(styles.ErrorStyle.Render("Error: Can only format .vint files"))
		os.Exit(1)
	}

	contents, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error: Failed to read file:", file))
		os.Exit(1)
	}

	// Parse the file using the lexer and parser
	l := lexer.NewWithFilename(string(contents), file)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Println(styles.ErrorStyle.Render("Error: Failed to parse file. Cannot format."))
		for _, msg := range p.Errors() {
			fmt.Println(styles.ErrorStyle.Render(msg))
		}
		os.Exit(1)
	}

	formatted := program.String()

	err = os.WriteFile(file, []byte(formatted), 0644)
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error: Failed to write formatted code to file."))
		os.Exit(1)
	}

	fmt.Println(styles.HelpStyle.Render("Formatted", file))
}

// runWithTrace executes a Vint file and writes the output of every pipeline
// stage (source → lexer → parser → evaluator) into a trace txt file.
func runWithTrace(file string) {
	if !strings.HasSuffix(file, ".vint") {
		fmt.Println(styles.ErrorStyle.Render("'" + file + "' is not a correct file type. Use '.vint'"))
		os.Exit(1)
	}

	contents, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error: vint Failed to read the file: ", file))
		os.Exit(1)
	}

	source := string(contents)
	var trace strings.Builder

	// Header
	trace.WriteString("=== VINT TRACE OUTPUT ===\n")
	trace.WriteString(fmt.Sprintf("File: %s\n", file))
	trace.WriteString(fmt.Sprintf("Date: %s\n", time.Now().Format(time.RFC3339)))
	trace.WriteString("\n")

	// Stage 1: Source Code
	trace.WriteString("=== STAGE 1: SOURCE CODE ===\n")
	trace.WriteString(source)
	if !strings.HasSuffix(source, "\n") {
		trace.WriteString("\n")
	}
	trace.WriteString("\n")

	// Stage 2: Lexer (Tokenization)
	trace.WriteString("=== STAGE 2: LEXER (Tokenization) ===\n")
	l := lexer.NewWithFilename(source, file)
	var tokens []token.Token
	for i := 1; ; i++ {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		trace.WriteString(fmt.Sprintf("  Token[%d]: Type=%-14s Literal=%-20q Line=%d Col=%d\n",
			i, tok.Type, tok.Literal, tok.Line, tok.Column))
		if tok.Type == token.EOF {
			break
		}
	}

	lexerErrors := l.Errors()
	if len(lexerErrors) > 0 {
		trace.WriteString(fmt.Sprintf("\nLexer Errors (%d):\n", len(lexerErrors)))
		for _, msg := range lexerErrors {
			trace.WriteString(fmt.Sprintf("  - %s\n", msg))
		}
	} else {
		trace.WriteString(fmt.Sprintf("\nTokens: %d (including EOF)\n", len(tokens)))
		trace.WriteString("Lexer Errors: None\n")
	}
	trace.WriteString("\n")

	// Stage 3: Parser (AST)
	trace.WriteString("=== STAGE 3: PARSER (Abstract Syntax Tree) ===\n")
	l2 := lexer.NewWithFilename(source, file)
	p := parser.New(l2)
	program := p.ParseProgram()

	parserErrors := p.Errors()
	if len(parserErrors) > 0 {
		trace.WriteString(fmt.Sprintf("Parser Errors (%d):\n", len(parserErrors)))
		for _, msg := range parserErrors {
			trace.WriteString(fmt.Sprintf("  - %s\n", msg))
		}
	} else {
		trace.WriteString("Parser Errors: None\n")
	}

	trace.WriteString(fmt.Sprintf("Statements: %d\n", len(program.Statements)))
	trace.WriteString("\nAST:\n")
	trace.WriteString(program.String())
	if astStr := program.String(); !strings.HasSuffix(astStr, "\n") {
		trace.WriteString("\n")
	}
	trace.WriteString("\n")

	// Stage 4: Evaluation
	trace.WriteString("=== STAGE 4: EVALUATION ===\n")

	if len(parserErrors) > 0 {
		trace.WriteString("Skipped: Parser errors prevent evaluation.\n")
	} else {
		scriptDir := filepath.Dir(file)
		if absDir, err := filepath.Abs(scriptDir); err == nil {
			evaluator.AddSearchPath(absDir)
		}

		env := object.NewEnvironment()
		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			trace.WriteString(fmt.Sprintf("Result Type: %s\n", evaluated.Type()))
			if errObj, ok := evaluated.(*object.Error); ok {
				trace.WriteString(fmt.Sprintf("Error: %s\n", errObj.Message))
			} else if evaluated.Type() != object.NULL_OBJ {
				trace.WriteString(fmt.Sprintf("Result Value: %s\n", evaluated.Inspect()))
			} else {
				trace.WriteString("Result Value: null\n")
			}
		} else {
			trace.WriteString("Result: <nil>\n")
		}
	}

	trace.WriteString("\n=== END OF TRACE ===\n")

	// Determine output file name
	outputFile := "vint_trace.txt"
	if len(os.Args) > 3 {
		outputFile = os.Args[3]
	}

	err = os.WriteFile(outputFile, []byte(trace.String()), 0644)
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error: Failed to write trace file: " + err.Error()))
		os.Exit(1)
	}

	fmt.Println(styles.HelpStyle.Render("Trace written to " + outputFile))
}
