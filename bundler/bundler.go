package bundler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/vintlang/vintlang/utils"
)

//Bundles the vintlang code to a go binary
func Bundle(vintFile string) error {
	fmt.Printf("📦 Starting build for '%s'\n", filepath.Base(vintFile))

	fmt.Print("🔍 Reading source file... ")
	data, err := os.ReadFile(vintFile)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	fmt.Println("✅")

	fmt.Print("📁 Creating temp build directory... ")
	tempDir, err := os.MkdirTemp("", "vint-bundle-*")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)
	fmt.Println("✅")

	escapedCode := strings.ReplaceAll(string(data), "`", "` + \"`\" + `")

	const goTemplate = `package main

import (
	"github.com/vintlang/vintlang/repl"
)

func main() {
	code := ` + "`{{.Code}}`" + `
	repl.Read(code)
}
`

	fmt.Print("⚙️  Generating Go code... ")
	mainPath := filepath.Join(tempDir, "main.go")
	mainFile, err := os.Create(mainPath)
	if err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}
	defer mainFile.Close()

	t := template.Must(template.New("main").Parse(goTemplate))
	if err := t.Execute(mainFile, map[string]string{"Code": escapedCode}); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	fmt.Println("✅")

	fmt.Print("📦 Initializing modules... ")
	goMod := `module vint-bundled

go 1.24

require github.com/vintlang/vintlang v0.2.0
`
	if err := os.WriteFile(filepath.Join(tempDir, "go.mod"), []byte(goMod), 0644); err != nil {
		return fmt.Errorf("failed to create go.mod: %w", err)
	}
	fmt.Println("✅")

	// Build binary
	binaryName := strings.TrimSuffix(filepath.Base(vintFile), ".vint")
	fmt.Printf("🔨 Building binary '%s'...\n", binaryName)

	spinner := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	done := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			select {
			case <-done:
				return
			default:
				fmt.Printf("\r%s Building...", spinner[i%len(spinner)])
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	buildCmd := fmt.Sprintf("cd %s && go mod tidy && go build -o %s", tempDir, binaryName)
	if err := utils.RunShell(buildCmd); err != nil {
		done <- true
		return fmt.Errorf("\n❌ Build failed: %w", err)
	}

	done <- true
	fmt.Printf("\r🎉 Build successful! Moving binary... ")

	finalBinary := filepath.Join(tempDir, binaryName)
	if err := os.Rename(finalBinary, filepath.Join(".", binaryName)); err != nil {
		return fmt.Errorf("\n❌ Failed to move binary: %w", err)
	}

	fmt.Println("✅")
	fmt.Printf("\n✨ Successfully created binary: ./%s\n", binaryName)
	return nil
}
