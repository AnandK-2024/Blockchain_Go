package network

import (
	"io"
	"net"
)

type RPC struct {
	from    net.Addr  // sent message fron transport
	payload io.Reader // data from server through rpc

}

//byte type message
type MessageType byte


// work as enum in go
const (
	MessageTypeTx        MessageType = 0x1
	MessageTypeBlock     MessageType = 0x2
	MessageTypeGetBlocks MessageType = 0x3
	MessageTypeStatus    MessageType = 0x4
	MessageTypeGetStatus MessageType = 0x5
	MessageTypeBlocks    MessageType = 0x6
)
