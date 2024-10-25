package service

import (
	"github.com/Scalingo/sclng-backend-test-v1/common"
	"github.com/Scalingo/sclng-backend-test-v1/internal/adapters/repositories"
	"github.com/Scalingo/sclng-backend-test-v1/internal/service/usecases"
	"net/http"
)

type Service struct {
	ListRepositories usecases.ListPublicRepositoriesHandler
}

func New(config *common.Config) *Service {

	client := http.Client{}
	repository := repositories.New(config, &client)

	return &Service{
		ListRepositories: usecases.NewListPublicRepositoriesHandler(repository),
	}
}
