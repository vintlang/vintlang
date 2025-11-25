package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {  // Ignore. linter error
	extensions := []string{".go", ".vint"}

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

	counterTXTFile, err := os.OpenFile("toolkit/count.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening count.txt: %v\n", err)
		os.Exit(1)
	}
	defer counterTXTFile.Close()
	_, err = counterTXTFile.WriteString(fmt.Sprintf("%d\n", totalLines))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to count.txt: %v\n", err)
		os.Exit(1)
	}
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