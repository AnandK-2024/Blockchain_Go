package core

import (
	"fmt"
	"testing"

	"github.com/AnandK-2024/Blockchain/crypto"
	"github.com/stretchr/testify/assert"
	// "github.com/holiman/uint256"
)

func TestTransaction(t *testing.T) {
	privkey := crypto.GeneratePrivatekey()
	tx := Transaction{
		data: []byte("Anand-->bob: 10ETH"),
	}
	err:=tx.sign(&privkey)
	fmt.Println("transaction signature", tx.signature)
	assert.Nil(t, err)
	assert.NotNil(t,tx.signature)
}
