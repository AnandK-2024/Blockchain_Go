package util

import (
	"github.com/AnandK-2024/Blockchain/core"
	"github.com/AnandK-2024/Blockchain/types"
)

// generate random byte of size
func RandomBytes(size int) []byte {
	return types.RandomByte(size)
}

func Randomhash() types.Hash {
	return types.Randomhash()
}

func NewRandomTransaction(datasize int) *core.Transaction {
	return core.NewTransaction(RandomBytes(datasize))
}

// func RandomBlock(height uint32) *core.Block {
// 	header := &Header{
// 		version:       1,
// 		prevblockHash: types.Randomhash(),
// 		DataHash:      types.Randomhash(),
// 		Timestamp:     time.Now().UnixNano(),
// 		Height:        height,
// 	}
// 	txs := Transaction{
// 		data: []byte("Anand --> bob: 5ETH"),
// 	}
// 	return NewBlock(header, []Transaction{txs})
// }
