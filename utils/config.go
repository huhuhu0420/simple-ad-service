package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	RestfulHost string `mapstructure:"RESTFUL_HOST"`
	RestfulPort string `mapstructure:"RESTFUL_PORT"`
	DbName      string `mapstructure:"DB_NAME"`
	DbUser      string `mapstructure:"DB_USER"`
	DbPassword  string `mapstructure:"DB_PASSWORD"`
	DbHost      string `mapstructure:"PG_HOST"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	var cfg Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	return &cfg, nil
}