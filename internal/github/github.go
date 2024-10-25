package github

import (
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/interfaces/public/http"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/service"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared/server"
)

func New(s *server.Server) *service.Service {
	application := service.New(s.Config)
	http.NewServer(s, application).Register()
	return application
}
