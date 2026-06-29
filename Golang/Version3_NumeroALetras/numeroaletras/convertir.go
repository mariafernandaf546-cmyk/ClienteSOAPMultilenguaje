package numeroaletras

import "fmt"

// Convertir convierte un número entero a letras.
func Convertir(numero int) string {

	if numero == 0 {
		return "cero"
	}

	if numero < 0 {
		return "menos " + Convertir(-numero)
	}

	if numero < 1000 {
		return centenas(numero)
	}

	if numero < 1000000 {

		miles := numero / 1000
		resto := numero % 1000

		var texto string

		if miles == 1 {
			texto = "mil"
		} else {
			texto = centenas(miles) + " mil"
		}

		if resto > 0 {
			texto += " " + centenas(resto)
		}

		return texto
	}

	if numero < 1000000000 {

		millones := numero / 1000000
		resto := numero % 1000000

		var texto string

		if millones == 1 {
			texto = "un millón"
		} else {
			texto = centenas(millones) + " millones"
		}

		if resto > 0 {

			if resto < 1000 {
				texto += " " + centenas(resto)
			} else {
				texto += " " + Convertir(resto)
			}
		}

		return texto
	}

	return fmt.Sprintf("%d", numero)
}

// ===============================
// CENTENAS
// ===============================

func centenas(n int) string {

	if n < 100 {
		return decenas(n)
	}

	switch {

	case n == 100:
		return "cien"

	case n < 200:
		return "ciento " + decenas(n-100)

	case n < 300:
		return "doscientos" + agregar(decenas(n-200))

	case n < 400:
		return "trescientos" + agregar(decenas(n-300))

	case n < 500:
		return "cuatrocientos" + agregar(decenas(n-400))

	case n < 600:
		return "quinientos" + agregar(decenas(n-500))

	case n < 700:
		return "seiscientos" + agregar(decenas(n-600))

	case n < 800:
		return "setecientos" + agregar(decenas(n-700))

	case n < 900:
		return "ochocientos" + agregar(decenas(n-800))

	default:
		return "novecientos" + agregar(decenas(n-900))
	}
}

func agregar(s string) string {

	if s == "" {
		return ""
	}

	return " " + s
}

// ===============================
// DECENAS
// ===============================

func decenas(n int) string {

	switch {

	case n < 10:
		return unidades(n)

	case n == 10:
		return "diez"

	case n == 11:
		return "once"

	case n == 12:
		return "doce"

	case n == 13:
		return "trece"

	case n == 14:
		return "catorce"

	case n == 15:
		return "quince"

	case n < 20:
		return "dieci" + unidades(n-10)

	case n == 20:
		return "veinte"

	case n < 30:

		switch n {
		case 21:
			return "veintiuno"
		case 22:
			return "veintidós"
		case 23:
			return "veintitrés"
		case 24:
			return "veinticuatro"
		case 25:
			return "veinticinco"
		case 26:
			return "veintiséis"
		case 27:
			return "veintisiete"
		case 28:
			return "veintiocho"
		case 29:
			return "veintinueve"
		}

	case n < 40:
		return combinar("treinta", n-30)

	case n < 50:
		return combinar("cuarenta", n-40)

	case n < 60:
		return combinar("cincuenta", n-50)

	case n < 70:
		return combinar("sesenta", n-60)

	case n < 80:
		return combinar("setenta", n-70)

	case n < 90:
		return combinar("ochenta", n-80)

	default:
		return combinar("noventa", n-90)
	}

	return ""
}

func combinar(prefijo string, unidad int) string {

	if unidad == 0 {
		return prefijo
	}

	return prefijo + " y " + unidades(unidad)
}

// ===============================
// UNIDADES
// ===============================

func unidades(n int) string {

	switch n {

	case 0:
		return ""

	case 1:
		return "uno"

	case 2:
		return "dos"

	case 3:
		return "tres"

	case 4:
		return "cuatro"

	case 5:
		return "cinco"

	case 6:
		return "seis"

	case 7:
		return "siete"

	case 8:
		return "ocho"

	case 9:
		return "nueve"
	}

	return ""
}

