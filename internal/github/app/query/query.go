package query

type (
	RepositoriesReader interface {
		ReadPublicRepositories() ([]Repository, error)
	}
)
