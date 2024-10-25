package usecases

import (
	"context"
	"fmt"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared/errs"
)

type GetPublicGithubRepositories struct {
	Filters RepositoriesFilters
}
type GetPublicGithubRepositoriesHandler struct {
	githubRepository RepositoriesReader
}

func NewGithubRepositoriesHandler(githubRepository RepositoriesReader) GetPublicGithubRepositoriesHandler {
	return GetPublicGithubRepositoriesHandler{
		githubRepository: githubRepository,
	}
}

func (h *GetPublicGithubRepositoriesHandler) Handle(ctx context.Context, params GetPublicGithubRepositories) ([]RepositoryDTO, error) {
	rs, err := h.githubRepository.ReadPublicRepositories(ctx, params.Filters)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrInternal{}, err)
	}
	return rs, nil
}
