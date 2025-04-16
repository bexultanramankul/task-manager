package storage

import (
	"fmt"
	"sync"
	"task-manager/internal/config"
	"task-manager/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func InitDB() {
	log := logger.Log

	once.Do(func() {
		cfg := config.GetConfig().Database

		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
		)

		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Database connection error: ", err)
		}

		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatal("Failed to get generic database object: ", err)
		}

		if err = sqlDB.Ping(); err != nil {
			log.Fatal("Database is unreachable: ", err)
		}

		log.Info("Connected to PostgreSQL with GORM")
	})
}

func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			logger.Log.Warn("Failed to get generic database object: ", err)
			return
		}

		if err := sqlDB.Close(); err != nil {
			logger.Log.Warn("Error closing storage: ", err)
		} else {
			logger.Log.Info("Database connection closed")
		}
	}
}
