package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestState(t *testing.T) {

	state := NewState()

	// Test Put function
	err := state.Put([]byte("key1"), []byte("value1"))
	if err != nil {
		t.Errorf("Error putting data into state: %v", err)
	}

	// Test Get function
	value, err := state.Get([]byte("key1"))
	if err != nil {
		t.Errorf("Error getting data from state: %v", err)
	} else if string(value) != "value1" {
		t.Errorf("Expected 'value1', got '%s'", value)
	}

	// Test Delete function
	err = state.Delete([]byte("key1"))
	if err != nil {
		t.Errorf("Error deleting data from state: %v", err)
	}

	// Test Get function after deletion
	_, err = state.Get([]byte("key1"))
	assert.NotNil(t, err)
}
