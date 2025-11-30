using System.ComponentModel.DataAnnotations;

namespace Gateway.DTOs.Request.Time;

public record TimeRequest
{
    [Required] public required DateTimeOffset StartTime { get; init; } = DateTimeOffset.UtcNow;
    
    public DateTimeOffset EndTime { get; init; } = DateTimeOffset.UtcNow;
}