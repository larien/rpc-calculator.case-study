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
        private static string url = "http://localhost:8080/main/calculadora";
        private static readonly HttpClient client = new HttpClient();
        
        public async static Task<string> Calculate()
        {
            //ExpressionModel expr = new ExpressionModel();

            //expr.Expression = "(1 + 4)";
            //expr.Value = "";

            //string json_expr = JsonMapper.ToJson(expr);


            //var content = new StringContent(json_expr, Encoding.UTF8, "application/json");
            //var response = await client.PostAsync(url, content);
            //var responseString = await response.Content.ReadAsStringAsync();

            var responseString = await client.GetStringAsync(url);


            itemData = JsonMapper.ToObject(responseString);
            return responseString;
        }
    }
}
