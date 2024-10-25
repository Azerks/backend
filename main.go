package main

import (
	"github.com/Scalingo/go-utils/logger"
	"github.com/Scalingo/sclng-backend-test-v1/common"
	"github.com/Scalingo/sclng-backend-test-v1/common/server"
	"github.com/Scalingo/sclng-backend-test-v1/internal"
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
	cfg, err := common.NewConfig()
	if err != nil {
		log.WithError(err).Error("Fail to initialize configuration")
		os.Exit(1)
	}

	mux := server.New(cfg, log)
	_ = internal.New(mux)

	return mux.Serve()
}
