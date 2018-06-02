package server

import (
	"go.uber.org/zap"
)

// Server struct
type Server struct {
	log *zap.Logger
}

// New Returns a new server
func New(c *Config, logger *zap.Logger) *Server {
	return &Server{
		logger,
	}
}

// Run server/service
func (s *Server) Run() {

	s.log.Info("Starting CLI Service...")
	StartRemoteCli(s.log)

	s.log.Info("Starting Server Runner...")
	for {
		//run
	}
}
