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

func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.HeaderData())
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

	if !b.Signature.Verify(b.Validator, b.HeaderData()) {
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

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}

func (b *Block) HeaderData() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(b.Header)

	return buf.Bytes()
}
