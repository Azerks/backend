package app

import (
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/app/query"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	GetGithubRepositories query.GithubRepositoriesHandler
}

func New() *App {
	return &App{
		Commands: Commands{},
		Queries: Queries{
			GetGithubRepositories: query.NewGithubRepositoriesHandler(),
		},
	}
}
