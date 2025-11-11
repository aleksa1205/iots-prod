package entities

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

type PaginatedSensorReading struct {
	Items           []SensorReading
	PageSize        int32
	PageNumber      int32
	HasPreviousPage bool
	HasNextPage     bool
	TotalItems      int64
}

func (d *SensorReading) UpdateFromEntity(other *SensorReading) {
	d.UsedKW = other.UsedKW
	d.GeneratedKW = other.GeneratedKW
	d.Temperature = other.Temperature
	d.ApparentTemperature = other.ApparentTemperature
	d.Humidity = other.Humidity
	d.Pressure = other.Pressure
	d.Time = other.Time
}
