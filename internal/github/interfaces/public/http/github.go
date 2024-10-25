package http

import (
	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/service/usecases"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) getRepositories() handlers.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		limit, err := strconv.Atoi(s.server.GetQueryParam(r, "limit"))
		if err != nil {
			limit = 0
		}
		filters := usecases.RepositoriesFilters{
			Limit:     limit,
			Languages: strings.Split(s.server.GetQueryParam(r, "language"), ","),
		}
		repositories, err := s.app.GetGithubRepositories.Handle(r.Context(), usecases.GetPublicGithubRepositories{
			Filters: filters,
		})
		if err != nil {
			s.server.Log.WithError(err)
			s.server.RespondErr(w, r, err)
			return err
		}
		s.server.Respond(w, r, http.StatusOK, repositories)
		return nil
	}
}
