package config

import (
	"github.com/Kolyan4ik99/blog-app/pkg/postgres"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

const (
	ProdEnv = "config/prod"
	QaEnv   = "config/qa"
)

type Config struct {
	DB postgres.Config
}

func Init(folder, file string) (*Config, error) {
	viper.AddConfigPath(folder)
	viper.SetConfigName(file)

	cfg := new(Config)

	err := envconfig.Process("db", &cfg.DB)
	if err != nil {
		return nil, err
	}
	return cfg, viper.ReadInConfig()
}
