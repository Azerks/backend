package service

import (
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/adapters/github"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared"
)

type App struct {
	GetGithubRepositories GetPublicGithubRepositoriesHandler
}

func New(config *shared.Config) *App {

	repository := github.New(config)

	return &App{
		GetGithubRepositories: NewGithubRepositoriesHandler(repository),
	}
}
