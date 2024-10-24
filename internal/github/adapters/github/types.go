package github

import (
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/service/usecase"
)

type GithubRepositoryModel struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Owner    struct {
		Login string `json:"login"`
	}
	LanguageURL string `json:"languages_url"`
}

func toGithubRepositoriesQuery(m GithubRepositoryModel, languages map[string]int) usecase.RepositoryDTO {
	return usecase.RepositoryDTO{
		FullName:   m.FullName,
		Owner:      m.Owner.Login,
		Repository: m.FullName,
		Language:   languages,
	}
}
