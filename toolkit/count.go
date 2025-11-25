package toolkit

import (
	"embed"
	"fmt"
)

//go:embed count.txt
var CountFile embed.FS

func PrintCodebaseLines() {
	line := GetCountFile()
	fmt.Println("TOTAL LINES OF CODE:", line)
}

func GetCountFile() string {
	content, err := CountFile.ReadFile("count.txt")
	if err != nil {
		return ""
	}
	return string(content)
}

// Lines counts and prints the total number of lines in files with specific extensions in the current directory and its subdirectories.
