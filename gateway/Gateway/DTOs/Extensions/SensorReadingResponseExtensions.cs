using Gateway.DTOs.Response.Sensor;

namespace Gateway.DTOs.Extensions;

public static class SensorReadingResponseExtensions
{
    public static SensorResponse ToResponse(this SensorReadingResponse response)
    {
        return new SensorResponse
        {
            ApparentTemperature = response.Data.ApparentTemperature,
            GeneratedKw = response.Data.GeneratedKw,
            Humidity = response.Data.Humidity,
            Pressure = response.Data.Pressure,
            Temperature = response.Data.Temperature,
            Time = DateTimeOffset.FromUnixTimeSeconds(response.Data.Time).UtcDateTime,
            UsedKw = response.Data.UsedKw,
            Id = response.Id
        };
    }
}