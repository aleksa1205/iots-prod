package services

import (
	"context"
	domain "data-manager/internal/entities"
	"data-manager/internal/repositories"
	"data-manager/internal/services/dtos"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type SensorReadingService struct {
	repository *repositories.SensorReadingRepository
}

func NewSensorReadingService(repository *repositories.SensorReadingRepository) *SensorReadingService {
	return &SensorReadingService{repository: repository}
}

func (s *SensorReadingService) GetAllSensors(ctx context.Context) ([]domain.SensorReading, error) {
	return s.repository.GetAll(ctx)
}

func (s *SensorReadingService) GetByID(ctx context.Context, id string) (*domain.SensorReading, error) {
	sensor, err := s.repository.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("SensorReadingService: sensor reading with id %s not found", id)
		}
		return nil, err
	}
	return sensor, nil
}

func (s *SensorReadingService) Create(ctx context.Context, request *dtos.SensorReadingRequest) (*domain.SensorReading, error) {
	sensor := dtos.ToDomain(request)

	err := s.repository.Create(ctx, sensor)
	if err != nil {
		return &domain.SensorReading{}, fmt.Errorf("SensorReadingService: Creating sensor reading with id %s failed", sensor.ID)
	}
	return sensor, nil
}
