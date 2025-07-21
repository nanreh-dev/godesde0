package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LLamaClosure() {
	tabladel2 := 2

	tabla2 := tabla(tabladel2)

	for i := 1; i <= 10; i++ {
		fmt.Println(tabla2())
	}
}
