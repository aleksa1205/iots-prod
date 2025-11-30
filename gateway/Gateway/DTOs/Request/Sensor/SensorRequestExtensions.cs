using Gateway.DTOs.Extensions;

namespace Gateway.DTOs.Request.Sensor;

public static class SensorRequestExtensions
{
    public static CreateSensorReadingRequest ToProto(this SensorRequest request)
    {
        return new CreateSensorReadingRequest
        {
            Data = request.ToProtoData()
        };
    }
}