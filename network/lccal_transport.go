package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC

	lock  sync.RWMutex
	peers map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s: couldn't send message to %s", t.addr, to)
	}

	t.lock.Lock()
	defer t.lock.Unlock()

	peer.consumeCh <- RPC{
		From:    t.addr,
		Payload: payload,
	}

	return nil
}

func (t *LocalTransport) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	if _, ok := t.peers[tr.Addr()]; ok {
		return fmt.Errorf("%s: already connected to %s", t.addr, tr.Addr())
	}

	t.peers[tr.Addr()] = tr.(*LocalTransport)

	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}
