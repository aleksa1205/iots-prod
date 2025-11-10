package handlers

import (
	"context"
	sensorpb "data-manager/internal/proto"
	"data-manager/internal/services"
	"data-manager/internal/services/dtos"

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
		response := dtos.ToProto(&sensor)

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

	response := dtos.ToProto(sensor)

	return response, nil
}

func (h *SensorReadingHandler) CreateSensor(ctx context.Context, request *sensorpb.CreateSensorReadingRequest) (*sensorpb.SensorReadingResponse, error) {
	sensor, err := h.service.Create(ctx, dtos.ToRequest(request))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed creating sensor: %v", err)
	}
	response := dtos.ToProto(sensor)
	return response, nil
}
