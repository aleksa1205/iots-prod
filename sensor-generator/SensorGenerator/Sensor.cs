using System.ComponentModel.DataAnnotations;

namespace SensorGenerator;

public record Sensor
{
    [Required] public required double UsedKw { get; init; }
    
    [Required] public required double GeneratedKw { get; init; }
    
    [Required] public required long Timestamp { get; init; }
    
    [Required] public required float Temperature { get; init; }
    
    [Required] public required float Humidity { get; init; }
    
    [Required] public required float Pressure { get; init; }
    
    public float? ApparentTemperature { get; init; }
}