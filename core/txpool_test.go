package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTxPool(t *testing.T) {
	// Create a new transaction pool

	txp := NewTxPool(10)

	// Create a sample transaction
	tx := NewTransaction([]byte("Anand-->bob:10ETH"))

	// Add the transaction to the pool
	txp.Add(tx)

	// Check if the transaction is contained in the pool
	if !txp.Contain(tx.Hash()) {
		t.Error("Expected transaction to be contained in the pool, but it is not")
	}

	// Get all pending transactions
	pending := txp.Pending()
	if len(pending) != 1 {
		t.Errorf("Expected 1 pending transaction, but got %d", len(pending))
	}

	// Clear all pending transactions
	txp.ClearPending()
	pending = txp.Pending()
	if len(pending) != 0 {
		t.Errorf("Expected 0 pending transactions after clearing, but got %d", len(pending))
	}
}

func TestTxMaxLength(t *testing.T) {
	p := NewTxPool(1)
	p.Add(NewRandomTransaction(10))
	// assert.Equal(t, 1, p.all.Count())

	p.Add(NewRandomTransaction(10))
	p.Add(NewRandomTransaction(10))
	p.Add(NewRandomTransaction(10))
	tx := NewRandomTransaction(100)
	p.Add(tx)
	assert.Equal(t, 1, p.all.count())
	assert.True(t, p.Contain(tx.Hash()))
}

func TestTxPoolAdd(t *testing.T) {
	p := NewTxPool(11)
	n := 10

	for i := 1; i <= n; i++ {
		tx := NewRandomTransaction(100)
		p.Add(tx)
		// cannot add twice
		p.Add(tx)

		assert.Equal(t, i, p.PendingCount())
		assert.Equal(t, i, p.pending.count())
		assert.Equal(t, i, p.all.count())
	}
}

func TestTxPoolMaxLength(t *testing.T) {
	maxLen := 10
	p := NewTxPool(maxLen)
	n := 100
	txx := []*Transaction{}

	for i := 0; i < n; i++ {
		tx := NewRandomTransaction(100)
		p.Add(tx)

		if i > n-(maxLen+1) {
			txx = append(txx, tx)
		}
	}

	assert.Equal(t, p.all.count(), maxLen)
	assert.Equal(t, len(txx), maxLen)

	for _, tx := range txx {
		assert.True(t, p.Contain(tx.Hash()))
	}
}

func TestTxSortedMapFirst(t *testing.T) {
	m := newTxMapSorter()
	first := NewRandomTransaction(100)
	m.add(first)
	m.add(NewRandomTransaction(10))
	m.add(NewRandomTransaction(10))
	m.add(NewRandomTransaction(10))
	m.add(NewRandomTransaction(10))
	assert.Equal(t, first, m.first())
}

func TestTxSortedMapAdd(t *testing.T) {
	m := newTxMapSorter()
	n := 100

	for i := 0; i < n; i++ {
		tx := NewRandomTransaction(100)
		m.add(tx)
		// cannot add the same twice
		m.add(tx)

		assert.Equal(t, m.count(), i+1)
		assert.True(t, m.Contain(tx.Hash()))
		assert.Equal(t, len(m.transactions), m.TxList.Len())
		txi, _ := m.get(tx.Hash())
		assert.Equal(t, txi, tx)
	}

	m.clear()
	assert.Equal(t, m.count(), 0)
	assert.Equal(t, len(m.transactions), 0)
	assert.Equal(t, m.TxList.Len(), 0)
}

func TestTxSortedMapRemove(t *testing.T) {
	m := newTxMapSorter()

	tx := NewRandomTransaction(100)
	m.add(tx)
	assert.Equal(t, m.count(), 1)

	m.remove(tx.Hash())
	assert.Equal(t, m.count(), 0)
	assert.False(t, m.Contain(tx.Hash()))
}
