using System.Net.Http.Json;
using Microsoft.Extensions.Options;
using SensorGenerator.Common;

namespace SensorGenerator;

public class GatewayClient
{
    private readonly HttpClient _client;
    private readonly IOptions<Gateway> _gatewayOptions;
    private readonly int _numberOfRetries;
    
    public GatewayClient(HttpClient client, IOptions<Gateway> gatewayOptions)
    {
        _client = client;
        _gatewayOptions = gatewayOptions;
        _numberOfRetries = 5;
    }

    public async Task SendAsync(IEnumerable<Sensor> sensors, CancellationToken cancellationToken)
    {
        var options = _gatewayOptions.Value;
        var attempt = 0;
        while (true)
        {
            try
            {
                var response =
                    await _client.PostAsJsonAsync($"{options.Address}/{options.Endpoint}", sensors, cancellationToken);
        
                response.EnsureSuccessStatusCode();
                await Task.Delay(options.BatchTimeout * 1000, cancellationToken);
            }
            catch (Exception e)
            {
                attempt++;
                Console.WriteLine(e);
                if (attempt > _numberOfRetries)
                {
                    throw;
                }
                await Task.Delay(10000, cancellationToken);
            }
        }
    }
}