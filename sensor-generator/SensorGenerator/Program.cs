using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using SensorGenerator;
using SensorGenerator.Common;

var host = Host
    .CreateDefaultBuilder()
    .ConfigureServices((context, services) =>
    {
        services.Configure<Gateway>(context.Configuration.GetSection(nameof(Gateway)));

        services.AddHttpClient<GatewayClient>();
        services.AddHostedService<ServiceWorker>();
    })
    .Build();
    
await host.RunAsync();