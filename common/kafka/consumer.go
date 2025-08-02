package kafka

import (
	"strings"

	kfk "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer struct {
	consumer *kfk.Consumer
}

func NewConsumer(brocker []string, topic []string) (*Consumer, error) {
	conf := &kfk.ConfigMap{
		"bootstrap.servers":  strings.Join(brocker, ","),
		"group.id":           "user-recommendations",
		"auto.offset.reset":  "earliest",
		"session.timeout.ms": 6000,
	}

	c, err := kfk.NewConsumer(conf)
	if err != nil {
		return nil, err
	}

	if err = c.SubscribeTopics(topic, nil); err != nil {
		return nil, err
	}

	return &Consumer{consumer: c}, nil
}
