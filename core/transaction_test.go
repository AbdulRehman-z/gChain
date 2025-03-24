package core

import (
	"testing"

	"github.com/AbdulRehman-z/goChain/crypto"
	"github.com/stretchr/testify/assert"
)

func TestTransactionSign(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	// pubKey := privKey.PublicKey()

	tx := &Transaction{
		Data: []byte("Hello,world"),
	}

	assert.Nil(t, tx.Sign(privKey))
}

func TestTransactionVerify(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()

	tx := &Transaction{
		Data: []byte("Hello,world"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.PublicKey = otherPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}
