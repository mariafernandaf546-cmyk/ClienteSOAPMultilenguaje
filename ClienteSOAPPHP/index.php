<?php

$wsdl = "https://www.dataaccess.com/webservicesserver/NumberConversion.wso?WSDL";

$cliente = new SoapClient($wsdl);

$respuesta = $cliente->NumberToWords(
    array(
        "ubiNum" => $_GET['n']
    )
);

echo $respuesta->NumberToWordsResult;

?>