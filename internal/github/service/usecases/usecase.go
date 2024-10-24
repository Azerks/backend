package usecases

type RepositoryDTO struct {
	FullName   string         `json:"full_name"`
	Owner      string         `json:"owner"`
	Repository string         `json:"repository"`
	Language   map[string]int `json:"languages"`
}

type RepositoriesFilters struct {
	Limit    int
	Language string
}

type (
	RepositoriesReader interface {
		ReadPublicRepositories(filters RepositoriesFilters) ([]RepositoryDTO, error)
	}
)
