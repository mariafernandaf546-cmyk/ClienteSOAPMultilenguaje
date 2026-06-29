<?php

require_once 'vendor/autoload.php';

use Stichoza\GoogleTranslate\GoogleTranslate;

$wsdl = "https://www.dataaccess.com/webservicesserver/NumberConversion.wso?WSDL";

$cliente = new SoapClient($wsdl);

$respuesta = $cliente->NumberToWords(
    array("ubiNum" => $_GET['n'])
);

$traductor = new GoogleTranslate('es', 'en');

echo $traductor->translate(
    $respuesta->NumberToWordsResult
);

?>