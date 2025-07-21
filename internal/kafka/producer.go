package kafka

import (
	"encoding/json"
	"strings"

	kfk "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
	producer *kfk.Producer
}

func NewProducer(broker []string) (*Producer, error) {
	conf := &kfk.ConfigMap{
		"bootstrap.servers":    strings.Join(broker, ","),
		"log.connection.close": "true",
		"debug":                "all",
	}

	p, err := kfk.NewProducer(conf)
	if err != nil {
		return nil, err
	}

	return &Producer{producer: p}, nil
}

func (p *Producer) SendMessage(topic string, value []any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	kafkaMsg := &kfk.Message{
		TopicPartition: kfk.TopicPartition{
			Topic:     &topic,
			Partition: kfk.PartitionAny,
		},
		Key:   nil,
		Value: data,
	}

	kafkaChan := make(chan kfk.Event)

	if err := p.producer.Produce(kafkaMsg, kafkaChan); err != nil {
		return err
	}

	e := <-kafkaChan
	switch ev := e.(type) {
	case *kfk.Message:
		return nil
	case *kfk.Error:
		return ev
	default:
		return kfk.NewError(kfk.ErrUnknown, "unknown error", false)
	}
}

func (p *Producer) Close() {
	p.producer.Flush(5000)
	p.producer.Close()
}
