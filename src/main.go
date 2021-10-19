package main

import "fmt"

// Functions CLASE 11

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(m1 string, m2 string, m3 string) {
	fmt.Println(m1, m2, m3)
}

func returnValue(a int) int {
	return a * 2
}

// Luego de los parametros, se declara lo que retorna tambien

func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument("Wena", "a", "a")

	value := returnValue(2)
	fmt.Println(value)

	// El _ es para ignorar un valor, en este caso la func retorna 2
	value1, _ := dobleReturn(2)
	fmt.Println(value1)
}
