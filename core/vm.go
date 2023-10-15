package core

import (
	"encoding/binary"
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
		vm.ip++
		if vm.ip > len(vm.data)-1 {
			break
		}
	}
	return nil
}

func (vm *VM) execute(opcode OpCode) error {
	switch opcode {
	case Store:
		var (
			key             = vm.stack.Pop().([]byte)
			value           = vm.stack.Pop()
			serializedValue []byte
		)
		switch val := value.(type) {
		case int:
			serializedValue = serializeInt64(int64(val))
		default:
			panic("TODO:Unknown type")

			vm.contractState.Put(key, serializedValue)

		}
	case PushInt:
		vm.stack.Push(int(vm.data[vm.ip+1]))
		// vm.ip++
	case PushByte:
		vm.stack.Push(byte(vm.data[vm.ip+1]))
		// vm.ip++
	case ADD:
		a := vm.stack.Pop().(int)
		b := vm.stack.Pop().(int)
		vm.stack.Push(int(a + b))
	case SUB:
		a := vm.stack.Pop().(int)
		b := vm.stack.Pop().(int)
		vm.stack.Push(int(a - b))

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
