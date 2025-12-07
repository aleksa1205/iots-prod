package services

import (
	"context"
	"data-manager/internal/dtos"
	domain "data-manager/internal/entities"
	"data-manager/internal/repositories"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type SensorReadingService struct {
	repository repositories.SensorReadingRepository
}

func NewSensorReadingService(repository repositories.SensorReadingRepository) SensorReadingServiceInterface {
	return &SensorReadingService{repository: repository}
}

func (s *SensorReadingService) GetAllSensors(ctx context.Context, pageSize int32, pageNumber int32) (*domain.PaginatedSensorReading, error) {
	items, totalItems, err := s.repository.GetAll(ctx, pageSize, pageNumber)
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
	return s.mustExist(ctx, id)
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
	sensor, err := s.mustExist(ctx, id)
	if err != nil {
		return nil, err
	}

	request.UpdateDomain(sensor)
	err = s.repository.Update(ctx, sensor)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "SensorReadingService: Updating sensor reading with id %s failed\nError: %w", id, err)
	}

	return sensor, nil
}

func (s *SensorReadingService) Delete(ctx context.Context, id string) error {
	sensor, err := s.mustExist(ctx, id)
	if err != nil {
		return err
	}
	return s.repository.Delete(ctx, sensor)
}

func (s *SensorReadingService) mustExist(ctx context.Context, id string) (entity *domain.SensorReading, err error) {
	sensor, err := s.repository.GetById(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "SensorReadingService/Delete: Sensor with id %s not found", id)
		}
		return nil, status.Errorf(codes.Internal, "SensorReadingService/Delete: Issue when deleting a record with %s id\nError: %w", id, err)
	}
	return sensor, nil
}

var _ SensorReadingServiceInterface = (*SensorReadingService)(nil)
