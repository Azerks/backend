package http

import (
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/service"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared/server"
)

type Server struct {
	server *server.Server
	app    *service.App
}

func NewServer(mux *server.Server, app *service.App) *Server {
	return &Server{
		server: mux,
		app:    app,
	}
}

func (s *Server) Register() {
	s.server.Router.HandleFunc("/v1/repositories", s.getRepositories())
}
