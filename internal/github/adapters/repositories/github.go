package repositories

import (
	"context"
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

func (r *Repository) ReadPublicRepositories(ctx context.Context, filters usecases.RepositoriesFilters) ([]usecases.RepositoryDTO, error) {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)
	req, err := r.generateRequest(r.config.GithubApiURI + "/repositories")
	if err != nil {
		return nil, err
	}

	response, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: error while fetching repositories: %w", getGithubErr(response), err)
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
	errChan := make(chan error, r.config.WorkersPoolSize)
	doneChan := make(chan struct{})

	for i := 0; i < r.config.WorkersPoolSize; i++ {
		rs := r.SplitWorkersData(i, repositories)
		go func() {
			r.worker(ctx, &wg, workerParams{
				repositories: rs,
				resultChan:   resultChan,
				errorChan:    errChan,
				filters:      filters,
			})
		}()
	}

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	go func() {
		select {
		case <-ctx.Done():
			return
		case err := <-errChan:
			cancel(err)
			return
		}
	}()

	select {
	case <-ctx.Done():
		close(resultChan)
		close(errChan)
		return nil, ctx.Err()
	case <-doneChan:
		close(resultChan)
		close(errChan)
	}

	repos := make([]usecases.RepositoryDTO, 0)
	for result := range resultChan {
		repos = append(repos, result)
	}

	return repos, nil
}

type workerParams struct {
	repositories []GithubRepositoryModel
	resultChan   chan<- usecases.RepositoryDTO
	errorChan    chan<- error
	filters      usecases.RepositoriesFilters
}

func (r *Repository) worker(ctx context.Context, wg *sync.WaitGroup, params workerParams) {
	wg.Add(1)
	defer wg.Done()

	for _, repo := range params.repositories {
		select {
		case <-ctx.Done():
			return
		default:
			req, err := r.generateRequest(repo.LanguageURL)
			if err != nil {
				params.errorChan <- err
				return
			}

			response, err := r.client.Do(req)
			if err != nil {
				params.errorChan <- fmt.Errorf("%w: error while fetching repositories: %w", getGithubErr(response), err)
				return
			}
			defer response.Body.Close()

			languages := map[string]int{}
			if err := json.NewDecoder(response.Body).Decode(&languages); err != nil {
				params.errorChan <- err
				return
			}

			repo := toGithubRepositoriesQuery(repo, languages)
			if !shouldBeInclude(repo, params.filters) {
				continue
			}

			params.resultChan <- repo
		}
	}
}

func (r *Repository) generateRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.config.GithubToken))
	return req, nil
}

func (r *Repository) SplitWorkersData(i int, repositories []GithubRepositoryModel) []GithubRepositoryModel {
	start := len(repositories) / r.config.WorkersPoolSize * i
	end := len(repositories) / r.config.WorkersPoolSize * (i + 1)

	if end > len(repositories) {
		end = len(repositories)
	}
	return repositories[start:end]
}

func shouldBeInclude(repo usecases.RepositoryDTO, filters usecases.RepositoriesFilters) bool {
	if len(filters.Languages) > 0 {
		for _, lang := range filters.Languages {
			if repo.Language[lang] == 0 {
				return false
			}
		}
	}

	return true
}
