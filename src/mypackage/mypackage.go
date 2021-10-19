package mypackage

import "fmt"

// Buena practica
// Nombre de la clase - lo que hace

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage - imprime un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}
