package builtins

import (
	"github.com/vintlang/vintlang/module"
	"github.com/vintlang/vintlang/object"
)

func init() {
	registerImportBuiltins()
}

func registerImportBuiltins() {
	RegisterBuiltin("import", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Function 'import()' requires exactly 1 argument, got %d", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("Argument to 'import()' must be a string, got %s", args[0].Type())
			}

			moduleName := args[0].(*object.String).Value

			if len(moduleName) == 0 {
				return newError("Module name cannot be empty")
			}

			// Check for built-in modules first
			if mod, exists := module.Mapper[moduleName]; exists {
				return mod
			}

			// TODO: For file-based modules, we'll need to implement the logic here
			// to avoid circular dependencies with import.go
			// For now, return an error indicating file-based imports aren't supported
			// in the builtin version yet
			return newError("Module '%s' not found in built-in modules. File-based module import not yet supported in import() builtin. Use 'import' statement instead for file modules.", moduleName)
		},
	})
}