package http

// "github.com/rs/zerolog/log"
// "github.com/spf13/cobra"
// "github.com/spf13/viper"
// "github.com/urfave/negroni"
// "github.com/yourusername/yourproject/config"
// "github.com/yourusername/yourproject/middleware"
// "github.com/yourusername/yourproject/routes"
// "github.com/yourusername/yourproject/utils"

func StartHttpServer(addr string, handler Handler) error {
	if handler == nil {
		handler = DefaultRouter
	}
	server := NewServer(addr, handler)
	return server.Serve()
}
