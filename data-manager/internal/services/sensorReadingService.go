package services

import (
	"context"
	domain "data-manager/internal/entities"
	"data-manager/internal/repositories"
	"data-manager/internal/services/dtos"
	"fmt"
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
		return nil, fmt.Errorf("SensorReadingService: Getting sensor reading with id %s failed: \n Error: %w", id, err)
	}
	if sensor == nil {
		return nil, fmt.Errorf("SensorReadingService: Sensor reading with id %s not found", id)
	}

	return sensor, nil
}

func (s *SensorReadingService) Create(ctx context.Context, request *dtos.SensorReadingRequest) (*domain.SensorReading, error) {
	sensor := dtos.ToDomain(request)
	err := s.repository.Create(ctx, sensor)

	if err != nil {
		return nil, fmt.Errorf("SensorReadingService: Creating sensor reading failed: \n Error: %w", err)
	}
	return sensor, nil
}

func (s *SensorReadingService) Update(ctx context.Context, id string, request *dtos.SensorReadingRequest) (*domain.SensorReading, error) {
	sensor, err := s.repository.GetByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("SensorReadingService: Updating sensor reading with id %s failed: \n Error: %w", id, err)
	}
	if sensor == nil {
		return nil, fmt.Errorf("SensorReadingService: Sensor reading with id %s not found", id)
	}

	sensor.GeneratedKW = request.GeneratedKW
	sensor.UsedKW = request.UsedKW
	sensor.Temperature = request.Temperature
	sensor.Humidity = request.Humidity
	sensor.ApparentTemperature = request.ApparentTemperature
	sensor.Pressure = request.Pressure

	err = s.repository.Update(ctx, sensor)

	if err != nil {
		return nil, fmt.Errorf("SensorReadingService: Updating sensor reading with id %s failed: \n Error: %w", id, err)
	}

	return sensor, nil
}

func (s *SensorReadingService) Delete(ctx context.Context, id string) error {
	sensor, err := s.repository.GetByID(ctx, id)

	if err != nil {
		return fmt.Errorf("SensorReadingService: Deleting sensor reading with id %s failed: \n Error: %w", id, err)
	}
	if sensor == nil {
		return fmt.Errorf("SensorReadingService: Sensor reading with id %s not found", id)
	}

	return s.repository.Delete(ctx, sensor)
}
