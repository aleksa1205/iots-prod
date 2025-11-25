using Gateway.Clients;
using Gateway.DTOs;
using Gateway.DTOs.Request.Id;
using Microsoft.AspNetCore.Mvc;

namespace Gateway.Controllers;

[ApiController]
[Route("[controller]")]
public class SensorReadingController(SensorReadingClient client) : ControllerBase
{
    
    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status500InternalServerError)]
    [HttpGet]
    public async Task<IActionResult> GetAll()
    {
        return Ok((await client.GetAllAsync()));
    }

    [HttpGet("{id}")]
    public async Task<IActionResult> Get([FromRoute] IdRequest request)
    {
        return Ok((await client.GetByIdAsync(request)));
    }
}