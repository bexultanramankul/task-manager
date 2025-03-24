package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func setEnv(key, value string) {
	if err := os.Setenv(key, value); err != nil {
		log.Fatalf("Ошибка установки %s: %v", key, err)
	}
}

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Ошибка загрузки конфигурации:", err)
	}

	setEnv("SERVER_PORT", viper.GetString("server.port"))
	setEnv("DB_USER", viper.GetString("database.user"))
	setEnv("DB_PASSWORD", viper.GetString("database.password"))
	setEnv("DB_NAME", viper.GetString("database.name"))
}
