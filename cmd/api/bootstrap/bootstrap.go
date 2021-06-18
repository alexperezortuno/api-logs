package bootstrap

import "github.com/alexperezortuno/api-logs/internal/platform/server"

const (
	host   = "api-logs"
	port   = 8085
	prefix = "log-api"
)

func Run() error {
	srv := server.New(host, port, prefix)
	return srv.Run()
}
