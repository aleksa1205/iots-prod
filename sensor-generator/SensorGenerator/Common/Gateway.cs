namespace SensorGenerator.Common;

public class Gateway
{
    public string Address { get; set; } = string.Empty;
    
    public string Endpoint { get; set; } = string.Empty;
    
    public int BatchSize { get; set; } = 0;
}