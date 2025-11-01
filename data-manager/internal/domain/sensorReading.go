package domain

type SensorReading struct {
	Id                  string `gorm:"primaryKey"`
	UsedKW              float64
	GeneratedKW         float64
	Time                int64
	Temperature         float32
	Humidity            float32
	Pressure            float32
	ApparentTemperature float32
}
