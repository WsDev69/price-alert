package quote

import (
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/config"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/quote"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/persistence/postgres"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/services/mb"
	"sync"
)

var topicCurrencies = "currencies"

type Service interface {
	QuoteConsumer() error
	UpdateCurrency(req quote.AddNewCurrencyRequest) error
}

type service struct {
	pg       *postgres.Client
	consumer mb.Service

	msgChan chan []byte

	urlNS    string
	urlQuote string
}

var (
	srv  service
	once = &sync.Once{}
)

// New returns service instance.
func New(pg *postgres.Client, consumer mb.Service, cfgNS config.NotificationService, cfgQuote config.QuoteService) Service {
	once.Do(func() {
		srv = service{
			pg:       pg,
			consumer: consumer,
		}
		srv.msgChan = make(chan []byte)
		srv.urlNS = cfgNS.URL + "/mail/send"
		srv.urlQuote = cfgQuote.URL + "/quotes"
		go srv.consumer.Consume([]string{topicCurrencies}, srv.msgChan)
		go srv.QuoteConsumer()
	})

	return srv
}
