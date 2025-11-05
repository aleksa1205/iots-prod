package domain

import (
	"context"
	"data-manager/internal/domain"
)

type SensorReadingRepository interface {
	GetAll(ctx context.Context) ([]domain.SensorReading, error)
	GetByID(ctx context.Context, id string) (domain.SensorReading, error)
}
