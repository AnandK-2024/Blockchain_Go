package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"

	"github.com/AnandK-2024/Blockchain/crypto"
	"github.com/AnandK-2024/Blockchain/types"
)

type Header struct {
	version uint32
	// hash of previous block header
	prevblockHash types.Hash
	// hash of transactions(merkle root)
	DataHash  types.Hash
	Timestamp int64
	Height    uint32
}

type Block struct {
	*Header      // don't work with copy of header , so taking pointer make much more efficient
	Transactions []Transaction
	// public key of block validator
	validator crypto.PublicKey
	// signature of block validator
	signature *crypto.Signature
	// hash of block header== hash of current block
	hash types.Hash
}

func NewBlock(h *Header, txs []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txs,
	}
}

func (b *Block) Hash() types.Hash {
	buf := &bytes.Buffer{}
	// NewEncoder returns a new encoder that will transmit on the io.Writer.
	enc := gob.NewEncoder(buf)
	// Encode transmits the data item represented by the empty interface value,
	// guaranteeing that all necessary type information has been transmitted first.
	err := enc.Encode(b.Header)
	// Passing a nil pointer to Encoder will panic, as they cannot be transmitted by gob.
	if err != nil {
		panic(err)
	}

	h := sha256.Sum256(buf.Bytes())
	return types.Hash(h)
}

func (b *Block) Sign(privkey *crypto.PrivateKey) error {
	hash := b.Hash()
	signature, err := privkey.SignMessage(hash[:])
	if err != nil {
		fmt.Println("unable to sign block with private key")
		panic(err)
	}
	b.validator = privkey.GeneratePublicKey()
	b.hash = hash
	b.signature = signature

	return nil
}

func (b *Block) Verify() error {
	if b.signature == nil {
		return fmt.Errorf("block has not signature")
	}
	sig := b.signature
	sucess := sig.Verify(b.validator, b.hash[:])
	if !sucess {
		return fmt.Errorf("invalid block header signature ")
	}
	return nil
}
