package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/AbdulRehman-z/goChain/crypto"
	"github.com/AbdulRehman-z/goChain/types"
)

type Header struct {
	Version       uint32
	Timestamp     uint64
	DataHash      types.Hash
	PrevBlockHash types.Hash
	Height        uint32
	Nonce         uint64
}

func (h *Header) Bytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(h)

	return buf.Bytes()
}

type Block struct {
	*Header
	Validator    crypto.PublicKey
	Signature    *crypto.Signature
	Transactions []Transaction

	hash types.Hash
}

func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txx,
	}
}

func (b *Block) AddTransaction(tx *Transaction) {
	b.Transactions = append(b.Transactions, *tx)
}

func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.Header.Bytes())
	if err != nil {
		return fmt.Errorf("failed to sign block header: %w", err)
	}

	b.Validator = privKey.PublicKey()
	b.Signature = sig

	return nil

}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block signature is nil")
	}

	if !b.Signature.Verify(b.Validator, b.Header.Bytes()) {
		return fmt.Errorf("block signature is invalid")
	}

	return nil
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Hash(hasher Hasher[*Header]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b.Header)
	}

	return b.hash
}
