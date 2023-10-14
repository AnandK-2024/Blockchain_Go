package core

import (
	// "fmt"

	// "fmt"
	"sort"

	"github.com/AnandK-2024/Blockchain/types"
)

// to sort transactions in mempool
type TxMapSorter struct {
	transactions []*Transaction
}

// Len implements sort.Interface.
func (s *TxMapSorter) Len() int {
	return len(s.transactions)
}

// Less implements sort.Interface.
func (s *TxMapSorter) Less(i int, j int) bool {
	return s.transactions[i].firstSeen < s.transactions[j].firstSeen
}

// Swap implements sort.Interface.
func (s *TxMapSorter) Swap(i int, j int) {
	s.transactions[i], s.transactions[j] = s.transactions[j], s.transactions[i]
}

type TxPool struct {
	Transactions map[types.Hash]*Transaction
}

// sort depends on time not on validator fee
func NewTxMapSorter(txMap map[types.Hash]*Transaction) *TxMapSorter {
	Txs := make([]*Transaction, len(txMap))
	i := 0
	for _, tx := range txMap {
		Txs[i] = tx
		i++
	}
	s := &TxMapSorter{Txs}
	sort.Sort(s)
	return s

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

// return all transactions of pool
func (txp *TxPool) Transaction() []*Transaction {
	sortertxs := NewTxMapSorter(txp.Transactions)
	return sortertxs.transactions
}
