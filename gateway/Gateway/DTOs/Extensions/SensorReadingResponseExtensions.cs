using Gateway.DTOs.Response;

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
            Time = response.Data.Time,
            UsedKw = response.Data.UsedKw,
            Id = response.Id
        };
    }
}