package types

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Hash [32]uint8

func (h Hash) IsZero() bool {
	for _, v := range h {
		if v != 0 {
			return false
		}
	}
	return true
}

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		err := fmt.Sprintf("lenght %d should be 32", len(b))
		panic(err)
	}

	var value [32]uint8
	for i := range 32 {
		value[i] = b[i]
	}
	return Hash(value)
}

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}
