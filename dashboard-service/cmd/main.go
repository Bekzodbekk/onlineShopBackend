package main

import (
	dashboardservice "dashboard-service/internal/dashboard-service"
	"dashboard-service/internal/pkg/config"
	"dashboard-service/internal/pkg/kafka/consumer"
	"dashboard-service/internal/pkg/postgres"
	"dashboard-service/internal/repository"
	"dashboard-service/internal/service"
	"log"
)

func main() {
	cfg, err := config.LOAD("../config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	queries := postgres.Queries(db)
	repo := repository.NewDashboardREPO(queries)
	repoForConsumer := repository.NewForConsumerDashboardRepo(queries)
	service := service.NewService(repo)

	consumerClient, err := consumer.NewConsumeInit(cfg, *repoForConsumer)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := consumerClient.ConsumeMessage()
		if err != nil {
			log.Fatal(err)
		}
	}()
	runGRPC := dashboardservice.NewRunGRPC(service)
	log.Printf("Dashboard Service Running on :%d port", cfg.DashboardServicePort)
	if err := runGRPC.RUN(cfg); err != nil {
		log.Fatal(err)
	}
}
