use strict;
use warnings;

use SOAP::Lite;
use LWP::UserAgent;
use JSON;
use URI::Escape;

print "Ingrese un numero: ";

my $numero = <STDIN>;
chomp($numero);

# ==============================
# Consumir el servicio SOAP
# ==============================

my $soap = SOAP::Lite
    -> uri('http://www.dataaccess.com/webservicesserver/')
    -> proxy('https://www.dataaccess.com/webservicesserver/NumberConversion.wso');

my $ingles = $soap
    -> call(
        'NumberToWords',
        SOAP::Data->name('ubiNum' => $numero)
    )
    -> result;

print "\nNumero en ingles:\n";
print "$ingles\n";

# ==============================
# Traducir con Google Translate
# ==============================

my $url =
"https://translate.googleapis.com/translate_a/single?client=gtx&sl=en&tl=es&dt=t&q="
. uri_escape($ingles);

my $ua = LWP::UserAgent->new;

my $response = $ua->get($url);

die "Error al traducir\n"
    unless $response->is_success;

my $json = decode_json($response->decoded_content);

my $espanol = $json->[0][0][0];

print "\nNumero en español:\n";
print "$espanol\n";