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
		return nil, fmt.Errorf("SensorReadingService/GetById: Sensor with id %s not found", id)
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
	sensor := dtos.ToDomain(request)
	err := s.repository.Create(ctx, sensor)

	if err != nil {
		return nil, fmt.Errorf("SensorReadingService/Create: Creating sensor reading failed\nError: %w", err)
	}
	return sensor, nil
}

func (s *SensorReadingService) Update(ctx context.Context, id string, request *dtos.SensorReadingRequest) (*domain.SensorReading, error) {
	sensor, err := s.repository.GetByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("SensorReadingService/Update: Sensor with id %s not found", id)
	}

	sensor.GeneratedKW = request.GeneratedKW
	sensor.UsedKW = request.UsedKW
	sensor.Temperature = request.Temperature
	sensor.Humidity = request.Humidity
	sensor.ApparentTemperature = request.ApparentTemperature
	sensor.Pressure = request.Pressure

	err = s.repository.Update(ctx, sensor)

	if err != nil {
		return nil, fmt.Errorf("SensorReadingService: Updating sensor reading with id %s failed\nError: %w", id, err)
	}

	return sensor, nil
}

func (s *SensorReadingService) Delete(ctx context.Context, id string) error {
	sensor, err := s.repository.GetByID(ctx, id)

	if err != nil {
		return fmt.Errorf("SensorReadingService/Delete: Sensor with id %s not found", id)
	}

	return s.repository.Delete(ctx, sensor)
}
