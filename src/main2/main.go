package main

import (
	cl "curso_golang_platzi/src/mispaquetes"
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

//ping, nos retorna un mensaje aleatorio o asignado

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2

}

func main() {

	a := 50
	b := &a //puntero, direccion donde esta guardado a
	// & acceder a la direccion de memoria
	// * acceder al valor de es direccion de memoria
	fmt.Println(b)
	fmt.Println(*b)
	//a y *b apuntan a la misma dirección de memorias
	*b = 100
	fmt.Println(a)

	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRAM()

	fmt.Println(myPc)
	myPc.duplicateRAM()

	fmt.Println(myPc)

	//Modificadores de acceso

	var misCelu cl.Celulares
	misCelu.Marca = "Xiaomi"
	misCelu.Version = 2018
	fmt.Println(misCelu)
}
