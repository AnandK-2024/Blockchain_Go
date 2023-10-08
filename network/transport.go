package network

type NetAddr string

type RPC struct {
	from    NetAddr // sent message fron transport
	payload []byte  // data from server through rpc

}

type Transport interface {
	consume() <-chan RPC
	connect(Transport) error
	send(NetAddr, []byte) error
	Addr() NetAddr
}
