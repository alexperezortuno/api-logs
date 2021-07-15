package bootstrap

import (
	"github.com/alexperezortuno/api-logs/internal/platform/server"
	"github.com/alexperezortuno/api-logs/tools/environment"
)

var params = environment.Server()

func Run() error {
	srv := server.New(params.Host, uint(params.Port), params.Context)
	return srv.Run()
}
