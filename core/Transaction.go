package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"math/rand"

	"github.com/AnandK-2024/Blockchain/crypto"
	"github.com/AnandK-2024/Blockchain/types"
)

type Transaction struct {
	data      []byte
	value     uint64
	from      crypto.PublicKey
	signature *crypto.Signature
	Nonce     uint64
}

/*
step1: make Transaction data
step2: sign transaction data using private key
step3: calculate hash transaction  data for using HAsh function
step4: verifier can verify transaction
*/

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		data:  data,
		Nonce: uint64(rand.Int63n(1000000000000000)),
	}
}

func NewRandomTransaction(datasize int) *Transaction {
	return NewTransaction(types.RandomByte(datasize))
}

func (tx *Transaction) Hash() types.Hash {
	buf := &bytes.Buffer{}
	// NewEncoder returns a new encoder that will transmit on the io.Writer.
	enc := gob.NewEncoder(buf)
	// Encode transmits the data item represented by the empty interface value,
	// guaranteeing that all necessary type information has been transmitted first.
	err := enc.Encode(tx)
	// Passing a nil pointer to Encoder will panic, as they cannot be transmitted by gob.
	if err != nil {
		panic(err)
	}

	h := sha256.Sum256(buf.Bytes())
	return types.Hash(h)
}

func (tx *Transaction) sign(privkey *crypto.PrivateKey) error {
	signature, err := privkey.SignMessage(tx.data)
	if err != nil {
		fmt.Println("unable to sign block with private key")
		return err
	}
	tx.from = privkey.GeneratePublicKey()
	tx.signature = signature
	return nil
}

func (tx *Transaction) Verify() error {
	if tx.signature == nil {
		return fmt.Errorf("tx has not signature")
	}
	sig := tx.signature
	sucess := sig.Verify(tx.from, tx.data)
	if !sucess {
		return fmt.Errorf("invalid tx signature ")
	}
	return nil
}

func (tx *Transaction) Decode(dec Decoder[*Transaction]) error {
	return dec.Decode(tx)
}

func (tx *Transaction) Encode(enc Encoder[*Transaction]) error {
	return enc.Encode(tx)
}
