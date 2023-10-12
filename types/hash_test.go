package types

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHashFromByte(t *testing.T) {
	// Test case 1: Valid input
	validBytes := make([]byte, 32)
	hash := HashFromByte(validBytes)
	fmt.Println("hash:=", hash)
	if len(hash) != 32 {
		t.Errorf("Expected hash length of 32, but got %d", len(hash))
	}

	// Test case 2: Invalid input
	invalidBytes := make([]byte, 31)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	_ = HashFromByte(invalidBytes)
}

func TestIsZero(t *testing.T) {
	// Test case 1: Non-zero hash
	nonZeroHash := HashFromByte(make([]byte, 32))
	nonZeroHash[0] = 1
	if nonZeroHash.IsZero() {
		t.Errorf("Expected non-zero hash to return false for isZero(), but got true")
	}

	// Test case 2: Zero hash
	zeroHash := Hash{}
	if !zeroHash.IsZero() {
		t.Errorf("Expected zero hash to return true for isZero(), but got false")
	}
}

func TestToSlice(t *testing.T) {
	// Test case 1: Valid hash
	validHash := HashFromByte(make([]byte, 32))
	slice := validHash.ToSlice()
	if len(slice) != 32 {
		t.Errorf("Expected slice length of 32, but got %d", len(slice))
	}

	// Test case 2: Invalid hash
	invalidHash := Hash{}
	slice = invalidHash.ToSlice()
	if len(slice) != 32 {
		t.Errorf("Expected slice length of 32, but got %d", len(slice))
	}
}

func TestString(t *testing.T) {
	// Test case 1: Valid hash
	validHash := HashFromByte(make([]byte, 32))
	expectedString := hex.EncodeToString(validHash.ToSlice())
	result := validHash.string()
	if result != expectedString {
		t.Errorf("Expected string '%s', but got '%s'", expectedString, result)
	}

	// Test case 2: Invalid hash
	invalidHash := Hash{}
	expectedString = hex.EncodeToString(invalidHash.ToSlice())
	result = invalidHash.string()
	if result != expectedString {
		t.Errorf("Expected string '%s', but got '%s'", expectedString, result)
	}
}

func TestRandomByte(t *testing.T) {
	// Test case 1: Valid size
	size := 32
	randomBytes := RandomByte(size)
	if len(randomBytes) != size {
		t.Errorf("Expected random byte length of %d, but got %d", size, len(randomBytes))
	}

	// Test case 2: Invalid size
	size = -1
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
	}()
	_ = RandomByte(size)
}

func TestRandomHash(t *testing.T) {
	randomHash := Randomhash()
	if len(randomHash) != 32 {
		t.Errorf("Expected random hash length of 32, but got %d", len(randomHash))
	}
}
