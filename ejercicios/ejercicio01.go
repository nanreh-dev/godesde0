package ejercicios

import "strconv"

func EsMayorOMenor(numero string) (int, string) {
	if num, err := strconv.Atoi(numero); err != nil {
		return 0, "no se pudo convertir el numero: " + err.Error()
	} else if num >= 100 {
		return num, "es mayor que 100"
	} else {
		return num, "es menor que 100"
	}
}
