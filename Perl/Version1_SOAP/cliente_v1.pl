use strict;
use warnings;
use SOAP::Lite;

print "Ingrese un numero: ";

my $numero = <STDIN>;
chomp($numero);

my $soap = SOAP::Lite
    -> uri('http://www.dataaccess.com/webservicesserver/')
    -> proxy('https://www.dataaccess.com/webservicesserver/NumberConversion.wso');

my $resultado = $soap
    -> call(
        'NumberToWords',
        SOAP::Data->name('ubiNum' => $numero)
    )
    -> result;

print "\nNumero en ingles:\n";
print "$resultado\n";