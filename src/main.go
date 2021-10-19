package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// Crea un metodo ping() dentro del struct myPC, instanciado de pc
func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50
	// el prefijo & me da la direccion en memoria de a
	b := &a

	fmt.Println(b)
	fmt.Println(*b)
	// El prefijo * es el contrario al &, es para acceder
	// al valor que la direccion tiene

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)

}
