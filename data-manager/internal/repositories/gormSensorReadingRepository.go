package repositories

import (
	"context"
	"data-manager/internal/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormSensorReadingRepository struct {
	db *gorm.DB
}

func NewSensorReadingRepository(db *gorm.DB) SensorReadingRepository {
	return &GormSensorReadingRepository{db: db}
}

func (r *GormSensorReadingRepository) DbWithCtx(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}

func (r *GormSensorReadingRepository) GetAll(ctx context.Context, pageSize int32, pageNumber int32) ([]entities.SensorReading, int64, error) {
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
		Limit(int(pageSize)).
		Offset(int(offset)).
		Find(&readings).Error

	return readings, totalItems, err
}

func (r *GormSensorReadingRepository) GetById(ctx context.Context, id string) (*entities.SensorReading, error) {
	var reading entities.SensorReading
	err := r.DbWithCtx(ctx).
		Model(&entities.SensorReading{}).
		First(&reading, "id = ?", id).Error
	return &reading, err
}

func (r *GormSensorReadingRepository) GetMin(ctx context.Context, start int64, end int64) (*entities.SensorReading, error) {
	var reading entities.SensorReading
	err := r.DbWithCtx(ctx).
		Model(&entities.SensorReading{}).
		Where("time >= ? AND time <= ?", start, end).
		Order("used_kw ASC").
		First(&reading).Error
	return &reading, err
}

func (r *GormSensorReadingRepository) GetMax(ctx context.Context, start int64, end int64) (*entities.SensorReading, error) {
	var reading entities.SensorReading
	err := r.DbWithCtx(ctx).
		Model(&entities.SensorReading{}).
		Where("time >= ? AND time <= ?", start, end).
		Order("used_kw DESC").
		First(&reading).Error
	return &reading, err
}

func (r *GormSensorReadingRepository) GetSum(ctx context.Context, start int64, end int64) (float64, error) {
	var sum float64
	err := r.DbWithCtx(ctx).
		Model(&entities.SensorReading{}).
		Where("time >= ? AND time <= ?", start, end).
		Select("SUM(used_kw)").
		Scan(&sum).Error
	return sum, err
}

func (r *GormSensorReadingRepository) GetAvg(ctx context.Context, start int64, end int64) (float64, error) {
	var avg float64
	err := r.DbWithCtx(ctx).
		Model(&entities.SensorReading{}).
		Where("time >= ? AND time <= ?", start, end).
		Select("AVG(used_kw)").
		Scan(&avg).Error
	return avg, err
}

func (r *GormSensorReadingRepository) Create(ctx context.Context, domain *entities.SensorReading) error {
	return r.DbWithCtx(ctx).
		Model(domain).
		Create(&domain).Error
}

func (r *GormSensorReadingRepository) BatchCreate(ctx context.Context, entityList []*entities.SensorReading) (int64, error) {
	if len(entityList) == 0 {
		return 0, nil
	}

	result := r.
		DbWithCtx(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		}).Create(&entityList)

	return result.RowsAffected, result.Error
}

func (r *GormSensorReadingRepository) Update(ctx context.Context, domain *entities.SensorReading) error {
	return r.DbWithCtx(ctx).
		Model(domain).
		Save(domain).Error
}

func (r *GormSensorReadingRepository) Delete(ctx context.Context, domain *entities.SensorReading) error {
	return r.DbWithCtx(ctx).
		Model(domain).
		Delete(&domain).Error
}

var _ SensorReadingRepository = (*GormSensorReadingRepository)(nil)
