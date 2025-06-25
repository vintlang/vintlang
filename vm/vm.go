package vm

import (
	"encoding/binary"
	"fmt"

	"github.com/vintlang/vintlang/code"
	"github.com/vintlang/vintlang/compiler"
	"github.com/vintlang/vintlang/object"
)

const StackSize = 2048

var (
	True  = &object.Boolean{Value: true}
	False = &object.Boolean{Value: false}
)

type VM struct {
	constants    []object.Object
	instructions []byte

	stack               []object.Object
	sp                  int // Always points to the next value. Top of the stack is stack[sp-1]
	lastPoppedStackElem object.Object
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,
		stack:        make([]object.Object, StackSize),
		sp:           0,
	}
}

func (vm *VM) LastPoppedStackElem() object.Object {
	return vm.lastPoppedStackElem
}

func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := vm.instructions[ip]

		switch op {
		case byte(code.OpConstant):
			constIndex := binary.BigEndian.Uint16(vm.instructions[ip+1:])
			ip += 2 // Move instruction pointer past the 2-byte operand

			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
		case byte(code.OpPop):
			vm.pop()
		case byte(code.OpAdd), byte(code.OpSub), byte(code.OpMul), byte(code.OpDiv):
			err := vm.executeBinaryOperation(op)
			if err != nil {
				return err
			}
		case byte(code.OpTrue):
			err := vm.push(True)
			if err != nil {
				return err
			}
		case byte(code.OpFalse):
			err := vm.push(False)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (vm *VM) executeBinaryOperation(op byte) error {
	right := vm.pop()
	left := vm.pop()

	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	var result int64

	switch op {
	case byte(code.OpAdd):
		result = leftVal + rightVal
	case byte(code.OpSub):
		result = leftVal - rightVal
	case byte(code.OpMul):
		result = leftVal * rightVal
	case byte(code.OpDiv):
		result = leftVal / rightVal
	default:
		return fmt.Errorf("unknown integer operator: %d", op)
	}

	return vm.push(&object.Integer{Value: result})
}

func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = o
	vm.sp++

	return nil
}

func (vm *VM) pop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	o := vm.stack[vm.sp-1]
	vm.sp--
	vm.lastPoppedStackElem = o
	return o
}
