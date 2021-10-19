package main

import (
	"fmt"
)

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c))

	// Range y close

	// Close, cierra el canal y no recibe mas a pesar de su capacidad
	close(c)

	// c <- "Mensaje3"

	// Range, para recorrer
	for message := range c {
		fmt.Println(message)
	}

	// Select, para cuando no se que canal respondera primero
	email1 := make(chan string)
	email2 := make(chan string)
	go message("Mensaje1 ", email1)
	go message("Mensaje2 ", email2)

	for i := 0; i < 2; i++ {
		select {
		// En case de que m1, venga del canal email1
		case m1 := <-email1:
			fmt.Println("Email recibido de", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de", m2)
		}
	}
}

// Nota: Lo ideal es cerrar un canal si se que no voy a recibir mas datos
