package core

import (
	"fmt"
	"testing"

	"github.com/AnandK-2024/Blockchain/crypto"
	"github.com/AnandK-2024/Blockchain/types"
	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
)

func TestSendNativeTransferTamper(t *testing.T) {
	privKeycoinbase := crypto.GeneratePrivatekey()
	pubkeycoinbase := privKeycoinbase.GeneratePublicKey()
	bc := newBlockchainWithGenesis(t, pubkeycoinbase.Address())
	fmt.Println(bc)
	signer := crypto.GeneratePrivatekey()

	block := randomBlock(t, uint32(1), types.Randomhash())
	assert.Nil(t, block.Sign(&signer))

	privKeyBob := crypto.GeneratePrivatekey()
	privKeyAlice := crypto.GeneratePrivatekey()
	amount := uint64(100)
	privkey := privKeyBob.GeneratePublicKey()

	accountBob := bc.accountState.CreateAccount(privkey.Address())
	accountBob.balance = amount

	tx := NewTransaction([]byte{})
	tx.From = privKeyBob.GeneratePublicKey()
	tx.To = privKeyAlice.GeneratePublicKey()
	tx.Value = amount
	tx.Sign(&privKeyBob)
	tx.Hash()

	hackerPrivKey := crypto.GeneratePrivatekey()
	tx.To = hackerPrivKey.GeneratePublicKey()

	block.AddTransaction(tx)
	privkey = hackerPrivKey.GeneratePublicKey()
	_, err := bc.accountState.GetAccount(privkey.Address())
	assert.Equal(t, err, ErrorAccountNotFound)
}

func TestSendNativeTransferInsuffientBalance(t *testing.T) {
	privKeycoinbase := crypto.GeneratePrivatekey()
	pubkeycoinbase := privKeycoinbase.GeneratePublicKey()
	bc := newBlockchainWithGenesis(t, pubkeycoinbase.Address())
	signer := crypto.GeneratePrivatekey()

	block := randomBlock(t, uint32(1), types.Randomhash())
	assert.Nil(t, block.Sign(&signer))

	privKeyBob := crypto.GeneratePrivatekey()
	privKeyAlice := crypto.GeneratePrivatekey()
	pubkeyAlice := privKeyAlice.GeneratePublicKey()
	amount := uint64(100)
	pubkeyBob := privKeyBob.GeneratePublicKey()
	accountBob := bc.accountState.CreateAccount(pubkeyBob.Address())
	accountBob.balance = uint64(99)

	tx := NewTransaction([]byte{})
	tx.From = privKeyBob.GeneratePublicKey()
	tx.To = privKeyAlice.GeneratePublicKey()
	tx.Value = amount
	tx.Sign(&privKeyBob)
	tx.Hash()
	fmt.Println("alice => \n", pubkeyAlice)
	fmt.Println("bob => \n", pubkeyBob)

	block.AddTransaction(tx)
	// assert.Nil(t, bc.AddBlock(block))

	_, err := bc.accountState.GetAccount(pubkeyAlice.Address())
	assert.NotNil(t, err)

	hash := tx.Hash()
	_, err = bc.GetTxByHash(hash)
	assert.NotNil(t, err)
}

func TestAddBlock(t *testing.T) {
	privKeycoinbase := crypto.GeneratePrivatekey()
	pubkeycoinbase := privKeycoinbase.GeneratePublicKey()
	bc := newBlockchainWithGenesis(t, pubkeycoinbase.Address())

	// lenBlocks := 5
	block := randomBlock(t, uint32(1), getPrevBlockHash(t, bc, uint32(0)))
	SignBlocktxs(t, block, privKeycoinbase)
	bc.Mine(block, &privKeycoinbase)
	// block.Sign(&privKeycoinbase)
	fmt.Println("transation of block:", block.Transactions)
	fmt.Println("merkel hash root ", block.DataHash, CalculateMerkleRoot(block.Transactions))
	assert.Nil(t, bc.AddBlock(block))
	// for i := 0; i < lenBlocks; i++ {

	// }

	// assert.Equal(t, bc.Height(), uint32(lenBlocks))
	// assert.Equal(t, len(bc.headers), lenBlocks+1)
	// assert.NotNil(t, bc.AddBlock(randomBlock(t, 89, types.Hash{})))
}

func TestNewBlockchain(t *testing.T) {
	privKeycoinbase := crypto.GeneratePrivatekey()
	pubkeycoinbase := privKeycoinbase.GeneratePublicKey()
	bc := newBlockchainWithGenesis(t, pubkeycoinbase.Address())
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHasBlock(t *testing.T) {
	privKeycoinbase := crypto.GeneratePrivatekey()
	pubkeycoinbase := privKeycoinbase.GeneratePublicKey()
	bc := newBlockchainWithGenesis(t, pubkeycoinbase.Address())
	assert.True(t, bc.HasBlock(0))
	assert.False(t, bc.HasBlock(1))
	assert.False(t, bc.HasBlock(100))
}

func TestGetBlock(t *testing.T) {
	privKeycoinbase := crypto.GeneratePrivatekey()
	pubkeycoinbase := privKeycoinbase.GeneratePublicKey()
	bc := newBlockchainWithGenesis(t, pubkeycoinbase.Address())
	lenBlocks := 100

	for i := 0; i < lenBlocks; i++ {
		block := randomBlock(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i)))
		SignBlocktxs(t, block, privKeycoinbase)
		bc.Mine(block, &privKeycoinbase)
		assert.Nil(t, bc.AddBlock(block))

		fetchedBlock, err := bc.GetBlock(block.Height)
		assert.Nil(t, err)
		assert.Equal(t, fetchedBlock, block)
	}
}

func TestGetHeader(t *testing.T) {
	privKeycoinbase := crypto.GeneratePrivatekey()
	pubkeycoinbase := privKeycoinbase.GeneratePublicKey()
	bc := newBlockchainWithGenesis(t, pubkeycoinbase.Address())
	lenBlocks := 1000

	for i := 0; i < lenBlocks; i++ {
		block := randomBlock(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i)))
		SignBlocktxs(t, block, privKeycoinbase)
		bc.Mine(block, &privKeycoinbase)
		assert.Nil(t, bc.AddBlock(block))
		header, err := bc.GetHeader(block.Height)
		assert.Nil(t, err)
		assert.Equal(t, header, block.Header)
	}
}

func TestAddBlockToHigh(t *testing.T) {
	privKeycoinbase := crypto.GeneratePrivatekey()
	pubkeycoinbase := privKeycoinbase.GeneratePublicKey()
	bc := newBlockchainWithGenesis(t, pubkeycoinbase.Address())
	block := randomBlock(t, 1, getPrevBlockHash(t, bc, uint32(0)))
	SignBlocktxs(t, block, privKeycoinbase)
	bc.Mine(block, &privKeycoinbase)
	assert.Nil(t, bc.AddBlock(block))
	assert.NotNil(t, bc.AddBlock(randomBlock(t, 3, types.Hash{})))
}

func newBlockchainWithGenesis(t *testing.T, coinbaseAddr types.Address) *Blockchain {
	b := randomBlock(t, 0, types.Randomhash())
	b.BlockHash = b.Hash()
	bc, err := NewBlockchian(log.NewNopLogger(), b, coinbaseAddr)
	assert.Nil(t, err)

	return bc
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevBlock, err := bc.GetBlock(height)
	assert.Nil(t, err)
	return prevBlock.BlockHash
}

func SignBlocktxs(t *testing.T, b *Block, privatekey crypto.PrivateKey) {
	for i := 0; i < len(b.Transactions); i++ {
		b.Transactions[i].Sign(&privatekey)
	}
}
