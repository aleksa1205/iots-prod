package dtos

type SensorReadingOverview struct {
	Time        int64
	ID          string
	UsedKW      float64
	GeneratedKW float64
}
