package application

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

//func (domain dmn.SensorReading) ToResponse() SensorReadingResponse {
//	return SensorReadingResponse{
//		ID:                  domain.ID,
//		UsedKW:              domain.UsedKW,
//		GeneratedKW:         domain.GeneratedKW,
//		Time:                domain.Time,
//		Temperature:         domain.Temperature,
//		Humidity:            domain.Humidity,
//		Pressure:            domain.Pressure,
//		ApparentTemperature: domain.ApparentTemperature,
//	}
//}
