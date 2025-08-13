package bundler

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDependencyAnalyzer(t *testing.T) {
	// Create temporary test files
	tempDir, err := os.MkdirTemp("", "bundler-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create main.vint
	mainContent := `import math_utils
import string_utils

print("Using math_utils.add:", math_utils.add(2, 3))
print("Using string_utils.concat:", string_utils.concat("Hello", "World"))
`
	mainFile := filepath.Join(tempDir, "main.vint")
	if err := os.WriteFile(mainFile, []byte(mainContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create math_utils.vint
	mathContent := `package math_utils {
    let add = func(a, b) {
        return a + b
    }
    
    let multiply = func(a, b) {
        return a * b
    }
}`
	mathFile := filepath.Join(tempDir, "math_utils.vint")
	if err := os.WriteFile(mathFile, []byte(mathContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create string_utils.vint
	stringContent := `package string_utils {
    let concat = func(a, b) {
        return a + " " + b
    }
}`
	stringFile := filepath.Join(tempDir, "string_utils.vint")
	if err := os.WriteFile(stringFile, []byte(stringContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Test dependency analysis
	analyzer := NewDependencyAnalyzer()
	bundle, err := analyzer.AnalyzeDependencies(mainFile)
	if err != nil {
		t.Fatalf("Failed to analyze dependencies: %v", err)
	}

	// Should find 3 files: main + 2 dependencies
	if len(bundle.Files) != 3 {
		t.Errorf("Expected 3 files, got %d", len(bundle.Files))
	}

	// Check that all expected files are present
	expectedFiles := map[string]bool{
		mainFile:   false,
		mathFile:   false,
		stringFile: false,
	}
	
	for filename := range bundle.Files {
		if _, exists := expectedFiles[filename]; exists {
			expectedFiles[filename] = true
		}
	}
	
	for filename, found := range expectedFiles {
		if !found {
			t.Errorf("Expected file %s not found in bundle", filename)
		}
	}
}

func TestStringProcessor(t *testing.T) {
	// Create a test bundle
	bundle := &FileBundle{
		MainFile: "/test/main.vint",
		Files: map[string]string{
			"/test/main.vint": `import utils
import time

print("Main starting")
let result = utils.helper("test")
print("Result:", result)`,
			"/test/utils.vint": `import json

package utils {
    let helper = func(input) {
        return "processed: " + input
    }
}`,
		},
	}

	processor := NewStringProcessor(bundle)
	processedCode, err := processor.ProcessBundle()
	if err != nil {
		t.Fatalf("Failed to process bundle: %v", err)
	}

	// Check that the processed code contains the package definition
	if !strings.Contains(processedCode, "package utils {") {
		t.Error("Processed code should contain package definition")
	}

	// Check that the import statement for 'utils' was removed from main
	lines := strings.Split(processedCode, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "import utils" {
			t.Error("Import statement for 'utils' should have been removed")
		}
	}

	// Check that other import statements are preserved
	if !strings.Contains(processedCode, "import time") {
		t.Error("Other import statements should be preserved")
	}
}

func TestBundledEvaluator(t *testing.T) {
	bundle := &FileBundle{
		MainFile: "/test/main.vint",
		Files: map[string]string{
			"/test/main.vint": `import helper
print("Testing bundled evaluator")
helper.test()`,
			"/test/helper.vint": `package helper {
    let test = func() {
        print("Helper function called")
    }
}`,
		},
	}

	evaluator := NewBundledEvaluator(bundle)
	goCode, err := evaluator.GenerateBundledCode("v0.1.0", "2023-01-01T00:00:00Z")
	if err != nil {
		t.Fatalf("Failed to generate Go code: %v", err)
	}

	// Check that the generated code contains expected elements
	if !strings.Contains(goCode, "package main") {
		t.Error("Generated code should contain 'package main'")
	}

	if !strings.Contains(goCode, "var BundlerVersion = \"v0.1.0\"") {
		t.Error("Generated code should contain bundler version")
	}

	if !strings.Contains(goCode, "package helper {") {
		t.Error("Generated code should contain the helper package")
	}

	if !strings.Contains(goCode, "repl.Read(processedCode)") {
		t.Error("Generated code should call repl.Read")
	}
}

func TestDependencyAnalyzerWithIncludes(t *testing.T) {
	// Create temporary test files
	tempDir, err := os.MkdirTemp("", "bundler-include-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create main.vint with include
	mainContent := `include "helper.vint"

print("Main file executing")
print("Message:", message)
`
	mainFile := filepath.Join(tempDir, "main.vint")
	if err := os.WriteFile(mainFile, []byte(mainContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create helper.vint (to be included)
	helperContent := `let message = "Hello from included file!"
print("Helper file loaded")
`
	helperFile := filepath.Join(tempDir, "helper.vint")
	if err := os.WriteFile(helperFile, []byte(helperContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Test dependency analysis
	analyzer := NewDependencyAnalyzer()
	bundle, err := analyzer.AnalyzeDependencies(mainFile)
	if err != nil {
		t.Fatalf("Failed to analyze dependencies: %v", err)
	}

	// Should find 2 files: main + included file
	if len(bundle.Files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(bundle.Files))
	}

	// Check that the helper file is marked as included
	if !bundle.IncludeFiles[helperFile] {
		t.Errorf("Helper file should be marked as included")
	}

	// Check that all expected files are present
	expectedFiles := map[string]bool{
		mainFile:   false,
		helperFile: false,
	}
	
	for filename := range bundle.Files {
		if _, exists := expectedFiles[filename]; exists {
			expectedFiles[filename] = true
		}
	}
	
	for filename, found := range expectedFiles {
		if !found {
			t.Errorf("Expected file %s not found in bundle", filename)
		}
	}
}

func TestStringProcessorWithIncludes(t *testing.T) {
	// Create a test bundle with both imports and includes
	bundle := &FileBundle{
		MainFile: "/test/main.vint",
		Files: map[string]string{
			"/test/main.vint": `import utils
include "helper.vint"

print("Main starting")
let result = utils.helper("test")
print("Result:", result)
print("Included message:", message)`,
			"/test/utils.vint": `package utils {
    let helper = func(input) {
        return "processed: " + input
    }
}`,
			"/test/helper.vint": `let message = "Hello from include!"
print("Helper included")`,
		},
		IncludeFiles: map[string]bool{
			"/test/helper.vint": true,
		},
	}

	processor := NewStringProcessor(bundle)
	processedCode, err := processor.ProcessBundle()
	if err != nil {
		t.Fatalf("Failed to process bundle: %v", err)
	}

	// Check that the processed code contains the package definition
	if !strings.Contains(processedCode, "package utils {") {
		t.Error("Processed code should contain package definition")
	}

	// Check that the included content is embedded directly (not in a package)
	if !strings.Contains(processedCode, `let message = "Hello from include!"`) {
		t.Error("Processed code should contain included content")
	}

	// Check that the import statement for 'utils' was removed from main
	lines := strings.Split(processedCode, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "import utils" {
			t.Error("Import statement for 'utils' should have been removed")
		}
		if strings.Contains(trimmed, `include "helper.vint"`) {
			t.Error("Include statement should have been removed")
		}
	}
}