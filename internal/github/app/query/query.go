package query

type (
	RepositoriesReader interface {
		ReadPublicRepositories() ([]RepositoryDTO, error)
	}
)
