using Gateway.DTOs.Request.Sensor;

namespace Gateway.DTOs.Extensions;

public static class SensorReadingDataExtensions
{
    public static SensorReadingData ToProtoData(this SensorRequest request)
    {
        return new SensorReadingData
        {
            ApparentTemperature = request.ApparentTemperature ?? 0,
            Temperature = request.Temperature,
            Humidity = request.Humidity,
            Pressure = request.Pressure,
            GeneratedKw = request.GeneratedKw,
            Time = request.Timestamp,
            UsedKw = request.UsedKw,
        };
    }
}