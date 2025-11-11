package handlers

import (
	"context"
	"data-manager/internal/entities"
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
		response := entities.ToProto(&sensor)

		if err := stream.Send(response); err != nil {
			return status.Errorf(codes.Internal, "Failed sending response: %v", err)
		}
	}
	return nil
}

func (h *SensorReadingHandler) GetSensorById(ctx context.Context, request *sensorpb.SensorReadingId) (*sensorpb.SensorReadingResponse, error) {
	sensor, err := h.service.GetByID(ctx, request.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed getting sensor:\n%v", err)
	}

	response := entities.ToProto(sensor)

	return response, nil
}

func (h *SensorReadingHandler) GetSensorByMinUsage(ctx context.Context, request *sensorpb.TimeRangeRequest) (*sensorpb.SensorReadingResponse, error) {
	sensor, err := h.service.GetMin(ctx, request.Start, request.End)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed getting sensor:\n%v", err)
	}

	response := entities.ToProto(sensor)

	return response, nil
}

func (h *SensorReadingHandler) GetSensorByMaxUsage(ctx context.Context, request *sensorpb.TimeRangeRequest) (*sensorpb.SensorReadingResponse, error) {
	sensor, err := h.service.GetMax(ctx, request.Start, request.End)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed getting sensor:\n%v", err)
	}

	response := entities.ToProto(sensor)

	return response, nil
}

func (h *SensorReadingHandler) GetSensorUsageAvg(ctx context.Context, request *sensorpb.TimeRangeRequest) (*sensorpb.NumericAggregationResponse, error) {
	value, err := h.service.GetAvg(ctx, request.Start, request.End)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed getting sensor:\n%v", err)
	}

	return &sensorpb.NumericAggregationResponse{Value: value}, nil
}

func (h *SensorReadingHandler) GetSensorUsageSum(ctx context.Context, request *sensorpb.TimeRangeRequest) (*sensorpb.NumericAggregationResponse, error) {
	value, err := h.service.GetSum(ctx, request.Start, request.End)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed getting sensor:\n%v", err)
	}

	return &sensorpb.NumericAggregationResponse{Value: value}, nil
}

func (h *SensorReadingHandler) CreateSensor(ctx context.Context, request *sensorpb.CreateSensorReadingRequest) (*sensorpb.SensorReadingResponse, error) {
	sensor, err := h.service.Create(ctx, dtos.ToRequest(request))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed creating sensor:\n%v", err)
	}

	response := entities.ToProto(sensor)
	return response, nil
}

func (h *SensorReadingHandler) UpdateSensor(ctx context.Context, request *sensorpb.UpdateSensorReadingRequest) (*sensorpb.SensorReadingResponse, error) {
	sensor, err := h.service.Update(ctx, request.Id, dtos.ToUpdateRequest(request))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed updating sensor:\n%v", err)
	}
	response := entities.ToProto(sensor)
	return response, nil
}

func (h *SensorReadingHandler) DeleteSensor(ctx context.Context, request *sensorpb.SensorReadingId) (*emptypb.Empty, error) {
	err := h.service.Delete(ctx, request.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed deleting sensor:\n%v", err)
	}

	return &emptypb.Empty{}, nil
}
