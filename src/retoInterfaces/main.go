package main

import (
	rf "curso_golang_platzi/src/retoInterfaces/retoInterface"
	f "fmt"
)

func main() {

	myRectangulo := rf.Rectangulo{Base: 10, Altura: 5}
	f.Println(myRectangulo)
	rf.Calcular(myRectangulo)

	myTrapecio := rf.Trapecio{Basemin: 5, Basemax: 10, Altura: 5}
	f.Println(myTrapecio)
	rf.Calcular(myTrapecio)

	myCirculo := rf.Circulo{Radio: 12.21}
	f.Println(myCirculo)
	rf.Calcular(myCirculo)
}
