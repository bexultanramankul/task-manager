package config

import (
	"task-manager/pkg/logger"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	SSLMode  string `mapstructure:"sslmode"`
}

var AppConfig Config

func LoadConfig() {
	log := logger.Log

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Warn("config.yaml not found, loading from environment variables")
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatal("Failed to parse configuration: ", err)
	}

	if AppConfig.Database.User == "" || AppConfig.Database.Password == "" {
		log.Fatal("Missing required storage configuration (user/password)")
	}

	log.Info("Configuration loaded successfully")
}
