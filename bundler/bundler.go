package bundler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/vintlang/vintlang/utils"
)

// Bundle compiles a Vint file into a standalone Go binary
func Bundle(vintFile string) error {
	// Read Vint source code
	data, err := os.ReadFile(vintFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Create Go template for the bundled code
	const goTemplate = `package main

import (
	"github.com/vintlang/vintlang/repl"
)

func main() {
	code := ` + "`{{.Code}}`" + `
	repl.Read(code)
}
`

	// Escape backticks in Vint code to prevent template parsing issues
	escaped := strings.ReplaceAll(string(data), "`", "`+\"`\"+`")

	// Create temporary Go file
	goFile := "bundled.go"
	f, err := os.Create(goFile)
	if err != nil {
		return fmt.Errorf("failed to create Go file: %w", err)
	}
	defer f.Close()

	// Generate Go code from template
	tmpl := template.Must(template.New("main").Parse(goTemplate))
	if err := tmpl.Execute(f, map[string]string{"Code": escaped}); err != nil {
		return fmt.Errorf("failed to generate Go code: %w", err)
	}

	// Build binary
	binaryName := strings.TrimSuffix(filepath.Base(vintFile), ".vint")
	buildCmd := fmt.Sprintf("go build -o %s %s", binaryName, goFile)
	if err := utils.RunShell(buildCmd); err != nil {
		return fmt.Errorf("failed to build binary: %w", err)
	}

	// Clean up temporary Go file
	os.Remove(goFile)

	return nil
}
