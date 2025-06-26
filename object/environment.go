package object

// Environment represents a variable/function scope in VintLang.
// Now supports function overloading: multiple functions with the same name but different signatures.
type Environment struct {
	store     map[string]Object      // For variables and non-function objects
	funcs     map[string][]*Function // For overloaded functions
	constants map[string]bool
	outer     *Environment
}

// NewEnvironment creates a new environment with support for function overloading.
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	f := make(map[string][]*Function)
	c := make(map[string]bool)
	return &Environment{store: s, funcs: f, constants: c, outer: nil}
}

// NewEnclosedEnvironment creates a new environment with an outer (parent) environment.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get returns a variable or function by name. For functions, returns the first overload (for backward compatibility).
func (e *Environment) Get(name string) (Object, bool) {
	if funcs, ok := e.funcs[name]; ok && len(funcs) > 0 {
		return funcs[0], true // Return the first overload for compatibility
	}
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Get(name)
	}
	return obj, ok
}

// GetAllFunctions returns all overloads for a function name, or nil if none exist.
func (e *Environment) GetAllFunctions(name string) []*Function {
	if funcs, ok := e.funcs[name]; ok && len(funcs) > 0 {
		return funcs
	}
	if e.outer != nil {
		return e.outer.GetAllFunctions(name)
	}
	return nil
}

// Define adds a variable or function to the environment. Functions are stored as overloads.
func (e *Environment) Define(name string, val Object) Object {
	if fn, ok := val.(*Function); ok {
		// Overload: append to the slice for this name
		e.funcs[name] = append(e.funcs[name], fn)
		return fn
	}
	if _, ok := e.store[name]; ok {
		return NewError("Identifier '" + name + "' has already been declared")
	}
	e.store[name] = val
	return val
}

// DefineConst adds a constant variable to the environment.
func (e *Environment) DefineConst(name string, val Object) Object {
	if _, ok := e.store[name]; ok {
		return NewError("Identifier '" + name + "' has already been declared")
	}
	e.constants[name] = true
	e.store[name] = val
	return val
}

// Assign updates the value of a variable in the environment.
func (e *Environment) Assign(name string, val Object) (Object, bool) {
	if e.constants[name] {
		return NewError("Cannot assign to constant '" + name + "'"), true
	}
	if _, ok := e.store[name]; ok {
		e.store[name] = val
		return val, true
	}
	if e.outer != nil {
		return e.outer.Assign(name, val)
	}
	return nil, false
}

// SetScoped sets a variable in the current scope only.
func (e *Environment) SetScoped(name string, val Object) Object {
	if e.constants[name] {
		return NewError("Cannot assign to constant '" + name + "'")
	}
	e.store[name] = val
	return val
}

// Del deletes a variable from the environment.
func (e *Environment) Del(name string) bool {
	_, ok := e.store[name]
	if ok {
		delete(e.store, name)
	}
	if _, ok := e.funcs[name]; ok {
		delete(e.funcs, name)
	}
	return true
}
