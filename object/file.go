package object

import (
	"os"
)

type File struct {
	Filename string
	Content  string
}

func (f *File) Type() ObjectType { return FILE_OBJ }
func (f *File) Inspect() string  { return f.Filename }

func (f *File) Method(method string, args []Object) Object {
	switch method {
	case "read":
		return f.read(args)
	case "write":
		return f.write(args)
	case "append":
		return f.append(args)
	default:
		return newError("Method '%s' is not supported for the file object.", method)
	}
}

func (f *File) read(args []Object) Object {
	if len(args) != 0 {
		return newError("Expected 0 arguments for 'read', but got %d.", len(args))
	}
	return &String{Value: f.Content}
}

func (f *File) write(args []Object) Object {
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

func (f *File) append(args []Object) Object {
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
