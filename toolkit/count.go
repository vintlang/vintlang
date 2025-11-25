package toolkit

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Lines() {
	extensions := []string{".go", ".vint", ".h", ".hpp", ".S"}

	totalLines := 0
	fileCount := 0

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if file has one of the target extensions
		hasExtension := false
		for _, ext := range extensions {
			if strings.HasSuffix(path, ext) {
				hasExtension = true
				break
			}
		}

		if !hasExtension {
			return nil
		}

		lines, err := CountLines(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", path, err)
			return nil // Continue walking
		}

		totalLines += lines
		fileCount++

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", totalLines)
	fmt.Fprintf(os.Stderr, "Files processed: %d\n", fileCount)
}

// CountLines counts the number of lines in a given file.
func CountLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0

	for scanner.Scan() {
		lines++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lines, nil
}
