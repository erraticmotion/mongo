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
