package repositories

import (
	"context"
	"data-manager/internal/entities"
	"fmt"

	"gorm.io/gorm"
)

type SensorReadingRepository struct {
	db *gorm.DB
}

func NewSensorReadingRepository(db *gorm.DB) *SensorReadingRepository {
	return &SensorReadingRepository{db: db}
}

func (r *SensorReadingRepository) DbWithCtx(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}

func (r *SensorReadingRepository) GetAll(ctx context.Context, pageSize int, pageNumber int) ([]entities.SensorReading, int64, error) {
	var readings []entities.SensorReading
	var totalItems int64

	err := r.DbWithCtx(ctx).
		Model(&entities.SensorReading{}).
		Count(&totalItems).Error
	if err != nil {
		return nil, 0, err
	}
	offset := pageSize * (pageNumber - 1)
	err = r.DbWithCtx(ctx).
		Model(&entities.SensorReading{}).
		Limit(pageSize).
		Offset(offset).
		Find(&readings).Error
	fmt.Println(readings)
	fmt.Println(totalItems)
	fmt.Println(offset)
	fmt.Println(pageSize)
	fmt.Println(pageNumber)
	return readings, totalItems, err
}

func (r *SensorReadingRepository) GetByID(ctx context.Context, id string) (*entities.SensorReading, error) {
	var reading entities.SensorReading
	err := r.DbWithCtx(ctx).
		Model(&entities.SensorReading{}).
		First(&reading, "id = ?", id).Error
	return &reading, err
}

func (r *SensorReadingRepository) GetMin(ctx context.Context, start int64, end int64) (*entities.SensorReading, error) {
	var reading entities.SensorReading
	err := r.DbWithCtx(ctx).
		Model(&entities.SensorReading{}).
		Where("time >= ? AND time <= ?", start, end).
		Order("used_kw ASC").
		First(&reading).Error
	return &reading, err
}

func (r *SensorReadingRepository) GetMax(ctx context.Context, start int64, end int64) (*entities.SensorReading, error) {
	var reading entities.SensorReading
	err := r.DbWithCtx(ctx).
		Model(&entities.SensorReading{}).
		Where("time >= ? AND time <= ?", start, end).
		Order("used_kw DESC").
		First(&reading).Error
	return &reading, err
}

func (r *SensorReadingRepository) GetSum(ctx context.Context, start int64, end int64) (float64, error) {
	var sum float64
	err := r.DbWithCtx(ctx).
		Model(&entities.SensorReading{}).
		Where("time >= ? AND time <= ?", start, end).
		Select("SUM(used_kw)").
		Scan(&sum).Error
	return sum, err
}

func (r *SensorReadingRepository) GetAvg(ctx context.Context, start int64, end int64) (float64, error) {
	var avg float64
	err := r.DbWithCtx(ctx).
		Model(&entities.SensorReading{}).
		Where("time >= ? AND time <= ?", start, end).
		Select("AVG(used_kw)").
		Scan(&avg).Error
	return avg, err
}

func (r *SensorReadingRepository) Create(ctx context.Context, domain *entities.SensorReading) error {
	return r.DbWithCtx(ctx).
		Model(domain).
		Create(&domain).Error
}

func (r *SensorReadingRepository) Update(ctx context.Context, domain *entities.SensorReading) error {
	return r.DbWithCtx(ctx).
		Model(domain).
		Save(domain).Error
}

func (r *SensorReadingRepository) Delete(ctx context.Context, domain *entities.SensorReading) error {
	return r.DbWithCtx(ctx).
		Model(domain).
		Delete(&domain).Error
}
