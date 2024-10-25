package http

import (
	"github.com/Scalingo/sclng-backend-test-v1/common/server"
	"github.com/Scalingo/sclng-backend-test-v1/internal/service"
)

type Server struct {
	server *server.Server
	app    *service.Service
}

func NewServer(mux *server.Server, app *service.Service) *Server {
	return &Server{
		server: mux,
		app:    app,
	}
}

func (s *Server) Register() {
	s.server.Router.HandleFunc("/v1/repositories", s.getRepositories())
}
