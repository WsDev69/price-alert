package mb

import (
	"time"

	"github.com/sirupsen/logrus"
)

func (srv kafkaMB) Consume(topics []string, msg chan []byte) error {
	err := srv.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	go srv.consume(topics, msg)

	return nil
}

func (srv kafkaMB) consume(topics []string, msg chan []byte) {
	if len(topics) == 0 {
		logrus.Debug("array of topics is empty")
		return
	}
	for i := range topics {
		srv.topicChan[topics[i]] = msg
	}

	for {
		msgKafka, err := srv.consumer.ReadMessage(5 * time.Second)
		if err != nil {
			logrus.Error("Consumer error:", err, msg)
			continue
		}

		for i := range topics {
			if c, ok := srv.topicChan[topics[i]]; ok {
				c <- msgKafka.Value
			}
		}

	}
}
