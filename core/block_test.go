package core

import (
	"fmt"
	"testing"
	"time"

	// "github.com/holiman/uint256"
	"github.com/AnandK-2024/Blockchain/types"
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
