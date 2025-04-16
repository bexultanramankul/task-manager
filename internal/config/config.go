package config

import (
	"task-manager/pkg/logger"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Auth     AuthConfig     `mapstructure:"auth"`
}

type ServerConfig struct {
	Port         string        `mapstructure:"port"`
	Host         string        `mapstructure:"host"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
	Env          string        `mapstructure:"env"`
}

type AuthConfig struct {
	JWTSecret           string        `mapstructure:"jwt_secret"`
	JWTAccessExpiresIn  time.Duration `mapstructure:"jwt_access_expires_in"`
	JWTRefreshExpiresIn time.Duration `mapstructure:"jwt_refresh_expires_in"`
}

type DatabaseConfig struct {
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	Name            string        `mapstructure:"name"`
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	SSLMode         string        `mapstructure:"sslmode"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

var appConfig *Config

func LoadConfig(path string) error {
	log := logger.Log

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	// Установка значений по умолчанию
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.read_timeout", "15s")
	viper.SetDefault("server.write_timeout", "15s")
	viper.SetDefault("server.idle_timeout", "60s")
	viper.SetDefault("server.env", "development")

	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("database.max_open_conns", 25)
	viper.SetDefault("database.max_idle_conns", 5)
	viper.SetDefault("database.conn_max_lifetime", "5m")

	viper.SetDefault("auth.jwt_access_expires_in", "15m")
	viper.SetDefault("auth.jwt_refresh_expires_in", "24h")

	// Чтение переменных окружения с префиксом TM_
	viper.SetEnvPrefix("tm")
	viper.AutomaticEnv()

	// Попытка чтения конфиг файла
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn("Config file not found, using defaults and environment variables")
		} else {
			return err
		}
	}

	appConfig = &Config{}
	if err := viper.Unmarshal(appConfig); err != nil {
		return err
	}

	// Валидация обязательных полей
	if appConfig.Auth.JWTSecret == "" {
		log.Fatal("JWT secret is required")
	}

	if appConfig.Database.User == "" || appConfig.Database.Password == "" {
		log.Fatal("Database credentials are required")
	}

	log.Info("Configuration loaded successfully")
	return nil
}

func GetConfig() *Config {
	return appConfig
}
