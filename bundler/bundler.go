package bund_test

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/vintlang/vintlang/utils"
)

const skeleton = `package main

import (
	"github.com/vintlang/vintlang/repl"
)

func main() {
	code := ` + "`{{.Code}}`" + `
	repl.Read(code)
}
`

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vintbuild <file.vint>")
		os.Exit(1)
	}

	vintFile := os.Args[1]

	data, err := os.ReadFile(vintFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	escaped := strings.ReplaceAll(string(data), "`", "`+\"`\"+`") // handle backticks in Vint code

	f, err := os.Create("compiled.go")
	if err != nil {
		fmt.Println("Error creating compiled.go:", err)
		os.Exit(1)
	}
	defer f.Close()

	tmpl := template.Must(template.New("main").Parse(skeleton))
	err = tmpl.Execute(f, map[string]string{"Code": escaped})
	if err != nil {
		fmt.Println("Template error:", err)
		os.Exit(1)
	}

	// Build binary
	binaryName := strings.TrimSuffix(vintFile, ".vint")
	buildCmd := fmt.Sprintf("go build -o %s compiled.go", binaryName)
	fmt.Println("Building:", buildCmd)
	err = utils.RunShell(buildCmd)
	if err != nil {
		fmt.Println("Build failed:", err)
		os.Exit(1)
	}

	fmt.Println("Built binary:", binaryName)
}
