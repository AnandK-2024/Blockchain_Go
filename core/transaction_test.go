package core

import (
	// "bytes"
	// "crypto/sha256"
	// "encoding/hex"
	"bytes"
	"fmt"
	"testing"

	"github.com/AnandK-2024/Blockchain/crypto"
	"github.com/stretchr/testify/assert"
	// // "github.com/holiman/uint256"
)

func TestTransaction(t *testing.T) {
	privkey := crypto.GeneratePrivatekey()
	tx := NewTransaction([]byte("Anand-->bob: 10ETH"))
	err := tx.Sign(&privkey)
	fmt.Println("transaction signature", tx.signature)
	assert.Nil(t, err)
	assert.NotNil(t, tx.signature)
}
func TestNewRandomTransaction(t *testing.T) {
	datasize := 10
	tx := NewRandomTransaction(datasize)

	assert.NotNil(t, tx)
	assert.Len(t, tx.data, datasize)
	assert.NotZero(t, tx.Nonce)
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
	err := transaction.Sign(&privKey)
	if err != nil {
		t.Errorf("Failed to sign transaction: %s", err)
	}

	// Verify the transaction signature
	err = transaction.Verify()
	if err != nil {
		t.Errorf("Failed to verify transaction signature: %s", err)
	}
}

func TestCalculateMerkleRoot(t *testing.T) {
	txs := []*Transaction{
		NewTransaction([]byte("data1")),
		NewTransaction([]byte("data2")),
		NewTransaction([]byte("data3")),
	}

	merkleRoot := CalculateMerkleRoot(txs)
	fmt.Println("merkle hash root of transaction", merkleRoot)
	assert.NotEmpty(t, merkleRoot)
}

func TestStringToHash(t *testing.T) {
	hashString := "0123456789abcdef"
	expectedHash, _ := StringToHash(hashString)

	hash, err := StringToHash(hashString)

	assert.NoError(t, err)
	assert.Equal(t, expectedHash, hash)
}

func TestTransaction_SetFirstSeen(t *testing.T) {
	tx := NewTransaction([]byte("test data"))
	firstSeen := int64(1234567890)

	tx.SetFirstSeen(firstSeen)

	assert.Equal(t, firstSeen, tx.FirstSeen())
}


func TestTxEncodeDecode(t *testing.T) {
	tx := NewRandomTransaction(100)
	// fmt.Println("new random transaction is:=", tx)
	buf := &bytes.Buffer{}
	assert.Nil(t, tx.Encode(NewGobTxEncoder(buf)))
	txDecode := new(Transaction)
	assert.Nil(t, txDecode.Decode(NewGobTxDecoder(buf)))
	fmt.Println("decoded transaction:", txDecode)
	// assert.Equal(t, tx, txDecode)

}
