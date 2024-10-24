package app

import (
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/app/query"
)

type GetPublicGithubRepositories struct {
	Filters query.RepositoriesFilters
}
type GetPublicGithubRepositoriesHandler struct {
	githubRepository query.RepositoriesReader
}

func NewGithubRepositoriesHandler(githubRepository query.RepositoriesReader) GetPublicGithubRepositoriesHandler {
	return GetPublicGithubRepositoriesHandler{
		githubRepository: githubRepository,
	}
}

func (h *GetPublicGithubRepositoriesHandler) Handle(params GetPublicGithubRepositories) ([]query.RepositoryDTO, error) {
	repositories, err := h.githubRepository.ReadPublicRepositories(params.Filters)
	if err != nil {
		return nil, err
	}
	return repositories, nil
}
