<?php

$formateador = new NumberFormatter(
    "es",
    NumberFormatter::SPELLOUT
);

$res = $formateador->format($_GET['n']);

echo $res;

?>