package object

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Filename string
	Content  string
}

func (f *File) Type() VintObjectType { return FILE_OBJ }
func (f *File) Inspect() string      { return f.Filename }

func (f *File) Method(method string, args []VintObject) VintObject {
	switch method {
	case "read":
		return f.read(args)
	case "write":
		return f.write(args)
	case "append":
		return f.append(args)
	case "exists":
		return f.exists(args)
	case "size":
		return f.size(args)
	case "delete":
		return f.delete(args)
	case "copy":
		return f.copy(args)
	case "move":
		return f.move(args)
	case "lines":
		return f.lines(args)
	case "extension":
		return f.extension(args)
	default:
		return newError("Method '%s' is not supported for the file object.", method)
	}
}

func (f *File) read(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("Expected 0 arguments for 'read', but got %d.", len(args))
	}
	return &String{Value: f.Content}
}

func (f *File) write(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("Expected 1 argument for 'write', but got %d.", len(args))
	}
	content, ok := args[0].(*String)
	if !ok {
		return newError("Argument for 'write' must be of type String.")
	}
	err := os.WriteFile(f.Filename, []byte(content.Value), 0644)
	if err != nil {
		return newError("Error writing to file: %s", err.Error())
	}
	f.Content = content.Value
	return &Boolean{Value: true}
}

func (f *File) append(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("Expected 1 argument for 'append', but got %d.", len(args))
	}
	content, ok := args[0].(*String)
	if !ok {
		return newError("Argument for 'append' must be of type String.")
	}
	file, err := os.OpenFile(f.Filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return newError("Error opening file for appending: %s", err.Error())
	}
	defer file.Close()
	_, err = file.WriteString(content.Value)
	if err != nil {
		return newError("Error appending to file: %s", err.Error())
	}
	f.Content += content.Value
	return &Boolean{Value: true}
}

func (f *File) exists(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("exists() expects 0 arguments, got %d", len(args))
	}

	_, err := os.Stat(f.Filename)
	return &Boolean{Value: err == nil}
}

func (f *File) size(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("size() expects 0 arguments, got %d", len(args))
	}

	info, err := os.Stat(f.Filename)
	if err != nil {
		return newError("Error getting file size: %s", err.Error())
	}

	return &Integer{Value: info.Size()}
}

func (f *File) delete(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("delete() expects 0 arguments, got %d", len(args))
	}

	err := os.Remove(f.Filename)
	if err != nil {
		return newError("Error deleting file: %s", err.Error())
	}

	return &Boolean{Value: true}
}

func (f *File) copy(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("copy() expects 1 argument, got %d", len(args))
	}

	destination, ok := args[0].(*String)
	if !ok {
		return newError("Destination must be a string")
	}

	src, err := os.Open(f.Filename)
	if err != nil {
		return newError("Error opening source file: %s", err.Error())
	}
	defer src.Close()

	dst, err := os.Create(destination.Value)
	if err != nil {
		return newError("Error creating destination file: %s", err.Error())
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return newError("Error copying file: %s", err.Error())
	}

	return &Boolean{Value: true}
}

func (f *File) move(args []VintObject) VintObject {
	if len(args) != 1 {
		return newError("move() expects 1 argument, got %d", len(args))
	}

	destination, ok := args[0].(*String)
	if !ok {
		return newError("Destination must be a string")
	}

	err := os.Rename(f.Filename, destination.Value)
	if err != nil {
		return newError("Error moving file: %s", err.Error())
	}

	f.Filename = destination.Value
	return &Boolean{Value: true}
}

func (f *File) lines(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("lines() expects 0 arguments, got %d", len(args))
	}

	content, err := os.ReadFile(f.Filename)
	if err != nil {
		return newError("Error reading file: %s", err.Error())
	}

	lines := strings.Split(string(content), "\n")
	elements := make([]VintObject, len(lines))
	for i, line := range lines {
		elements[i] = &String{Value: line}
	}

	return &Array{Elements: elements}
}

func (f *File) extension(args []VintObject) VintObject {
	if len(args) != 0 {
		return newError("extension() expects 0 arguments, got %d", len(args))
	}

	ext := filepath.Ext(f.Filename)
	return &String{Value: ext}
}
