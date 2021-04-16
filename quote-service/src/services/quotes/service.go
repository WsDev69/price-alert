package quotes

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync"

	"github.com/wsdev69/price-alert/quote-service/v0.0.1/src/config"
	"github.com/wsdev69/price-alert/quote-service/v0.0.1/src/persistence/redis"
	"github.com/wsdev69/price-alert/quote-service/v0.0.1/src/services/mb"
)

const (
	currenciesKey = "currencies"
	wsActionSub = "SubAdd"
	wsTypeTicker = "2"

	coinbaseSubscribeFormat = "2~Coinbase~%s~%s"
)

type Service interface {
	AddCurrencyPair(fromSymbol, toSymbol string) error
	Start() error
}

type service struct {
	messageBrokerSrv mb.Service
	redisCli         redis.RedisCli
	conn             *websocket.Conn

	close chan struct{}
}

var (
	srv  Service
	once = &sync.Once{}
)

// New returns service instance.
func New(messageBrokerSrv mb.Service,
	redisCli redis.RedisCli,
	config config.CryptoCompare) Service {
	once.Do(func() {
		c, _, err := websocket.DefaultDialer.Dial(config.URL+config.APIKey, nil)
		if err != nil {
			logrus.Fatal("dial:", err)
		}
		srv = service{
			messageBrokerSrv: messageBrokerSrv,
			redisCli:         redisCli,
			conn:             c,
			close:            make(chan struct{}),
		}
	})
	return srv
}
