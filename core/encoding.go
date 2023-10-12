package core

import (
	"crypto/elliptic"
	"encoding/gob"
	"io"
)

type Encoder[T any] interface {
	Encode(T) error
}

type Decoder[T any] interface {
	Decode(T) error
}

type GobTxEncoder struct {
	w io.Writer
}

func NewGobTxEncoder(w io.Writer) *GobTxEncoder {
	gob.Register(elliptic.P256())
	gob.Register(Transaction{}) // Register the Transaction struct
	// gob.RegisterName("Transaction", Transaction{}) // Register the Transaction struct with a type ID

	return &GobTxEncoder{
		w: w,
	}
}

//encoding a transaction object into a binary format
func (g *GobTxEncoder) Encode(tx *Transaction) error {
	return gob.NewEncoder(g.w).Encode(tx)
}

// // EncodeTransaction encodes a transaction into a byte slice using the gob encoding format.
// func EncodeTransaction(tx *Transaction) ([]byte, error) {
// 	var buf bytes.Buffer
// 	enc := gob.NewEncoder(&buf)
// 	err := enc.Encode(tx)
// 	if err != nil {
// 		return nil, fmt.Errorf("error encoding transaction: %s", err)
// 	}
// 	return buf.Bytes(), nil
// }

type GobtxDecoder struct {
	r io.Reader
}

func NewGobTxDecoder(r io.Reader) *GobtxDecoder {
	gob.Register(elliptic.P256())
	gob.Register(Transaction{}) // Register the Transaction struct
	// gob.RegisterName("Transaction", Transaction{}) // Register the Transaction struct with a type ID

	return &GobtxDecoder{
		r: r,
	}
}

//decode transaction from a binary format
func (g *GobtxDecoder) Decode(tx *Transaction) error {
	return gob.NewDecoder(g.r).Decode(tx)
}

// // DecodeTransaction decodes a transaction from a byte slice using the gob encoding format.
// func DecodeTransaction(data []byte) (*Transaction, error) {
// 	var tx Transaction
// 	buf := bytes.NewBuffer(data)
// 	dec := gob.NewDecoder(buf)
// 	err := dec.Decode(&tx)
// 	if err != nil {
// 		return nil, fmt.Errorf("error decoding transaction: %s", err)
// 	}
// 	return &tx, nil
// }

type GobBlockEncoder struct {
	w io.Writer
}

func NewGobBlockEncoder(w io.Writer) *GobBlockEncoder {
	gob.Register(elliptic.P256())
	return &GobBlockEncoder{
		w: w,
	}
}

//encoding a Block object into a binary format
func (g *GobBlockEncoder) Encode(b *Block) error {
	return gob.NewEncoder(g.w).Encode(b)
}

type GobBlockDecoder struct {
	r io.Reader
}

func NewGobBlockDecoder(r io.Reader) *GobBlockDecoder {
	gob.Register(elliptic.P256())
	return &GobBlockDecoder{
		r: r,
	}
}

//decode block from a binary format
func (g *GobBlockDecoder) Decode(b *Block) error {
	return gob.NewDecoder(g.r).Decode(b)
}
