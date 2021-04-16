package mb

import (
	"fmt"
	"os"
	"sync"

	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/config"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	_ "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka/librdkafka_vendor"
)

type Service interface {
	Consume(topic []string, msg chan []byte) error
}

type kafkaMB struct {
	consumer  *kafka.Consumer
	topicChan map[string]chan []byte
}

var (
	srv  Service
	once = &sync.Once{}
)

// Load - create a kafka client
func Load(cfg config.Kafka) (err error) {
	once.Do(func() {
		p, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": cfg.BootstrapServer,
			"auto.offset.reset": "earliest",
			"group.id":          cfg.GroupID,
		})
		if err != nil {
			fmt.Printf("couldn't initialize kafka consumer, err : %s", err.Error())
			os.Exit(-1)
		}
		srv = &kafkaMB{consumer: p, topicChan: make(map[string]chan []byte)}
	})

	return
}

func GetMBService() Service {
	return srv
}
