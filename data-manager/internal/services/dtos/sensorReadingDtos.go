package dtos

import (
	"data-manager/internal/entities"

	"github.com/google/uuid"
)

type SensorReadingRequest struct {
	UsedKW              float64
	GeneratedKW         float64
	Time                int64
	Temperature         float32
	Humidity            float32
	Pressure            float32
	ApparentTemperature float32
}

type SensorReadingResponse struct {
	ID                  string
	UsedKW              float64
	GeneratedKW         float64
	Time                int64
	Temperature         float32
	Humidity            float32
	Pressure            float32
	ApparentTemperature float32
}

type SensorReadingOverview struct {
	ID          string
	UsedKW      float64
	GeneratedKW float64
}

func (r *SensorReadingRequest) ToDomain() *entities.SensorReading {
	return &entities.SensorReading{
		ID:                  uuid.New().String(),
		UsedKW:              r.UsedKW,
		GeneratedKW:         r.GeneratedKW,
		Time:                r.Time,
		Temperature:         r.Temperature,
		ApparentTemperature: r.ApparentTemperature,
		Pressure:            r.Pressure,
		Humidity:            r.Humidity,
	}
}

func (r *SensorReadingRequest) UpdateDomain(entity *entities.SensorReading) {
	entity.UpdateFromEntity(&entities.SensorReading{
		UsedKW:              r.UsedKW,
		GeneratedKW:         r.GeneratedKW,
		Time:                r.Time,
		Temperature:         r.Temperature,
		ApparentTemperature: r.ApparentTemperature,
		Pressure:            r.Pressure,
		Humidity:            r.Humidity})
}

func ToOverview(r *entities.SensorReading) *SensorReadingOverview {
	return &SensorReadingOverview{
		ID:          r.ID,
		UsedKW:      r.UsedKW,
		GeneratedKW: r.GeneratedKW,
	}
}
