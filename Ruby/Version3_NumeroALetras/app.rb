require 'sinatra'
require 'numbers_and_words'
require 'i18n'

# 🌟 Configuración obligatoria para activar el español en la gema del profesor
I18n.enforce_available_locales = false
I18n.locale = :es

get '/' do
  numero = params['n']
  return 'Uso: http://localhost:4567/?n=25' unless numero

  # El método del profesor ahora sí responderá en español
  numero.to_i.to_words(lang: :es)
end