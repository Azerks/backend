package server

import (
	"encoding/json"
	"fmt"
	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared"
	"github.com/sirupsen/logrus"
	"log/slog"
	"net/http"
)

type Server struct {
	Router *handlers.Router
	config *shared.Config
}

func New(config *shared.Config, log logrus.FieldLogger) *Server {
	router := handlers.NewRouter(log)
	return &Server{
		Router: router,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.Router)
}

func (s *Server) Respond(w http.ResponseWriter, _ *http.Request, status int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		slog.Error(
			"error encoding response", err.Error(),
			slog.String("status", fmt.Sprintf("%d", status)),
		)
	}
}

func (s *Server) RespondErr(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		return
	}

}
