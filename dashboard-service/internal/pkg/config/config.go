package config

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Postgres struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type Kafka struct {
	Host          string
	Port          int
	TopicDebts    string
	TopicProducts string
}

type Config struct {
	Postgres Postgres
	Kafka    Kafka

	DashboardServiceHost string
	DashboardServicePort int
}

func LOAD(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		Postgres: Postgres{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetInt("postgres.port"),
			Database: viper.GetString("postgres.dbname"),
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
		},

		Kafka: Kafka{
			Host:          viper.GetString("kafka.host"),
			Port:          viper.GetInt("kafka.port"),
			TopicDebts:    viper.GetString("kafka.topic-debts"),
			TopicProducts: viper.GetString("kafka.topic-products"),
		},

		DashboardServiceHost: viper.GetString("service.host"),
		DashboardServicePort: viper.GetInt("service.port"),
	}
	return &cfg, nil
}
