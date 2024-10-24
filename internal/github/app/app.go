package app

import (
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/adapters/github"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/app/query"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	GetGithubRepositories query.GetPublicGithubRepositoriesHandler
}

func New() *App {

	repository := github.New()

	return &App{
		Commands: Commands{},
		Queries: Queries{
			GetGithubRepositories: query.NewGithubRepositoriesHandler(repository),
		},
	}
}
