package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocaltransport(NetAddr("A"))
	trb := NewLocaltransport(NetAddr("B"))
	err := tra.connect(trb)
	assert.NoError(t, err)
	err = trb.connect(tra)
	assert.NoError(t, err)
	assert.Equal(t, trb, tra.peers[trb.Addr()])
	assert.Equal(t, trb, tra.peers[trb.Addr()])
}

func TestSendMessage(t *testing.T) {
	tra := NewLocaltransport(NetAddr("A"))
	trb := NewLocaltransport(NetAddr("B"))
	tra.connect(trb)
	trb.connect(tra)

	payload := []byte("Hello, World!")
	assert.Nil(t, tra.sendMessage(trb.Addr(), payload))
	rpc := <-trb.consume()
	assert.Equal(t, tra.Addr(), rpc.from)
	assert.Equal(t, payload, rpc.payload)
}
