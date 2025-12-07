package handlers

import (
	"context"
	sensorpb "data-manager/internal/proto"
	"data-manager/internal/services"
	"fmt"
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SensorReadingHandler struct {
	sensorpb.UnimplementedSensorReadingServiceServer
	service services.SensorReadingServiceInterface
}

var _ sensorpb.SensorReadingServiceServer = (*SensorReadingHandler)(nil)

func NewSensorReadingHandler(service services.SensorReadingServiceInterface) *SensorReadingHandler {
	return &SensorReadingHandler{
		UnimplementedSensorReadingServiceServer: sensorpb.UnimplementedSensorReadingServiceServer{},
		service:                                 service}
}

func (h *SensorReadingHandler) GetSensors(ctx context.Context, pag *sensorpb.PaginationRequest) (*sensorpb.PaginationSensorReadingResponse, error) {
	result, err := h.service.GetAllSensors(ctx, pag.PageSize, pag.PageNumber)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed getting sensors: %v", err)
	}

	sensorResults := make([]*sensorpb.SensorReadingResponse, len(result.Items))
	for i, sensor := range result.Items {
		sensorResults[i] = sensorpb.ToProto(&sensor)
	}

	return &sensorpb.PaginationSensorReadingResponse{
		Items:           sensorResults,
		PageSize:        result.PageSize,
		PageNumber:      result.PageNumber,
		HasPreviousPage: result.HasPreviousPage,
		HasNextPage:     result.HasNextPage,
		TotalItems:      int32(result.TotalItems),
	}, nil
}

func (h *SensorReadingHandler) GetSensorById(ctx context.Context, request *sensorpb.SensorReadingId) (*sensorpb.SensorReadingResponse, error) {
	sensor, err := h.service.GetByID(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return sensorpb.ToProto(sensor), nil
}

func (h *SensorReadingHandler) CreateSensor(ctx context.Context, request *sensorpb.CreateSensorReadingRequest) (*sensorpb.SensorReadingResponse, error) {
	sensor, err := h.service.Create(ctx, request.ToRequest())
	if err != nil {
		return nil, err
	}
	
	return sensorpb.ToProto(sensor), nil
}

func (h *SensorReadingHandler) UpdateSensor(ctx context.Context, request *sensorpb.UpdateSensorReadingRequest) (*sensorpb.SensorReadingResponse, error) {
	sensor, err := h.service.Update(ctx, request.Id, request.ToUpdateRequest())
	if err != nil {
		return nil, err
	}
	return sensorpb.ToProto(sensor), nil
}

func (h *SensorReadingHandler) DeleteSensor(ctx context.Context, request *sensorpb.SensorReadingId) (*emptypb.Empty, error) {
	err := h.service.Delete(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h *SensorReadingHandler) GetSensorByMinUsage(ctx context.Context, request *sensorpb.TimeRangeRequest) (*sensorpb.SensorReadingResponse, error) {
	sensor, err := h.service.GetMin(ctx, request.Start, request.End)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed getting min sensor:\n%v", err)
	}

	return sensorpb.ToProto(sensor), nil
}

func (h *SensorReadingHandler) GetSensorByMaxUsage(ctx context.Context, request *sensorpb.TimeRangeRequest) (*sensorpb.SensorReadingResponse, error) {
	sensor, err := h.service.GetMax(ctx, request.Start, request.End)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed getting max sensor:\n%v", err)
	}

	return sensorpb.ToProto(sensor), nil
}

func (h *SensorReadingHandler) GetSensorUsageAvg(ctx context.Context, request *sensorpb.TimeRangeRequest) (*sensorpb.NumericAggregationResponse, error) {
	value, err := h.service.GetAvg(ctx, request.Start, request.End)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed getting avg usage of sensors:\n%v", err)
	}

	return &sensorpb.NumericAggregationResponse{Value: value}, nil
}

func (h *SensorReadingHandler) GetSensorUsageSum(ctx context.Context, request *sensorpb.TimeRangeRequest) (*sensorpb.NumericAggregationResponse, error) {
	value, err := h.service.GetSum(ctx, request.Start, request.End)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed getting sum usage of sensors:\n%v", err)
	}

	return &sensorpb.NumericAggregationResponse{Value: value}, nil
}

func (h *SensorReadingHandler) StreamSensorReadings(stream sensorpb.SensorReadingService_StreamSensorReadingsServer) error {
	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				fmt.Println("Stream finished")
				break
			}
			return status.Errorf(codes.Internal, "Failed streaming sensor readings: %v", err)
		}
		_, err = h.service.Create(stream.Context(), req.ToRequest())
		if err != nil {
			return status.Errorf(codes.Internal, "Failed streaming sensor readings: %v", err)
		}
	}
	return stream.SendAndClose(&emptypb.Empty{})
}
