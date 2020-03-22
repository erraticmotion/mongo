# NET + Mongo

## Dot Net Core 3.1

```console
wget -q https://packages.microsoft.com/config/ubuntu/18.04/packages-microsoft-prod.deb
sudo dpkg -i packages-microsoft-prod.deb
sudo add-apt-repository universe
sudo apt-get install apt-transport-https
sudo apt-get update
sudo apt-get install dotnet-sdk-3.1
```

## Connect
```console
dotnet new console -n connect
dotnet add package MongoDB.Driver --version 2.10.2
```

```csharp
namespace Connect
{
    using System;
    using MongoDB.Driver;

    public class Program
    {
        public static void Main(string[] args)
        {
            var client = new MongoClient("mongodb://localhost:27017");
            Console.WriteLine($"{client.Settings.Server}");
        }
    }
}
```

## Connect
```console
dotnet build
dotnet run
```