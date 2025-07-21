package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var texto string

func CreaTablas() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese número: ")
		if scanner.Scan() {
			numero, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El número ingresado es incorrecto")
				continue
			}
			for i := 1; i <= 10; i++ {
				texto += fmt.Sprintf("%d x %d = %d\n", numero, i, numero*i)
			}
		}
		break
	}
	return texto
}
