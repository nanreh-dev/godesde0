package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Bs. As."

	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonatos := map[string]int{
		"barcelona": 20,
		"boca":      5,
		"chivas":    4,
		"river":     5,
	}

	fmt.Println(campeonatos)

	for equipo, puntaje := range campeonatos {
		fmt.Printf("Equipo: %s, Puntos: %d\n", equipo, puntaje)
	}

	delete(campeonatos, "boca")

	fmt.Println(campeonatos)

	puntaje, existe := campeonatos["juventus"]

	fmt.Printf("Puntaje: %d, existe: %t", puntaje, existe)

	puntaje2, existe2 := campeonatos["river"]

	fmt.Printf("\nPuntaje: %d, existe: %t", puntaje2, existe2)
}
