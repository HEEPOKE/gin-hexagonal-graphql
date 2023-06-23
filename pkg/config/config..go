package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Cfg *Config
)

type Config struct {
	DB_HOST     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_PORT     string
	DB_SSL      string
	DB_TIMEZONE string
	PORT        int
	PRIVATE_KEY string
	PUBLIC_KEY  string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	config := &Config{
		DB_HOST:     viper.GetString("DB_HOST"),
		DB_NAME:     viper.GetString("DB_NAME"),
		DB_USER:     viper.GetString("DB_USER"),
		DB_PASSWORD: viper.GetString("DB_PASSWORD"),
		DB_PORT:     viper.GetString("DB_PORT"),
		DB_SSL:      viper.GetString("DB_SSL"),
		DB_TIMEZONE: viper.GetString("DB_TIMEZONE"),
		PORT:        viper.GetInt("PORT"),
		PRIVATE_KEY: viper.GetString("PRIVATE_KEY"),
		PUBLIC_KEY:  viper.GetString("PUBLIC_KEY"),
	}

	Cfg = config

	return config, nil
}
