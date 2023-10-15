package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := NewStack(100)

	// Test pushing items onto the stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	// fmt.Println(stack.data)

	// Test popping items from the stack
	value := stack.Pop()
	if value != 3 {
		t.Errorf("Expected 3, got %d", value)
	}

	value = stack.Pop()
	if value != 2 {
		t.Errorf("Expected 2, got %d", value)
	}

	value = stack.Pop()
	if value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}

	// Test popping from an empty stack
	err := stack.Pop()
	assert.Nil(t, err)

}
