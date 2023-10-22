package core

import (
	"bytes"
	// "encoding/gob"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/AnandK-2024/Blockchain/crypto"

	// "github.com/holiman/uint256"
	"github.com/AnandK-2024/Blockchain/types"
	"github.com/stretchr/testify/assert"
)

func RandomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevblockHash: types.Randomhash(),
		DataHash:      types.Randomhash(),
		Timestamp:     time.Now().UnixNano(),
		Height:        height,
	}
	txs := Transaction{
		Data: []byte("Anand --> bob: 5ETH"),
	}
	return NewBlock(header, []*Transaction{&txs})
}

func randomBlock(t *testing.T, height uint32, prevhash types.Hash) *Block {
	header := &Header{
		Version:       1,
		PrevblockHash: prevhash,
		DataHash:      types.Randomhash(),
		Timestamp:     time.Now().UnixNano(),
		Height:        height,
	}
	txs := &Transaction{
		Data: []byte{0x0, 0x05},
	}
	return NewBlock(header, []*Transaction{txs})
}

func TestHashBlock(t *testing.T) {
	genesisBlock := RandomBlock(0)
	fmt.Println("prev hash", genesisBlock.PrevblockHash)
	fmt.Println("current blockhash", genesisBlock.Hash())

}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivatekey()
	b := RandomBlock(100)

	assert.Nil(t, b.Sign(&privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivatekey()
	b := RandomBlock(100)

	assert.Nil(t, b.Sign(&privKey))
	// assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivatekey()
	b.Validator = otherPrivKey.GeneratePublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}
func TestBlockHash(t *testing.T) {
	// Create a sample block
	block := &Block{
		Header: &Header{
			Version:       1,
			PrevblockHash: types.Hash{},
			DataHash:      types.Hash{},
			Timestamp:     0,
			Height:        0,
		},
		Transactions: []*Transaction{},
	}

	// Calculate the hash for the block
	hash := block.Hash()
	fmt.Println("current block details:", block)
	fmt.Println("hash of current block:", hash)

	// Assert that the hash is not nil
	assert.NotNil(t, hash, "Block hash should not be nil")
}
func TestBlockAddTransaction(t *testing.T) {
	// Create a sample block
	block := &Block{
		Header: &Header{
			Version:       1,
			PrevblockHash: types.Hash{},
			DataHash:      types.Hash{},
			Timestamp:     0,
			Height:        0,
		},
		Transactions: []*Transaction{NewRandomTransaction(4), NewRandomTransaction(400), NewRandomTransaction(40)},
	}

	// Create a sample transaction
	transaction := NewRandomTransaction(456)

	// Add the transaction to the block
	block.AddTransaction(transaction)
	fmt.Println("merkle root of transaction", block.DataHash)

	// Assert that the block's transaction list is not empty
	assert.NotEmpty(t, block.Transactions, "Block's transaction list should not be empty")
}

func TestEncodeDecode(t *testing.T) {
	// Create a sample block
	block := &Block{
		Header: &Header{
			Version:       1,
			PrevblockHash: types.Hash{},
			DataHash:      types.Hash{},
			Timestamp:     0,
			Height:        0,
		},
		Transactions: []*Transaction{NewRandomTransaction(4), NewRandomTransaction(400), NewRandomTransaction(40)},
	}

	// Create a buffer to encode the block
	var buf bytes.Buffer
	encoder := NewGobBlockEncoder(&buf)

	// Encode the block
	err := encoder.Encode(block)
	if err != nil {
		t.Errorf("Error encoding block: %v", err)
	}

	decorder := NewGobBlockDecoder(&buf)
	// Create a new block to decode into
	decodedBlock := &Block{}

	// Create a buffer from the encoded data
	// dec := gob.NewDecoder(bytes.NewReader(buf.Bytes()))

	// Decode the block
	err = decodedBlock.Decode(decorder)
	if err != nil {
		t.Errorf("Error decoding block: %v", err)
	}
	assert.Equal(t, block, decodedBlock)
	// Compare the original block and the decoded block
	if !reflect.DeepEqual(block, decodedBlock) {
		t.Errorf("Decoded block does not match original block")
	}
}
