package query

type (
	RepositoriesReader interface {
		ReadPublicRepositories(filters RepositoriesFilters) ([]RepositoryDTO, error)
	}
)
