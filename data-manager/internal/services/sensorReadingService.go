package services

import (
	"context"
	"data-manager/internal/dtos"
	domain "data-manager/internal/entities"
	lmqtt "data-manager/internal/mqtt"
	sensorpb "data-manager/internal/proto"
	"data-manager/internal/repositories"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type SensorReadingService struct {
	repository repositories.SensorReadingRepository
	broker     *lmqtt.SensorMqttClient
	topic      string
	batchSize  int32
}

func NewSensorReadingService(repository repositories.SensorReadingRepository, broker *lmqtt.SensorMqttClient, topic string) SensorReadingServiceInterface {
	return &SensorReadingService{repository: repository, broker: broker, topic: topic, batchSize: 10}
}

func (s *SensorReadingService) GetAllSensors(ctx context.Context, pageSize int32, pageNumber int32) (*domain.PaginatedSensorReading, error) {
	items, totalItems, err := s.repository.GetAll(ctx, pageSize, pageNumber)
	if err != nil {
		return nil, fmt.Errorf("SensorReadingService:GetAll: Issue when fetching records\nError: %w", err)
	}

	return &domain.PaginatedSensorReading{
		Items:           items,
		PageSize:        pageSize,
		PageNumber:      pageNumber,
		HasPreviousPage: pageNumber > 1,
		HasNextPage:     (int64(pageNumber) * int64(pageSize)) < totalItems,
		TotalItems:      totalItems,
	}, nil
}

func (s *SensorReadingService) GetByID(ctx context.Context, id string) (*domain.SensorReading, error) {
	return s.mustExist(ctx, id)
}

func (s *SensorReadingService) GetMin(ctx context.Context, start int64, end int64) (*domain.SensorReading, error) {
	sensor, err := s.repository.GetMin(ctx, start, end)

	if err != nil {
		return nil, fmt.Errorf("SensorReadingService/GetMin: Sensor with start %d and end %d not found", start, end)
	}

	return sensor, nil
}

func (s *SensorReadingService) GetMax(ctx context.Context, start int64, end int64) (*domain.SensorReading, error) {
	sensor, err := s.repository.GetMax(ctx, start, end)

	if err != nil {
		return nil, fmt.Errorf("SensorReadingService/GetMax: Sensor with start %d and end %d not found", start, end)
	}

	return sensor, nil
}

func (s *SensorReadingService) GetSum(ctx context.Context, start int64, end int64) (float64, error) {
	value, err := s.repository.GetSum(ctx, start, end)

	if err != nil {
		return 0.0, fmt.Errorf("SensorReadingService/GetSum: Sensor with start %d and end %d not found", start, end)
	}

	return value, nil
}

func (s *SensorReadingService) GetAvg(ctx context.Context, start int64, end int64) (float64, error) {
	value, err := s.repository.GetAvg(ctx, start, end)

	if err != nil {
		return 0.0, fmt.Errorf("SensorReadingService/GetAvg: Sensor with start %d and end %d not found", start, end)
	}

	return value, nil
}

func (s *SensorReadingService) Create(ctx context.Context, request *dtos.SensorReadingRequest) (*domain.SensorReading, error) {
	sensor := request.ToDomain()
	err := s.repository.Create(ctx, sensor)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "SensorReadingService/Create: Creating sensor reading failed\nError: %w", err)
	}

	err = s.publishSensorReadingToTopic(sensor)
	if err != nil {
		fmt.Printf("SensorReadingService/Create: Publishing sensor reading failed\nError: %w", err)
	}

	return sensor, nil
}

func (s *SensorReadingService) BatchCreate(ctx context.Context, recv func() (*sensorpb.CreateSensorReadingRequest, error)) (int64, error) {
	batch := make([]*domain.SensorReading, 0, s.batchSize)
	var totalInserted int64

	flush := func() error {
		if len(batch) == 0 {
			return nil
		}
		inserted, err := s.repository.BatchCreate(ctx, batch)
		if err != nil {
			return err
		}

		for _, reading := range batch {
			payload, err := json.Marshal(reading)
			if err != nil {
				return err
			}
			if err := s.broker.Publish(payload); err != nil {
				return err
			}
		}
		totalInserted += inserted
		batch = batch[:0]
		return nil
	}

	for {
		proto, err := recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		entity := proto.ToRequest().ToDomain()
		batch = append(batch, entity)

		if len(batch) >= int(s.batchSize) {
			if err := flush(); err != nil {
				return 0, err
			}
		}
	}

	if err := flush(); err != nil {
		return 0, err
	}

	return totalInserted, nil
}

func (s *SensorReadingService) Update(ctx context.Context, id string, request *dtos.SensorReadingRequest) (*domain.SensorReading, error) {
	sensor, err := s.mustExist(ctx, id)
	if err != nil {
		return nil, err
	}

	request.UpdateDomain(sensor)
	err = s.repository.Update(ctx, sensor)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "SensorReadingService: Updating sensor reading with id %s failed\nError: %w", id, err)
	}

	return sensor, nil
}

func (s *SensorReadingService) Delete(ctx context.Context, id string) error {
	sensor, err := s.mustExist(ctx, id)
	if err != nil {
		return err
	}
	return s.repository.Delete(ctx, sensor)
}

func (s *SensorReadingService) mustExist(ctx context.Context, id string) (entity *domain.SensorReading, err error) {
	sensor, err := s.repository.GetById(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "SensorReadingService/Delete: Sensor with id %s not found", id)
		}
		return nil, status.Errorf(codes.Internal, "SensorReadingService/Delete: Issue when deleting a record with %s id\nError: %w", id, err)
	}
	return sensor, nil
}

func (s *SensorReadingService) publishSensorReadingToTopic(sensor *domain.SensorReading) error {
	payload, err := json.Marshal(dtos.ToOverview(sensor))
	if err != nil {
		log.Println("Failed to marshal to overview:", err)
	}

	return s.broker.Publish(payload)
}

var _ SensorReadingServiceInterface = (*SensorReadingService)(nil)
