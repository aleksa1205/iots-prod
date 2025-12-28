package dtos

type AlertEvent struct {
	Type    string
	Time    int64
	Reading SensorReadingOverview
}
