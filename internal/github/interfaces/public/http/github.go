package http

import (
	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/app"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github/app/query"
	"net/http"
	"strconv"
)

func (s *Server) getRepositories() handlers.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		limit, err := strconv.Atoi(s.server.GetQueryParam(r, "limit"))
		if err != nil {
			limit = 0
		}
		filters := query.RepositoriesFilters{
			Limit:    limit,
			Language: s.server.GetQueryParam(r, "language"),
			License:  s.server.GetQueryParam(r, "license"),
		}
		repositories, err := s.app.GetGithubRepositories.Handle(app.GetPublicGithubRepositories{
			Filters: filters,
		})
		if err != nil {
			s.server.Log.WithError(err)
			s.server.RespondErr(w, r, err)
		}
		s.server.Respond(w, r, http.StatusOK, repositories)
		return nil
	}
}
