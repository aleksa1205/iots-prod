package dtos

import (
	"data-manager/internal/entities"

	"github.com/google/uuid"
)

func ToResponse(request entities.SensorReading) SensorReadingResponse {
	return SensorReadingResponse{
		ID:                  request.ID,
		UsedKW:              request.UsedKW,
		GeneratedKW:         request.GeneratedKW,
		Time:                request.Time,
		Temperature:         request.Temperature,
		ApparentTemperature: request.ApparentTemperature,
		Pressure:            request.Pressure,
		Humidity:            request.Humidity,
	}
}

func ToDomain(request SensorReadingRequest) entities.SensorReading {
	return entities.SensorReading{
		ID:                  uuid.New().String(),
		UsedKW:              request.UsedKW,
		GeneratedKW:         request.GeneratedKW,
		Time:                request.Time,
		Temperature:         request.Temperature,
		ApparentTemperature: request.ApparentTemperature,
		Pressure:            request.Pressure,
		Humidity:            request.Humidity,
	}
}
