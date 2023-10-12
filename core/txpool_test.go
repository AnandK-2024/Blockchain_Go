package core

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTxMaxLength(t *testing.T) {
	p := NewTxPool()
	// p.Add(utils.NewRandomTransaction(10))
	// assert.Equal(t, 1, p.all.Count())

	p.Add(NewRandomTransaction(10))
	p.Add(NewRandomTransaction(20))
	p.Add(NewRandomTransaction(30))
	tx := NewRandomTransaction(100)
	p.Add(tx)
	assert.Equal(t, 4, p.Len())

}

func TestTxPoolAdd(t *testing.T) {
	p := NewTxPool()
	n := 10

	for i := 1; i <= n; i++ {
		tx := NewRandomTransaction(100)
		p.Add(tx)
		// cannot add twice
		p.Add(tx)

		assert.Equal(t, i, p.Len())

	}
	p.Flush()
	assert.Equal(t, 0, p.Len())
}

func TestSortTransaction(t *testing.T) {
	// create new trasaction pool
	p := NewTxPool()

	// set length of transaction
	txlen := 100
	for i := 0; i < txlen; i++ {
		tx := NewTransaction([]byte(strconv.FormatInt(int64(i), 10)))
		tx.SetFirstSeen(int64(i * rand.Intn(1000)))
		assert.Nil(t, p.Add(tx))
	}
	assert.Equal(t, 100, p.Len())
	/// write test for sort transaction
	// p=NewTxMapSorter(&p.Transactions)
}

// func TestTxPoolMaxLength(t *testing.T) {
// 	maxLen := 10
// 	p := NewTxPool(maxLen)
// 	n := 100
// 	txx := []*core.Transaction{}

// 	for i := 0; i < n; i++ {
// 		tx := util.NewRandomTransaction(100)
// 		p.Add(tx)

// 		if i > n-(maxLen+1) {
// 			txx = append(txx, tx)
// 		}
// 	}

// 	assert.Equal(t, p.all.Count(), maxLen)
// 	assert.Equal(t, len(txx), maxLen)

// 	for _, tx := range txx {
// 		assert.True(t, p.Contains(tx.Hash(core.TxHasher{})))
// 	}
// }

// func TestTxSortedMapFirst(t *testing.T) {
// 	m := NewTxSortedMap()
// 	first := util.NewRandomTransaction(100)
// 	m.Add(first)
// 	m.Add(util.NewRandomTransaction(10))
// 	m.Add(util.NewRandomTransaction(10))
// 	m.Add(util.NewRandomTransaction(10))
// 	m.Add(util.NewRandomTransaction(10))
// 	assert.Equal(t, first, m.First())
// }

// func TestTxSortedMapAdd(t *testing.T) {
// 	m := NewTxSortedMap()
// 	n := 100

// 	for i := 0; i < n; i++ {
// 		tx := util.NewRandomTransaction(100)
// 		m.Add(tx)
// 		// cannot add the same twice
// 		m.Add(tx)

// 		assert.Equal(t, m.Count(), i+1)
// 		assert.True(t, m.Contains(tx.Hash(core.TxHasher{})))
// 		assert.Equal(t, len(m.lookup), m.txx.Len())
// 		assert.Equal(t, m.Get(tx.Hash(core.TxHasher{})), tx)
// 	}

// 	m.Clear()
// 	assert.Equal(t, m.Count(), 0)
// 	assert.Equal(t, len(m.lookup), 0)
// 	assert.Equal(t, m.txx.Len(), 0)
// }

// func TestTxSortedMapRemove(t *testing.T) {
// 	m := NewTxSortedMap()

// 	tx := util.NewRandomTransaction(100)
// 	m.Add(tx)
// 	assert.Equal(t, m.Count(), 1)

// 	m.Remove(tx.Hash(core.TxHasher{}))
// 	assert.Equal(t, m.Count(), 0)
// 	assert.False(t, m.Contains(tx.Hash(core.TxHasher{})))
// }
