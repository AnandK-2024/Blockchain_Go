package network

import "github.com/AnandK-2024/Blockchain/core"

type StatusMessage struct {
	// id of server
	ID      string
	Version uint32
	Height  uint32
}

type BlockMessage struct {
	Blocks []*core.Block
}

type GetStatusMessage struct{}