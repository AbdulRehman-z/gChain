package core

import (
	"testing"

	"github.com/AbdulRehman-z/goChain/crypto"
	"github.com/AbdulRehman-z/goChain/types"
	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis() (*Blockchain, error) {
	genesisBlock := randomBlock(0, types.Hash{})

	// Sign the genesis block
	privateKey := crypto.GeneratePrivateKey()
	err := genesisBlock.Sign(privateKey)
	if err != nil {
		return nil, err
	}

	return NewBlockchain(genesisBlock)
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)

	prevBlockHash := BlockHasher{}.Hash(prevHeader)

	return prevBlockHash
}

func TestBlockChain(t *testing.T) {
	bc, err := newBlockchainWithGenesis()
	assert.Nil(t, err)
	assert.NotNil(t, bc.validator)
}

func TestAddBlock(t *testing.T) {
	bc, err := newBlockchainWithGenesis()
	assert.Nil(t, err)
	assert.NotNil(t, bc.validator)

	lenBlock := 1000
	for height := range lenBlock {
		block := randomBlockWithSig(t, uint32(height+1), getPrevBlockHash(t, bc, uint32(height+1)))
		bc.AddBlock(block)
	}

	assert.Equal(t, bc.Height(), uint32(lenBlock))
	assert.Equal(t, len(bc.headers), lenBlock+1)
	assert.NotNil(t, bc.AddBlock(randomBlockWithSig(t, 99, types.Hash{})))
}

func TestHasBlock(t *testing.T) {
	bc, err := newBlockchainWithGenesis()
	assert.Nil(t, err)
	assert.True(t, bc.HasBlock(0))
	assert.False(t, bc.HasBlock(100))
}

func TestAddBlockTooHigh(t *testing.T) {
	bc, err := newBlockchainWithGenesis()
	assert.Nil(t, err)

	assert.Nil(t, bc.AddBlock(randomBlockWithSig(t, 1, getPrevBlockHash(t, bc, uint32(1)))))
	assert.NotNil(t, bc.AddBlock(randomBlockWithSig(t, 1001, types.Hash{})))
}

func TestGetHeader(t *testing.T) {
	bc, err := newBlockchainWithGenesis()
	assert.Nil(t, err)
	lenBlock := 1000

	for height := range lenBlock {
		block := randomBlockWithSig(t, uint32(height+1), getPrevBlockHash(t, bc, uint32(height+1)))
		assert.Nil(t, bc.AddBlock(block))
		header, err := bc.GetHeader(block.Height)
		assert.Nil(t, err)
		assert.Equal(t, header, block.Header)
	}
}
