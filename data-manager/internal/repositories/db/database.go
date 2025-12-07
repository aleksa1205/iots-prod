package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(connectionString string) (*gorm.DB, error) {
	const maxRetries = 5
	var db *gorm.DB
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if err == nil {
			return db, nil
		}

		time.Sleep(time.Duration(attempt) * time.Second)
	}

	return nil, fmt.Errorf("Connect: Could not open DB connection after %v retries: %w", maxRetries, err)
}
