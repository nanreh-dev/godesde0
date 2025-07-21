package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 0, 55}
var matriz [20][30]int

func MuestraArreglo() {
	tabla[7] = 33
	tabla[2] = 99

	tabla2 := [10]string{"hernan", "claudia"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[5][24] = 15
	fmt.Println(matriz)
}
