require 'sinatra'
require 'nokogiri'
require 'net/http'

set :protection, :except => [:json_csrf, :csrf]

# Algoritmo matemático para convertir CUALQUIER número del 0 al 99 a palabras
def numero_a_letras(num)
  unidades = %w[cero uno dos tres cuatro cinco seis siete ocho nueve]
  diez_a_diecinueve = %w[diez once doce trece catorce quince dieciséis diecisiete dieciocho diecinueve]
  decenas = %w[[] [] veinte treinta cuarenta cincuenta sesenta setenta ochenta noventa]
  
  n = num.to_i
  return "Número fuera de rango (0-99)" if n < 0 || n > 99

  return unidades[n] if n < 10
  return diez_a_diecinueve[n - 10] if n < 20
  
  if n < 30
    return "veinte" if n == 20
    return "veinti#{unidades[n % 10]}"
  end
  
  if n % 10 == 0
    return decenas[n / 10]
  else
    return "#{decenas[n / 10]} y #{unidades[n % 10]}"
  end
end

post '/fake_soap' do
  request_body = request.body.read
  
  # Extracción clásica basada en el formato limpio enviado por C++
  numero = request_body[%r{<ubiNum>(.*?)</ubiNum>}, 1].to_s.strip
  idioma = request_body[%r{<idioma>(.*?)</idioma>}, 1].to_s.strip.downcase

  resultado = ""

  if idioma == "es"
    resultado = numero_a_letras(numero)
  elsif idioma == "en"
    begin
      # Consumo del WebService SOAP externo real (DataAccess) para inglés
      uri = URI('https://www.dataaccess.com/webservicesserver/NumberConversion.wso')
      req = Net::HTTP::Post.new(uri.path, { 'Content-Type' => 'text/xml; charset=utf-8' })
      req.body = "<?xml version=\"1.0\" encoding=\"utf-8\"?><soap:Envelope xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\"><soap:Body><NumberToWords xmlns=\"http://www.dataaccess.com/webservicesserver/\"><ubiNum>#{numero}</ubiNum></NumberToWords></soap:Body></soap:Envelope>"
      
      http = Net::HTTP.new(uri.hostname, uri.port)
      http.use_ssl = true
      http.open_timeout = 5
      http.read_timeout = 5
      
      res = http.request(req)
      res_doc = Nokogiri::XML(res.body)
      res_doc.remove_namespaces!
      resultado = res_doc.xpath('//NumberToWordsResult').text.strip
    rescue => e
      resultado = "Error en servicio externo: #{e.message}"
    end
  else
    resultado = "Idioma no soportado (Use 'es' o 'en')"
  end

  # Respuesta en formato SOAP estándar hacia C++
  content_type 'text/xml; charset=utf-8'
  "<?xml version=\"1.0\" encoding=\"utf-8\"?>\r\n<soap:Envelope xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">\r\n  <soap:Body>\r\n    <NumberToWordsResult>#{resultado}</NumberToWordsResult>\r\n  </soap:Body>\r\n</soap:Envelope>"
end