package services

import (
	"github.com/wsdev69/price-alert/notififcation-service/v0.0.1/src/config"
	"github.com/wsdev69/price-alert/notififcation-service/v0.0.1/src/services/mail"
)

// Service provides concrete services
type Service interface {
	GetMail() mail.Service
}

// serviceRepo combines all services
type serviceRepo struct {
	mail      mail.Service
}

var srv serviceRepo

// Load initializes services.
func Load(cfg *config.Configuration) error {
	srv.mail = mail.New(cfg.SMTP)

	return nil
}

// Get returns service repository
func Get() Service {
	return &srv
}

func (s serviceRepo) GetMail() mail.Service {
	return s.mail
}


