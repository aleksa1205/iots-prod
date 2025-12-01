using Gateway.Common;
using Gateway.DTOs.Extensions;
using Gateway.DTOs.Request.Filtering;
using Gateway.DTOs.Request.Id;
using Gateway.DTOs.Request.Sensor;
using Gateway.DTOs.Request.Time;
using Gateway.DTOs.Response;
using Gateway.DTOs.Response.Sensor;

namespace Gateway.Clients;

public class SensorReadingClient
{
    private readonly global::SensorReadingService.SensorReadingServiceClient _client;
    
    public SensorReadingClient(global::SensorReadingService.SensorReadingServiceClient client)
    {
        _client = client;
    }

    public async Task<PageOfResponse<SensorResponse>> GetAllAsync(PageParams pageParams)
    {
        var request = new PaginationRequest
        {
            PageNumber = pageParams.Page ?? Constants.Request.PageNumber,
            PageSize = pageParams.PageSize ?? Constants.Request.PageSize
        };

        return (await _client.GetSensorsAsync(request)).ToResponse();
    }

    public async Task<SensorResponse> GetByIdAsync(IdRequest request)
    {
        var id = new SensorReadingId
        {
            Id = request.Id.ToString()
        };
        return (await _client.GetSensorByIdAsync(id)).ToResponse();
    }

    public async Task<SensorResponse> CreateAsync(SensorRequest request)
        => (await _client
            .CreateSensorAsync(request.ToProto()))
            .ToResponse();

    public async Task<SensorResponse> UpdateAsync(IdRequest id, SensorRequest request)
    {
        var proto = new UpdateSensorReadingRequest
        {
            Id = id.Id.ToString(),
            Data = request.ToProtoData()
        };

    return (await _client.UpdateSensorAsync(proto)).ToResponse();
    }

    public async Task DeleteAsync(IdRequest request)
    {
        var id = new SensorReadingId
        {
            Id = request.Id.ToString()
        };
        await _client.DeleteSensorAsync(id);
    }

    public async Task<SensorResponse> GetByMinUsage(TimeRequest request)
        => (await _client
            .GetSensorByMinUsageAsync(request.ToProto()))
            .ToResponse();

    public async Task<SensorResponse> GetByMaxUsage(TimeRequest request)
        => (await _client
            .GetSensorByMaxUsageAsync(request.ToProto()))
            .ToResponse();
    

    public async Task<double> GetAvgUsage(TimeRequest request)
        => (await _client
            .GetSensorUsageAvgAsync(request.ToProto())).Value;

    public async Task<double> GetSumUsage(TimeRequest request)
        => (await _client
                .GetSensorUsageSumAsync(request.ToProto())).Value;

    public async Task Stream(IEnumerable<SensorRequest> readings)
    {
        using var call = _client.StreamSensorReadings();
        
        foreach (var reading in readings)
        {
            var proto = reading.ToProto();
            await call.RequestStream.WriteAsync(proto);
        }
        
        await call.RequestStream.CompleteAsync();
        await call.ResponseAsync;
    }
}