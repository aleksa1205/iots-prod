using System.ComponentModel.DataAnnotations;

namespace Gateway.DTOs.Response;

public record PageOfResponse<T>
{
    [Required] public required IEnumerable<T> Items { get; init; }
    
    [Required] public required int PageNumber { get; init; }
    
    [Required] public required int PageSize { get; init; }
    
    [Required] public required bool HasPreviousPage { get; init; }
    
    [Required] public required bool HasNextPage { get; init; }
    
    [Required] public required int TotalItems { get; init; }
}