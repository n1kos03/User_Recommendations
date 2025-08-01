package kafka

import (
	"encoding/json"
	"log/slog"
	"strings"

	kfk "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/n1kos03/User_Recommendations/common/models"
	"github.com/n1kos03/User_Recommendations/services/recommendations/handlers"
	// "github.com/n1kos03/User_Recommendations/services/recommendations/handlers"
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

func (c *Consumer) Start(r *handlers.RecommendationService) {
	go func() {
		for {
			msg, err := c.consumer.ReadMessage(-1)
			if err != nil {
				slog.Error("Error reading message from Kafka", "Error: ", err)
				continue
			}

			switch *msg.TopicPartition.Topic {
			case "user_updates":
				var event models.UserMessageEvent
				slog.Info("Received message from Kafka", "Message: ", string(msg.Value))
				if err := json.Unmarshal(msg.Value, &event); err != nil {
					slog.Error("Error unmarshalling message from Kafka", "Error: ", err)
					continue
				}
				r.HandleUserUpdateEvent(event)
			case "product_updates":
				// TODO: Handle product updates
				continue
			}
		}
	}()
}
