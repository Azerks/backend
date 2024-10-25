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
