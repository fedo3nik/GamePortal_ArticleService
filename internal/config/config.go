package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresURL string `envconfig:"POSTGRES_URL"`
	Host        string `envconfig:"ARTICLE_SERVICE_HOST"`
	Port        string `envconfig:"ARTICLE_SERVICE_PORT"`
	GrpcPort    string `envconfig:"GRPC_PORT"`
}

func NewConfig() (*Config, error) {
	var c Config

	err := envconfig.Process("articleservice", &c)
	if err != nil {
		log.Printf("Process config file error: %v\n", err)
		return nil, err
	}

	return &c, nil
}
