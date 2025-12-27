package services

import (
	"context"
	"data-manager/internal/dtos"
	domain "data-manager/internal/entities"
	sensorpb "data-manager/internal/proto"
)

type SensorReadingServiceInterface interface {
	GetAllSensors(ctx context.Context, pageSize, pageNumber int32) (*domain.PaginatedSensorReading, error)

	GetByID(ctx context.Context, id string) (*domain.SensorReading, error)

	GetMin(ctx context.Context, start, end int64) (*domain.SensorReading, error)

	GetMax(ctx context.Context, start, end int64) (*domain.SensorReading, error)

	GetSum(ctx context.Context, start, end int64) (float64, error)

	GetAvg(ctx context.Context, start, end int64) (float64, error)

	Create(ctx context.Context, request *dtos.SensorReadingRequest) (*domain.SensorReading, error)

	BatchCreate(ctx context.Context, recv func() (*sensorpb.CreateSensorReadingRequest, error)) (int64, error)

	Update(ctx context.Context, id string, request *dtos.SensorReadingRequest) (*domain.SensorReading, error)

	Delete(ctx context.Context, id string) error
}
