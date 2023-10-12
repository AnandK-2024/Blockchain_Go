// package main

// import (
// 	"fmt"
// 	// "github.com/AnandK-2024/Blockchain/network"
// )

// func main() {
// 	fmt.Println("Hello, World!")
// 	// Create a TCP address server-A
// 	// tcpAddrA := &net.TCPAddr{
// 	// 	IP:   net.ParseIP("192.0.2.1"),
// 	// 	Port: 80,
// 	// }
// 	// transport := network.NewLocaltransport(tcpAddrA)

// 	// // Create a TCP address serverB
// 	// tcpAddrB := &net.TCPAddr{
// 	// 	IP:   net.ParseIP("10.0.0.1"),
// 	// 	Port: 8080,
// 	// }
// 	// remote := network.NewLocaltransport(tcpAddrB)
// 	// transport.connect(remote)
// 	// remote.connect(transport)

// 	// opts := network.serveropts{
// 	// 	transport: []network.Transport{transport},
// 	// }
// 	// s := network.Newserver(opts)
// 	// s.start()

// }

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/AnandK-2024/Blockchain/core"
)

// type Transaction struct {
// 	Data  []byte
// 	Value uint64
// 	// From      crypto.PublicKey
// 	// Signature *crypto.Signature
// 	Nonce uint64
// }

// EncodeTransaction encodes a transaction into a byte slice using the gob encoding format.
func EncodeTransaction(tx *core.Transaction) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(tx)
	if err != nil {
		return nil, fmt.Errorf("error encoding transaction: %s", err)
	}
	return buf.Bytes(), nil
}

// DecodeTransaction decodes a transaction from a byte slice using the gob encoding format.
func DecodeTransaction(data []byte) (*core.Transaction, error) {
	var tx core.Transaction
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&tx)
	if err != nil {
		return nil, fmt.Errorf("error decoding transaction: %s", err)
	}
	return &tx, nil
}

func main() {
	// Create a sample transaction
	tx := &core.Transaction{
		// Data:  []byte("Test transaction"),
		// Value: 100,
		// Set other fields of the transaction as needed
	}

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
}
