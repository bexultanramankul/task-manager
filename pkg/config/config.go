package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Ошибка загрузки конфигурации:", err)
	}

	os.Setenv("SERVER_PORT", viper.GetString("server.port"))
	os.Setenv("DB_USER", viper.GetString("database.user"))
	os.Setenv("DB_PASSWORD", viper.GetString("database.password"))
	os.Setenv("DB_NAME", viper.GetString("database.name"))
}
