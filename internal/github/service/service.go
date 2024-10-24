package service

import (
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/adapters/github"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/service/usecase"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared"
)

type App struct {
	GetGithubRepositories usecase.GetPublicGithubRepositoriesHandler
}

func New(config *shared.Config) *App {

	repository := github.New(config)

	return &App{
		GetGithubRepositories: usecase.NewGithubRepositoriesHandler(repository),
	}
}
