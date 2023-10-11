package core

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/AnandK-2024/Blockchain/types"
)

func TestGobTxEncoderAndDecoder(t *testing.T) {
	// // Create a buffer to store the encoded transaction
	// buf := &bytes.Buffer{}
	// privKey := crypto.GeneratePrivatekey()
	// Create a transaction to encode
	tx := &Transaction{
		data:  []byte("Test transaction"),
		value: 100,
		// Set other fields of the transaction as needed
	}
	// if err := tx.sign(&privKey); err != nil {
	// 	t.Errorf("transaction couldn't signed")
	// }

	// // Create a GobTxEncoder with the buffer
	// encoder := NewGobTxEncoder(buf)

	// fmt.Println("encoded transaction:", encoder)

	// // Encode the transaction
	// err := encoder.Encode(tx)
	// if err != nil {
	// 	t.Errorf("Error encoding transaction: %s", err)
	// }

	// // Create a GobTxDecoder with the same buffer
	// decoder := NewGobTxDecoder(buf)

	// // Create a new transaction to decode into
	// decodedTx := &Transaction{}

	// // Decode the transaction
	// err = decoder.Decode(decodedTx)
	// if err != nil {
	// 	t.Errorf("Error decoding transaction: %s", err)
	// }

	// fmt.Println("original tx data and decoded data", tx.data, decodedTx.data)
	// fmt.Println("original tx value and decoded value", tx.value, decodedTx.value)

	// Create a sample transaction

	// Encode the transaction
	encodedTx, err := EncodeTransaction(tx)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the transaction
	decodedTx, err := DecodeTransaction(encodedTx)
	if err != nil {
		log.Fatal(err)
	}

	// Print the original and decoded transactions
	fmt.Println("Original Transaction:", tx)
	fmt.Println("Decoded Transaction:", decodedTx)
	// Compare the original transaction and the decoded transaction
	if !bytes.Equal(tx.data, decodedTx.data) || tx.value != decodedTx.value {
		t.Errorf("Decoded transaction does not match original transaction")
	}
}

func TestGobBlockEncoderAndDecoder(t *testing.T) {
	// Create a buffer to store the encoded block
	buf := &bytes.Buffer{}

	// Create a block to encode
	block := &Block{
		Header: &Header{
			version:       1,
			prevblockHash: types.Hash{},
			DataHash:      types.Hash{},
			Timestamp:     1633968000,
			Height:        1,
		},
		Transactions: []Transaction{},
		// Set other fields of the block as needed
	}

	// Create a GobBlockEncoder with the buffer
	encoder := NewGobBlockEncoder(buf)

	// Encode the block
	err := encoder.Encode(block)
	if err != nil {
		t.Errorf("Error encoding block: %s", err)
	}

	// Create a GobBlockDecoder with the same buffer
	decoder := NewGobBlockDecoder(buf)

	// Create a new block to decode into
	decodedBlock := &Block{}

	// Decode the block
	err = decoder.Decode(decodedBlock)
	if err != nil {
		t.Errorf("Error decoding block: %s", err)
	}

	// Compare the original block and the decoded block
	// You can add more comparisons based on your specific block structure
	if block.Header.version != decodedBlock.Header.version || block.Header.Timestamp != decodedBlock.Header.Timestamp {
		t.Errorf("Decoded block does not match original block")
	}
}
