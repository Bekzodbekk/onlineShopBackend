package consumer

import (
	"context"
	"dashboard-service/internal/pkg/config"
	"dashboard-service/internal/repository"
	"encoding/json"
	"fmt"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

type IConsumeInit interface {
	ConsumeMessage() error
	Close() error
}

type ConsumeInit struct {
	clientProduct *kgo.Client
	clientDebt    *kgo.Client
	topicDebt     string
	topicProduct  string
	reportRepo    repository.DashboardREPO
}

func NewConsumeInit(cfg *config.Config, repo repository.DashboardREPO) (IConsumeInit, error) {

	address := fmt.Sprintf("%s:%d", cfg.Kafka.Host, cfg.Kafka.Port)
	clientProduct, err := kgo.NewClient(
		kgo.SeedBrokers(address),
		kgo.ConsumeTopics(cfg.Kafka.TopicProducts),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtEnd()), // Eng oxirgi offsetdan boshlash
	)
	if err != nil {
		return nil, err
	}
	clientDebt, err := kgo.NewClient(
		kgo.SeedBrokers(address),
		kgo.ConsumeTopics(cfg.Kafka.TopicDebts),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtEnd()),
	)
	if err != nil {
		return nil, err
	}
	return &ConsumeInit{
		clientDebt:    clientDebt,
		clientProduct: clientProduct,
		topicDebt:     cfg.Kafka.TopicDebts,
		topicProduct:  cfg.Kafka.TopicProducts,
		reportRepo:    repo,
	}, nil
}

func (c *ConsumeInit) ConsumeMessage() error {
	ctx := context.Background()

	// Channel to capture errors from goroutines
	errCh := make(chan error, 2)

	// Goroutine to handle messages from product topic
	go func() {
		for {
			fetches := c.clientProduct.PollFetches(ctx)
			if errs := fetches.Errors(); len(errs) > 0 {
				for _, err := range errs {
					log.Print(err)
				}
				errCh <- fmt.Errorf("error consuming product messages: %v", errs)
				return
			}

			fetches.EachPartition(func(partition kgo.FetchTopicPartition) {
				for _, record := range partition.Records {
					fmt.Printf("Product Message - Key: %s, Value: %s\n", string(record.Key), string(record.Value))

					key := string(record.Key)
					value := record.Value

					if key == "commandProductsForDashboard" {
						kafkaResp := repository.ProductDashboard{}
						err := json.Unmarshal(value, &kafkaResp)
						if err != nil {
							log.Printf("Error unmarshalling product message: %v", err)
							continue
						}

						err = c.reportRepo.SellProductInsertDashboard(ctx, &kafkaResp)
						if err != nil {
							log.Printf("Error inserting product data: %v", err)
						}
					}
				}
			})
		}
	}()

	// Goroutine to handle messages from debt topic
	go func() {
		for {
			fetches := c.clientDebt.PollFetches(ctx)
			if errs := fetches.Errors(); len(errs) > 0 {
				for _, err := range errs {
					log.Print(err)
				}
				errCh <- fmt.Errorf("error consuming debt messages: %v", errs)
				return
			}

			fetches.EachPartition(func(partition kgo.FetchTopicPartition) {
				for _, record := range partition.Records {
					fmt.Printf("Debt Message - Key: %s, Value: %s\n", string(record.Key), string(record.Value))

					key := string(record.Key)
					value := record.Value

					if key == "addOrPaymentDebt" {
						kafkaResp := repository.DebtDashboard{}
						err := json.Unmarshal(value, &kafkaResp)
						if err != nil {
							log.Printf("Error unmarshalling debt message: %v", err)
							continue
						}

						err = c.reportRepo.AddDebtInsertDashboard(ctx, &kafkaResp)
						if err != nil {
							log.Printf("Error inserting debt data: %v", err)
						}
					}
				}
			})
		}
	}()

	// Wait for errors or block forever
	err := <-errCh
	return err
}

func (c *ConsumeInit) Close() error {
	c.clientDebt.Close()
	c.clientProduct.Close()
	log.Print("Kafka client closed successfully")
	return nil
}
