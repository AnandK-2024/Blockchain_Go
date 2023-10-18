package network

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

// purpose: for local testing of network
// 			responsible for connection peer to servers by TCP/UDP

type LocalTransport struct {
	//net:
	// type Addr interface {
	// 	Network() string // name of the network (for example, "tcp", "udp")
	// 	String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
	// }
	addr      net.Addr
	consumeCh chan RPC                     //represents a channel used for consuming RPC messages.
	lock      sync.RWMutex                 //represents a read-write mutex used for synchronization.
	peers     map[net.Addr]*LocalTransport //represents a map that associates network addresses with corresponding "LocalTransport" instances.
}

// take address and make channel and return locl transport
func NewLocaltransport(addr net.Addr) *LocalTransport {
	return &LocalTransport{
		addr: addr,
		//creates a buffered channel of type RPC with a capacity of 1024
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[net.Addr]*LocalTransport),
	}
}

// return the local tranport channel
func (t *LocalTransport) consume() <-chan RPC {
	return t.consumeCh

}

// connect with local channel
func (t *LocalTransport) connect(tr *LocalTransport) error {
	// trans := tr.(*LocalTransport)
	//acquires a write lock on the "local" mutex of the "LocalTransport" instance.
	t.lock.Lock()
	//This line defers the release of the write lock until the surrounding function returns
	defer t.lock.Unlock()
	t.peers[tr.Addr()] = tr
	return nil

}

func (t *LocalTransport) Addr() net.Addr {
	return t.addr
}

// send message with payload to server

func (t *LocalTransport) sendMessage(to net.Addr, payload []byte) error {
	// acquires a read lock on the "local" mutex of the "LocalTransport" instance.
	t.lock.RLock()
	// defers the release of the read lock until the surrounding function returns.
	defer t.lock.RUnlock()

	// can't send message to itself
	if t.Addr() == to {
		return nil
	}

	peer, ok := t.peers[to]

	if !ok {
		return fmt.Errorf("%s: could not send message to unknown peer %s", t.addr, to)

	}

	//from: sender of the message
	// payload: message that need send through channel to reciever
	peer.consumeCh <- RPC{
		from:    t.addr,
		payload: bytes.NewReader(payload),
	}

	return nil
}

func (t *LocalTransport) Broadcast(payload []byte) error {

	// send message to all peers
	for _, peer := range t.peers {
		err := t.sendMessage(peer.Addr(), payload)
		if err != nil {
			return err
		}

	}
	return nil
}
