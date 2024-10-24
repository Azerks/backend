package github

import (
	"encoding/json"
	"fmt"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/app/query"
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

func (r *Repository) ReadPublicRepositories() ([]query.RepositoryDTO, error) {
	response, err := http.Get(r.config.GithubApiURI + "/repositories")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	repositories := make([]GithubRepositoryModel, 0)
	if err := json.NewDecoder(response.Body).Decode(&repositories); err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	resultChan := make(chan query.RepositoryDTO, len(repositories))

	for i := 0; i < r.config.Workers; i++ {
		repos := repositories[len(repositories)/r.config.Workers*i : len(repositories)/r.config.Workers*(i+1)]
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := worker(repos, resultChan)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	wg.Wait()
	close(resultChan)

	repos := make([]query.RepositoryDTO, 0)
	for result := range resultChan {
		repos = append(repos, result)
	}

	return repos, nil
}

func worker(i []GithubRepositoryModel, resultChan chan<- query.RepositoryDTO) error {
	for _, repo := range i {
		response, err := http.Get(repo.LanguageURL)
		if err != nil {
			return err
		}

		languages := map[string]int{}
		if err := json.NewDecoder(response.Body).Decode(&languages); err != nil {
			return err
		}

		resultChan <- toGithubRepositoriesQuery(repo, languages)

		err = response.Body.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
