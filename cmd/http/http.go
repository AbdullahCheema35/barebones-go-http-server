package http

// "github.com/rs/zerolog/log"
// "github.com/spf13/cobra"
// "github.com/spf13/viper"
// "github.com/urfave/negroni"

var supportedHttpProtocol = "HTTP/1.0"

func StartHttpServer(addr string, handler Handler) error {
	if handler == nil {
		handler = DefaultRouter
	}
	server := NewServer(addr, handler)
	return server.Serve()
}
