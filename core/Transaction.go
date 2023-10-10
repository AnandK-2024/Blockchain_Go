package core

import (
	"fmt"

	"github.com/AnandK-2024/Blockchain/crypto"
)

type Transaction struct {
	data      []byte
	value     uint64
	publickey crypto.PublicKey
	signature *crypto.Signature
	Nonce     uint64
}

func (tx *Transaction) sign(privkey *crypto.PrivateKey) error {
	signature, err := privkey.SignMessage(tx.data)
	if err != nil {
		fmt.Println("unable to sign block with private key")
		return err
	}
	tx.publickey = privkey.GeneratePublicKey()
	tx.signature = signature
	return nil
}

func (tx *Transaction) Verify() error {
	if tx.signature == nil {
		return fmt.Errorf("tx has not signature")
	}
	sig := tx.signature
	sucess := sig.Verify(tx.publickey, tx.data)
	if !sucess {
		return fmt.Errorf("invalid tx signature ")
	}
	return nil
}
