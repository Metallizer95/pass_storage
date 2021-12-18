package httpserver

import "time"

type Option func(*Server)

func Port(port string) Option {
	return func(s *Server) {
		s.server.Addr = port
	}
}

func ReadTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = t
	}
}

func WriteTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = t
	}
}

func ShutdownTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = t
	}
}
