package core

import (
	"fmt"
	"testing"

	"github.com/AnandK-2024/Blockchain/crypto"
	"github.com/AnandK-2024/Blockchain/types"
	"github.com/stretchr/testify/assert"
	// "github.com/holiman/uint256"
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

	data1 := []byte("expected hash")
	databyte := types.HashFromByte(data1)
	// Calculate the expected hash
	expectedHash := types.Hash(databyte)

	// // Mock the behavior of gob encoding
	// encodeMock := func(enc Encoder[*Transaction]) error {
	// 	return enc.Encode(transaction)
	// }

	// // Override the Encode method with the mock
	// transaction.Encode = encodeMock

	// Call the Hash function
	hash := transaction.Hash()
	fmt.Println("hash of transaction data", hash)

	// Verify that the calculated hash matches the expected hash
	if hash != expectedHash {
		t.Errorf("Expected hash to be %s, but got %s", expectedHash, hash)
	}
}

// func TestTransactionSignAndVerify(t *testing.T) {
// 	data := []byte("test data")
// 	privKey, _ := crypto.GenerateKeyPair()

// 	// Create a new transaction
// 	transaction := NewTransaction(data)

// 	// Sign the transaction
// 	err := transaction.sign(privKey)
// 	if err != nil {
// 		t.Errorf("Failed to sign transaction: %s", err)
// 	}

// 	// Verify the transaction signature
// 	err = transaction.Verify()
// 	if err != nil {
// 		t.Errorf("Failed to verify transaction signature: %s", err)
// 	}
// }
