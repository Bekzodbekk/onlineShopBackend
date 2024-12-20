package config

import "github.com/spf13/viper"

type ServiceConfig struct {
	Host string
	Port int
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Kafka struct {
	Host          string
	Port          int
	TopicProducts string
	TopicDebts    string
}

type Config struct {
	Service  ServiceConfig
	Postgres PostgresConfig
	Kafka    Kafka
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		Service: ServiceConfig{
			Host: viper.GetString("service.host"),
			Port: viper.GetInt("service.port"),
		},
		Postgres: PostgresConfig{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetInt("postgres.port"),
			Database: viper.GetString("postgres.dbname"),
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
		},
		Kafka: Kafka{
			Host:          viper.GetString("kafka.host"),
			Port:          viper.GetInt("kafka.port"),
			TopicProducts: viper.GetString("kafka.topic-products"),
			TopicDebts:    viper.GetString("kafka.topic-debts"),
		},
	}

	return &cfg, nil
}
