package server

import "github.com/Useurmind/solar-controller/pkg/config"

type Server struct {
	config *config.Config
}

func NewServer(cfg *config.Config) *Server{
	return &Server{
		config: cfg,
	}
}

func (s *Server) Run() error {
	return nil
}