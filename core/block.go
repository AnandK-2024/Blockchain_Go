package core

import (
	"io"
	"encoding/binary"

	"github.com/AnandK-2024/Blockchain/types"
)

type Header struct {
	version   uint32
	prevblock types.Hash
	Timestamp uint64 //creates a slice of bytes with a length of 32
	Height    uint32
	Nonce     uint64
}

type Block struct {
	Header
	Transactions []Transaction
}

func (h Header) EncodeBinary(w io.Writer) error {
	if err:=binary.write(w, binary.littleEndian, h.version); err!=nil {return nil}
	if err:=binary.write(w, binary.littleEndian, h.prevblock); err!=nil {return nil}
	if err:=binary.write(w, binary.littleEndian, h.Timestamp); err!=nil {return nil}
	if err:=binary.write(w, binary.littleEndian, h.Height); err!=nil {return nil}
	if err:=binary.write(w, binary.littleEndian, h.Nonce); err!=nil {return nil}
	return binary.write(w, binary.littleEndian, h.Nonce)

}

func (h Header) DecodeBinary(w io.Reader) error {
	if err:=binary.Read(w, binary.littleEndian, h.version); err!=nil {return nil}
	if err:=binary.Read(w, binary.littleEndian, h.prevblock); err!=nil {return nil}
	if err:=binary.Read(w, binary.littleEndian, h.Timestamp); err!=nil {return nil}
	if err:=binary.Read(w, binary.littleEndian, h.Height); err!=nil {return nil}
	if err:=binary.Read(w, binary.littleEndian, h.Nonce); err!=nil {return nil}
	return binary.Read(w, binary.littleEndian, h.Nonce)
}
