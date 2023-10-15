package network

import "net"

type NetAddr string


type Transport interface {
	consume() <-chan RPC
	connect(Transport) error
	send(net.Addr, []byte) error
	Addr() net.Addr
	Broadcast([]byte) error
}
