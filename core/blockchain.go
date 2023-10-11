package core

import (
	"fmt"
	"sync"
)

type Blockchain struct {
	//used to provide synchronization for concurrent access to shared resources.
	lock      sync.RWMutex
	header    []*Header
	validator validator
}

func (B *Blockchain) Height() uint32 {
	//When a goroutine calls Lock(), it will block until it can acquire the lock
	B.lock.RLock()

	//to release the lock, allowing other goroutines to acquire it.
	defer B.lock.RUnlock()
	return uint32(len(B.header) - 1)
}

// func (B *Blockchain) GetBlock(height uint32) *Block {

// }

// func (B *Blockchain) GetBlockHash(height uint32) types.Hash {
// 	if height > B.Height() {
// 		fmt.Errorf("given height %d is too high ", height)
// 	}
// 	return
// }

func (B *Blockchain) AddBlock(block *Block) error {
	B.lock.Lock()
	defer B.lock.Unlock()
	if err := B.validator.ValidateBlock(block); err != nil {
		return fmt.Errorf("block verification test failed")
	}
	B.AddBlockWithoutValidation(block)

	return nil
}

func (B *Blockchain) SetValidator(v validator) {
	B.validator = v
}

func (B *Blockchain) HasBlock(height uint32) bool {
	return height <= B.Height()
}

func (B *Blockchain) AddBlockWithoutValidation(b *Block) error {

	B.header = append(B.header, b.Header)
	return nil
}

func (B *Blockchain) AddGensisBlock(b *Block) error {

	return nil
}

func (B *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height > B.Height() {
		return nil, fmt.Errorf("given height %d is too high ", height)
	}
	B.lock.RLock()
	defer B.lock.RUnlock()
	return B.header[height], nil
}
