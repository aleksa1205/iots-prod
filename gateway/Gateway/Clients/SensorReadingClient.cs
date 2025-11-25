using Gateway.DTOs.Extensions;
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

    public async Task<PageOfResponse<SensorResponse>> GetAllAsync(int pageNumber = 1, int pageSize = 10)
    {
        var request = new PaginationRequest
        {
            PageNumber = pageNumber,
            PageSize = pageSize
        };

        return (await _client.GetSensorsAsync(request)).ToResponse();
    }

    public async Task<SensorResponse> GetByIdAsync(IdRequest request)
    {
        var id = new SensorReadingId
        {
            Id = request.Id.ToString()
        };
        Console.WriteLine(id);
        return (await _client.GetSensorByIdAsync(id)).ToResponse();
    }
}