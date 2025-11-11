package mappers

import (
	"data-manager/internal/entities"
	sensorpb "data-manager/internal/proto"
	"data-manager/internal/services/dtos"

	"github.com/google/uuid"
)

func ToDomain(request *dtos.SensorReadingRequest) *entities.SensorReading {
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

func ToRequest(request *sensorpb.CreateSensorReadingRequest) *dtos.SensorReadingRequest {
	return &dtos.SensorReadingRequest{
		UsedKW:              request.Data.UsedKw,
		GeneratedKW:         request.Data.GeneratedKw,
		Temperature:         request.Data.Temperature,
		ApparentTemperature: request.Data.ApparentTemperature,
		Pressure:            request.Data.Pressure,
		Humidity:            request.Data.Humidity,
		Time:                request.Data.Time,
	}
}

func ToUpdateRequest(request *sensorpb.UpdateSensorReadingRequest) *dtos.SensorReadingRequest {
	return &dtos.SensorReadingRequest{
		UsedKW:              request.Data.UsedKw,
		GeneratedKW:         request.Data.GeneratedKw,
		Temperature:         request.Data.Temperature,
		ApparentTemperature: request.Data.ApparentTemperature,
		Pressure:            request.Data.Pressure,
		Humidity:            request.Data.Humidity,
		Time:                request.Data.Time,
	}
}

func ToProto(domain *entities.SensorReading) *sensorpb.SensorReadingResponse {
	return &sensorpb.SensorReadingResponse{
		Id: domain.ID,
		Data: &sensorpb.SensorReadingData{
			UsedKw:              domain.UsedKW,
			GeneratedKw:         domain.GeneratedKW,
			Temperature:         domain.Temperature,
			ApparentTemperature: domain.ApparentTemperature,
			Pressure:            domain.Pressure,
			Humidity:            domain.Humidity,
			Time:                domain.Time,
		},
	}
}
