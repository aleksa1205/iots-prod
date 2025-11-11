package entities

import sensorpb "data-manager/internal/proto"

type SensorReading struct {
	ID                  string
	UsedKW              float64
	GeneratedKW         float64
	Time                int64
	Temperature         float32
	Humidity            float32
	Pressure            float32
	ApparentTemperature float32
}

func ToProto(domain *SensorReading) *sensorpb.SensorReadingResponse {
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
