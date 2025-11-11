package entities

import "data-manager/internal/services/dtos"

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

type PaginatedResponse struct {
	Items           []SensorReading
	PageSize        int32
	PageNumber      int32
	HasPreviousPage bool
	HasNextPage     bool
	TotalItems      int64
}

func (domain *SensorReading) Update(request *dtos.SensorReadingRequest) {
	domain.UsedKW = request.UsedKW
	domain.GeneratedKW = request.GeneratedKW
	domain.Temperature = request.Temperature
	domain.ApparentTemperature = request.ApparentTemperature
	domain.Humidity = request.Humidity
	domain.Pressure = request.Pressure
}
