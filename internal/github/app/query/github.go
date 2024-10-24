package query

type GetPublicGithubRepositories struct {
}
type GetPublicGithubRepositoriesHandler struct {
	githubRepository RepositoriesReader
}

func NewGithubRepositoriesHandler(githubRepository RepositoriesReader) GetPublicGithubRepositoriesHandler {
	return GetPublicGithubRepositoriesHandler{
		githubRepository: githubRepository,
	}
}

func (h *GetPublicGithubRepositoriesHandler) Handle() ([]Repository, error) {
	repositories, err := h.githubRepository.ReadPublicRepositories()
	if err != nil {
		return nil, err
	}
	return repositories, nil
}
