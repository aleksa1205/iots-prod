namespace Gateway.DTOs.Request.Time;

public static class TimeRequestExtensions
{
    public static TimeRangeRequest ToProto(this TimeRequest request)
    {
        return new TimeRangeRequest
        {
            Start = request.StartTime.ToUnixTimeSeconds(),
            End = request.EndTime.ToUnixTimeSeconds(),
        };
    }
}