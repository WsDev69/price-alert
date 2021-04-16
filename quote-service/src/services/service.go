package services

import (
	"github.com/sirupsen/logrus"
	"github.com/wsdev69/price-alert/quote-service/v0.0.1/src/config"
	"github.com/wsdev69/price-alert/quote-service/v0.0.1/src/persistence/redis"
	"github.com/wsdev69/price-alert/quote-service/v0.0.1/src/services/mb"
	"github.com/wsdev69/price-alert/quote-service/v0.0.1/src/services/quotes"
)

// Service provides concrete services
type Service interface {
	GetMessageBroker() mb.Service
	GetQuote() quotes.Service
}

// serviceRepo combines all services
type serviceRepo struct {
	messageBroker mb.Service
	quoteSrv      quotes.Service
}

var srv serviceRepo

// Load initializes services.
func Load(redisCli redis.RedisCli, cfg *config.Configuration) error {
	err := mb.Load(cfg.Kafka)
	if err != nil {
		logrus.Error("couldn't initialize kafka service", err)
		return err
	}

	srv.messageBroker = mb.GetMBService()
	srv.quoteSrv = quotes.New(srv.messageBroker, redisCli, cfg.CryptoCompare)

	err = srv.quoteSrv.Start()
	if err != nil {
		return err
	}

	return nil
}

// Get returns service repository
func Get() Service {
	return &srv
}

func (s serviceRepo) GetMessageBroker() mb.Service {
	return s.messageBroker
}

func (s serviceRepo) GetQuote() quotes.Service {
	return s.quoteSrv
}

