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

func (repository *SensorReadingRepository) GetMin(ctx context.Context, start int64, end int64) (*entities.SensorReading, error) {
	var sensorReading entities.SensorReading
	err := repository.db.WithContext(ctx).Where("time >= ? AND time <= ?", start, end).Order("used_kw ASC").First(&sensorReading).Error

	return &sensorReading, err
}

func (repository *SensorReadingRepository) GetMax(ctx context.Context, start int64, end int64) (*entities.SensorReading, error) {
	var sensorReading entities.SensorReading
	err := repository.db.WithContext(ctx).Where("time >= ? AND time <= ?", start, end).Order("used_kw DESC").First(&sensorReading).Error

	return &sensorReading, err
}

func (repository *SensorReadingRepository) GetSum(ctx context.Context, start int64, end int64) (float64, error) {
	var sum float64
	err := repository.db.WithContext(ctx).Model(&entities.SensorReading{}).Select("sum(used_kw)").Scan(&sum).Error

	return sum, err
}

func (repository *SensorReadingRepository) GetAvg(ctx context.Context, start int64, end int64) (float64, error) {
	var avg float64
	err := repository.db.WithContext(ctx).Model(&entities.SensorReading{}).Select("avg(used_kw)").Scan(&avg).Error

	return avg, err
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
