package core

import (
	// "bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/AnandK-2024/Blockchain/crypto"
	"github.com/stretchr/testify/assert"
	// // "github.com/holiman/uint256"
)

func TestTransaction(t *testing.T) {
	privkey := crypto.GeneratePrivatekey()
	tx := NewTransaction([]byte("Anand-->bob: 10ETH"))
	err := tx.sign(&privkey)
	fmt.Println("transaction signature", tx.signature)
	assert.Nil(t, err)
	assert.NotNil(t, tx.signature)
}

func TestTransactionHash(t *testing.T) {
	data := []byte("test data")
	transaction := NewTransaction(data)

	// Call the Hash function
	hash := transaction.Hash()
	fmt.Println("hash of transaction data and length of hash", hash, len(hash))

	assert.Equal(t, 32, len(hash))
}

func TestTransactionSignAndVerify(t *testing.T) {
	data := []byte("test data")
	privKey := crypto.GeneratePrivatekey()
	// pubkey:=privKey.GeneratePublicKey()

	// Create a new transaction
	transaction := NewTransaction(data)

	// Sign the transaction
	err := transaction.sign(&privKey)
	if err != nil {
		t.Errorf("Failed to sign transaction: %s", err)
	}

	// Verify the transaction signature
	err = transaction.Verify()
	if err != nil {
		t.Errorf("Failed to verify transaction signature: %s", err)
	}
}

// mekle root test
func TestCalculateMerkleRoot(t *testing.T) {
	// Create some example transactions
	txs := []Transaction{
		{
			data:      []byte("Transaction 1"),
			value:     100,
			from:      crypto.PublicKey{},
			signature: &crypto.Signature{},
			Nonce:     1,
		},
		{
			data:      []byte("Transaction 2"),
			value:     200,
			from:      crypto.PublicKey{},
			signature: &crypto.Signature{},
			Nonce:     2,
		},
		{
			data:      []byte("Transaction 3"),
			value:     300,
			from:      crypto.PublicKey{},
			signature: &crypto.Signature{},
			Nonce:     3,
		},
	}

	// Calculate the expected Merkle root manually
	hashes := make([]string, len(txs))
	for i, tx := range txs {
		h := tx.Hash()
		hashes[i] = hex.EncodeToString(h[:])
	}

	for len(hashes) > 1 {
		if len(hashes)%2 != 0 {
			hashes = append(hashes, hashes[len(hashes)-1])
		}

		nextLevel := make([]string, len(hashes)/2)

		for i := 0; i < len(hashes); i += 2 {
			concatenated := hashes[i] + hashes[i+1]
			hash := sha256.Sum256([]byte(concatenated))
			nextLevel[i/2] = hex.EncodeToString(hash[:])
		}

		hashes = nextLevel
	}

	expectedMerkleRoot := hashes[0]

	// Calculate the Merkle root using the function
	calculatedMerkleRoot := CalculateMerkleRoot(txs)

	// Compare the expected and calculated Merkle roots
	if calculatedMerkleRoot != expectedMerkleRoot {
		t.Errorf("Expected Merkle root: %s, but got: %s", expectedMerkleRoot, calculatedMerkleRoot)
	}
}

// func TestTxEncodeDecode(t *testing.T) {
// 	tx := NewRandomTransaction(100)
// 	fmt.Println("new random transaction is:=", tx)
// 	buf := &bytes.Buffer{}
// 	assert.Nil(t, tx.Encode(NewGobTxEncoder(buf)))
// 	txDecode := new(Transaction)
// 	assert.Nil(t, txDecode.Decode(NewGobTxDecoder(buf)))
// 	fmt.Println("decoded transaction:", txDecode)
// 	assert.Equal(t, tx, txDecode)

// }
