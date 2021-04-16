package services

import (
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/config"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/persistence/postgres"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/services/alert"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/services/mb"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/services/quote"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/services/user"
)

// Service provides concrete services
type Service interface {
	User() user.Service
	Alert() alert.Service
	Quote() quote.Service
}

// serviceRepo combines all services
type serviceRepo struct {
	user  user.Service
	alert alert.Service
	quote quote.Service
}

var srv serviceRepo

// Load initializes services.
func Load(postgresCli *postgres.Client, kafkaSrv mb.Service,
	cfg *config.Configuration) error {
	userService := user.New(postgresCli)
	quoteService := quote.New(postgresCli, kafkaSrv, cfg.NotificationService, cfg.QuoteService)
	alertService := alert.New(postgresCli, quoteService)

	srv.alert = alertService
	srv.user = userService
	srv.quote = quoteService

	return nil
}

// Get returns service repository
func Get() Service {
	return &srv
}

func (s *serviceRepo) User() user.Service {
	return s.user
}

func (s *serviceRepo) Alert() alert.Service {
	return s.alert
}

func (s *serviceRepo) Quote() quote.Service {
	return s.quote
}
