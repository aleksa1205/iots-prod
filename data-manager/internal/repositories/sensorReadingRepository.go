package repositories

import (
	"context"
	"data-manager/internal/entities"

	"gorm.io/gorm"
)

type SensorReadingRepository struct {
	db *gorm.DB
}

func NewSensorReadingRepository(db *gorm.DB) *SensorReadingRepository {
	return &SensorReadingRepository{db: db}
}

func (repository *SensorReadingRepository) GetAll(ctx context.Context) ([]entities.SensorReading, error) {
	var sensorReadings []entities.SensorReading
	result := repository.db.WithContext(ctx).Find(&sensorReadings)

	return sensorReadings, result.Error
}

func (repository *SensorReadingRepository) GetByID(ctx context.Context, id string) (entities.SensorReading, error) {
	var sensorReading entities.SensorReading
	result := repository.db.WithContext(ctx).First(&sensorReading, id)

	return sensorReading, result.Error
}
