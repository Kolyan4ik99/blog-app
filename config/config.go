package config

import (
	"github.com/spf13/viper"
)

const (
	ProdEnv = "config/prod"
	QaEnv   = "config/qa"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}

func Init(folder, file string) (*Config, error) {
	viper.AddConfigPath(folder)
	viper.SetConfigName(file)

	cfg := new(Config)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(cfg)
	return cfg, err
}
