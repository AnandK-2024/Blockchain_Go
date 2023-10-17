package core

import (
	// "fmt"

	// "fmt"

	"fmt"
	"sync"

	"github.com/AnandK-2024/Blockchain/types"
)

type TxPool struct {
	// repersent all transaction in tx pool
	all *TxMapSorter
	// pending repersent pending transaction in tx pool
	pending *TxMapSorter
	// maxLength repersent max length capacity of pool handle transaction
	maxLength int
}

// create new transaction pool
func NewTxPool(maxLength int) *TxPool {
	return &TxPool{
		all:       newTxMapSorter(),
		pending:   newTxMapSorter(),
		maxLength: maxLength,
	}
}

// add transaction in tx pool
func (txp *TxPool) Add(tx *Transaction) {

	// first remove the oldest transaction from pending pool
	if txp.pending.count() == txp.maxLength {
		oldesttx := txp.pending.first()
		txp.pending.remove(oldesttx.Hash())
	}
	// check tx are already exist or not
	if ok := txp.pending.Contain(tx.Hash()); !ok {
		txp.pending.add(tx)
		txp.all.add(tx)
	}
}

// check transaction or avilable or not in tx pool
func (txp *TxPool) Contain(hash types.Hash) bool {
	return txp.all.Contain(hash)
}

// get all pending transactions
func (txp *TxPool) Pending() []*Transaction {
	return txp.pending.TxList.Data
}

// count all pending transaction
func (txp *TxPool) PendingCount() int {
	return txp.pending.count()
}

// clear all pending transaction
func (txp *TxPool) ClearPending() {
	txp.pending.clear()
}

// to sort transactions in mempool
type TxMapSorter struct {
	lock         sync.RWMutex
	transactions map[types.Hash]*Transaction
	TxList       *types.List[*Transaction]
}

func newTxMapSorter() *TxMapSorter {
	return &TxMapSorter{
		lock:         sync.RWMutex{},
		transactions: make(map[types.Hash]*Transaction),
		TxList:       types.NewList[*Transaction](),
	}
}

// get first transaction
func (s *TxMapSorter) first() *Transaction {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.TxList.Get(0)
}

func (s *TxMapSorter) add(tx *Transaction) {
	hash := tx.Hash()
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.transactions[hash]
	if !ok {
		s.transactions[hash] = tx
		s.TxList.Insert(tx)
	}
}

func (s *TxMapSorter) get(hash types.Hash) (*Transaction, error) {
	tx, ok := s.transactions[hash]
	s.lock.RLock()
	defer s.lock.RUnlock()
	if !ok {
		fmt.Errorf("transaction not found for given %d hash:", hash)
	}
	return tx, nil
}

func (s *TxMapSorter) remove(hash types.Hash) {
	tx, ok := s.transactions[hash]
	s.lock.Lock()
	defer s.lock.Unlock()
	if ok {
		delete(s.transactions, hash)
		index := s.TxList.GetIndex(tx)
		s.TxList.Pop(index)
	}
}

func (s *TxMapSorter) count() int {
	return s.TxList.Len()
}

func (s *TxMapSorter) Contain(hash types.Hash) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if _, ok := s.transactions[hash]; ok {
		return true
	}
	return false
}

func (s *TxMapSorter) clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.transactions = make(map[types.Hash]*Transaction)
	s.TxList.Clear()
}
