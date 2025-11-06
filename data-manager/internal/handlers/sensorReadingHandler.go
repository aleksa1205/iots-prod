package handlers

import (
	sensorpb "data-manager/internal/proto"
	"data-manager/internal/services"
)

type SensorReadingHandler struct {
	sensorpb.UnimplementedSensorReadingServiceServer
	service services.SensorReadingService
}

var _ sensorpb.SensorReadingServiceServer = (*SensorReadingHandler)(nil)

func NewSensorReadingHandler(service services.SensorReadingService) *SensorReadingHandler {
	return &SensorReadingHandler{
		UnimplementedSensorReadingServiceServer: sensorpb.UnimplementedSensorReadingServiceServer{},
		service:                                 service}
}
