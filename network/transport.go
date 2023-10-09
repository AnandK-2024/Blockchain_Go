package network

import "net"

type NetAddr string

type RPC struct {
	from    net.Addr // sent message fron transport
	payload []byte  // data from server through rpc

}

type Transport interface {
	consume() <-chan RPC
	connect(Transport) error
	send(net.Addr, []byte) error
	Addr() net.Addr
	Broadcast([]byte) error
}
