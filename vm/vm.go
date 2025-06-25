package vm

import (
	"github.com/vintlang/vintlang/compiler"
	"github.com/vintlang/vintlang/object"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions []byte

	stack []object.Object
	sp    int // Always points to the next value. Top of the stack is stack[sp-1]
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,
		stack:        make([]object.Object, StackSize),
		sp:           0,
	}
}

func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) Run() error {
	// the main fetch-decode-execute cycle will go here
	return nil
}
