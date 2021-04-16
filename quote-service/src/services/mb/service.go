package mb

import (
	"github.com/sirupsen/logrus"
	"github.com/wsdev69/price-alert/quote-service/v0.0.1/src/config"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"os"
	"sync"
)

type Service interface {
	Publish(topic string, message []byte) error
}

type kafkaMB struct {
	producer *kafka.Producer
}

var (
	srv  Service
	once = &sync.Once{}
)

// Load - create a kafka client
func Load(cfg config.Kafka) (err error) {
	once.Do(func() {
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": cfg.BootstrapServer,
		})
		if err != nil {
			logrus.Fatal("couldn't initialize kafka producer, err : %s", err.Error())
			os.Exit(-1)
		}
		srv = &kafkaMB{producer: p}
	})

	return
}

func GetMBService() Service {
	return srv
}
