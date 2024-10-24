package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/service/usecases"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared"
	"net/http"
	"sync"
)

type Repository struct {
	client *http.Client
	config *shared.Config
}

func New(config *shared.Config, client *http.Client) *Repository {
	return &Repository{
		config: config,
		client: client,
	}
}

func (r *Repository) ReadPublicRepositories(filters usecases.RepositoriesFilters) ([]usecases.RepositoryDTO, error) {
	req, err := http.NewRequest("GET", r.config.GithubApiURI+"/repositories", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.config.GithubToken))

	response, err := r.client.Do(req)
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
	resultChan := make(chan usecases.RepositoryDTO, len(repositories))

	for i := 0; i < r.config.WorkersPoolSize; i++ {
		repos := repositories[len(repositories)/r.config.WorkersPoolSize*i : len(repositories)/r.config.WorkersPoolSize*(i+1)]
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := worker(r, repos, resultChan, filters)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()
	close(resultChan)

	repos := make([]usecases.RepositoryDTO, 0)
	for result := range resultChan {
		repos = append(repos, result)
	}

	return repos, nil
}

func worker(r *Repository, i []GithubRepositoryModel, resultChan chan<- usecases.RepositoryDTO, filters usecases.RepositoriesFilters) error {
	for _, repo := range i {
		req, err := http.NewRequest("GET", repo.LanguageURL, nil)
		if err != nil {
			return err
		}
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.config.GithubToken))

		response, err := r.client.Do(req)
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

func shouldBeInclude(repo usecases.RepositoryDTO, filters usecases.RepositoriesFilters) bool {
	if filters.Language != "" && repo.Language[filters.Language] == 0 {
		return false
	}

	return true
}
