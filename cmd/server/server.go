package server

import (
	"log"
	"net"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	// "github.com/rs/zerolog/log"
	// "github.com/spf13/cobra"
	// "github.com/spf13/viper"
	// "github.com/urfave/negroni"
	// "github.com/yourusername/yourproject/config"
	// "github.com/yourusername/yourproject/middleware"
	// "github.com/yourusername/yourproject/routes"
	// "github.com/yourusername/yourproject/utils"
)

// var log = logrus.New()
// var router = mux.NewRouter()

type Server struct {
	addr    string
	network string
	router  *mux.Router
	log     *logrus.Logger
}

func NewServer(addr string) *Server {
	return &Server{
		addr:    addr,
		network: "tcp",
		router:  nil,
		log:     nil,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen(s.network, s.addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Printf("Listening on %s", s.addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
