package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"golang-chat-controller/config"
)

type Kafka struct {
	cfg *config.Config

	consumer *kafka.Consumer
}

func NewKafka(cfg *config.Config) (*Kafka, error) {
	k := &Kafka{cfg: cfg}

	var err error

	if k.consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.URL,
		"group.id":          cfg.Kafka.GroupID,
		"auto.offset.reset": "latest", // 최근값만 읽겠다.
	}); err != nil {
		return nil, err
	} else {
		return k, nil
	}
}

func (k *Kafka) Poll(timeoutMs int) kafka.Event {
	return k.consumer.Poll(timeoutMs)
}

func (k *Kafka) RegisterSubTopic(topic string) error {
	if err := k.consumer.Subscribe(topic, nil); err != nil {
		return err
	} else {
		return nil
	}
}
