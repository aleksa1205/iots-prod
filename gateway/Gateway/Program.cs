using FluentValidation;
using FluentValidation.AspNetCore;
using Gateway.Clients;
using Gateway.Common;
using Gateway.Filters;
using Gateway.OpenApi;
using Microsoft.Extensions.Options;
using GatewayOptions = Gateway.Common.Gateway;

var builder = WebApplication.CreateBuilder(args);

builder.Services
    .AddOpenApi()
    .AddSwaggerGen()
    .AddValidatorsFromAssembly(typeof(Program).Assembly, includeInternalTypes: true)
    .AddFluentValidationAutoValidation()
    .AddControllers(options =>
    {
        options.Filters.Add<RpcExceptionFilter>();
    });
builder.Services.Configure<DataManager>(builder.Configuration.GetSection(nameof(DataManager)));
builder.Services.Configure<GatewayOptions>(builder.Configuration.GetSection(typeof(GatewayOptions).Name));

builder.Services.AddGrpcClient<SensorReadingService.SensorReadingServiceClient>((sp, options) =>
{
    var configuration = sp.GetRequiredService<IOptions<DataManager>>().Value;
    options.Address = new Uri(configuration.Address);
});

builder.Services.AddSingleton<SensorReadingClient>();
builder.Services.AddHostedService<OpenApiFetcher>();

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.MapOpenApi();
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();
app.MapControllers();

app.Run();
