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
    public static class CalcController
    {
        private static JsonData itemData;
        private static string url = "http://172.16.59.218:25000/main";
        private static readonly HttpClient client = new HttpClient();

        public async static Task<string> CalculateExpression(string expression)
        {
            string result = string.Empty;

            ExpressionModel model = new ExpressionModel();
            model.Expression = expression;

            string json_model = JsonMapper.ToJson(model);

            var content = new StringContent(json_model, Encoding.UTF8, "application/json");

            try
            {
                var response = await client.PostAsync(url + "/calculadora", content);
                var responseString = await response.Content.ReadAsStringAsync();

                itemData = JsonMapper.ToObject(responseString);

                try
                {
                    result = itemData["Result"].ToString();
                }
                catch
                {
                    result = "KEY RETRIEVE ERROR";
                }
            }
            catch
            {
                result = "THE ANSWER IS A LIE";
            }

            return result;
        }

        public async static Task<string> CalculateFunction(string function, string value)
        {
            string result = string.Empty;

            FunctionModel model = new FunctionModel();

            model.Function = function;
            model.Value = value;

            string json_expr = JsonMapper.ToJson(model);

            var content = new StringContent(json_expr, Encoding.UTF8, "application/json");

            try
            {
                var response = await client.PostAsync(url + "/funcao", content);
                var responseString = await response.Content.ReadAsStringAsync();

                itemData = JsonMapper.ToObject(responseString);

                try
                {
                    result = itemData["Result"].ToString();
                }
                catch
                {
                    result = "KEY RETRIEVE ERROR";
                }
            }
            catch
            {
                result = "THE ANSWER IS A LIE";
            }

            return result;
        }
    }
}
