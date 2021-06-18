package server

import (
	"fmt"
	"github.com/alexperezortuno/api-logs/internal/platform/server/handler/health"
	"github.com/alexperezortuno/api-logs/internal/platform/server/handler/logger"
	"github.com/alexperezortuno/api-logs/internal/platform/server/middleware/logging"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint, prefix string) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}

	srv.registerRoutes(prefix)
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes(prefix string) {
	s.engine.Use(logging.Middleware())

	s.engine.GET(fmt.Sprintf("/%s/%s", prefix, "/health"), health.CheckHandler())
	s.engine.POST(fmt.Sprintf("/%s/%s", prefix, "/log"), logger.CreateHandler())
}