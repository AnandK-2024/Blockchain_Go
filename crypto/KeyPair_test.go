package crypto

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivatekey()

	if privKey.key == nil {
		t.Errorf("Private key generation failed")
	}
}

func TestGeneratePublicKey(t *testing.T) {
	privKey := GeneratePrivatekey()
	pubKey := privKey.GeneratePublicKey()

	if pubKey == nil {
		t.Errorf("Public key generation failed")
	}
}

func TestSignMessage(t *testing.T) {
	privKey := GeneratePrivatekey()
	pubKey := privKey.GeneratePublicKey()

	fmt.Println("key pair:=", privKey, pubKey)

	message := []byte("Hello, world!")
	hash := sha256.Sum256(message)

	fmt.Println("message hash:=", hash)

	signature, err := privKey.SignMessage(hash[:])
	if err != nil {
		t.Errorf("Failed to sign the message")

	}

	fmt.Println("signed hash:", signature)

	if signature.R == nil || signature.S == nil {
		t.Errorf("Failed to sign the message")
	}
}

func TestPublicKeyAddress(t *testing.T) {
	privKey := GeneratePrivatekey()
	pubKey := privKey.GeneratePublicKey()

	address := pubKey.Address()
	fmt.Println("privkey,pubkey,address:", privKey, pubKey, address)

	if len(address) == 0 {
		t.Errorf("Failed to generate address from public key")
	}
}

func TestSignAndVerifyMessage(t *testing.T) {
	privKey := GeneratePrivatekey()
	pubKey := privKey.GeneratePublicKey()

	message := []byte("Hello, world!")
	hash := sha256.Sum256(message)

	signature, err := privKey.SignMessage(hash[:])
	if err != nil {
		t.Errorf("failed to sign message")
	}

	if !signature.Verify(pubKey, hash[:]) {
		t.Error("Message verification failed")
	}
}
