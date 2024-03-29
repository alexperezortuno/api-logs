package environment

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

type ServerValues struct {
	Host    string
	Port    int
	Context string
}

func env() {
	env := os.Getenv("APP_ENV")

	if env == "" || env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func Server() ServerValues {
	env()
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	host := os.Getenv("APP_HOST")
	context := os.Getenv("APP_CONTEXT")

	if err != nil {
		log.Printf("error parsing port")
		port = 8185
	}

	if host == "" {
		host = "api-logs"
	}

	if context == "" {
		context = "log-api"
	}

	return ServerValues{
		Host:    host,
		Context: context,
		Port:    port,
	}
}
