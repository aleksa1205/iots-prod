package dtos

import (
	"data-manager/internal/entities"
	sensorpb "data-manager/internal/proto"

	"github.com/google/uuid"
)

func ToDomain(request *SensorReadingRequest) *entities.SensorReading {
	return &entities.SensorReading{
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

func ToProto(response *entities.SensorReading) *sensorpb.SensorReadingResponse {
	return &sensorpb.SensorReadingResponse{
		Id: response.ID,
		Data: &sensorpb.SensorReadingData{
			UsedKw:              response.UsedKW,
			GeneratedKw:         response.GeneratedKW,
			Temperature:         response.Temperature,
			ApparentTemperature: response.ApparentTemperature,
			Pressure:            response.Pressure,
			Humidity:            response.Humidity,
			Time:                response.Time,
		},
	}
}

func ToRequest(request *sensorpb.CreateSensorReadingRequest) *SensorReadingRequest {
	return &SensorReadingRequest{
		UsedKW:              request.Data.UsedKw,
		GeneratedKW:         request.Data.GeneratedKw,
		Temperature:         request.Data.Temperature,
		ApparentTemperature: request.Data.ApparentTemperature,
		Pressure:            request.Data.Pressure,
		Humidity:            request.Data.Humidity,
		Time:                request.Data.Time,
	}
}

func ToUpdateRequest(request *sensorpb.UpdateSensorReadingRequest) *SensorReadingRequest {
	return &SensorReadingRequest{
		UsedKW:              request.Data.UsedKw,
		GeneratedKW:         request.Data.GeneratedKw,
		Temperature:         request.Data.Temperature,
		ApparentTemperature: request.Data.ApparentTemperature,
		Pressure:            request.Data.Pressure,
		Humidity:            request.Data.Humidity,
		Time:                request.Data.Time,
	}
}
