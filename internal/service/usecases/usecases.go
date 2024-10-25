package usecases

import "context"

type (
	RepositoryDTO struct {
		FullName   string         `json:"full_name"`
		Owner      string         `json:"owner"`
		Repository string         `json:"repository"`
		Language   map[string]int `json:"languages"`
	}

	RepositoriesFilters struct {
		Limit     int
		Languages []string
	}

	RepositoryReader interface {
		ReadPublicRepositories(ctx context.Context, filters RepositoriesFilters) ([]RepositoryDTO, error)
	}
)
