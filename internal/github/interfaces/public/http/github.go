package http

import (
	"github.com/Scalingo/go-handlers"
	"net/http"
)

func (s *Server) getRepositories() handlers.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		repositories, err := s.app.Queries.GetGithubRepositories.Handle()
		if err != nil {
			s.server.Log.WithError(err)
			s.server.RespondErr(w, r, err)
		}
		s.server.Respond(w, r, http.StatusOK, repositories)
		return nil
	}
}
