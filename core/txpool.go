package core

import (
	// "fmt"

	"github.com/AnandK-2024/Blockchain/types"
)

type TxPool struct {
	Transactions map[types.Hash]*Transaction
}

// create new transaction pool
func NewTxPool() *TxPool {
	return &TxPool{
		Transactions: make(map[types.Hash]*Transaction),
	}

}

// add transaction to txpool
func (txp *TxPool) Add(tx *Transaction) error {
	hash := tx.Hash()
	if txp.Has(hash) {
		// return fmt.Errorf("transacton already added")
		return nil
	}
	txp.Transactions[hash] = tx
	return nil

}

// check tx is excited or not by hash
func (txp *TxPool) Has(hash types.Hash) bool {
	_, ok := txp.Transactions[hash]
	return ok

}

// check the length of tx pool
func (txp *TxPool) Len() int {
	return len(txp.Transactions)

}

// flush the transaction pool
func (txp *TxPool) Flush() {
	txp.Transactions = make(map[types.Hash]*Transaction)
}
