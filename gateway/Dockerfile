FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build
WORKDIR /app

COPY Gateway/Gateway.csproj ./Gateway/
RUN dotnet restore ./Gateway/Gateway.csproj

COPY Gateway/ ./Gateway/

RUN dotnet build ./Gateway/Gateway.csproj -c Release -o /app/build

FROM build AS publish
RUN dotnet publish './Gateway/Gateway.csproj' -c Release -o /app/publish

FROM mcr.microsoft.com/dotnet/aspnet:9.0
EXPOSE 8081
WORKDIR /app
COPY --from=publish /app/publish .

ENTRYPOINT ["dotnet", "Gateway.dll"]