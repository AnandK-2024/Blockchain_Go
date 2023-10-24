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
type GetBlocksMessage struct {
	From uint32
	// If To is 0 the maximum blocks will be returned.
	To uint32
}
