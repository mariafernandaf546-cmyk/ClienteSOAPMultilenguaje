require 'sinatra'
require 'nokogiri'
require 'net/http'

set :protection, :except => [:json_csrf, :csrf]

# Traductor local a Español
def numero_a_letras_es(num)
  unidades = %w[cero uno dos tres cuatro cinco seis siete ocho nueve]
  diez_a_diecinueve = %w[diez once doce trece catorce quince diecisiete dieciocho diecinueve]
  diez_a_diecinueve.insert(6, "dieciséis") 
  decenas = %w[[] [] veinte treinta cuarenta cincuenta sesenta setenta ochenta noventa]
  
  n = num.to_i
  return "Número fuera de rango" if n < 0 || n > 99
  return unidades[n] if n < 10
  return diez_a_diecinueve[n - 10] if n < 20
  if n < 30
    return "veinte" if n == 20
    return "veinti#{unidades[n % 10]}"
  end
  return decenas[n / 10] if n % 10 == 0
  "#{decenas[n / 10]} y #{unidades[n % 10]}"
end

# Traductor local a Inglés (¡Para salvar la versión 1 de caídas de red!)
def numero_a_letras_en(num)
  unidades = %w[zero one two three four five six seven eight nine]
  diez_a_diecinueve = %w[ten eleven twelve thirteen fourteen fifteen sixteen seventeen eighteen nineteen]
  decenas = %w[[] [] twenty thirty forty fifty sixty seventy eighty ninety]
  
  n = num.to_i
  return "Number out of range" if n < 0 || n > 99
  return unidades[n] if n < 10
  return diez_a_diecinueve[n - 10] if n < 20
  return decenas[n / 10] if n % 10 == 0
  "#{decenas[n / 10]}-#{unidades[n % 10]}"
end

post '/fake_soap' do
  request_body = request.body.read
  
  numero = request_body[%r{<ubiNum>(.*?)</ubiNum>}, 1].to_s.strip
  idioma_raw = request_body[%r{<idioma>(.*?)</idioma>}, 1]
  idioma = idioma_raw ? idioma_raw.strip.downcase : "en"

  resultado = if idioma == "es"
                numero_a_letras_es(numero)
              else
                numero_a_letras_en(numero)
              end

  content_type 'text/xml; charset=utf-8'
  "<?xml version=\"1.0\" encoding=\"utf-8\"?>\r\n<soap:Envelope xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">\r\n  <soap:Body>\r\n    <NumberToWordsResult>#{resultado}</NumberToWordsResult>\r\n  </soap:Body>\r\n</soap:Envelope>"
end