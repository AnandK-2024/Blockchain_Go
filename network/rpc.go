package network

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io"
	"net"

	"github.com/AnandK-2024/Blockchain/core"
	"github.com/sirupsen/logrus"
)

type RPC struct {
	From    net.Addr  // sent message fron transport
	Payload io.Reader // data from server through rpc

}
type RPCProcessor interface {
	ProcessMessage(*DecodedMessage) error
}

type RPCDecodeFunc func(RPC) (*DecodedMessage, error)

// byte type message
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

type Message struct {
	Header MessageType
	Data   []byte
}

func NewMessage(t MessageType, data []byte) *Message {
	return &Message{
		Header: t,
		Data:   data,
	}
}

// convert mes to bytes
func (msg *Message) Byte() []byte {
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(msg)
	return buf.Bytes()
}

type DecodedMessage struct {
	From net.Addr
	Data any
}

func DefaultRPCDecodeFunc(rpc RPC) (*DecodedMessage, error) {
	msg := Message{}
	if err := gob.NewDecoder(rpc.Payload).Decode(&msg); err != nil {
		return nil, fmt.Errorf("failed to decode message from %s to %s", rpc.From, err)
	}
	logrus.WithFields(logrus.Fields{
		"from":  rpc.From,
		"types": msg.Header,
	}).Debug("new incoming message")
	switch msg.Header {
	case MessageTypeTx:
		tx := new(core.Transaction)
		if err := tx.Decode(core.NewGobTxDecoder(bytes.NewReader(msg.Data))); err != nil {
			return nil, err
		}

		return &DecodedMessage{
			From: rpc.From,
			Data: tx,
		}, nil

	case MessageTypeBlock:
		block := new(core.Block)
		if err := block.Decode(core.NewGobBlockDecoder(bytes.NewReader(msg.Data))); err != nil {
			return nil, err
		}
		return &DecodedMessage{
			From: rpc.From,
			Data: block,
		}, nil
	case MessageTypeStatus:
		statusMessage := new(StatusMessage)
		if err := gob.NewDecoder(bytes.NewReader(msg.Data)).Decode(statusMessage); err != nil {
			return nil, err
		}
		return &DecodedMessage{
			From: rpc.From,
			Data: statusMessage,
		}, nil

	case MessageTypeGetStatus:
		return &DecodedMessage{
			From: rpc.From,
			Data: &GetStatusMessage{},
		}, nil

	case MessageTypeGetBlocks:
		getBlocks := new(GetBlocksMessage)
		if err := gob.NewDecoder(bytes.NewReader(msg.Data)).Decode(getBlocks); err != nil {
			return nil, err
		}

		return &DecodedMessage{
			From: rpc.From,
			Data: getBlocks,
		}, nil

	case MessageTypeBlocks:
		blocks := new(BlockMessage)
		if err := gob.NewDecoder(bytes.NewReader(msg.Data)).Decode(blocks); err != nil {
			return nil, err
		}

		return &DecodedMessage{
			From: rpc.From,
			Data: blocks,
		}, nil
	default:
		return nil, fmt.Errorf("invalid message header %x", msg.Header)
	}
}

func init() {
	gob.Register(elliptic.P256())
}
