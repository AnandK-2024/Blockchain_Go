package network

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/AnandK-2024/Blockchain/crypto"
	"github.com/AnandK-2024/Blockchain/types"
	"github.com/go-kit/log"

	// "github.com/sirupsen/logrus"

	"github.com/AnandK-2024/Blockchain/api"
	"github.com/AnandK-2024/Blockchain/core"
)

var defaultBlockTime = 5 * time.Second

type Serveropts struct {
	SeedNodes []string
	// provide api listen at port "apilisenaddr"
	ApiListenAddr string
	// address of tcp connection
	ListnerAdd    string
	TCPTransport  *TCPTransport
	Logger        log.Logger
	blockTime     time.Duration
	ID            string
	Privatekey    *crypto.PrivateKey
	RPCProcessor  RPCProcessor
	RPCDecodeFunc RPCDecodeFunc
}

type Server struct {
	peerCh chan *TCPPeer
	// map of tcp peer
	peer         map[net.Addr]*TCPPeer
	TCPTransport *TCPTransport
	Serveropts
	lock        sync.Mutex
	blockTime   time.Duration
	mempool     *core.TxPool
	chain       *core.Blockchain
	isValidator bool
	rpcCh       chan RPC
	quitCh      chan struct{}
	txchan      chan *core.Transaction
}

// make new server
func Newserver(opts Serveropts) (*Server, error) {
	if opts.blockTime == time.Duration(0) {
		opts.blockTime = defaultBlockTime
	}
	if opts.Logger == nil {
		opts.Logger = log.NewLogfmtLogger(os.Stderr)
		opts.Logger = log.With(opts.Logger, "addr", opts.ID)
	}
	if opts.RPCDecodeFunc == nil {
		opts.RPCDecodeFunc = DefaultRPCDecodeFunc
	}
	var coinbasePubkey = crypto.PublicKey{}
	if opts.Privatekey != nil {
		coinbasePubkey = opts.Privatekey.GeneratePublicKey()
	}

	chain, err := core.NewBlockchian(opts.Logger, genesisBlock(), coinbasePubkey.Address())
	if err != nil {
		return nil, err
	}
	// Channel being used to communicate between the JSON RPC server
	// and the node that will process this message.
	txchan := make(chan *core.Transaction)
	if len(opts.ApiListenAddr) > 0 {
		// start api server
		apiserver := api.NewServer(opts.ApiListenAddr, chain, txchan)
		go apiserver.Start()

		opts.Logger.Log("msg", "JSON api server runningn at =>", "port:", opts.ApiListenAddr)
	}

	peerCh := make(chan *TCPPeer)
	tr := NewTCPTransport(opts.ListnerAdd, peerCh)

	s := &Server{
		TCPTransport: tr,
		peerCh:       peerCh,
		peer:         make(map[net.Addr]*TCPPeer),
		Serveropts:   opts,
		blockTime:    opts.blockTime,
		mempool:      core.NewTxPool(1000),
		chain:        chain,
		isValidator:  opts.Privatekey != nil,
		rpcCh:        make(chan RPC),
		quitCh:       make(chan struct{}, 1),
		txchan:       txchan,
	}
	s.TCPTransport.PeerCh = peerCh

	if s.RPCProcessor == nil {
		s.RPCProcessor = s
	}
	if s.isValidator {
		go s.ValidatorLoop()
	}

	return s, nil
}
func (s *Server) sendGetStatusMessage(peer *TCPPeer) error {
	var (
		getStatusMsg = new(GetStatusMessage)
		buf          = new(bytes.Buffer)
	)
	if err := gob.NewEncoder(buf).Encode(getStatusMsg); err != nil {
		return err
	}

	msg := NewMessage(MessageTypeGetStatus, buf.Bytes())
	return peer.Send(msg.Byte())
}

func (s *Server) ValidatorLoop() {
	ticker := time.NewTicker(s.blockTime)
	s.Logger.Log("msg", "starting validatorLoop..", "Blocktime", s.blockTime)
	for {
		fmt.Println("creating new block")
		if err := s.CreateNewBlock(); err != nil {
			s.Logger.Log("error in ceating new block", err.Error())
		}
		<-ticker.C
	}
}

func (s *Server) bootstrap() {
	for _, addr := range s.SeedNodes {
		fmt.Println("trying to connect to:", addr)
		go func(addr string) {
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				fmt.Printf("\ncouldn't able to connect to %+v", conn)
				return
			}
			s.peerCh <- &TCPPeer{
				connection: conn,
			}

		}(addr)
	}
}

// func (s *Server) HandleTransaction(tx *core.Transaction) error {
// 	if err := tx.Verify(); err != nil {
// 		return err
// 	}
// 	logrus.WithFields(logrus.Fields{
// 		"hash": tx.Hash(),
// 	}).Info("adding new transaction to mempool")

// 	s.mempool.Add(tx)
// 	fmt.Println("transaction added in mempool")

// 	return nil

// }

func (s *Server) Start() {
	// start tcp transport
	go s.TCPTransport.Start()
	time.Sleep(1 * time.Second)
	s.bootstrap()

free:
	for {
		select {
		case tx := <-s.txchan:
			if err := s.ProcessTransaction(tx); err != nil {
				s.Logger.Log("process tx error", err)
			}
		case peer := <-s.peerCh:
			s.peer[peer.connection.RemoteAddr()] = peer
			go peer.readLoop(s.rpcCh)
			if err := s.sendGetStatusMessage(peer); err != nil {
				s.Logger.Log("err", err)
				continue
			}

			s.Logger.Log("msg", "peer added to the server", "outgoing", peer.Outgoing, "addr", peer.connection.RemoteAddr())

		case rpc := <-s.rpcCh:
			msg, err := s.RPCDecodeFunc(rpc)
			if err != nil {
				s.Logger.Log("RPC error", rpc)
				continue
			}
			if err := s.RPCProcessor.ProcessMessage(msg); err != nil {
				if err != core.ErrBlockKnown {
					s.Logger.Log("error", err)
				}
			}

		case <-s.quitCh:
			break free
		}

	}
	fmt.Println("server is shuting down")
}

// process all type of data
func (s *Server) ProcessMessage(msg *DecodedMessage) error {
	switch t := msg.Data.(type) {
	case *core.Transaction:
		return s.ProcessTransaction(t)
	case *core.Block:
		return s.ProcessBlock(t)
		// if s.isValidator {
		// }
	case *BlockMessage:
		return s.ProcessBlockMessage(msg.From, t)
	case *GetStatusMessage:
		s.processGetStatusMessage(msg.From, t)
	case *StatusMessage:
	case *GetBlocksMessage:

	}

	return nil

}

func (s *Server) processGetStatusMessage(from net.Addr, data *GetStatusMessage) error {
	s.Logger.Log("msg", "recieve get status message", "from", from)
	statusMessage := &StatusMessage{
		ID:     s.ID,
		Height: s.chain.Height(),
	}
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(statusMessage); err != nil {
		return err
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	peer, ok := s.peer[from]
	if !ok {
		return fmt.Errorf("peer %s not known", peer.connection.RemoteAddr())
	}
	msg := NewMessage(MessageTypeStatus, buf.Bytes())
	return peer.Send(msg.Byte())

}

func (s *Server) ProcessBlockMessage(from net.Addr, data *BlockMessage) error {
	fmt.Println("!!!!!block recieved!!!!!")
	s.Logger.Log("msg", "!!!!!block recieved!!!!!", "from", from)
	// return index and value of block
	for _, block := range data.Blocks {
		if err := s.chain.AddBlock(block); err != nil {
			s.Logger.Log("error", err.Error())
			return err
		}

	}
	return nil

}

// process block in mempool
func (s *Server) ProcessBlock(b *core.Block) error {
	// add block in blockchain
	if err := s.chain.AddBlock(b); err != nil {
		s.Logger.Log("error", err.Error())
		return err
	}

	// broadcast the block
	go s.BroadcastBlock(b)
	return nil
}

// process transaction in mempool
func (s *Server) ProcessTransaction(tx *core.Transaction) error {
	// fmt.Println("start ProcessTransaction through recieve data from api")
	// find hash of transaction
	hash := tx.Hash()
	// check mempool contains this tx or not. if tx not avilable in mempool then process this transaction
	if ok := s.mempool.Contain(hash); ok {
		// transaction already in mempool
		return nil
	}
	// verify the transaction
	if err := tx.Verify(); err != nil {
		// transaction not verified
		return err
	}
	s.Logger.Log(
		"msg", "adding new tx into mempool",
		"hash", hash,
		"mempoolPending", s.mempool.PendingCount(),
	)
	// broadcast the transactions
	go s.BroadcastTx(tx)
	s.mempool.Add(tx)
	return nil
}

// broadcast the data to all peers
func (s *Server) Broadcast(payload []byte) error {

	// Lock() is used to acquire a write lock, which means it allows only one thread to access the shared resource exclusively.
	s.lock.Lock()
	defer s.lock.Unlock()
	// broadcast the message to all peers
	for netAddr, peer := range s.peer {
		if err := peer.Send(payload); err != nil {
			fmt.Printf("peer send error ==> addr %s and [err: %s]\n", netAddr, err)
		}
	}
	return nil

}

// broadcast transactions in network
func (s *Server) BroadcastTx(tx *core.Transaction) error {
	// enocde the block
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}
	//convert blocks in msg struct
	msg := NewMessage(MessageTypeTx, buf.Bytes())

	// broadcast the msg bytes
	return s.Broadcast(msg.Byte())

}

// broadcast block in network
func (s *Server) BroadcastBlock(b *core.Block) error {
	fmt.Println("block broadcasting.........................................")
	// enocde the block
	buf := &bytes.Buffer{}
	if err := b.Encode(core.NewGobBlockEncoder(buf)); err != nil {
		return err
	}
	//convert blocks in msg struct
	msg := NewMessage(MessageTypeBlock, buf.Bytes())

	// broadcast the msg bytes
	return s.Broadcast(msg.Byte())

}

// create new block
func (s *Server) CreateNewBlock() error {
	// get current header
	prevHeader, err := s.chain.GetHeader(s.chain.Height())
	// currentHeader.Height
	if err != nil {
		return err
	}

	// get all pending from Transaction pool
	PendingTx := s.mempool.Pending()

	// get block from previous header and pending transaction
	block, err := core.NewBlockFromPrevHeader(*prevHeader, PendingTx)
	if err != nil {
		return nil
	}
	// mine the block
	s.chain.Mine(block, s.Privatekey)

	// sign the block
	block.Sign(s.Privatekey)

	//add block in blockchain
	s.lock.Lock()
	defer s.lock.Unlock()
	if err := s.chain.AddBlock(block); err != nil {
		return err
	}

	//clear pending tx pool
	// pending pool of tx should only reflect on validator nodes.
	// for other validators still have these pending transactions
	s.mempool.ClearPending()

	// broadcast block
	go s.BroadcastBlock(block)

	// return nil
	return nil
}

// create genesis block
func genesisBlock() *core.Block {
	header := &core.Header{
		Version:   uint32(1),
		DataHash:  types.Randomhash(),
		Height:    0,
		Timestamp: time.Now().Unix(),
	}

	privk := crypto.GeneratePrivatekey()
	tx := core.NewCompleteTx(privk.GeneratePublicKey(), uint64(100), nil, uint64(1))
	tx.Sign(&privk)
	b := core.NewBlock(header, []*core.Transaction{tx})
	// b.Transactions = append(b.Transactions, tx)
	if err := b.Sign(&privk); err != nil {
		panic(err)
	}
	b.SetHash()
	return b
}
