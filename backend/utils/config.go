package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	RestfulHost string `mapstructure:"RESTFUL_HOST"`
	RestfulPort string `mapstructure:"RESTFUL_PORT"`
	DbDatabase  string `mapstructure:"DB_DATABASE"`
	DbUser      string `mapstructure:"POSTGRES_USER"`
	DbPassword  string `mapstructure:"POSTGRES_PASSWORD"`
	DbHost      string `mapstructure:"DB_HOST"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("dotenv")
	viper.AutomaticEnv()

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	return &cfg, nil
}
