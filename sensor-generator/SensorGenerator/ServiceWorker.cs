using System.ComponentModel;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Options;
using SensorGenerator.Common;

namespace SensorGenerator;

public class ServiceWorker : BackgroundService
{
    private readonly GatewayClient _client;
    private readonly IOptions<Gateway> _gatewayOptions;

    public ServiceWorker(GatewayClient client, IOptions<Gateway> gatewayOptions)
    {
        _gatewayOptions = gatewayOptions;
        _client = client;
    }

    protected override async Task ExecuteAsync(CancellationToken stoppingToken)
    {
        string filePath = "Dataset.csv";
        if (!File.Exists(filePath))
        {
            Console.WriteLine("Problem connecting to the IoT sensors");
            return;
        }

        using var reader = new StreamReader(filePath);
        var batch = new List<Sensor>();

        await reader.ReadLineAsync(stoppingToken);
        string? line;
        int count = 0;
        while ((line = await reader.ReadLineAsync(stoppingToken)) is not null)
        {
            var values = line.Split(',');
            Console.WriteLine(line);
            
            batch.Add(new Sensor
            {
                Timestamp = long.Parse(values[0]),
                UsedKw = Double.Parse(values[1]),
                GeneratedKw = Double.Parse(values[2]),
                Temperature = float.Parse(values[3]),
                Humidity = float.Parse(values[4]),
                ApparentTemperature = float.Parse(values[5]),
                Pressure = float.Parse(values[6]),
            });

            if (batch.Count >= _gatewayOptions.Value.BatchSize)
            {
                await _client.SendAsync(batch, stoppingToken);
                Console.WriteLine($"Sent batch {count} of {batch.Count}");
                batch.Clear();
            }
        }

        if (batch.Any())
        {
            await _client.SendAsync(batch, stoppingToken);
            Console.WriteLine($"Sent {batch.Count} of {batch.Count} items");
        }
        
        Console.WriteLine("Finished sending sensor data.");
    }
}