FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build
WORKDIR /app

COPY SensorGenerator/SensorGenerator.csproj ./SensorGenerator/
RUN dotnet restore ./SensorGenerator/SensorGenerator.csproj

COPY SensorGenerator/ ./SensorGenerator/
RUN dotnet build ./SensorGenerator/SensorGenerator.csproj -c Release -o /app/build

FROM build AS publish
RUN dotnet publish './SensorGenerator/SensorGenerator.csproj' -c Release -o /app/publish

FROM mcr.microsoft.com/dotnet/aspnet:9.0
EXPOSE 8082
WORKDIR /app

COPY SensorGenerator/Dataset.csv ./Dataset.csv

COPY --from=publish /app/publish .

ENTRYPOINT ["dotnet", "SensorGenerator.dll"]