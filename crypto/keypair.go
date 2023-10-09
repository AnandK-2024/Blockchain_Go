package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/AnandK-2024/Blockchain/types"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}
type PublicKey struct {
	key *ecdsa.PublicKey
}

type Signature struct {
	s, r *big.Int
}

// Step 1: Create a random private key
func GeneratePrivatekey() PrivateKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Failed to generate private key:", err)
	}
	return PrivateKey{
		key: privKey,
	}
}

// / Step 2: Make public key from private key
func (p PrivateKey) GeneratePublicKey() PublicKey {

	return PublicKey{
		key: &p.key.PublicKey,
	}

}

func (k *PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(k.key, k.key.X, k.key.Y)
}

func (k *PublicKey) address() types.Address {
	hash := sha256.Sum256(k.ToSlice())
	return types.AddressFromByte(hash[len(hash)-20:])
}

func (p *PrivateKey) SignMessage(hash []byte) (*Signature,error) {
	r, s, err := ecdsa.Sign(rand.Reader, p.key, hash)

	if err != nil {
		fmt.Printf("error signing: %s", err)
		panic(err)
	}

	return &Signature{
		r: r,
		s: s,
	},nil
}

func (Sig *Signature) verify(pub PublicKey, hash []byte) bool{
	return ecdsa.Verify(pub.key,hash,Sig.r,Sig.s)
}
