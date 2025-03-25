package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis() (*Blockchain, error) {
	genesisBlock := randomBlock(0)
	return NewBlockchain(genesisBlock)
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
		b := randomBlockWithSig(t, uint32(height+1))
		bc.AddBlock(b)
	}

	assert.Equal(t, bc.Height(), uint32(lenBlock))
	assert.Equal(t, len(bc.headers), lenBlock+1)
	assert.NotNil(t, bc.AddBlock(randomBlockWithSig(t, 99)))
}

func TestHasBlock(t *testing.T) {
	bc, err := newBlockchainWithGenesis()
	assert.Nil(t, err)
	assert.True(t, bc.HasBlock(0))
}
