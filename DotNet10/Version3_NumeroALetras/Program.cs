using Humanizer;

var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/", (string? n) =>
{
    if (string.IsNullOrEmpty(n))
    {
        return "Uso: http://localhost:5002/?n=123";
    }

    if (!int.TryParse(n, out int numero))
    {
        return "Número inválido";
    }

    return numero.ToWords(new System.Globalization.CultureInfo("es"));
});

app.Run("http://localhost:5002");