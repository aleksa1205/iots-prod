package sensorpb

import (
	"data-manager/internal/entities"
	"data-manager/internal/services/dtos"
)

func (r *CreateSensorReadingRequest) ToRequest() *dtos.SensorReadingRequest {
	return &dtos.SensorReadingRequest{
		UsedKW:              r.Data.UsedKw,
		GeneratedKW:         r.Data.GeneratedKw,
		Temperature:         r.Data.Temperature,
		ApparentTemperature: r.Data.ApparentTemperature,
		Pressure:            r.Data.Pressure,
		Humidity:            r.Data.Humidity,
		Time:                r.Data.Time,
	}
}

func (r *UpdateSensorReadingRequest) ToUpdateRequest() *dtos.SensorReadingRequest {
	return &dtos.SensorReadingRequest{
		UsedKW:              r.Data.UsedKw,
		GeneratedKW:         r.Data.GeneratedKw,
		Temperature:         r.Data.Temperature,
		ApparentTemperature: r.Data.ApparentTemperature,
		Pressure:            r.Data.Pressure,
		Humidity:            r.Data.Humidity,
		Time:                r.Data.Time,
	}
}

func ToProto(d *entities.SensorReading) *SensorReadingResponse {
	return &SensorReadingResponse{
		Id: d.ID,
		Data: &SensorReadingData{
			UsedKw:              d.UsedKW,
			GeneratedKw:         d.GeneratedKW,
			Temperature:         d.Temperature,
			ApparentTemperature: d.ApparentTemperature,
			Pressure:            d.Pressure,
			Humidity:            d.Humidity,
			Time:                d.Time,
		},
	}
}
