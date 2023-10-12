package core

import (
	"testing"

	"github.com/AnandK-2024/Blockchain/crypto"
	"github.com/stretchr/testify/assert"
)

func TestAccounState(t *testing.T) {
	state := NewAccountState()

	privkey := crypto.GeneratePrivatekey()
	pubkey := privkey.GeneratePublicKey()
	address := pubkey.Address()
	account := state.CreateAccount(address)

	assert.Equal(t, account.address, address)
	assert.Equal(t, account.balance, uint64(0))

	fetchedAccount, err := state.GetAccount(address)
	assert.Nil(t, err)
	assert.Equal(t, fetchedAccount, account)
}

func TestTransferFailInsufficientBalance(t *testing.T) {
	state := NewAccountState()

	privkey := crypto.GeneratePrivatekey()
	pubkey := privkey.GeneratePublicKey()
	addressBob := pubkey.Address()
	accountBob := state.CreateAccount(addressBob)

	privkey = crypto.GeneratePrivatekey()
	pubkey = privkey.GeneratePublicKey()
	addressAlice := pubkey.Address()
	accountAlice := state.CreateAccount(addressAlice)

	accountBob.balance = 99

	amount := uint64(100)
	assert.NotNil(t, state.Transfer(addressBob, addressAlice, amount))
	assert.Equal(t, accountAlice.balance, uint64(0))
}

func TestTransferSuccessEmpyToAccount(t *testing.T) {
	state := NewAccountState()

	privkey := crypto.GeneratePrivatekey()
	pubkey := privkey.GeneratePublicKey()
	addressBob := pubkey.Address()
	accountBob := state.CreateAccount(addressBob)

	privkey = crypto.GeneratePrivatekey()
	pubkey = privkey.GeneratePublicKey()
	addressAlice := pubkey.Address()
	accountAlice := state.CreateAccount(addressAlice)

	accountBob.balance = 100

	amount := uint64(100)
	assert.Nil(t, state.Transfer(addressBob, addressAlice, amount))
	assert.Equal(t, accountAlice.balance, amount)
}
