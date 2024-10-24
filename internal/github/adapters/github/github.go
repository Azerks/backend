package github

import (
	"encoding/json"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/app/query"
	"net/http"
)

type Repository struct {
	uri string
}

func New() *Repository {
	return &Repository{
		uri: "https://api.github.com",
	}
}

func (r *Repository) ReadPublicRepositories() ([]query.Repository, error) {
	response, err := http.Get(r.uri + "/repositories")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	repositories := make([]GithubRepositoryModel, 0)
	if err := json.NewDecoder(response.Body).Decode(&repositories); err != nil {
		return nil, err
	}

	return toGithubRepositoriesQuery(repositories), nil
}
