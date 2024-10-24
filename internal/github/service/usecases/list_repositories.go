package usecases

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

func (h *GetPublicGithubRepositoriesHandler) Handle(params GetPublicGithubRepositories) ([]RepositoryDTO, error) {
	repositories, err := h.githubRepository.ReadPublicRepositories(params.Filters)
	if err != nil {
		return nil, err
	}
	return repositories, nil
}
