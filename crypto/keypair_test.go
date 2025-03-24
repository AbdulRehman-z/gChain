package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivatePublicKeySigSuccess(t *testing.T) {

	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	msg := []byte("Hello, world!")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	sig.Verify(pubKey, msg)
	assert.True(t, sig.Verify(pubKey, msg))
}

func TestPrivatePublicKeySigFail(t *testing.T) {

	privKey := GeneratePrivateKey()
	msg := []byte("Hello, world!")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivKey := GeneratePrivateKey()
	pubKey := otherPrivKey.PublicKey()

	assert.False(t, sig.Verify(pubKey, msg))
}
