using Gateway.Common;
using Gateway.DTOs.Extensions;
using Gateway.DTOs.Request;
using Gateway.DTOs.Request.Filtering;
using Gateway.DTOs.Request.Id;
using Gateway.DTOs.Response;

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
    {
        var proto = request.ToProto();
        return (await _client.CreateSensorAsync(proto)).ToResponse();
    }

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
}