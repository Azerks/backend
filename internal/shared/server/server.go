package server

import (
	"encoding/json"
	"fmt"
	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-utils/errors/v2"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	Router *handlers.Router
	Config *shared.Config
	Log    logrus.FieldLogger
}

func New(config *shared.Config, log logrus.FieldLogger) *Server {
	router := handlers.NewRouter(log)
	return &Server{
		Router: router,
		Config: config,
		Log:    log,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.Config.Port), s.Router)
}

func (s *Server) Respond(w http.ResponseWriter, _ *http.Request, status int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		s.Log.WithError(err)
	}
}

func (s *Server) RespondErr(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		return
	}

	s.Respond(w, r, http.StatusInternalServerError, nil)
}

func (s *Server) Decode(_ http.ResponseWriter, r *http.Request, v interface{}) error {
	if r.ContentLength == 0 {
		return nil
	}

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return errors.New(r.Context(), "Invalid input cannot decode to json")
	}

	return nil
}
