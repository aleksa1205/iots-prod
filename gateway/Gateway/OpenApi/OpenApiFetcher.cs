using Microsoft.Extensions.Options;

namespace Gateway.OpenApi;

public class OpenApiFetcher : BackgroundService
{
    private readonly string _swaggerUrl;
    
    public OpenApiFetcher(IOptions<Common.Gateway> options)
    {
        _swaggerUrl = options.Value.Address;
    }
    
    protected override async Task ExecuteAsync(CancellationToken stoppingToken)
    {
        await Task.Delay(3000, stoppingToken);
        
        var httpClient = new HttpClient();
        var json = await httpClient.GetStringAsync($"{_swaggerUrl}/swagger/v1/swagger.json", stoppingToken);
        await File.WriteAllTextAsync("OpenApi/openapi.json", json, stoppingToken);
    }
}