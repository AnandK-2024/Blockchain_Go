package core

import (
	"fmt"
	"sync"

	"github.com/AnandK-2024/Blockchain/types"
	"github.com/go-kit/log"
)

type Blockchain struct {
	logger log.Logger

	//used to provide synchronization for concurrent access to shared resources.
	lock sync.RWMutex
	// contain list of headers
	headers []*Header
	// contain list of blocks
	blocks []*Block
	// contain blocks with header hash
	blockStore map[types.Hash]*Block

	//contain transaction with tx hash
	txStore map[types.Hash]*Transaction

	// contain Accountstate of users
	accountState *AccountState

	// validator
	validator validator
}

// create new blockchain with genesis block
func NewBlockchian(l log.Logger, genesis *Block, CoinbaseAddr types.Address) (*Blockchain, error) {
	account := NewAccountState()
	account.CreateAccount(CoinbaseAddr)

	bc := &Blockchain{
		headers:      []*Header{},
		blocks:       []*Block{},
		logger:       l,
		blockStore:   make(map[types.Hash]*Block),
		txStore:      make(map[types.Hash]*Transaction),
		accountState: account,
	}

	bc.validator = NewBlockValidator(bc)

	err := bc.AddGensisBlock(genesis)
	return bc, err
}

// get height of the blockchian
func (B *Blockchain) Height() uint32 {
	//When a goroutine calls Lock(), it will block until it can acquire the lock
	B.lock.RLock()

	//to release the lock, allowing other goroutines to acquire it.
	defer B.lock.RUnlock()
	return uint32(len(B.headers) - 1)
}

// add the block in blockchain :
func (B *Blockchain) AddBlock(block *Block) error {
	B.lock.Lock()
	defer B.lock.Unlock()
	if err := B.validator.ValidateBlock(block); err != nil {
		return fmt.Errorf("block verification test failed")
	}
	B.AddBlockWithoutValidation(block)

	return nil
}

// set the validator of blockchain
func (B *Blockchain) SetValidator(v validator) {
	B.validator = v
}

// check block is present or not for given height
func (B *Blockchain) HasBlock(height uint32) bool {
	return height <= B.Height()
}

// add block without validation
func (B *Blockchain) AddBlockWithoutValidation(b *Block) error {

	B.headers = append(B.headers, b.Header)
	return nil
}

// add genesis block in blockchain
func (B *Blockchain) AddGensisBlock(b *Block) error {

	return nil
}

// get header of blockchain
func (B *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height > B.Height() {
		return nil, fmt.Errorf("given height %d is too high ", height)
	}
	B.lock.RLock()
	defer B.lock.RUnlock()
	return B.headers[height], nil
}

// get block by hash in blockchain
func (b *Blockchain) GetBlockByHash(hash types.Hash) (*Block, error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	block, ok := b.blockStore[hash]
	if !ok {
		return nil, fmt.Errorf("block with %d hash not found", hash)
	}
	return block, nil
}

// get block by height of blockchain
func (b *Blockchain) GetBlock(height uint32) (*Block, error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.blocks[height], nil
}

// get transactoin of blockchain by hash
func (b *Blockchain) GetTxByHash(hash types.Hash) (*Transaction, error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	tx, ok := b.txStore[hash]
	if !ok {
		return nil, fmt.Errorf("transaction with given hash %d not found ", hash)
	}
	return tx, nil
}
