package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"

	"encoding/hex"

	// "errors"
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
	hash      types.Hash
	// first seen is the timestamp when tx is seen locally
	firstSeen int64
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
	if !tx.hash.IsZero() {
		return tx.hash
	}
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
	tx.hash = h
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

func CalculateMerkleRoot(txs []Transaction) string {
	// Convert the transactions to their hash representation
	hashes := make([]string, len(txs))
	for i, tx := range txs {
		h := tx.Hash()
		hashes[i] = hex.EncodeToString(h[:])
	}

	// Calculate the Merkle root
	for len(hashes) > 1 {
		// If the number of hashes is odd, duplicate the last hash
		if len(hashes)%2 != 0 {
			hashes = append(hashes, hashes[len(hashes)-1])
		}

		// Create a new slice to store the next level of hashes
		nextLevel := make([]string, len(hashes)/2)

		// Calculate the hash of each pair of hashes
		for i := 0; i < len(hashes); i += 2 {
			concatenated := hashes[i] + hashes[i+1]
			hash := sha256.Sum256([]byte(concatenated))
			nextLevel[i/2] = hex.EncodeToString(hash[:])
		}

		// Replace the current level of hashes with the next level
		hashes = nextLevel
	}

	// Return the Merkle root
	return hashes[0]
}

func StringToHash(hashString string) (types.Hash, error) {
	var hash types.Hash

	// Decode the string to bytes
	bytes, err := hex.DecodeString(hashString)
	if err != nil {
		return hash, err
	}

	// Copy the bytes to the hash variable
	copy(hash[:], bytes)

	return hash, nil
}

func (tx *Transaction) SetFirstSeen(t int64){
	tx.firstSeen=t
}

func (tx *Transaction) FirstSeen() int64{
	return tx.firstSeen
}

// enoding Transaction
func (tx *Transaction) Encode(enc Encoder[*Transaction]) error{
	return enc.Encode(tx)
}

//decoding transaction  
func (tx *Transaction) Decode(dec Decoder[*Transaction]) error{
	return dec.Decode(tx)
}