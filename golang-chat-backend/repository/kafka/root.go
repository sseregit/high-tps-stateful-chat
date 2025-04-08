package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"websocket-high-tps-chat/config"
)

type Kafka struct {
	cfg *config.Config

	producer *kafka.Producer
}

func NewKafka(cfg *config.Config) (*Kafka, error) {
	k := &Kafka{cfg: cfg}

	var err error

	if k.producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.URL,
		"client.id":         cfg.Kafka.ClientID,
		"acks":              "all",
	}); err != nil {
		return nil, err
	} else {
		return k, nil
	}
}

func (k *Kafka) PublishEvent(topic string, value []byte, ch chan kafka.Event) (kafka.Event, error) {
	if err := k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
	}, ch); err != nil {
		return nil, err
	} else {
		return <-ch, nil
	}

}
