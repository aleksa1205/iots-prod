package application

import (
	"context"
	domain "data-manager/internal/domain"
	interfaces "data-manager/internal/domain/interfaces"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type SensorReadingService struct {
	repository interfaces.SensorReadingRepository
}

func NewSensorReadingService(repository interfaces.SensorReadingRepository) *SensorReadingService {
	return &SensorReadingService{repository: repository}
}
func (s *SensorReadingService) GetAllSensors(ctx context.Context) ([]domain.SensorReading, error) {
	return s.repository.GetAll(ctx)
}

func (s *SensorReadingService) GetByID(ctx context.Context, id string) (domain.SensorReading, error) {
	sensor, err := s.repository.GetByID(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.SensorReading{}, fmt.Errorf("SensorReadingService: sensor reading with id %s not found", id)
		}
		return domain.SensorReading{}, err
	}
	return sensor, nil
}
