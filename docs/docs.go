package docs

import (
	"embed"
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


func GetDocsItem() []Item{
	f,err :=Docs.ReadDir(".")
	if err != nil {
		return nil
	}
	var items []Item
	for _,file := range f{
		if !file.IsDir(){
			items = append(items,Item{
				title:    file.Name(),
				desc:     "Vintlang documentation file",
				filename: file.Name(),
			})
		}
	}
	return items
}