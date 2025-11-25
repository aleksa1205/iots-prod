package services

import (
	"context"
	domain "data-manager/internal/entities"
	"data-manager/internal/repositories"
	"data-manager/internal/services/dtos"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type SensorReadingService struct {
	repository *repositories.SensorReadingRepository
}

func NewSensorReadingService(repository *repositories.SensorReadingRepository) *SensorReadingService {
	return &SensorReadingService{repository: repository}
}

func (s *SensorReadingService) GetAllSensors(ctx context.Context, pageSize int32, pageNumber int32) (*domain.PaginatedSensorReading, error) {
	items, totalItems, err := s.repository.GetAll(ctx, int(pageSize), int(pageNumber))
	if err != nil {
		return nil, fmt.Errorf("SensorReadingService:GetAll: Issue when fetching records\nError: %w", err)
	}

	return &domain.PaginatedSensorReading{
		Items:           items,
		PageSize:        pageSize,
		PageNumber:      pageNumber,
		HasPreviousPage: pageNumber > 1,
		HasNextPage:     (int64(pageNumber) * int64(pageSize)) < totalItems,
		TotalItems:      totalItems,
	}, nil
}

func (s *SensorReadingService) GetByID(ctx context.Context, id string) (*domain.SensorReading, error) {
	sensor, err := s.repository.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "SensorReadingService:GetByID: Sensor with id %s not found", id)
		}
		return nil, status.Errorf(codes.Internal, "SensorReadingService/GetById: Issue when fetching a record with %s id\nError: %v", id, err)
	}

	return sensor, nil
}

func (s *SensorReadingService) GetMin(ctx context.Context, start int64, end int64) (*domain.SensorReading, error) {
	sensor, err := s.repository.GetMin(ctx, start, end)

	if err != nil {
		return nil, fmt.Errorf("SensorReadingService/GetMin: Sensor with start %d and end %d not found", start, end)
	}

	return sensor, nil
}

func (s *SensorReadingService) GetMax(ctx context.Context, start int64, end int64) (*domain.SensorReading, error) {
	sensor, err := s.repository.GetMax(ctx, start, end)

	if err != nil {
		return nil, fmt.Errorf("SensorReadingService/GetMax: Sensor with start %d and end %d not found", start, end)
	}

	return sensor, nil
}

func (s *SensorReadingService) GetSum(ctx context.Context, start int64, end int64) (float64, error) {
	value, err := s.repository.GetSum(ctx, start, end)

	if err != nil {
		return 0.0, fmt.Errorf("SensorReadingService/GetSum: Sensor with start %d and end %d not found", start, end)
	}

	return value, nil
}

func (s *SensorReadingService) GetAvg(ctx context.Context, start int64, end int64) (float64, error) {
	value, err := s.repository.GetAvg(ctx, start, end)

	if err != nil {
		return 0.0, fmt.Errorf("SensorReadingService/GetAvg: Sensor with start %d and end %d not found", start, end)
	}

	return value, nil
}

func (s *SensorReadingService) Create(ctx context.Context, request *dtos.SensorReadingRequest) (*domain.SensorReading, error) {
	sensor := request.ToDomain()
	err := s.repository.Create(ctx, sensor)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "SensorReadingService/Create: Creating sensor reading failed\nError: %w", err)
	}
	return sensor, nil
}

func (s *SensorReadingService) Update(ctx context.Context, id string, request *dtos.SensorReadingRequest) (*domain.SensorReading, error) {
	sensor, err := s.repository.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "SensorReadingService/Update: Sensor with id %s not found", id)
		}
		return nil, status.Errorf(codes.Internal, "SensorReadingService/Update: Issue when updating a record with %s id\nError: %v", id, err)
	}

	request.UpdateDomain(sensor)
	err = s.repository.Update(ctx, sensor)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "SensorReadingService: Updating sensor reading with id %s failed\nError: %w", id, err)
	}

	return sensor, nil
}

func (s *SensorReadingService) Delete(ctx context.Context, id string) error {
	sensor, err := s.repository.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Errorf(codes.NotFound, "SensorReadingService/Delete: Sensor with id %s not found", id)
		}
		return status.Errorf(codes.Internal, "SensorReadingService/Delete: Issue when deleting a record with %s id\nError: %w", id, err)
	}

	return s.repository.Delete(ctx, sensor)
}
