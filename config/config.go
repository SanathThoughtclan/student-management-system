package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	URI      string
	Database string
}

type JWTConfig struct {
	SecretKey string `mapstructure:"secret_key"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	config := &Config{}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	return config
}
