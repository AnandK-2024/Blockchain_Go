package core

// import (
// 	"fmt"
// 	"reflect"
// 	"testing"

// 	"github.com/AnandK-2024/Blockchain/crypto"
// )

// func TestTransactionEncodingDecoding(t *testing.T) {
// 	// privKey:=crypto.GeneratePrivatekey()
// 	// Create a sample transaction
// 	tx := &Transaction{
// 		data:      []byte("sample data"),
// 		value:     100,
// 		from:      crypto.PublicKey{},
// 		signature: &crypto.Signature{},
// 		to:        crypto.PublicKey{},
// 		Nonce:     123,
// 		firstSeen: 1634720000,
// 	}

// 	// Encode the transaction
// 	var buf bytes.Buffer
// 	encoder := NewGobTxEncoder(&buf)
// 	err := encoder.Encode(tx)
// 	if err != nil {
// 		t.Errorf("Error encoding transaction: %s", err)
// 	}

// 	// Decode the transaction
// 	decoder := NewGobTxDecoder(&buf)
// 	decodedTx := &Transaction{}
// 	err = decoder.Decode(decodedTx)
// 	if err != nil {
// 		t.Errorf("Error decoding transaction: %s", err)
// 	}

// 	fmt.Println(tx, "decoded transaction\n", decodedTx)
// 	// Compare the original and decoded transactions
// 	// if !reflect.DeepEqual(tx, decodedTx) {
// 	// 	t.Errorf("Decoded transaction does not match the original")
// 	// }
// }

// func TestBlockEncodingDecoding(t *testing.T) {
// 	// Create a sample block
// 	block := &Block{
// 		Header: &Header{
// 			Version:       1,
// 			prevblockHash: types.Hash{},
// 			DataHash:      types.Hash{},
// 			Timestamp:     1634720000,
// 			Height:        1,
// 		},
// 		Transactions: []*Transaction{},
// 		validator:    crypto.PublicKey{},
// 		signature:    &crypto.Signature{},
// 		hash:         types.Hash{},
// 	}

// 	// Encode the block
// 	var buf bytes.Buffer
// 	encoder := NewGobBlockEncoder(&buf)
// 	err := encoder.Encode(block)
// 	if err != nil {
// 		t.Errorf("Error encoding block: %s", err)
// 	}

// 	// Decode the block
// 	decoder := NewGobBlockDecoder(&buf)
// 	decodedBlock := &Block{}
// 	err = decoder.Decode(decodedBlock)
// 	if err != nil {
// 		t.Errorf("Error decoding block: %s", err)
// 	}

// 	// Compare the original and decoded blocks
// 	if !reflect.DeepEqual(block, decodedBlock) {
// 		t.Errorf("Decoded block does not match the original")
// 	}
// }

// func TestTransactionEncodingDecoding(t *testing.T) {
// 	// Create a sample transaction
// 	tx := &Transaction{
// 		data:      []byte("sample data"),
// 		value:     100,
// 		from:      crypto.PublicKey{},
// 		signature: &crypto.Signature{},
// 		to:        crypto.PublicKey{},
// 		Nonce:     123,
// 		firstSeen: 1634720000,
// 	}

// 	// Encode the transaction
// 	encoded, err := EncodeTransaction(tx)
// 	if err != nil {
// 		t.Errorf("Error encoding transaction: %s", err)
// 	}

// 	// Decode the transaction
// 	decodedTx, err := DecodeTransaction(encoded)
// 	if err != nil {
// 		t.Errorf("Error decoding transaction: %s", err)
// 	}

// 	fmt.Println(tx,decodedTx)

// 	// Compare the original and decoded transactions
// 	if !reflect.DeepEqual(tx, decodedTx) {
// 		t.Errorf("Decoded transaction does not match the original")
// 	}
// }
