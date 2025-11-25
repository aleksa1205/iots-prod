using System.ComponentModel.DataAnnotations;
using Microsoft.AspNetCore.Mvc;

namespace Gateway.DTOs.Request.Id;

public record IdRequest
{
    [FromRoute(Name = "id")]
    [Required] public required string Id { get; init; }
}