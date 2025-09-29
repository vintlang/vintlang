package main

import (
	"fmt"
	"os"
	"strings"
)

// ANSI color codes for terminal output
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Bold   = "\033[1m"
)

// ColorizedError formats an error message with colors
func ColorizedError(errorType, message, sourceLine string, line, column int) string {
	var builder strings.Builder
	
	// Error header with color
	builder.WriteString(fmt.Sprintf("%s%s[ERROR]%s %sLine %d:%d:%s ", 
		Bold, Red, Reset, Yellow, line, column, Reset))
	
	// Error message
	builder.WriteString(message)
	builder.WriteString("\n")
	
	// Source line with highlighting
	if sourceLine != "" {
		builder.WriteString(fmt.Sprintf("%s    %s%s\n", Blue, sourceLine, Reset))
		if column > 0 {
			builder.WriteString(fmt.Sprintf("%s    %s^%s\n", 
				Red, strings.Repeat(" ", column-1), Reset))
		}
	}
	
	return builder.String()
}

// CheckColorSupport checks if terminal supports colors
func CheckColorSupport() bool {
	term := os.Getenv("TERM")
	return strings.Contains(term, "color") || strings.Contains(term, "256") || term == "xterm"
}

func main() {
	if CheckColorSupport() {
		fmt.Println("✅ Terminal supports colors!")
		
		// Demo colorized error
		demo := ColorizedError(
			"SYNTAX", 
			"Illegal character '?' - single '?' is not a valid operator, did you mean '??'?",
			"let x = 5 ? 10",
			2, 11,
		)
		fmt.Print(demo)
	} else {
		fmt.Println("❌ Terminal does not support colors")
	}
}