package docs

import (
	"embed"
	"io/fs"
	"strings"
)

//go:embed *
var Docs embed.FS

type Item struct {
	title, desc, filename string
}

func (i Item) Title() string       { return i.title }
func (i Item) Filename() string       { return i.filename }
func (i Item) Description() string { return i.desc }
func (i Item) FilterValue() string { return i.title }

func NewDocsItem(title,desc,filename string) Item {
	return Item{
		title:    title,
		desc:     desc,
		filename: filename,
	}
}


func GetDocsItem() []Item{
	f,err :=Docs.ReadDir(".")
	if err != nil {
		return nil
	}
	var items []Item
	for _,file := range f{
		if !file.IsDir() && strings.HasSuffix(file.Name(),".md"){
			items = append(items,Item{
				title:    file.Name(),
				desc:     GetDocDescription(file),
				filename: file.Name(),
			})
		}
	}
	return items
}

func GetDocDescription(file fs.DirEntry) string {
	content, err := Docs.ReadFile(file.Name())
	if err != nil {
		return "No description available."
	}
	lines := string(content)
	firstLine := ""
	for _, line := range SplitLines(lines) {
		trimmed := TrimWhitespace(line)
		if trimmed != "" {
			firstLine = trimmed
			break
		}
	}
	return strings.ReplaceAll(firstLine, "#", "")
}

func SplitLines(s string) []string {
	var lines []string
	currentLine := ""
	for _, r := range s {
		if r == '\n' {
			lines = append(lines, currentLine)
			currentLine = ""
		} else {
			currentLine += string(r)
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}

func TrimWhitespace(s string) string {
	start := 0
	end := len(s) - 1
	for start <= end && (s[start] == ' ' || s[start] == '\t') {
		start++
	}
	for end >= start && (s[end] == ' ' || s[end] == '\t') {
		end--
	}
	if start > end {
		return ""
	}
	return s[start : end+1]
}