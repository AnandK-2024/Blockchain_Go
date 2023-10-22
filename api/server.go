package api

import (
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/AnandK-2024/Blockchain/core"
	"github.com/AnandK-2024/Blockchain/types"
	"github.com/labstack/echo/v4"
)

type Block struct {
	Version           uint32     `json:"version"`
	Height            uint32     `json:"height"`
	PreviousBlockHash string     `json:"prevBlockHash"`
	Hash              string     `json:"hash"`
	DataHash          string     `json:"datahash"`
	Timestamp         int64      `json:"timestamp"`
	Validator         string     `json:"validator"`
	Signature         string     `json:"signature"`
	TxResponse        TxResponse `json:"txresponse"`
}

type TxResponse struct {
	TxCount uint     `json:"txCount"`
	Hash    []string `json:"txsHashes"`
}

func intoJsonBlock(block *core.Block) Block {

	txResponse := TxResponse{
		TxCount: uint(len(block.Transactions)),
		Hash:    make([]string, len(block.Transactions)),
	}
	for i := 0; i < len(block.Transactions); i++ {
		txResponse.Hash[i] = block.Transactions[i].Hash().String()
	}
	return Block{
		Version:           block.Version,
		Height:            block.Height,
		PreviousBlockHash: block.PrevblockHash.String(),
		Hash:              block.BlockHash.String(),
		DataHash:          block.DataHash.String(),
		Timestamp:         block.Timestamp,
		Validator:         block.Validator.Address().String(),
		Signature:         block.Signature.String(),
		TxResponse:        txResponse,
	}
}

type Server struct {
	ListenAddr string
	Bc         *core.Blockchain
}

func (s *Server) Start() error {
	// New creates an instance of Echo.
	e := echo.New()
	e.GET("/tx/:hash", s.handleGetTx)
	e.GET("/block/:hashorId", s.handleGetBlock)
	return e.Start(s.ListenAddr)
}

// input: hash of tranaction
// return : transaction
func (s *Server) handleGetTx(c echo.Context) error {
	hash := c.Param("hash")
	b, err := hex.DecodeString(hash)
	if err != nil {
		// JSON sends a JSON response with status code.
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tx, err := s.Bc.GetTxByHash(types.HashFromByte(b))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())

	}
	return c.JSON(http.StatusOK, tx)
}

// get block by height
func (s *Server) handleGetBlock(c echo.Context) error {
	hash := c.Param("hashorId")
	// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	height, err := strconv.Atoi(hash)
	if err == nil {
		// hash is converted to int:means height was given
		block, err := s.Bc.GetBlock(uint32(height))
		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusOK, intoJsonBlock(block))
	}
	b, err := hex.DecodeString(hash)
	if err != nil {
		// JSON sends a JSON response with status code.
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	block, err := s.Bc.GetBlockByHash(types.HashFromByte(b))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, intoJsonBlock(block))

}

func handlePostTx(c echo.Context) error {

	return nil
}
