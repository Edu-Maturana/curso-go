package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {

	// Ultima en ejecutar
	defer wg.Done()

	fmt.Println(text)
}

func main() {
	// Acumula conjunto de goroutines y las ejecuta
	var wg sync.WaitGroup

	// Escribe Hello y anade una goroutine
	fmt.Println("Hello")
	wg.Add(1)

	// Aqui se ejecuta esa goroutine, con el puntero de wg
	go say("world", &wg)

	go func(text string) {
		fmt.Println(text)
	}("Adios")
	// Esperar a que todas las rutinas de wg se ejecuten
	wg.Wait()
}
