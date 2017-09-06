using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading.Tasks;
using LitJson;
using System.Collections;
using System.Net.Http;

namespace CalculadoraRPC
{
    public class Person
    {
        // C# 3.0 auto-implemented properties
        public string Name { get; set; }
        public int Age { get; set; }
        public DateTime Birthday { get; set; }
    }

    public static class CalcController
    {
        private static JsonData itemData;
        private static string url = "http://httpbin.org/post";
        private static readonly HttpClient client = new HttpClient();
        
        public async static Task<string> Calculate()
        {
            Person person = new Person();

            person.Name = "William Shakespeare";
            person.Age = 51;
            person.Birthday = new DateTime(1564, 4, 26);

            string json_person = JsonMapper.ToJson(person);


            var content = new StringContent(json_person, Encoding.UTF8, "application/json");
            var response = await client.PostAsync(url, content);
            var responseString = await response.Content.ReadAsStringAsync();


            itemData = JsonMapper.ToObject(responseString);
            return itemData["json"]["Name"].ToString();
        }
    }
}
