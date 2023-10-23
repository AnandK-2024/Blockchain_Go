package types

import (
	"encoding/hex"
	"fmt"
)

type Address [20]uint8

// to convert string to bytes
func (a Address) ToSlice() []byte {
	// reates a slice of bytes with a length of 32: dynamic slice byte
	b := make([]byte, 20)
	for i := 0; i < 20; i++ {
		b[i] = a[i]
	}
	return b

}

func AddressFromByte(b []byte) Address {
	if len(b) != 20 {
		msg := fmt.Sprintf("given bytes with length %d should be 20", len(b))
		// panic built-in function stops normal execution of the current goroutine
		panic(msg)
	}
	var value [20]uint8
	for i := 0; i < 20; i++ {
		value[i] = b[i]
	}
	return Address(value)
}

func (a Address) String() string {
	return hex.EncodeToString(a.ToSlice())
}
