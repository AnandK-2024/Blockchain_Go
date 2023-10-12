package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/AnandK-2024/Blockchain/crypto"

	// "github.com/holiman/uint256"
	"github.com/AnandK-2024/Blockchain/types"
	"github.com/stretchr/testify/assert"
)

func RandomBlock(height uint32) *Block {
	header := &Header{
		version:       1,
		prevblockHash: types.Randomhash(),
		DataHash:      types.Randomhash(),
		Timestamp:     time.Now().UnixNano(),
		Height:        height,
	}
	txs := Transaction{
		data: []byte("Anand --> bob: 5ETH"),
	}
	return NewBlock(header, []Transaction{txs})
}

func TestHashBlock(t *testing.T) {
	genesisBlock := RandomBlock(0)
	fmt.Println("prev hash", genesisBlock.prevblockHash)
	fmt.Println("current blockhash", genesisBlock.Hash())

}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivatekey()
	b := RandomBlock(100)

	assert.Nil(t, b.Sign(&privKey))
	assert.NotNil(t, b.signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivatekey()
	b := RandomBlock(100)

	assert.Nil(t, b.Sign(&privKey))
	// assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivatekey()
	b.validator = otherPrivKey.GeneratePublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}
