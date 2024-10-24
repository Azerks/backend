package main

import (
	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/sclng-backend-test-v1/internal/github"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared"
	"github.com/Scalingo/sclng-backend-test-v1/internal/shared/server"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	log := logger.Default()
	log.Info("Initializing app")

	if err := run(log); err != nil {
		log.WithError(err).Error("Fail to run the application")
		os.Exit(1)
	}
}

func run(log logrus.FieldLogger) error {
	cfg, err := shared.NewConfig()
	if err != nil {
		log.WithError(err).Error("Fail to initialize configuration")
		os.Exit(1)
	}

	mux := server.New(cfg, log)
	_ = github.New(mux)

	return mux.Serve()
}

//func main() {
//	log := logger.Default()
//	log.Info("Initializing app")
//	cfg, err := newConfig()
//	if err != nil {
//		log.WithError(err).Error("Fail to initialize configuration")
//		os.Exit(1)
//	}
//
//	log.Info("Initializing routes")
//	router := handlers.NewRouter(log)
//	router.HandleFunc("/ping", pongHandler)
//	// Initialize web server and configure the following routes:
//	// GET /repos
//	// GET /stats
//
//	log = log.WithField("port", cfg.Port)
//	log.Info("Listening...")
//	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
//	if err != nil {
//		log.WithError(err).Error("Fail to listen to the given port")
//		os.Exit(2)
//	}
//}
//
//func pongHandler(w http.ResponseWriter, r *http.Request, _ map[string]string) error {
//	log := logger.Get(r.Context())
//	w.Header().Add("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//
//	err := json.NewEncoder(w).Encode(map[string]string{"status": "pong"})
//	if err != nil {
//		log.WithError(err).Error("Fail to encode JSON")
//	}
//	return nil
//}
