package main

import "fmt"

func say(text string, c chan<- string) {

	// Nombre del canal <- lo que ingresaremos a el
	c <- text
}

// A la inversa con salida, mismo de arriba

// func say(text string, c <-chan string) {

// 	// Nombre del canal <- lo que ingresaremos a el
// 	text = <- c
// }

func main() {
	// Crea chan  datatype limit
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	// El simbolo de linea 8 inverso, en vez de ingresar dato, devuelve la salida
	fmt.Println(<-c)
}
