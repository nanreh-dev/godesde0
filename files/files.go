package files

import (
	"fmt"
	"os"

	"github.com/hcamara/godesde0/ejercicios"
)

var filename string = "./files/txt/tablas.txt"

func GrabaTabla() {
	texto := ejercicios.CreaTablas()

	archivo, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error al crear el archivo" + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	texto := ejercicios.CreaTablas()

	if !Append(texto) {
		fmt.Println("Error al concatenar el contenido")
	}
}

func Append(texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error al agregar contenido: " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error al escribir el contenido")
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error al leer el archivo: " + err.Error())
		return
	}

	fmt.Println(string(archivo))
}
