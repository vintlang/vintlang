package bundler

import (
	"fmt"
	"strings"
)

// BundledEvaluator provides an evaluator that works with bundled files
type BundledEvaluator struct {
	bundle *FileBundle
}

// NewBundledEvaluator creates a new bundled evaluator
func NewBundledEvaluator(bundle *FileBundle) *BundledEvaluator {
	return &BundledEvaluator{
		bundle: bundle,
	}
}

// GenerateBundledCode generates Go code that includes all bundled files
func (be *BundledEvaluator) GenerateBundledCode(bundlerVersion, buildTime string) (string, error) {
	if be.bundle == nil || len(be.bundle.Files) == 0 {
		return "", fmt.Errorf("no files in bundle")
	}

	// Process the bundle to extract packages and modify imports
	processor := NewPackageProcessor(be.bundle)
	processed, err := processor.ProcessBundle()
	if err != nil {
		return "", fmt.Errorf("failed to process bundle: %w", err)
	}

	// Escape the processed content
	escapedProcessedContent := strings.ReplaceAll(processed.ProcessedCode, "`", "` + \"`\" + `")

	// Generate the Go code template
	goTemplate := fmt.Sprintf(`package main

import (
	"flag"
	"fmt"

	"github.com/vintlang/vintlang/repl"
)

var BundlerVersion = "%s"
var BuildTime = "%s"

func main() {
	bundledDetails := flag.Bool("i", false, "Show the bundle details of the app")
	flag.Parse()
	
	if *bundledDetails {
		fmt.Printf("[Bundler Version: %%s | Build Time: %%s]\n", BundlerVersion, BuildTime)
		return
	}
	
	// Processed code with embedded packages and modified imports
	processedCode := `+"`%s`"+`
	repl.Read(processedCode)
}
`, bundlerVersion, buildTime, escapedProcessedContent)

	return goTemplate, nil
}