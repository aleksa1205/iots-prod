using Gateway.Clients;
using Gateway.DTOs.Request.Filtering;
using Gateway.DTOs.Request.Id;
using Gateway.DTOs.Request.Sensor;
using Gateway.DTOs.Request.Time;
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

    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [ProducesResponseType(StatusCodes.Status500InternalServerError)]
    [HttpPost]
    public async Task<IActionResult> Create(SensorRequest request)
    {
        return Ok((await client.CreateAsync(request)));
    }

    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [ProducesResponseType(StatusCodes.Status404NotFound)]
    [ProducesResponseType(StatusCodes.Status500InternalServerError)]
    [HttpPut("{id}")]
    public async Task<IActionResult> Update([FromRoute] IdRequest request, SensorRequest requestData)
    {
        return Ok((await client.UpdateAsync(request, requestData)));
    }

    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [ProducesResponseType(StatusCodes.Status404NotFound)]
    [ProducesResponseType(StatusCodes.Status500InternalServerError)]
    [HttpDelete("{id}")]
    public async Task<IActionResult> Delete([FromRoute] IdRequest request)
    {
        await client.DeleteAsync(request);
        return NoContent();
    }

    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [ProducesResponseType(StatusCodes.Status500InternalServerError)]
    [HttpGet("min")]
    public async Task<IActionResult> GetByMinUsage([FromQuery] TimeRequest request)
    {
        return Ok(await client.GetByMinUsage(request));
    }
    
    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [ProducesResponseType(StatusCodes.Status500InternalServerError)]
    [HttpGet("max")]
    public async Task<IActionResult> GetByMaxUsage([FromQuery] TimeRequest request)
    {
        return Ok(await client.GetByMaxUsage(request));
    }

    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [ProducesResponseType(StatusCodes.Status500InternalServerError)]
    [HttpGet("avg")]
    public async Task<IActionResult> GetByAvgUsage([FromQuery] TimeRequest request)
    {
        return Ok(await client.GetAvgUsage(request));
    }

    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [ProducesResponseType(StatusCodes.Status500InternalServerError)]
    [HttpGet("sum")]
    public async Task<IActionResult> GetSum([FromQuery] TimeRequest request)
    {
        return Ok(await client.GetSumUsage(request));
    }

    [HttpPost("stream")]
    [ProducesResponseType(StatusCodes.Status200OK)]
    [ProducesResponseType(StatusCodes.Status400BadRequest)]
    [ProducesResponseType(StatusCodes.Status500InternalServerError)]
    public async Task<IActionResult> Stream([FromBody] IEnumerable<SensorRequest> readings)
    {
        if (readings is null || !readings.Any())
        {
            return BadRequest("No readings found");
        }
        
        await client.Stream(readings);
        return Ok();
    }
}