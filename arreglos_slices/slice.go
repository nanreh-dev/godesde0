package arreglos_slices

import "fmt"

var tabla3 []int = []int{3, 7, 9}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestraSlice() {
	fmt.Println(tabla3)

	porcion := arreglo[3:]  //creado desde arreglo desde pos 3 hasta final
	porcion2 := arreglo[:5] //desde el principio hasta la posicion 5
	porcion3 := arreglo[6:7]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

	nums := make([]int, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
}
