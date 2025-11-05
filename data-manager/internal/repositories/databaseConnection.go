package repositories

import (
	"data-manager/internal/entities"
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDatabase(connectionString string) (*sql.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, fmt.Errorf("CoonectToDatabase: Failed to initiate a connection with database: %w", err)
	}

	sqlDb, _ := db.DB()

	if err := sqlDb.Ping(); err != nil {
		return nil, fmt.Errorf("CoonectToDatabase: Failed to ping database: %w", err)
	}

	err = db.AutoMigrate(&entities.SensorReading{})
	if err != nil {
		return nil, fmt.Errorf("CoonectToDatabase: Failed to auto-migrate database: %w", err)
	}

	return sqlDb, nil
}
