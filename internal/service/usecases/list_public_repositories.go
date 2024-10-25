package usecases

import (
	"context"
	"fmt"
	"github.com/Scalingo/sclng-backend-test-v1/common/errs"
)

type ListPublicRepositories struct {
	Filters RepositoriesFilters
}
type ListPublicRepositoriesHandler struct {
	repository RepositoryReader
}

func NewListPublicRepositoriesHandler(repository RepositoryReader) ListPublicRepositoriesHandler {
	return ListPublicRepositoriesHandler{
		repository: repository,
	}
}

func (h *ListPublicRepositoriesHandler) Handle(ctx context.Context, params ListPublicRepositories) ([]RepositoryDTO, error) {
	rs, err := h.repository.ReadPublicRepositories(ctx, params.Filters)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errs.ErrInternal{}, err)
	}
	return rs, nil
}
