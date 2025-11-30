using Gateway.DTOs.Response;
using Gateway.DTOs.Response.Sensor;

namespace Gateway.DTOs.Extensions;

public static class PaginatedSensorReadingResponseExtensions
{
    public static PageOfResponse<SensorResponse> ToResponse(this PaginationSensorReadingResponse response)
    {
        return new PageOfResponse<SensorResponse>
        {
            Items = response.Items.Select(item => item.ToResponse()),
            PageNumber = response.PageNumber,
            PageSize = response.PageSize,
            HasNextPage = response.HasNextPage,
            HasPreviousPage = response.HasPreviousPage,
            TotalItems = response.TotalItems,
        };
    }
}