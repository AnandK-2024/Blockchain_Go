package core

import (
	"encoding/binary"
	"fmt"
)

type VM struct {
	data          []byte // contain all opcode
	ip            int    // instruction pointer
	stack         *Stack // contains all oprands
	contractState *state
}

func NewVM(data []byte, contractstate state) *VM {
	return &VM{
		stack:         NewStack(1024),
		contractState: &contractstate,
		data:          data,
		ip:            0,
	}
}

func (vm *VM) Run() error {
	for {
		opcode := OpCode(vm.data[vm.ip])
		if err := vm.execute(opcode); err != nil {
			return err
		}
		if vm.ip > len(vm.data)-1 {
			break
		}
	}
	return nil
}

func (vm *VM) execute(opcode OpCode) error {
	switch opcode {
	case Store:
		fmt.Println("store opcode called")

		var (
			key            = vm.stack.Pop().([]byte)
			value          = vm.stack.Pop()
			serializevalue []byte
		)
		switch v := value.(type) {
		case int:
			serializevalue = serializeInt64(int64(v))
		default:
			panic("TODO: unknown type")
		}
		vm.contractState.Put(key, serializevalue)

		vm.ip = vm.ip + 1
	case PushInt:
		fmt.Println("push int opcode called")
		vm.stack.Push(int(vm.data[vm.ip+1]))
		vm.ip = vm.ip + 2
	case PushByte:
		fmt.Println("push byte opcode called")
		vm.stack.Push(byte(vm.data[vm.ip+1]))
		vm.ip = vm.ip + 2
	case ADD:
		fmt.Println("Add opcode called")
		a := vm.stack.Pop().(int)
		b := vm.stack.Pop().(int)
		vm.stack.Push(int(a + b))
		vm.ip = vm.ip + 1
	case SUB:
		fmt.Println("Sub opcode called")
		a := vm.stack.Pop().(int)
		b := vm.stack.Pop().(int)
		vm.stack.Push(int(b - a))
		vm.ip = vm.ip + 1
	case MUL:
		fmt.Println("MUL opcode called")
		a := vm.stack.Pop().(int)
		b := vm.stack.Pop().(int)
		vm.stack.Push(int(b * a))
		vm.ip = vm.ip + 1
	default:
		fmt.Println("opcode called in default:", opcode)

	}
	return nil
}

func serializeInt64(val int64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(val))
	return buf
}

func deserializeInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(b))
}
