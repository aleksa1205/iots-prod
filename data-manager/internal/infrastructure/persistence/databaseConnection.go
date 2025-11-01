package persistence

import (
	cfg "data-manager/config"
	"data-manager/internal/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDatabase() error {
	connectionString := cfg.GetEnvOrPanic("DB_CONNECTION_STRING")

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return fmt.Errorf("CoonectToDatabase: Failed to initiate a connection with database: %w", err)
	}

	sqlDb, _ := db.DB()

	if err := sqlDb.Ping(); err != nil {
		return fmt.Errorf("CoonectToDatabase: Failed to ping database: %w", err)
	}

	err = db.AutoMigrate(&domain.SensorReading{})
	if err != nil {
		return fmt.Errorf("CoonectToDatabase: Failed to auto-migrate database: %w", err)
	}

	return nil
}
