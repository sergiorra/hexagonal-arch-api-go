package bootstrap

import "github.com/sergiorra/hexagonal-arch-api-go/internal/platform/server"

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
