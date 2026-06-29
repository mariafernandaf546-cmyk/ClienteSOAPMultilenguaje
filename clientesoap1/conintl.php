<?php

if (!isset($_GET['n'])) {
    die("Ingrese un número.");
}

$formateador = new NumberFormatter(
    "es",
    NumberFormatter::SPELLOUT
);

echo $formateador->format((int)$_GET['n']);

?>