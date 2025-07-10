package main

import (
	"fmt"
	"runtime"

	"github.com/hcamara/godesde0/ejercicios"
	"github.com/hcamara/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(1579)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." {
		fmt.Println("Esto no es windows")
	} else {
		fmt.Println("Esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es windows")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	esMenor, msj := ejercicios.EsMayorOMenor("1000")

	fmt.Print(esMenor)
	fmt.Printf(" %s\n", msj)

	esMayor, msj := ejercicios.EsMayorOMenor("10")

	fmt.Print(esMayor)
	fmt.Printf(" %s\n", msj)

	esError, msj := ejercicios.EsMayorOMenor("numero")

	fmt.Print(esError)
	fmt.Printf(" %s", msj)
}
