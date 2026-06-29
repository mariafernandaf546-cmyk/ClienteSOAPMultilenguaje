package main

import (
	"fmt"

	"version3local/numeroaletras"
)

func main() {

	var numero int

	fmt.Print("Ingrese un número: ")
	fmt.Scan(&numero)

	fmt.Println()
	fmt.Println("Número en español:")
	fmt.Println(numeroaletras.Convertir(numero))

}