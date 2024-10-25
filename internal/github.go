package internal

import (
	"github.com/Scalingo/sclng-backend-test-v1/common/server"
	"github.com/Scalingo/sclng-backend-test-v1/internal/interfaces/public/http"
	"github.com/Scalingo/sclng-backend-test-v1/internal/service"
)

func New(s *server.Server) *service.Service {
	serv := service.New(s.Config)
	http.NewServer(s, serv).Register()
	return serv
}
