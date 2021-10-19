package main

import "fmt"

// FMT CLASE 10

func main() {
	// Variables
	helloMessage := "Hello"
	worldMessage := "world"

	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500

	// %s = indicar q recibe un string
	// %d = indicar q recibe un integer
	// %v = valor flexible, para cuando no sabemos el dato
	fmt.Printf("%s tiene mas de  %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de  %d cursos\n", nombre, cursos)

	// Sprintf - genera el string pero lo guarda, no imprime
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
