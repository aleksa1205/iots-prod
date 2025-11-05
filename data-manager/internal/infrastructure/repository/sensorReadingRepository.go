package infrastructure

import (
	"context"
	"data-manager/internal/domain"

	"gorm.io/gorm"
)

type SensorReadingRepository struct {
	db *gorm.DB
}

func NewSensorReadingRepository(db *gorm.DB) *SensorReadingRepository {
	return &SensorReadingRepository{db: db}
}

func (repository *SensorReadingRepository) GetAll(ctx context.Context) ([]domain.SensorReading, error) {
	var sensorReadings []domain.SensorReading
	result := repository.db.WithContext(ctx).Find(&sensorReadings)

	return sensorReadings, result.Error
}

func (repository *SensorReadingRepository) GetByID(ctx context.Context, id string) (domain.SensorReading, error) {
	var sensorReading domain.SensorReading
	result := repository.db.WithContext(ctx).First(&sensorReading, id)

	return sensorReading, result.Error
}
