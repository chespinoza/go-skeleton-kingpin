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

// Run the server
func (s *Server) Run() {

	s.log.Info("Starting Server...")
}
