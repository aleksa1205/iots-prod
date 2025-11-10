package dtos

import (
	"data-manager/internal/entities"
	sensorpb "data-manager/internal/proto"

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

func ToProto(response entities.SensorReading) *sensorpb.SensorReadingResponse {
	return &sensorpb.SensorReadingResponse{
		Id:                  response.ID,
		UsedKw:              response.UsedKW,
		GeneratedKw:         response.GeneratedKW,
		Time:                response.Time,
		Temperature:         response.Temperature,
		ApparentTemperature: response.ApparentTemperature,
		Pressure:            response.Pressure,
		Humidity:            response.Humidity,
	}
}
