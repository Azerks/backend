package github

import (
	"encoding/json"
	"fmt"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/service/usecase"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared"
	"net/http"
	"sync"
)

type Repository struct {
	config *shared.Config
}

func New(config *shared.Config) *Repository {
	return &Repository{
		config: config,
	}
}

func (r *Repository) ReadPublicRepositories(filters usecase.RepositoriesFilters) ([]usecase.RepositoryDTO, error) {
	req, err := http.NewRequest("GET", r.config.GithubApiURI+"/repositories", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.config.GithubToken))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	repositories := make([]GithubRepositoryModel, 0)
	if err := json.NewDecoder(response.Body).Decode(&repositories); err != nil {
		return nil, err
	}

	if filters.Limit != 0 {
		repositories = repositories[:filters.Limit]
	}

	var wg sync.WaitGroup
	resultChan := make(chan usecase.RepositoryDTO, len(repositories))

	for i := 0; i < r.config.WorkersPoolSize; i++ {
		repos := repositories[len(repositories)/r.config.WorkersPoolSize*i : len(repositories)/r.config.WorkersPoolSize*(i+1)]
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := worker(r.config, repos, resultChan, filters)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()
	close(resultChan)

	repos := make([]usecase.RepositoryDTO, 0)
	for result := range resultChan {
		repos = append(repos, result)
	}

	return repos, nil
}

func worker(config *shared.Config, i []GithubRepositoryModel, resultChan chan<- usecase.RepositoryDTO, filters usecase.RepositoriesFilters) error {
	for _, repo := range i {
		req, err := http.NewRequest("GET", repo.LanguageURL, nil)
		if err != nil {
			return err
		}
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.GithubToken))

		response, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		languages := map[string]int{}
		if err := json.NewDecoder(response.Body).Decode(&languages); err != nil {
			return err
		}

		repo := toGithubRepositoriesQuery(repo, languages)
		if !shouldBeInclude(repo, filters) {
			continue
		}

		resultChan <- repo
	}
	return nil
}

func shouldBeInclude(repo usecase.RepositoryDTO, filters usecase.RepositoriesFilters) bool {
	if filters.Language != "" && repo.Language[filters.Language] == 0 {
		return false
	}

	return true
}
