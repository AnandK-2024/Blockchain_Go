package network

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	// Create a TCP address server-A
	tcpAddrA := &net.TCPAddr{
		IP:   net.ParseIP("192.0.2.1"),
		Port: 80,
	}
	tra := NewLocaltransport(tcpAddrA)
	// Create a TCP address serverB
	tcpAddrB := &net.TCPAddr{
		IP:   net.ParseIP("10.0.0.1"),
		Port: 8080,
	}
	trb := NewLocaltransport(tcpAddrB)
	err := tra.connect(trb)
	assert.NoError(t, err)
	err = trb.connect(tra)
	assert.NoError(t, err)
	assert.Equal(t, trb, tra.peers[trb.Addr()])
	assert.Equal(t, trb, tra.peers[trb.Addr()])
}

func TestSendMessage(t *testing.T) {
	// Create a TCP address server-A
	tcpAddrA := &net.TCPAddr{
		IP:   net.ParseIP("192.0.2.1"),
		Port: 80,
	}
	tra := NewLocaltransport(tcpAddrA)
	// Create a TCP address serverB
	tcpAddrB := &net.TCPAddr{
		IP:   net.ParseIP("10.0.0.1"),
		Port: 8080,
	}
	trb := NewLocaltransport(tcpAddrB)
	err := tra.connect(trb)
	assert.NoError(t, err)
	err = trb.connect(tra)
	assert.NoError(t, err)

	payload := []byte("Hello, World!")
	assert.Nil(t, tra.sendMessage(trb.Addr(), payload))
	rpc := <-trb.consume()
	fmt.Println(tra)
	fmt.Println("new server")
	fmt.Println(trb)
	fmt.Println(rpc)
	assert.Equal(t, tra.Addr(), rpc.from)
	reader := rpc.payload
	actual := make([]byte, len(payload))
	_, _ = reader.Read(actual)
	assert.Equal(t, payload, actual)
}
