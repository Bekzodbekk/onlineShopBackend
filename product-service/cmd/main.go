package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	app "product-service/internal/app"
	config "product-service/internal/config"
	"product-service/internal/kafka/consumer"
	kafka "product-service/internal/kafka/producer"
	service "product-service/internal/service"
	"product-service/internal/storage"
	mongo "product-service/internal/storage/mongodb"
	"product-service/logger"
	"syscall"
	"time"
)

func main() {

	logger.InitLog()

	cfg, err := config.Load("./config/config.yaml")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Configuration loaded")

	db, err := mongo.NewConnection(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Connected to MongoDB")

	repo := storage.NewProductRepository(db)

	producer, err := kafka.NewProducerInit(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer producer.Close()

	srv := service.NewProductService(repo, producer)

	application := app.New(*srv, cfg.Service.Port)
	consume, err := consumer.NewConsumeInit(cfg, *db)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		consume.ConsumeMessage()
	}()
	go func() {
		application.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sig := <-stop

	logger.Info("Received signal: ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	application.Stop()
	<-ctx.Done()
	logger.Info("Graceful shutdown complete.")
}
