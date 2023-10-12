package network

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	// Create a TCP address server-A
	tcpAddrA := &net.TCPAddr{
		IP:   net.ParseIP("192.0.2.1"),
		Port: 80,
	}
	local := NewLocaltransport(tcpAddrA)
	// Create a TCP address serverB
	tcpAddrB := &net.TCPAddr{
		IP:   net.ParseIP("10.0.0.1"),
		Port: 8080,
	}
	remote := NewLocaltransport(tcpAddrB)
	err := local.connect(remote)
	assert.NoError(t, err)
	// err = remote.connect(local)
	// assert.NoError(t, err)
	// assert.Equal(t, remote, local.peers[remote.Addr()])
	// assert.Equal(t, local, remote.peers[local.Addr()])

	go func() {
		for {
			remote.sendMessage(local.Addr(), []byte("I love go lang"))
			time.Sleep(1 * time.Second)
		}
	}()

	opt := serveropts{
		//
	}

	s := Newserver(opt)
	s.start()

}
