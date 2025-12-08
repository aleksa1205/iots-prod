package dtos

import "time"

type AlterEvent struct {
	Type    string
	Time    int64
	Reading SensorReadingOverview
}

func CreateAlertEvent(reading SensorReadingOverview) AlterEvent {
	return AlterEvent{
		Type:    "SENSOR READING ALERT",
		Time:    time.Now().Unix(),
		Reading: reading,
	}
}
