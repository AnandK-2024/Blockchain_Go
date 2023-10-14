package core

import (
	"fmt"
	"sync"
	"time"

	"github.com/AnandK-2024/Blockchain/crypto"
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
	// B.lock.Lock()
	// defer B.lock.Unlock()
	if err := B.ValidateBlock(block); err != nil {
		fmt.Println("block verification test failed with given below reason")
		return err
	}
	B.addBlockWithoutValidation(block)

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
func (B *Blockchain) addBlockWithoutValidation(b *Block) error {

	B.headers = append(B.headers, b.Header)
	B.blocks = append(B.blocks, b)
	return nil
}

// add genesis block in blockchain
func (B *Blockchain) AddGensisBlock(genesis *Block) error {
	B.addBlockWithoutValidation(genesis)
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
func (B *Blockchain) GetBlockByHash(hash types.Hash) (*Block, error) {
	B.lock.Lock()
	defer B.lock.Unlock()
	block, ok := B.blockStore[hash]
	if !ok {
		return nil, fmt.Errorf("block with %d hash not found", hash)
	}
	return block, nil
}

// get block by height of blockchain
func (B *Blockchain) GetBlock(height uint32) (*Block, error) {
	// B.lock.Lock()
	// defer B.lock.Unlock()
	return B.blocks[height], nil
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

// check validation of the block before finalize by verifier
func (B *Blockchain) ValidateBlock(b *Block) error {
	fmt.Println("validation of block is starting................")
	// validate block height
	if B.HasBlock(b.Height) {
		return fmt.Errorf("Blockchian already contain block %d", b.Height)
	}

	// validate current height of block in blockchain:
	// current height=height( blockchain )+1
	if b.Height != B.Height()+1 {
		return fmt.Errorf("block height: %d is too high--> current height is: %d", b.Height, B.Height())
	}

	// validate prev header of the block(must present in blockchian)
	prevBlock, err := B.GetBlock(b.Height - 1)
	if err != nil {
		return err
	}

	// validate prevhash of block
	hash := prevBlock.Hash()
	if b.prevblockHash != hash {
		return fmt.Errorf("Hash of previous block: %d is invalid", b.prevblockHash)
	}

	// verify the block signature
	if err := b.Verify(); err != nil {
		return fmt.Errorf("invalid block signature")
	}

	// validate timestamp of block
	// should be less than current time

	if Time := b.Header.Timestamp; Time > int64(time.Now().UnixMicro()) {
		return fmt.Errorf("timestamp of current block must less than curent timestamp")
	}

	// validate merklehashroot of transactions
	hashstring := CalculateMerkleRoot(b.Transactions)
	merklehash, _ := StringToHash(hashstring)
	if merklehash != b.DataHash {
		return fmt.Errorf("invalid merkle hash root ")
	}
	fmt.Println("validation of block successfull")
	return nil
}

// mine block by validator/miner
func (B *Blockchain) Mine(b *Block, privkey *crypto.PrivateKey) (types.Hash, error) {

	// calculate merkle root hash and set to data hash of block header
	if err := b.CalculateMerkleRoot(); err != nil {
		return types.Hash{0}, err
	}

	// set timestamp of block during mine
	b.Timestamp = int64(time.Now().UnixMicro())

	// calculate hash of block and return
	hash := b.Hash()
	b.hash = hash

	//sign the block by miner/validator
	b.Sign(privkey)

	return hash, nil
}
