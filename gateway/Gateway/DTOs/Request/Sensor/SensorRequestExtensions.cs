using Gateway.DTOs.Extensions;

namespace Gateway.DTOs.Request;

public static class SensorRequestExtensions
{
    public static CreateSensorReadingRequest ToProto(this SensorRequest sensorRequest)
    {
        return new CreateSensorReadingRequest
        {
            Data = sensorRequest.ToProtoData()
        };
    }
}