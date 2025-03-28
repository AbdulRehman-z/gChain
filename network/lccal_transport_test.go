package network

import (
	"testing"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	// assert.Equal(t, tra.peers[trb.Addr()], trb)
	// assert.Equal(t, trb.peers[tra.Addr()], tra)
}

func TestConsume(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	// assert.Equal(t, tra.peers[trb.Addr()], trb)
	// assert.Equal(t, trb.peers[tra.Addr()], tra)

	// msg := []byte("Hello, World!")
	// // assert.Nil(t, tra.SendMessage(trb.Addr(), msg), msg)

	// rpc := <-trb.Consume()
	// assert.Equal(t, rpc.Payload, msg)
	// assert.Equal(t, rpc.From, tra.Addr())

}
