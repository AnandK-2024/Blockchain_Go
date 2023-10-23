package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/AnandK-2024/Blockchain/types"
	// "github.com/gogo/protobuf/test/data"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}
type PublicKey []byte

type Signature struct {
	S *big.Int
	R *big.Int
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
	return elliptic.MarshalCompressed(p.key.PublicKey, p.key.PublicKey.X, p.key.PublicKey.Y)

}

// func (k *PublicKey) ToSlice() []byte {
// 	return elliptic.MarshalCompressed(k.key, k.key.X, k.key.Y)
// }

func (k PublicKey) Address() types.Address {
	hash := sha256.Sum256(k)
	return types.AddressFromByte(hash[len(hash)-20:])
}

func (p PrivateKey) SignMessage(hash []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, p.key, hash)

	if err != nil {
		fmt.Printf("error signing: %s", err)
		panic(err)
	}

	return &Signature{
		R: r,
		S: s,
	}, nil
}

func (Sig Signature) Verify(pub PublicKey, hash []byte) bool {
	x, y := elliptic.UnmarshalCompressed(elliptic.P256(), pub)
	key := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	return ecdsa.Verify(key, hash, Sig.R, Sig.S)
}

func (Sig Signature) String() string {
	b := append(Sig.S.Bytes(), Sig.R.Bytes()...)
	return hex.EncodeToString(b)
}
