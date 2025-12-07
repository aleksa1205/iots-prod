package repositories

import (
	"context"
	"data-manager/internal/entities"
)

type SensorReadingRepository interface {
	GetAll(ctx context.Context, pageSize int32, pageNumber int32) ([]entities.SensorReading, int64, error)

	GetById(ctx context.Context, id string) (*entities.SensorReading, error)

	Create(ctx context.Context, entity *entities.SensorReading) error

	Update(ctx context.Context, entity *entities.SensorReading) error

	Delete(ctx context.Context, entity *entities.SensorReading) error

	GetMin(ctx context.Context, start int64, end int64) (*entities.SensorReading, error)

	GetMax(ctx context.Context, start int64, end int64) (*entities.SensorReading, error)

	GetSum(ctx context.Context, start int64, end int64) (float64, error)

	GetAvg(ctx context.Context, start int64, end int64) (float64, error)
}
