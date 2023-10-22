package core

import (
	"crypto/elliptic"
	"encoding/gob"
	"io"
	// "github.com/AnandK-2024/Blockchain/crypto"
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
	return &GobTxEncoder{
		w: w,
	}
}

// encoding a transaction object into a binary format
func (g *GobTxEncoder) Encode(tx *Transaction) error {
	return gob.NewEncoder(g.w).Encode(tx)
}

type GobtxDecoder struct {
	r io.Reader
}

func NewGobTxDecoder(r io.Reader) *GobtxDecoder {
	return &GobtxDecoder{
		r: r,
	}
}

// decode transaction from a binary format
func (g *GobtxDecoder) Decode(tx *Transaction) error {
	return gob.NewDecoder(g.r).Decode(tx)
}

type GobBlockEncoder struct {
	w io.Writer
}

func NewGobBlockEncoder(w io.Writer) *GobBlockEncoder {
	return &GobBlockEncoder{
		w: w,
	}
}

// encoding a Block object into a binary format
func (g *GobBlockEncoder) Encode(b *Block) error {
	return gob.NewEncoder(g.w).Encode(b)
}

type GobBlockDecoder struct {
	r io.Reader
}

func NewGobBlockDecoder(r io.Reader) *GobBlockDecoder {
	return &GobBlockDecoder{
		r: r,
	}
}

// decode block from a binary format
func (g *GobBlockDecoder) Decode(b *Block) error {
	return gob.NewDecoder(g.r).Decode(b)
}

func init() {
	gob.Register(elliptic.P256())
	gob.Register(Transaction{})

}
