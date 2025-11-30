package repositories

import (
	"data-manager/internal/entities"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDatabase(connectionString string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for attempt := 1; attempt <= 5; attempt++ {
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err == nil {
			break
		}
		time.Sleep(time.Duration(attempt) * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("CoonectToDatabase: Failed to initiate a connection with database: %w", err)
	}

	err = db.AutoMigrate(&entities.SensorReading{})
	if err != nil {
		return nil, fmt.Errorf("CoonectToDatabase: Failed to auto-migrate database: %w", err)
	}

	return db, nil
}
