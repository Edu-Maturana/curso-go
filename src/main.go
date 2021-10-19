package main

import "fmt"

// MAPS || DICCIONARIOS || OBJETOS

func main() {
	//           llave  valor
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer map: Output Jose 14, Pepito 20
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	value, ok := m["Josep"]
	fmt.Println(value, ok)
}
