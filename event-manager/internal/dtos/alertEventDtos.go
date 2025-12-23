package dtos

import "time"

const (
	GenerateOverflow string = "GENERATE_OVERFLOW"
	UsedOverflow     string = "USED_OVERFLOW"
)

type EventThreshold struct {
	GenerateKw float64
	UsedKw     float64
}

type AlertEvent struct {
	Type    string
	Time    int64
	Reading SensorReadingOverview
}

func CreateSensorReadingAlert(reading SensorReadingOverview, eventType string) AlertEvent {
	return AlertEvent{
		Type:    eventType,
		Time:    time.Now().Unix(),
		Reading: reading,
	}
}
