package server

import (
	"encoding/json"
	"fmt"
	"github.com/Scalingo/go-handlers"
	"github.com/Scalingo/go-utils/errors/v2"
	"github.com/Scalingo/sclng-backend-test-v1/common"
	"github.com/Scalingo/sclng-backend-test-v1/common/errs"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	Router *handlers.Router
	Config *common.Config
	Log    logrus.FieldLogger
}

func New(config *common.Config, log logrus.FieldLogger) *Server {
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

	errs.HTTP(w, r, err)
}

func (s *Server) GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
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
