package main

import (
	"fmt"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/config"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/handlers"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/persistence/postgres"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/services"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/services/mb"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/validator"
	"log"
	"net/http"
	"os"
)

// defaultConfigPath defines a path to JSON-config file
const defaultConfigPath = "config.json"

func main() {
	err := config.Load(defaultConfigPath)
	if err != nil {
		log.Fatalf("Failed to initialize Config: %s", err.Error())
	}

	// setup log-level
	logLevel, err := logrus.ParseLevel(config.Config.LogPreset)
	if err != nil {
		logrus.Fatal("cannot parse log level", err)
	}
	logrus.SetLevel(logLevel)

	corsInstance := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:*", "https://localhost:*",
			"http://127.0.0.1:*", "https://127.0.0.1:*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodPut, http.MethodPatch},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
	})

	server := &http.Server{
		Addr:    config.Config.ListenURL,
		Handler: corsInstance.Handler(handlers.NewRouter()),
	}

	if err := postgres.Load(&config.Config.Postgres, logrus.New()); err != nil {
		logrus.Fatal(fmt.Sprintf("cannot connect to the Postgres server with config [%+v]: %v",
			config.Config.Postgres, err))
	}

	if err := mb.Load(config.Config.Kafka); err != nil {
		logrus.Fatal(fmt.Sprintf("cannot initalize kafka connection [%+v]: %v",
			config.Config.Postgres, err))
	}

	if err = validator.Load(); err != nil {
		logrus.Fatal(fmt.Sprintf("cannot initialize validator: %v", err))
	}

	if err := services.Load(postgres.GetDB(), mb.GetMBService(), &config.Config); err != nil {
		logrus.Fatal("cannot load services", err)
	}

	err = server.ListenAndServe()
	if err != nil {
		logrus.Error("Failed to initialize HTTP server", "error", err)
		os.Exit(1)
	}
}
