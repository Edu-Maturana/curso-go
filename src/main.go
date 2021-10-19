package main

import "fmt"

// Condicional if CLASE 15

func main() {
	modulo := 4 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// O tambien

	switch chumbi := 4 % 2; chumbi {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condicion
	value := 10
	switch {
	case value > 100:
		fmt.Println("Mayor a 100")
	case value < 0:
		fmt.Println("Menor que 0")
	default:
		fmt.Println("No condition")
	}

}
