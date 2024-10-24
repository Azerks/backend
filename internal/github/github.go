package github

import (
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/app"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/interfaces/public/http"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared/server"
)

func NewGithubService(s *server.Server) *app.App {
	application := app.New(s.Config)
	http.NewServer(s, application).Register()
	return application
}
