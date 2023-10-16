package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVM(t *testing.T) {
	data := []byte{0x0, 0x05, 0x0, 0xa, 0x1, 0x0, 0xb, 0x3, 0x0, 0xf, 0x2, 0x0, 0x5}
	s := state{}
	contractState := s.NewState()
	vm := NewVM(data, *contractState)
	assert.Nil(t, vm.Run())
	fmt.Println("vm result after execution:", vm.stack.peek())

	// valueBytes, err := contractState.Get([]byte("FOO"))
	// value := deserializeInt64(valueBytes)
	// assert.Nil(t, err)
	// assert.Equal(t, value, int64(5))
}

// func TestExecute(t *testing.T) {
// 	contractState := state{}
// 	vm := NewVM([]byte{0x0, 0x1, 0xe}, contractState)

// 	// Test Store opcode
// 	vm.stack.Push([]byte("key"))
// 	vm.stack.Push(42)
// 	err := vm.execute(Store)
// 	if err != nil {
// 		t.Errorf("Error executing Store opcode: %v", err)
// 	}

// 	// Test PushInt opcode
// 	vm.ip = 1
// 	err = vm.execute(PushInt)
// 	if err != nil {
// 		t.Errorf("Error executing PushInt opcode: %v", err)
// 	}

// 	// Test ADD opcode
// 	vm.ip = 2
// 	vm.stack.Push(10)
// 	vm.stack.Push(20)
// 	err = vm.execute(ADD)
// 	if err != nil {
// 		t.Errorf("Error executing ADD opcode: %v", err)
// 	}

// 	// Test SUB opcode
// 	vm.ip = 3
// 	vm.stack.Push(30)
// 	vm.stack.Push(15)
// 	err = vm.execute(SUB)
// 	if err != nil {
// 		t.Errorf("Error executing SUB opcode: %v", err)
// 	}

// 	// Add more test cases for other opcodes...

// 	// 	// Verify the final state of the stack and contractState
// 	// 	expectedStack := []interface{}{42, 1, 30, 15}
// 	// 	if !compareStacks(vm.stack, expectedStack) {
// 	// 		t.Errorf("Unexpected stack state. Expected: %v, Got: %v", expectedStack, vm.stack.items)
// }

// 	expectedContractState := map[string][]byte{
// 		"key": serializeInt64(42),
// 	}
// 	if !compareContractStates(vm.contractState, expectedContractState) {
// 		t.Errorf("Unexpected contract state. Expected: %v, Got: %v", expectedContractState, vm.contractState.data)
// 	}
// }

// func compareStacks(stack *Stack, expected []interface{}) bool {
// 	if len(stack.items) != len(expected) {
// 		return false
// 	}
// 	for i, item := range stack.items {
// 		if item != expected[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func compareContractStates(contractState *state, expected map[string][]byte) bool {
// 	if len(contractState.data) != len(expected) {
// 		return false
// 	}
// 	for key, value := range contractState.data {
// 		expectedValue, ok := expected[key]
// 		if !ok || !compareByteSlices(value, expectedValue) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func compareByteSlices(a, b []byte) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	for i, val := range a {
// 		if val != b[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
