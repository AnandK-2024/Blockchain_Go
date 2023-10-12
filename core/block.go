package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/AnandK-2024/Blockchain/crypto"
	"github.com/AnandK-2024/Blockchain/types"
)

type Header struct {
	version uint32
	// hash of previous block header
	prevblockHash types.Hash
	// hash of transactions(merkle root)
	DataHash  types.Hash
	Timestamp int64
	Height    uint32
}

type Block struct {
	*Header      // don't work with copy of header , so taking pointer make much more efficient
	Transactions []Transaction
	// public key of block validator
	validator crypto.PublicKey
	// signature of block validator
	signature *crypto.Signature
	// hash of block header== hash of current block
	hash types.Hash
}

// get bytes of header of block
func (h *Header) Bytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(h)

	return buf.Bytes()
}

// create new block with header and array of transaction
func NewBlock(h *Header, txs []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txs,
	}
}

// create new block with previous header and transactions
func NewBlockFromPrevHeader(prevHeader Header,txs [] Transaction) (*Block,error){
	merklerootstring:=CalculateMerkleRoot(txs)
	merkleroothash,err:=StringToHash(merklerootstring)
	if err != nil {
		return nil, fmt.Errorf("error in converting string to merklehash")
	}
	header:=&Header{
		version: 1,
		prevblockHash:prevHeader.prevblockHash ,
		DataHash: merkleroothash,
		Timestamp: time.Now().UnixNano(),
		Height: prevHeader.Height,
	}
	return NewBlock(header,txs),nil

}

// calculate data hash as merkle tree hash root of transaction
func (b *Block) CalculateMerkleRoot() error {
	if len(b.Transactions) == 0 {
		return fmt.Errorf("No transaction avilable!! Add transactions into block")
	}
	merklestring := CalculateMerkleRoot(b.Transactions)
	merkleHashRoot, err := StringToHash(merklestring)
	if err != nil {
		return fmt.Errorf("error in converting string to merklehash")
	}
	b.DataHash = merkleHashRoot
	return nil
}

// add transaction in block
func(b *Block) AddTransaction(tx Transaction){
	b.Transactions=append(b.Transactions, tx)
	b.CalculateMerkleRoot()

}

// add transactions in block
func (b *Block) AddTransactions(txs []Transaction){
	for i:=0;i<len(txs);i++{
		b.Transactions=append(b.Transactions, txs[i])
	}
	b.CalculateMerkleRoot()
}

// calculate hash of block
func (b *Block) Hash() types.Hash {
	buf := &bytes.Buffer{}
	// NewEncoder returns a new encoder that will transmit on the io.Writer.
	enc := gob.NewEncoder(buf)
	// Encode transmits the data item represented by the empty interface value,
	// guaranteeing that all necessary type information has been transmitted first.
	err := enc.Encode(b.Header)
	// Passing a nil pointer to Encoder will panic, as they cannot be transmitted by gob.
	if err != nil {
		panic(err)
	}

	h := sha256.Sum256(buf.Bytes())
	return types.Hash(h)
}

// validator will sign the block
func (b *Block) Sign(privkey *crypto.PrivateKey) error {
	hash := b.Hash()
	signature, err := privkey.SignMessage(hash[:])
	if err != nil {
		fmt.Println("unable to sign block with private key")
		panic(err)
	}
	b.validator = privkey.GeneratePublicKey()
	b.hash = hash
	b.signature = signature

	return nil
}

// verifier verify the signature the block
func (b *Block) Verify() error {
	if b.signature == nil {
		return fmt.Errorf("block has not signature")
	}

	// validate signature
	sig := b.signature
	sucess := sig.Verify(b.validator, b.hash[:])
	if !sucess {
		return fmt.Errorf("invalid block header signature ")
	}
	// verify all transactions of block
	for _, tx := range b.Transactions {
		if err := tx.Verify(); err != nil {
			return err
		}
	}

	return nil
}

// validator can validate the block before voting / finalization


// enoding block
func (b *Block) Encode(enc Encoder[*Block]) error {
	return enc.Encode(b)
}

//decoding block
func (b *Block) Decode(dec Decoder[*Block]) error {
	return dec.Decode(b)
}
