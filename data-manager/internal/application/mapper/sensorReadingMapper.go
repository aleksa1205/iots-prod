package application

import (
	application "data-manager/internal/application/dtos"
	"data-manager/internal/domain"

	"github.com/google/uuid"
)

func ToDomain(request application.SensorReadingRequest) domain.SensorReading {
	return domain.SensorReading{
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

func ToResponse(request domain.SensorReading) application.SensorReadingResponse {
	return application.SensorReadingResponse{
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
