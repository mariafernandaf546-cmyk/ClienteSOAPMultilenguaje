use strict;
use warnings;
use utf8;
use open ':std', ':encoding(UTF-8)';

use Lingua::Num2Word qw(cardinal);

print "Ingrese un numero: ";

my $numero = <STDIN>;
chomp($numero);

my $texto = cardinal('spa', $numero);

print "\nNumero en español:\n";
print "$texto\n";