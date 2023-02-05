package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Repository provides access for configuring and starting the HTTP API server.
type Repository interface {
	Engine() *gin.Engine
	Run() error
}

// Adapter is a structure that contains the necessary dependencies for creating
// and running a HTTP API server.
type service struct {
	conf   *Configuration
	engine *gin.Engine
}

// Configuration is a structure for configuring a HTTP API server.
type Configuration struct {
	Env  string
	Host string
	Port uint16
}

// New creates an adapter given a proxy service and a configuration structure.
func Init(conf *Configuration) Repository {
	s := &service{
		conf:   conf,
		engine: gin.New(),
	}

	if conf.Env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	s.engine.SetTrustedProxies(nil)

	return s
}

// Engine returns the underlying server engine used by the HTTP API server.
func (s *service) Engine() *gin.Engine {
	return s.engine
}

// Run starts the HTTP API server. It returns an error if the server panics.
func (s *service) Run() error {
	return s.engine.Run(fmt.Sprintf("%s:%d", s.conf.Host, s.conf.Port))
}
