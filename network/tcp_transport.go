package network

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

type TCPPeer struct {
	connection net.Conn
	Outgoing   bool
}

// send message data to peers
func (tp *TCPPeer) Send(message []byte) error {
	if _, err := tp.connection.Write(message); err != nil {
		return err
	}
	return nil
}
func (p *TCPPeer) readLoop(rpcCh chan RPC) {
	buf := make([]byte, 4096)
	for {
		n, err := p.connection.Read(buf)
		if err == io.EOF {
			continue
		}
		if err != nil {
			fmt.Printf("read error: %s", err)
			continue
		}

		msg := buf[:n]
		// fmt.Println("msg recieved from peer", rpcCh, "by this peer:", p.connection.LocalAddr())
		rpcCh <- RPC{
			From:    p.connection.RemoteAddr(),
			Payload: bytes.NewReader(msg),
		}
	}
}

type TCPTransport struct {
	PeerCh     chan *TCPPeer
	listenAddr string
	// A Listener is a generic network listener for stream-oriented protocols.
	// Multiple goroutines may invoke methods on a Listener simultaneously.
	listner net.Listener
}

// create new tcp network
func NewTCPTransport(addr string, peer chan *TCPPeer) *TCPTransport {

	return &TCPTransport{
		PeerCh:     peer,
		listenAddr: addr,
	}
}
func (tcp *TCPTransport) readLoop(peer *TCPPeer) {
	buf := make([]byte, 2048)
	for {
		// Read reads n data and error from the connection.
		n, err := peer.connection.Read(buf)

		// EOF is the error returned by Read when no more input is available.
		if err == io.EOF {
			continue
		}
		if err != nil {
			fmt.Println("error in read data from peer\n", err)
			continue
		}
		msg := buf[:n]
		fmt.Printf("\nincoming data from peer is:%s\n", string(msg))
		// tcp.PeerCh <- peer
		// fmt.Println("val recieved from peer chan==>", tcp.PeerCh)
		// rpcCh <- RPC{
		// 	from:    tp.connection.RemoteAddr(),
		// 	payload: bytes.NewReader(msg),
		// }
	}
}

func (tcp *TCPTransport) acceptLoop() {
	for {
		// Accept waits for and returns the next connection to the listener.
		connection, err := tcp.listner.Accept()
		if err != nil {
			fmt.Println("accept error from:", connection)
			continue
		}
		fmt.Println("connection are:==>", connection)
		peer := &TCPPeer{
			connection: connection,
		}
		tcp.PeerCh <- peer
		// fmt.Printf("details of peer:\n", connection.LocalAddr(), connection.RemoteAddr())
		// go tcp.readLoop(peer)

	}
}

// start tcp network listening to port
func (tcp *TCPTransport) Start() error {
	fmt.Println("server start")
	listner, err := net.Listen("tcp", tcp.listenAddr)
	if err != nil {
		return err
	}
	tcp.listner = listner
	go tcp.acceptLoop()
	fmt.Printf("\ntcp start listening at port:%+v\n", tcp.listenAddr)
	return nil
}
