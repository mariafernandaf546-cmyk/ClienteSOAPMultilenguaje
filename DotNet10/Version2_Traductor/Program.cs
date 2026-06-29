using System.Text;
using System.Text.RegularExpressions;
using System.Net.Http;

var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", async (string? n) =>
{
    if (string.IsNullOrEmpty(n))
    {
        return "Uso: http://localhost:5001/?n=100";
    }

    using HttpClient client = new();

    string soapEnvelope = $@"<?xml version=""1.0"" encoding=""utf-8""?>
<soap:Envelope xmlns:xsi=""http://www.w3.org/2001/XMLSchema-instance""
xmlns:xsd=""http://www.w3.org/2001/XMLSchema""
xmlns:soap=""http://schemas.xmlsoap.org/soap/envelope/"">
<soap:Body>
<NumberToWords xmlns=""http://www.dataaccess.com/webservicesserver/"">
<ubiNum>{n}</ubiNum>
</NumberToWords>
</soap:Body>
</soap:Envelope>";

    var soapContent = new StringContent(
        soapEnvelope,
        Encoding.UTF8,
        "text/xml"
    );

    soapContent.Headers.Add(
        "SOAPAction",
        "\"http://www.dataaccess.com/webservicesserver/NumberToWords\""
    );

    var soapResponse = await client.PostAsync(
        "https://www.dataaccess.com/webservicesserver/NumberConversion.wso",
        soapContent
    );

    string xml = await soapResponse.Content.ReadAsStringAsync();

    Match match = Regex.Match(
        xml,
        @"<m:NumberToWordsResult>(.*?)</m:NumberToWordsResult>"
    );

    if (!match.Success)
    {
        return "Error al obtener resultado SOAP";
    }

    string ingles = match.Groups[1].Value.Trim();

    Console.WriteLine("INGLES: " + ingles);

    string url =
        $"https://api.mymemory.translated.net/get?q={Uri.EscapeDataString(ingles)}&langpair=en|es";

    string json = await client.GetStringAsync(url);

    Console.WriteLine("JSON:");
    Console.WriteLine(json);

    Match traduccion = Regex.Match(
        json,
        @"""translatedText"":""(.*?)"""
    );

    string español = traduccion.Success
        ? traduccion.Groups[1].Value
        : ingles;

    Console.WriteLine("ESPAÑOL: " + español);

    return español;
});

app.Run("http://localhost:5001");