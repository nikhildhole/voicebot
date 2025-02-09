package server

import (
	"net"
	"sync"
	"voicebot/pkg/logger"
)

type Server struct {
	address string
	handler func(net.Conn) // Injected dependency
}

func New(address string, handler func(net.Conn)) *Server {
	return &Server{address: address, handler: handler}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		logger.Error("Failed to start server: %v", err)
		return err
	}
	defer listener.Close()

	logger.Info("Server listening on %s", s.address)

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Error accepting connection: %v", err)
			continue
		}

		logger.Info("New client connected: %s", conn.RemoteAddr())

		wg.Add(1)
		go func() {
			defer wg.Done()
			s.handler(conn) // Injected function instead of hardcoding handler
		}()
	}
}
