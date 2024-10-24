package http

import (
	"github.com/Scalingo/go-handlers"
	"net/http"
)

func (s *Server) getRepositories() handlers.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		//s.server.Respond()
		return nil
	}
}
