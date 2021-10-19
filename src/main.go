package main

import "fmt"

// Condicional if CLASE 15

func main() {
	// Defer: Antes de morirse la funcion
	//        se ejecuta la linea 10
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue, cuando una condition
		// me interesa, la freno pero continuo
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// Break, sentido contrario de continue
		if i == 8 {
			fmt.Println("break")
			break
		}
	}
}
