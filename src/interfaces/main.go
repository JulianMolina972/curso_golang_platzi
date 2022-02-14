package main

import "fmt"

type figuras2D interface {
	area() float64
}
type cuadrado struct {
	base float64
}
type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista de interfaces
	//[1, "bo", true, 122.123]

	myInterface := []interface{}{"Hello", 12, 212.21, true}
	fmt.Println(myInterface...)
}

/*
interface{}es el tipo de interfaz vacía

[]interface{}es una porción de tipo interfaz vacía

interface{}{}es un literal compuesto de tipo de interfaz vacío

[]interface{}{}es una porción de tipo literales compuestos de interfaz vacía

*/
