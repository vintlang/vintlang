package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	c := make(map[string]bool)
	return &Environment{store: s, constants: c, outer: nil}
}

type Environment struct {
	store     map[string]Object
	constants map[string]bool
	outer     *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Define(name string, val Object) Object {
	if _, ok := e.store[name]; ok {
		return NewError("Identifier '" + name + "' has already been declared")
	}
	e.store[name] = val
	return val
}

func (e *Environment) DefineConst(name string, val Object) Object {
	if _, ok := e.store[name]; ok {
		return NewError("Identifier '" + name + "' has already been declared")
	}
	e.constants[name] = true
	e.store[name] = val
	return val
}

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

func (e *Environment) SetScoped(name string, val Object) Object {
	if e.constants[name] {
		return NewError("Cannot assign to constant '" + name + "'")
	}
	e.store[name] = val
	return val
}

func (e *Environment) Del(name string) bool {
	_, ok := e.store[name]
	if ok {
		delete(e.store, name)
	}
	return true
}
