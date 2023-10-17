package types

import (
	"testing"
	// "github.com/AnandK-2024/Blockchain/types"
)

func TestList(t *testing.T) {
	// Create a new list
	list := NewList[int]()

	// Insert values into the list
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)

	// Check the length of the list
	if list.Len() != 3 {
		t.Errorf("Expected list length to be 3, but got %d", list.Len())
	}

	// Get values from the list by index
	val := list.Get(0)
	if val != 10 {
		t.Errorf("Expected value at index 0 to be 10, but got %d", val)
	}

	// Get index by value
	index := list.GetIndex(20)
	if index != 1 {
		t.Errorf("Expected index of value 20 to be 1, but got %d", index)
	}

	// Check if a value is contained in the list
	contains := list.Contain(30)
	if !contains {
		t.Error("Expected list to contain value 30, but it does not")
	}

	// Remove a value from the list
	list.Remove(20)

	// Check if the value was removed
	contains = list.Contain(20)
	if contains {
		t.Error("Expected list to not contain value 20, but it does")
	}

	// Clear the list
	list.Clear()

	// Check if the list is empty
	if list.Len() != 0 {
		t.Error("Expected list to be empty, but it is not")
	}
}
