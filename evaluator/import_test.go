package evaluator

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/vintlang/vintlang/object"
)

// TestCircularImportDetectedAcrossChain verifies that circular imports are
// detected through a nested chain: a.vint → b.vint → a.vint.
func TestCircularImportDetectedAcrossChain(t *testing.T) {
	dir := t.TempDir()
	addSearchPath(dir)
	t.Cleanup(func() { searchPaths = nil; importedModules = make(map[string]bool) })

	os.WriteFile(filepath.Join(dir, "a.vint"), []byte(`import b`), 0644)
	os.WriteFile(filepath.Join(dir, "b.vint"), []byte(`import a`), 0644)

	result := testEval(`import a`)
	if result == nil || result.Type() != object.ERROR_OBJ {
		t.Fatalf("expected circular import error, got %v", result)
	}
	if !strings.Contains(result.Inspect(), "Circular import detected") {
		t.Fatalf("expected 'Circular import detected' in error, got: %s", result.Inspect())
	}
}

// TestImportFailureUnmarksModule verifies that a failed import properly
// removes the module from importedModules so a subsequent independent import
// of the same name does not falsely report a circular import.
func TestImportFailureUnmarksModule(t *testing.T) {
	dir := t.TempDir()
	addSearchPath(dir)
	t.Cleanup(func() { searchPaths = nil; importedModules = make(map[string]bool) })

	// Write a file that will cause a runtime error on evaluation.
	os.WriteFile(filepath.Join(dir, "badmod.vint"), []byte(`let x = unknownIdent`), 0644)

	// First import should fail.
	result := testEval(`import badmod`)
	if result == nil || result.Type() != object.ERROR_OBJ {
		t.Fatalf("expected an import error, got %v", result)
	}

	// After failure, importedModules should NOT contain stale entry.
	if importedModules["badmod"] {
		t.Fatal("importedModules still contains 'badmod' after failed import")
	}

	// A second import of the same module should NOT report circular import.
	result2 := testEval(`import badmod`)
	if result2 != nil && result2.Type() == object.ERROR_OBJ {
		if strings.Contains(result2.Inspect(), "Circular import detected") {
			t.Fatalf("false circular import error after prior failure: %s", result2.Inspect())
		}
	}
}

// TestSuccessfulNestedImportChain verifies that importing a chain a → b → c
// works without false circular import errors, and that importedModules is
// clean after the chain completes.
func TestSuccessfulNestedImportChain(t *testing.T) {
	dir := t.TempDir()
	addSearchPath(dir)
	t.Cleanup(func() { searchPaths = nil; importedModules = make(map[string]bool) })

	os.WriteFile(filepath.Join(dir, "c.vint"), []byte(`let val = 1`), 0644)
	os.WriteFile(filepath.Join(dir, "b.vint"), []byte(`import c`), 0644)
	os.WriteFile(filepath.Join(dir, "a.vint"), []byte(`import b`), 0644)

	result := testEval(`import a`)
	if result != nil && result.Type() == object.ERROR_OBJ {
		t.Fatalf("unexpected error in nested import chain: %s", result.Inspect())
	}

	// After the chain completes, no modules should remain marked.
	for mod, marked := range importedModules {
		if marked {
			t.Errorf("importedModules still contains '%s' after successful chain", mod)
		}
	}
}
