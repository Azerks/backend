package github

import "github.com/Scalingo/sclng-backend-test-v1/internal/github/app/query"

type GithubRepositoryModel struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Owner    struct {
		Login string `json:"login"`
	}
}

type GithubLanguageModel struct {
	Languages map[string]int `json:"languages"`
}

func toGithubRepositoriesQuery(m []GithubRepositoryModel) []query.Repository {
	repositories := make([]query.Repository, 0)
	for _, model := range m {
		repositories = append(repositories, query.Repository{
			FullName:       model.FullName,
			Owner:          model.Owner.Login,
			RepositoryName: model.FullName,
			Language:       nil,
		})
	}
	return repositories
}

//func toGithubLanguagesQuery(m GithubLanguageModel) query.Languages {
//	return query.Languages{}
//}
