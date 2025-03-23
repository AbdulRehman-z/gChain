package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/AbdulRehman-z/goChain/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	h := &Header{
		Version:       10,
		Timestamp:     uint64(time.Now().UnixNano()),
		PrevBlockHash: types.RandomHash(),
		Height:        1,
		Nonce:         4,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hdec := &Header{}
	assert.Nil(t, hdec.DecodeBinary(buf))
	assert.Equal(t, h, hdec)
}

func TestBlock_Encode_Decode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:       10,
			Timestamp:     uint64(time.Now().UnixNano()),
			PrevBlockHash: types.RandomHash(),
			Height:        1,
			Nonce:         21214,
		},
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDec := &Block{}
	assert.Nil(t, bDec.DecodeBinary(buf))
	assert.Equal(t, bDec, b)

	fmt.Println(bDec)

}

func TestBlockHash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:       10,
			Timestamp:     uint64(time.Now().UnixNano()),
			PrevBlockHash: types.RandomHash(),
			Height:        1,
			Nonce:         21214,
		},
		Transactions: nil,
	}

	hash := b.Hash()
	assert.False(t, hash.IsZero())
}
