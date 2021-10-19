package main

import "fmt"

func main() {
	// Declarando constantes
	const pi float64 = 3.14
	const pi2 = 3.1414

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declarando variables enteras
	base := 12
	// |^| Se crea aqui, no es asignacion

	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values

	var a int     // Output: 0
	var b float64 // Output: 0
	var c string  // Output: " "
	var d bool    // Output: false

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Mult
	result = y * x
	fmt.Println("Mult:", result)

	// Div
	result = y / x
	fmt.Println("Div:", result)

	// Mod
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

}
