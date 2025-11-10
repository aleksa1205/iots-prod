package handlers

import (
	"context"
	sensorpb "data-manager/internal/proto"
	"data-manager/internal/services"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SensorReadingHandler struct {
	sensorpb.UnimplementedSensorReadingServiceServer
	service *services.SensorReadingService
}

var _ sensorpb.SensorReadingServiceServer = (*SensorReadingHandler)(nil)

func NewSensorReadingHandler(service *services.SensorReadingService) *SensorReadingHandler {
	return &SensorReadingHandler{
		UnimplementedSensorReadingServiceServer: sensorpb.UnimplementedSensorReadingServiceServer{},
		service:                                 service}
}

func (h *SensorReadingHandler) GetSensors(req *emptypb.Empty, stream sensorpb.SensorReadingService_GetSensorsServer) error {
	result, err := h.service.GetAllSensors(stream.Context())
	if err != nil {
		return status.Errorf(codes.Internal, "Failed getting sensors: %v", err)
	}

	for _, sensor := range result {
		response := &sensorpb.SensorReadingResponse{
			Id:                  sensor.ID,
			UsedKw:              sensor.UsedKW,
			GeneratedKw:         sensor.GeneratedKW,
			Time:                sensor.Time,
			Temperature:         sensor.Temperature,
			Humidity:            sensor.Humidity,
			Pressure:            sensor.Pressure,
			ApparentTemperature: sensor.ApparentTemperature,
		}

		if err := stream.Send(response); err != nil {
			return status.Errorf(codes.Internal, "Failed sending response: %v", err)
		}
	}
	return nil
}

func (h *SensorReadingHandler) GetSensorById(ctx context.Context, request *sensorpb.SensorReadingGetByIdRequest) (*sensorpb.SensorReadingResponse, error) {
	sensor, err := h.service.GetByID(ctx, request.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed getting sensor: %v", err)
	}

	response := &sensorpb.SensorReadingResponse{
		Id:                  sensor.ID,
		UsedKw:              sensor.UsedKW,
		GeneratedKw:         sensor.GeneratedKW,
		Time:                sensor.Time,
		Temperature:         sensor.Temperature,
		Humidity:            sensor.Humidity,
		Pressure:            sensor.Pressure,
		ApparentTemperature: sensor.ApparentTemperature,
	}

	return response, nil
}
