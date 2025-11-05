package domain

import "data-manager/internal/domain"

type SensorReadingRepository interface {
	GetAll() ([]domain.SensorReading, error)
	GetByID(id string) (domain.SensorReading, error)
}
