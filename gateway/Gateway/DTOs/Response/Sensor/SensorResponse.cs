using System.ComponentModel.DataAnnotations;

namespace Gateway.DTOs.Response.Sensor;

public record SensorResponse
{
    [Required] public required string Id { get; init; }
    
    [Required] public required double UsedKw { get; init; }
    
    [Required] public required double GeneratedKw { get; init; }
    
    [Required] public required DateTime Time { get; init; }
    
    [Required] public float Temperature { get; init; }
    
    [Required] public float Humidity { get; init; }
    
    [Required] public float Pressure { get; init; }
    
    public float ApparentTemperature { get; init; }
};