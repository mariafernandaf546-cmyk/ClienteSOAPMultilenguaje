using System.Net.Http;
using Microsoft.AspNetCore.Builder;

var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", async (string? n) =>
{
    if (string.IsNullOrEmpty(n))
        return "Uso: http://localhost:5000/?n=10";

    string soapRequest =
$@"<?xml version=""1.0"" encoding=""utf-8""?>
<soap:Envelope
xmlns:xsi=""http://www.w3.org/2001/XMLSchema-instance""
xmlns:xsd=""http://www.w3.org/2001/XMLSchema""
xmlns:soap=""http://schemas.xmlsoap.org/soap/envelope/"">
<soap:Body>
<NumberToWords xmlns=""http://www.dataaccess.com/webservicesserver/"">
<ubiNum>{n}</ubiNum>
</NumberToWords>
</soap:Body>
</soap:Envelope>";

    using var client = new HttpClient();

    var content = new StringContent(
        soapRequest,
        System.Text.Encoding.UTF8,
        "text/xml"
    );

    content.Headers.Add(
        "SOAPAction",
        "\"http://www.dataaccess.com/webservicesserver/NumberToWords\""
    );

    var response = await client.PostAsync(
        "https://www.dataaccess.com/webservicesserver/NumberConversion.wso",
        content
    );

    var xml = await response.Content.ReadAsStringAsync();

    int inicio = xml.IndexOf("<m:NumberToWordsResult>");
    int fin = xml.IndexOf("</m:NumberToWordsResult>");

    if (inicio >= 0 && fin >= 0)
    {
        inicio += "<m:NumberToWordsResult>".Length;
        return xml.Substring(inicio, fin - inicio);
    }

    return xml;
});

app.Run("http://localhost:5000");