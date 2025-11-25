namespace Gateway.DTOs.Request.Filtering;

public record PageParams
{
    public int? Page { get; init; }
    
    public int? PageSize { get; init; }
}