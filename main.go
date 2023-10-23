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
	"fmt"
	"net"
	"net/http"
	"time"

	// "github.com/AnandK-2024/Blockchain/api"
	"github.com/AnandK-2024/Blockchain/core"
	"github.com/AnandK-2024/Blockchain/crypto"
	"github.com/AnandK-2024/Blockchain/network"
	"github.com/AnandK-2024/Blockchain/types"
	"github.com/go-kit/log"
)

func main() {
	// bc := NewBlockchian()
	// server := api.Server{
	// 	ListenAddr: ":8080",
	// 	Bc:         bc,
	// }
	// go server.Start()
	// time.Sleep(1 * time.Second)
	ValidatorPrivkey := crypto.GeneratePrivatekey()
	RemoteNode := makeServer("Remote transport", nil, ":7000", nil, "")
	go RemoteNode.Start()

	localNode := makeServer("local transport", &ValidatorPrivkey, ":8000", []string{":7000"}, ":8081")
	go localNode.Start()

	// RemoteNodeB := makeServer("RemoteB transport", nil, ":9000", nil, "")
	// go RemoteNodeB.Start()

	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	LateNode := makeServer("LateNode transport", nil, ":6000", []string{":6000", ":7000"}, "")
	// 	go LateNode.Start()
	// }()
	time.Sleep(1 * time.Second)
	if err := SendTx(ValidatorPrivkey); err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)

	// for i := 0; i < 10; i++ {
	// 	time.Sleep(1 * time.Second)

	// }

	// go tcptester()
	select {}

}

func NewBlockchian() *core.Blockchain {
	privKeycoinbase := crypto.GeneratePrivatekey()
	pubkeycoinbase := privKeycoinbase.GeneratePublicKey()
	bc := newBlockchainWithGenesis(pubkeycoinbase.Address())

	// lenBlocks := 5

	for i := 0; i < 10; i++ {
		block := randomBlock(uint32(i+1), getPrevBlockHash(bc, uint32(i)))
		SignBlocktxs(block, privKeycoinbase)
		block.Sign(&privKeycoinbase)
		bc.Mine(block, &privKeycoinbase)
		fmt.Printf("hash of block:%s\n", block.Hash())
		bc.AddBlock(block)
		time.Sleep(1 * time.Second)

	}
	fmt.Println("height of block:", bc.Height())
	return bc
}

func makeServer(id string, pk *crypto.PrivateKey, Listenaddr string, seednode []string, apilistenaddr string) *network.Server {
	opts := network.Serveropts{
		ApiListenAddr: apilistenaddr,
		ID:            id,
		ListnerAdd:    Listenaddr,
		Privatekey:    pk,
		SeedNodes:     seednode,
		// TCPTransport:  network.NewTCPTransport(":8080", make(chan *network.TCPPeer)),
	}
	s, err := network.Newserver(opts)
	if err != nil {
		panic(err)
	}
	return s
}

func SendTx(FromPrivatekey crypto.PrivateKey) error {
	toprivkey := crypto.GeneratePrivatekey()
	tx := core.NewTransaction(nil)
	tx.Value = 100
	tx.To = toprivkey.GeneratePublicKey()
	if err := tx.Sign(&FromPrivatekey); err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "http://localhost:8081/tx", buf)
	if err != nil {
		return err
	}
	client := http.Client{}
	_, err = client.Do(req)
	fmt.Println("tx req has been sent through api.")
	return err

}

func tcptester() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	data := []byte{0x0, 0x05, 0x0, 0xa, 0x1, 0x0, 0xb, 0x3, 0x0, 0xf, 0x2, 0x0, 0x5}
	// tx := core.NewTransaction(data)
	// privKey := crypto.GeneratePrivatekey()
	// tx.Sign(&privKey)
	// buf := &bytes.Buffer{}
	// if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
	// 	panic(err)
	// }
	// msg := network.NewMessage(network.MessageTypeTx, buf.Bytes())
	conn.Write([]byte(data))
}

func newBlockchainWithGenesis(coinbaseAddr types.Address) *core.Blockchain {
	b := randomBlock(0, types.Randomhash())
	b.BlockHash = b.Hash()
	bc, err := core.NewBlockchian(log.NewNopLogger(), b, coinbaseAddr)
	if err != nil {
		panic(err)
	}
	return bc
}

func getPrevBlockHash(bc *core.Blockchain, height uint32) types.Hash {
	prevBlock, err := bc.GetBlock(height)
	if err != nil {
		panic(err)
	}
	return prevBlock.BlockHash
}

func SignBlocktxs(b *core.Block, privatekey crypto.PrivateKey) {
	for i := 0; i < len(b.Transactions); i++ {
		b.Transactions[i].Sign(&privatekey)
		fmt.Printf("hash of transactions:%s\n", b.Transactions[i].Hash())
	}
}

func randomBlock(height uint32, prevhash types.Hash) *core.Block {
	header := &core.Header{
		Version:       1,
		PrevblockHash: prevhash,
		DataHash:      types.Randomhash(),
		Timestamp:     time.Now().UnixNano(),
		Height:        height,
	}
	txs := &core.Transaction{
		Data:      []byte{0x0, 0x05},
		Timestamp: time.Now().UnixMicro(),
		From:      crypto.PublicKey{},
		To:        crypto.PublicKey{},
	}
	return core.NewBlock(header, []*core.Transaction{txs})
}
