package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"product-service/internal/config"
	"product-service/internal/storage/mongodb"
	"product-service/logger"

	"github.com/twmb/franz-go/pkg/kgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IConsumeInit interface {
	ConsumeMessage() error
	Close() error
}

type ConsumeInit struct {
	client     *kgo.Client
	topicDebts string
	Mongo      mongodb.Mongo
}

func NewConsumeInit(cfg *config.Config, mongo mongodb.Mongo) (IConsumeInit, error) {
	address := fmt.Sprintf("%s:%d", cfg.Kafka.Host, cfg.Kafka.Port)
	client, err := kgo.NewClient(
		kgo.SeedBrokers(address),
		kgo.ConsumeTopics(cfg.Kafka.TopicDebts),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtEnd()), // Eng oxirgi offsetdan boshlash
	)
	if err != nil {
		logger.Error("Failed to create Kafka client", "error", err)
		return nil, err
	}
	logger.Info("Kafka consumer initialized successfully")
	return &ConsumeInit{
		client:     client,
		topicDebts: cfg.Kafka.TopicDebts,
		Mongo:      mongo,
	}, nil
}

func (c *ConsumeInit) ConsumeMessage() error {
	ctx := context.Background()

	for {
		fetches := c.client.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			for _, err := range errs {
				logger.Error("Error consuming messages", "error", err)
			}
			return fmt.Errorf("error consuming messages: %v", errs)
		}

		fetches.EachPartition(func(partition kgo.FetchTopicPartition) {
			for _, record := range partition.Records {
				if string(record.Key) == "commandForProductAddDebtColor" {
					logger.Info(fmt.Sprintf("Received message - Key: %s, Value: %s", string(record.Key), string(record.Value)))
					kafkaResp := make(map[string]string)

					if err := json.Unmarshal(record.Value, &kafkaResp); err != nil {
						continue
					}

					productId := kafkaResp["id"]
					colorKey := kafkaResp["color"]

					id, err := primitive.ObjectIDFromHex(productId)
					if err != nil {
						logger.Error("Invalid product ID", "error", err)
						continue
					}

					filter := bson.M{"_id": id}
					update := bson.M{"$inc": bson.M{
						"count":              -1,
						"colors." + colorKey: -1,
					}}

					if _, err := c.Mongo.Collection.UpdateOne(ctx, filter, update); err != nil {
						logger.Error("Failed to update MongoDB", "error", err)
						continue
					}
					logger.Info("MongoDB updated successfully for commandForProductAddDebtColor")
				}
			}
		})
	}
}

func (c *ConsumeInit) Close() error {
	c.client.Close()
	logger.Info("Kafka client closed successfully")
	return nil
}
