using Gateway.Clients;
using Gateway.DTOs.Request.Filtering;
using Gateway.DTOs.Request.Id;
using Grpc.Core;
using Microsoft.AspNetCore.Mvc;

namespace Gateway.Controllers;

[ApiController]
[Route("[controller]")]
public class SensorReadingController(SensorReadingClient client) : ControllerBase
{
    
    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [ProducesResponseType(StatusCodes.Status500InternalServerError)]
    [HttpGet]
    public async Task<IActionResult> GetAll([FromQuery] PageParams request)
    {
        return Ok((await client.GetAllAsync(request)));
    }

    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [ProducesResponseType(StatusCodes.Status404NotFound)]
    [ProducesResponseType(StatusCodes.Status500InternalServerError)]
    [HttpGet("{id}")]
    public async Task<IActionResult> Get([FromRoute] IdRequest request)
    {
        return Ok((await client.GetByIdAsync(request)));
    }
}