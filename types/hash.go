package types

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Hash [32]uint8 

func HashFromByte(b []byte) Hash{
	if len(b)!=32{
		msg:=fmt.Sprintf("given bytes with length %d should be 32", len(b))
		// panic built-in function stops normal execution of the current goroutine
		panic(msg)
	}
	var value [32]uint8
	for i:=0;i<32;i++{
		value[i]=b[i]
	}
	return Hash(value)
}

// to check given hash value is zero or not
func (h Hash) isZero() bool{
	for i:=0;i<32;i++{
		if h[i]!=0{
			return false
		}
	}
	return true
}

// to convert string to bytes
func (h Hash) ToSlice() []byte{
	// reates a slice of bytes with a length of 32: dynamic slice byte
	b:=make([]byte,32)
	for i:=0;i<32;i++{
		b[i]=h[i]
	}
	return b;

}


// convert hash hex value to string
func (h Hash) string() string{
	return hex.EncodeToString(h.ToSlice())

}

// to generate random byte of size
func RandomByte(size int)[]byte{
random:=make([]byte,size)
rand.Read(random)
return random
}

// to generate hash of random byte
func Randomhash() Hash{
	return HashFromByte(RandomByte(32))
}
