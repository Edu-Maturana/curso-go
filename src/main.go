package main

import "fmt"

//Structs
/*
Tipo de dato - nombre de la clase - struct {}
*/
type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
