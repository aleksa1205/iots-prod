package application

type SensorReadingRequest struct {
	UsedKW              float64
	GeneratedKW         float64
	Time                int64
	Temperature         float32
	Humidity            float32
	Pressure            float32
	ApparentTemperature float32
}

//func (dto SensorReadingRequest) ToDomain() domain.SensorReading {
//	return domain.SensorReading{
//		ID:                  uuid.New().String(),
//		UsedKW:              dto.UsedKW,
//		GeneratedKW:         dto.GeneratedKW,
//		Time:                dto.Time,
//		Temperature:         dto.Temperature,
//		Humidity:            dto.Humidity,
//		Pressure:            dto.Pressure,
//		ApparentTemperature: dto.ApparentTemperature,
//	}
//}
