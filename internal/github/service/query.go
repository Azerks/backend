package service

type (
	RepositoriesReader interface {
		ReadPublicRepositories() ([]RepositoryDTO, error)
	}
)
