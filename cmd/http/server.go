package http

import (
	"log"
	"net"
)

type Server struct {
	Addr    string
	Network string // tcp default
	Handler Handler
	// log     *logrus.Logger
}

func NewServer(addr string, handler Handler) *Server {
	return &Server{
		Addr:    addr,
		Network: "tcp",
		Handler: handler,
	}
}

func (s *Server) Serve() error {
	listener, err := net.Listen(s.Network, s.Addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Printf("Listening on %s", s.Addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
