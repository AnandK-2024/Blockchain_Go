package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"

	"github.com/AnandK-2024/Blockchain/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct{}

func Hash(h *Header) types.Hash {
	buf := &bytes.Buffer{}
	//// NewEncoder returns a new encoder that will transmit on the io.Writer.
	enc := gob.NewEncoder(buf)
	// Encode transmits the data item represented by the empty interface value,
	// guaranteeing that all necessary type information has been transmitted first.
	err := enc.Encode(h)
	// Passing a nil pointer to Encoder will panic, as they cannot be transmitted by gob.
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(buf.Bytes())
	return types.Hash(hash)
}
