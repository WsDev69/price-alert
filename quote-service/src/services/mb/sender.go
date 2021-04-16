package mb

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func (srv kafkaMB) Publish(topic string, message []byte) error {
	var kafkaMsg = &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:          message,

	}
	return srv.producer.Produce(kafkaMsg, nil)
}

