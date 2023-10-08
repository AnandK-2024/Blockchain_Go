package network

import (
	"fmt"
	"sync"
)

// purpose: for local testing of network
// 			responsible for connection peer to servers by TCP/UDP

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC                    //represents a channel used for consuming RPC messages.
	local     sync.RWMutex                //represents a read-write mutex used for synchronization.
	peers     map[NetAddr]*LocalTransport //represents a map that associates network addresses with corresponding "LocalTransport" instances.
}

// take address and make channel and return locl transport
func NewLocaltransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

// return the local tranport channel
func (t *LocalTransport) consume() <-chan RPC {
	return t.consumeCh

}

// connect with local channel
func (t *LocalTransport) connect(tr *LocalTransport) error {
	//acquires a write lock on the "local" mutex of the "LocalTransport" instance.
	t.local.Lock()
	//This line defers the release of the write lock until the surrounding function returns
	defer t.local.Unlock()
	t.peers[t.Addr()] = tr
	return nil

}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}

// send message with payload to server

func (t *LocalTransport) sendMessage(to NetAddr, payload []byte) error {
	// acquires a read lock on the "local" mutex of the "LocalTransport" instance.
	t.local.RLock()
	// defers the release of the read lock until the surrounding function returns.
	defer t.local.RUnlock()

	peer, ok := t.peers[to]

	if !ok {
		return fmt.Errorf("%s: couldn't send message to %s ", t.addr, to)

	}
	peer.consumeCh <- RPC{
		from:    to,
		payload: payload,
	}

	return nil
}
