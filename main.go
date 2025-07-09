package main

import (
	"fmt"

	"github.com/hcamara/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoATexto(1579)
	fmt.Println(estado)
	fmt.Println(texto)
}
