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

// logError logs errors to a file for debugging
func logError(err error) {
	f, ferr := os.OpenFile("bundler.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if ferr != nil {
		fmt.Fprintf(os.Stderr, "Failed to open log file: %v\n", ferr)
		return
	}
	defer f.Close()
	timestamp := time.Now().Format(time.RFC3339)
	f.WriteString(fmt.Sprintf("[%s] %v\n", timestamp, err))
}

// Bundles the vintlang code to a go binary
func Bundle(args []string) error {
	if len(args) == 0 {
		err := fmt.Errorf("no input file provided to bundler")
		logError(err)
		return err
	}
	vintFile := args[0]
	fmt.Println(len(vintFile))

	// if *name != "" {
	// 	fmt.Printf("🔧 Custom binary name set to '%s'\n", *name)
	// }
	fmt.Printf("📦 Starting Bundle for '%s'\n", filepath.Base(vintFile))

	fmt.Print("🔍 Reading source file... ")
	data, err := os.ReadFile(vintFile)
	if err != nil {
		err = fmt.Errorf("failed to read file '%s': %w", vintFile, err)
		logError(err)
		return err
	}
	fmt.Println("✅")

	fmt.Print("📁 Creating temp Bundle directory... ")
	tempDir, err := os.MkdirTemp("", "vint-bundle-*")
	if err != nil {
		err = fmt.Errorf("failed to create temp dir: %w", err)
		logError(err)
		return err
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
		err = fmt.Errorf("failed to create main.go in temp dir '%s': %w", tempDir, err)
		logError(err)
		return err
	}
	defer mainFile.Close()

	t := template.Must(template.New("main").Parse(goTemplate))
	if err := t.Execute(mainFile, map[string]string{"Code": escapedCode}); err != nil {
		err = fmt.Errorf("failed to execute template for main.go: %w", err)
		logError(err)
		return err
	}
	fmt.Println("✅")

	fmt.Print("📦 Initializing modules... ")
	goMod := `module vint-bundled

go 1.24

require github.com/vintlang/vintlang v0.2.0
`
	if err := os.WriteFile(filepath.Join(tempDir, "go.mod"), []byte(goMod), 0644); err != nil {
		err = fmt.Errorf("failed to create go.mod in temp dir '%s': %w", tempDir, err)
		logError(err)
		return err
	}
	fmt.Println("✅")

	// Bundle binary
	binaryName := strings.TrimSuffix(filepath.Base(vintFile), ".vint")

	fmt.Println(args)
	if len(args) == 3 {
		binaryName = args[2]
	}
	fmt.Printf("🔨 Bundling binary '%s'...\n", binaryName)

	spinner := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	done := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			select {
			case <-done:
				return
			default:
				fmt.Printf("\r%s Bundling...", spinner[i%len(spinner)])
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	BundleCmd := fmt.Sprintf("cd %s && go mod tidy && go build -o %s", tempDir, binaryName)
	if err := utils.RunShell(BundleCmd); err != nil {
		err = fmt.Errorf("bundle command failed: %w", err)
		logError(err)
		done <- true
		return fmt.Errorf("\n❌ Bundle failed: %w", err)
	}

	done <- true
	fmt.Printf("\r🎉 Bundle successful! Moving binary... ")

	finalBinary := filepath.Join(tempDir, binaryName)
	if err := os.Rename(finalBinary, filepath.Join(".", binaryName)); err != nil {
		err = fmt.Errorf("failed to move binary from '%s' to current directory: %w", finalBinary, err)
		logError(err)
		return fmt.Errorf("\n❌ Failed to move binary: %w", err)
	}

	fmt.Println("✅")
	fmt.Printf("\n✨ Successfully created binary: ./%s\n", binaryName)
	return nil
}
