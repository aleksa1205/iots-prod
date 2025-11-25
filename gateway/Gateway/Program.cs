

using FluentValidation;
using FluentValidation.AspNetCore;
using Gateway.Clients;

var builder = WebApplication.CreateBuilder(args);

builder.Services
    .AddOpenApi()
    .AddSwaggerGen()
    .AddValidatorsFromAssembly(typeof(Program).Assembly, includeInternalTypes: true)
    .AddFluentValidationAutoValidation()
    .AddControllers();

builder.Services.AddGrpcClient<SensorReadingService.SensorReadingServiceClient>(o =>
{
    o.Address = new Uri("http://localhost:4880");
});

builder.Services.AddSingleton<SensorReadingClient>();

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
