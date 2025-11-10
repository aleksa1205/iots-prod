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

func (repository *SensorReadingRepository) GetByID(ctx context.Context, id string) (*entities.SensorReading, error) {
	var sensorReading entities.SensorReading
	result := repository.db.WithContext(ctx).First(&sensorReading, "id = ?", id)

	return &sensorReading, result.Error
}

func (repository *SensorReadingRepository) Create(ctx context.Context, domain *entities.SensorReading) error {
	return repository.db.WithContext(ctx).Create(&domain).Error
}

func (repository *SensorReadingRepository) Update(ctx context.Context, domain *entities.SensorReading) error {
	return repository.db.WithContext(ctx).Where("id = ?", domain.ID).Updates(domain).Error
}

func (repository *SensorReadingRepository) Delete(ctx context.Context, domain *entities.SensorReading) error {
	return repository.db.WithContext(ctx).Delete(&domain).Error
}
