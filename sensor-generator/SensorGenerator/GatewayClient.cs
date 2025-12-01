using System.Net.Http.Json;
using Microsoft.Extensions.Options;
using SensorGenerator.Common;

namespace SensorGenerator;

public class GatewayClient
{
    private readonly HttpClient _client;
    private readonly IOptions<Gateway> _gatewayOptions;
    
    public GatewayClient(HttpClient client, IOptions<Gateway> gatewayOptions)
    {
        _client = client;
        _gatewayOptions = gatewayOptions;
    }

    public async Task SendAsync(IEnumerable<Sensor> sensors, CancellationToken cancellationToken)
    {
        var attempt = 0;
        while (true)
        {
            try
            {
                var response =
                    await _client.PostAsJsonAsync($"{_gatewayOptions.Value.Address}/{_gatewayOptions.Value.Endpoint}", sensors, cancellationToken);
        
                response.EnsureSuccessStatusCode();
            }
            catch (Exception e)
            {
                attempt++;
                Console.WriteLine(e);
                if (attempt > 5)
                {
                    throw;
                }
                await Task.Delay(10000, cancellationToken);
            }

        }

    }
}