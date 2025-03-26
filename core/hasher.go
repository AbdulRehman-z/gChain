package core

import (
	"crypto/sha256"

	"github.com/AbdulRehman-z/goChain/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct{}

func (BlockHasher) Hash(header *Header) types.Hash {
	if header == nil {
		panic("Cannot hash nil Header")
	}

	h := sha256.Sum256(header.Bytes())
	return types.Hash(h)
}
