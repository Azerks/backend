package usecase

type (
	RepositoriesReader interface {
		ReadPublicRepositories(filters RepositoriesFilters) ([]RepositoryDTO, error)
	}
)
