package main

import (
	"fmt"
	pk "std/curso-go/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	pk.PrintMessage("Hola platzi")
}
