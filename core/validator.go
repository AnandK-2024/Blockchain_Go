package core

import (
	"fmt"
	"time"
)

type validator interface {
	ValidateBlock(b *Block) error
}

type Blockvalidator struct {
	B *Blockchain
}

// set block validator in blockchain
func NewBlockValidator(b *Blockchain) * Blockvalidator{
	return &Blockvalidator{
		B: b,
	}
}
// validator will validate the block before finalize
func (v *Blockvalidator) ValidateBlock(b *Block) error {
	// validate block height
	if v.B.HasBlock(b.Height) {
		return fmt.Errorf("Blockchian already contain block %d", b.Height)
	}

	// validate current height of block in blockchain:
	// current height=height( blockchain )+1
	if b.Height != v.B.Height()+1 {
		return fmt.Errorf("block height: %d is too high--> current height is: %d", b.Height, v.B.Height())
	}

	// validate prev header of the block(must present in blockchian)
	prevHeader, err := v.B.GetHeader(b.Height - 1)
	if err != nil {
		return err
	}

	// validate prevhash of block
	hash := Hash(prevHeader)
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

	return nil
}
