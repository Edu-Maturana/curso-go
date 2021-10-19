package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GBs de RAM , %d GBs de disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	myPC := pc{ram: 16, brand: "msi", disk: 100}
	fmt.Println(myPC)
}
