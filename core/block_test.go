package core

import (
	"testing"
	"time"

	"github.com/AbdulRehman-z/goChain/crypto"
	"github.com/AbdulRehman-z/goChain/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32, prevBlockHash types.Hash) *Block {
	h := &Header{
		Version:       10,
		Timestamp:     uint64(time.Now().UnixNano()),
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Nonce:         4,
	}

	// tx := Transaction{
	// 	Data: []byte("hello world"),
	// }

	b := NewBlock(h, []Transaction{})

	return b
}

func randomBlockWithSig(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	b := randomBlock(height, prevBlockHash)
	privateKey := crypto.GeneratePrivateKey()
	tx := randomTxWithSig(t)
	b.AddTransaction(tx)

	assert.Nil(t, b.Sign(privateKey))

	return b
}

func TestSignBlock(t *testing.T) {
	b := randomBlock(1, types.Hash{})
	privateKey := crypto.GeneratePrivateKey()

	assert.Nil(t, b.Sign(privateKey))
}

func TestVerifyBlock(t *testing.T) {
	b := randomBlock(1, types.Hash{})
	privateKey := crypto.GeneratePrivateKey()

	assert.Nil(t, b.Sign(privateKey))
	assert.Nil(t, b.Verify())

	otherPrivateKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivateKey.PublicKey()

	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}
