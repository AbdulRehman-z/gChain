package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts
	rpcCh  chan RPC
	quitCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()

	ticker := time.NewTicker(1 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("received rpc from %v\n", rpc)
		case <-ticker.C:
			fmt.Println("tick")
		case <-s.quitCh:
			break free
		}
	}

	fmt.Println("server stopped")

}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func() {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}()
	}
}
