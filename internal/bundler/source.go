package bundler

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// vintlangSourceRoot locates the vintlang module source tree at runtime.
// The bundler runs inside the vintlang binary, so the compiled-in path of
// this file points back at the source tree (a checkout or the module cache).
// Set VINTLANG_SOURCE to override the detection.
func vintlangSourceRoot() (string, error) {
	if env := os.Getenv("VINTLANG_SOURCE"); env != "" {
		if _, err := os.Stat(filepath.Join(env, "internal")); err == nil {
			return env, nil
		}
		return "", fmt.Errorf("VINTLANG_SOURCE=%s does not contain a vintlang source tree", env)
	}
	_, file, _, ok := runtime.Caller(0)
	if ok {
		root := filepath.Dir(filepath.Dir(filepath.Dir(file)))
		if _, err := os.Stat(filepath.Join(root, "internal")); err == nil {
			return root, nil
		}
	}
	return "", fmt.Errorf("could not locate the vintlang source tree; run the bundler from a vintlang checkout or set VINTLANG_SOURCE")
}

// copyRuntime copies the vintlang interpreter (internal/) into dst so the
// bundled module can import the internal packages directly instead of
// depending on the published module, which cannot be imported due to Go's
// internal package visibility rules and lags behind the current source.
func copyRuntime(dst string) error {
	root, err := vintlangSourceRoot()
	if err != nil {
		return err
	}
	src := filepath.Join(root, "internal")
	if err := copyDir(src, filepath.Join(dst, "internal")); err != nil {
		return fmt.Errorf("failed to copy interpreter runtime: %w", err)
	}
	return nil
}

// copyDir recursively copies src into dst, skipping Go test files.
func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		if strings.HasSuffix(info.Name(), "_test.go") {
			return nil
		}
		return copyFile(path, target)
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}
