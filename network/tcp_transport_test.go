package network

import (
	"net"
	"testing"
	"time"
)

// func TestTCPPeer_Send(t *testing.T) {

// 	// Create a TCP connection for testing
// 	conn, err := net.Dial("tcp", ":8080")
// 	if err != nil {
// 		t.Fatalf("Failed to create TCP connection: %v", err)
// 	}
// 	defer conn.Close()

// 	// Create a TCPPeer instance
// 	peer := &TCPPeer{
// 		connection: conn,
// 	}

// 	// Test sending a message
// 	message := []byte("Hello, World!")
// 	err = peer.Send(message)
// 	if err != nil {
// 		t.Errorf("Failed to send message: %v", err)
// 	}
// }

func TestAcceptLoop(t *testing.T) {
	// Create a TCPTransport instance
	peerCh := make(chan *TCPPeer)
	tcp := NewTCPTransport("localhost:1234", peerCh)

	// Create a mock listener
	mockListener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		t.Fatalf("Failed to create mock listener: %v", err)
	}
	defer mockListener.Close()

	// Start the acceptLoop in a goroutine
	go tcp.Start()

	// Create a mock connection
	mockConn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Fatalf("Failed to create mock connection: %v", err)
	}
	defer mockConn.Close()

	// Accept the mock connection
	mockListener.(*net.TCPListener).SetDeadline(time.Now().Add(1 * time.Second))
	conn, err := mockListener.Accept()
	if err != nil {
		t.Fatalf("Failed to accept mock connection: %v", err)
	}

	// Verify that the peer is sent to the PeerCh channel
	select {
	case peer := <-peerCh:
		if peer.connection != conn {
			t.Errorf("Received unexpected peer connection. Expected: %v, Got: %v", conn, peer.connection)
		}
	default:
		t.Error("No peer received on PeerCh channel")
	}
}

// func TestTCPTransport_Start(t *testing.T) {
// 	// Create a TCPTransport instance
// 	transport := NewTCPTransport("localhost:1234", make(chan *TCPPeer))

// 	// Start the TCPTransport
// 	err := transport.Start()
// 	if err != nil {
// 		t.Fatalf("Failed to start TCPTransport: %v", err)
// 	}

// 	// Create a TCP connection for testing
// 	conn, err := net.Dial("tcp", "localhost:1234")
// 	if err != nil {
// 		t.Fatalf("Failed to create TCP connection: %v", err)
// 	}
// 	defer conn.Close()

// 	// Ensure that a TCPPeer is received on the PeerCh channel
// 	select {
// 	case peer := <-transport.PeerCh:
// 		// Test the received TCPPeer
// 		if peer.connection != conn {
// 			t.Errorf("Received TCPPeer has incorrect connection")
// 		}
// 	default:
// 		t.Errorf("No TCPPeer received on PeerCh channel")
// 	}
// }
