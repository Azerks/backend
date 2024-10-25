package service

import (
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/adapters/repositories"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/service/usecases"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared"
	"net/http"
)

type Service struct {
	GetGithubRepositories usecases.GetPublicGithubRepositoriesHandler
}

func New(config *shared.Config) *Service {

	client := http.Client{}
	repository := repositories.New(config, &client)

	return &Service{
		GetGithubRepositories: usecases.NewGithubRepositoriesHandler(repository),
	}
}
