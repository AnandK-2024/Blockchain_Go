package network

import (
	"fmt"
	"time"

	"github.com/AnandK-2024/Blockchain/crypto"
	"github.com/sirupsen/logrus"

	"github.com/AnandK-2024/Blockchain/core"
	// "golang.org/x/text/cases"
)

type serveropts struct {
	Transports []Transport
	blockTime  time.Duration
	privatekey *crypto.PrivateKey
}

type server struct {
	serveropts
	blockTime   time.Duration
	mempool     *core.TxPool
	isValidator bool
	rpcCh       chan RPC
	quitCh      chan struct{}
}

func Newserver(opts serveropts) *server {
	return &server{
		serveropts:  opts,
		blockTime:   opts.blockTime,
		mempool:     core.NewTxPool(),
		isValidator: opts.privatekey != nil,
		rpcCh:       make(chan RPC),
		quitCh:      make(chan struct{}, 1),
	}
}

func (s *server) HandleTransaction(tx *core.Transaction) error {
	if err := tx.Verify(); err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"hash": tx.Hash(),
	}).Info("adding new transaction to mempool")

	if err := s.mempool.Add(tx); err != nil {
		return err
	}
	fmt.Println("transaction added in mempool")

	return nil

}

func (s *server) start() {
	s.initTransport()
	ticker := time.NewTicker(s.blockTime)

	//This line labels a loop with the name free.(using "break free" break all nested loop till free keyword other wise outer loop will contniue excuted in case of nested loop)
free:

	//This line starts an infinite loop. The loop will continue indefinitely until a break statement is encountered.
	for {

		//This line starts a select statement, which allows the goroutine to wait for multiple channel operations simultaneously.
		select {

		//his line listens for a value from the s.rpcCh channel. If a value is received, it is assigned to the variable rpc.
		case rpc := <-s.rpcCh:
			fmt.Println(" +ve", rpc)

		// This line listens for a value from the s.quitCh channel
		case <-s.quitCh:
			break free

		case <-ticker.C:
			if s.isValidator {

				s.CreateNewBlock()
			}
		}

		fmt.Println("server shutdown")
	}
}

func (s *server) CreateNewBlock() {
	
	fmt.Println("creating new block ....")
}

func (s *server) initTransport() {

	// This line starts a loop that iterates over the s.transport slice.
	for _, tr := range s.Transports {

		// excute parrallely with other function
		go func(tr Transport) {

			//This line starts a loop that iterates over the channel returned by the consume method of the tr variable.
			for rpc := range tr.consume() {
				s.rpcCh <- rpc
			}
		}(tr)

	}

}
