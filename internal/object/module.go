package object

import "github.com/vintlang/vintlang/internal/ast"

type ModuleFunction func(args []VintObject, defs map[string]VintObject) VintObject

type Module struct {
	Name       string
	Functions  map[string]ModuleFunction
	FuncTypes  map[string][]ast.Type // param types per function, nil for untyped
	FuncReturns map[string]ast.Type  // return type per function, nil for unknown
	Variables  map[string]VintObject
	Submodules map[string]*Module
	Doc        string
	Version    string
	Author     string
}

// Returns a new Module
func NewModule(name string, functions map[string]ModuleFunction) *Module {
	if functions == nil {
		functions = make(map[string]ModuleFunction)
	}
	return &Module{
		Name:        name,
		Functions:   functions,
		FuncTypes:   make(map[string][]ast.Type),
		FuncReturns: make(map[string]ast.Type),
		Variables:   make(map[string]VintObject),
		Submodules:  make(map[string]*Module),
	}
}

// Register a function at runtime
func (m *Module) RegisterFunction(name string, fn ModuleFunction) {
	m.Functions[name] = fn
}

// RegisterTypedFunction registers a function with parameter and return type info.
// paramTypes can be nil for untyped params (no enforcement).
// returnType can be nil for unknown/void return.
func (m *Module) RegisterTypedFunction(name string, fn ModuleFunction, paramTypes []ast.Type, returnType ast.Type) {
	m.Functions[name] = fn
	if paramTypes != nil {
		m.FuncTypes[name] = paramTypes
	}
	if returnType != nil {
		m.FuncReturns[name] = returnType
	}
}

// Register a variable at runtime
func (m *Module) RegisterVariable(name string, value VintObject) {
	m.Variables[name] = value
}

// Register a submodule at runtime
func (m *Module) RegisterSubmodule(name string, mod *Module) {
	m.Submodules[name] = mod
}

// List all function names
func (m *Module) ListFunctions() []string {
	keys := make([]string, 0, len(m.Functions))
	for k := range m.Functions {
		keys = append(keys, k)
	}
	return keys
}

// List all variable names
func (m *Module) ListVariables() []string {
	keys := make([]string, 0, len(m.Variables))
	for k := range m.Variables {
		keys = append(keys, k)
	}
	return keys
}

// List all submodule names
func (m *Module) ListSubmodules() []string {
	keys := make([]string, 0, len(m.Submodules))
	for k := range m.Submodules {
		keys = append(keys, k)
	}
	return keys
}

func (m *Module) Type() VintObjectType {
	switch m.Name {
	case "time":
		return TIME_OBJ
	case "json":
		return JSON_OBJ
	default:
		return MODULE_OBJ
	}
}
func (m *Module) Inspect() string { return "Module: " + m.Name }
