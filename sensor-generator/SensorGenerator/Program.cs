using SensorGenerator;
using Timer = System.Timers.Timer;

Console.WriteLine("Collecting data from connected IoT sensors...");

var projectRoot = Directory.GetParent(
    Directory.GetCurrentDirectory())
    .Parent
    .Parent
    .FullName;
string filePath = $"{projectRoot}/Dataset.csv";

if (!File.Exists(filePath))
{
    Console.WriteLine("Problem connecting to the IoT sensors");
    return;
}

var reader = new StreamReader(File.OpenRead(filePath));
reader.ReadLine();
while (!reader.EndOfStream)
{
    var line = reader.ReadLine();
    var values = line.Split(',');
    Console.WriteLine(line);
    var sensor = new Sensor
    {
        Timestamp = long.Parse(values[0]),
        UsedKw = Double.Parse(values[1]),
        GeneratedKw = Double.Parse(values[2]),
        Temperature = float.Parse(values[3]),
        Humidity = float.Parse(values[4]),
        ApparentTemperature = float.Parse(values[5]),
        Pressure = float.Parse(values[6]),
    };
    Thread.Sleep(1500);
}

reader.Close();

Console.WriteLine("All sensor data successfully collected and processed.");