package main

import (
	"fmt"
	"math"
)

func normalFuntion(message1 string) {
	fmt.Println(message1)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)

}
func areaTrapecio1(basemin1, basemax1, alturaTrapecio1 float32) float32 {
	return ((basemax1 + basemin1) / 2) * alturaTrapecio1
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}
func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma: ", result)

	//Resta
	result = y - x
	fmt.Println("Resta: ", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación:", result)

	//División
	result = y / x
	fmt.Println("División:", result)

	//Módulo ó residuo

	result = y % x
	fmt.Println("Modulo: ", result)

	//Incremental
	x++
	fmt.Println("Incremental: ", x)

	//Decremental
	x--
	fmt.Println("Decremental: ", x)

	//Retos

	// Area de un rectangulo

	const base1 int = 13
	const altura1 int = 6
	areaRectangulo := base * altura

	fmt.Println("Area del rectangulo: ", areaRectangulo)

	// Área de un trapecio

	const basemin int = 5
	const basemax int = 10
	const alturaTrapecio int = 4

	areaTrapecio := ((basemax + basemin) / 2) * alturaTrapecio

	fmt.Println("Area del trapecio", areaTrapecio)

	// Área de un circulo

	const radio float64 = 3.891291

	areaCirculo := pi * math.Pow(radio, 2)

	fmt.Println(areaCirculo)

	//Tipos de datos primitivos

	/*Numeros enteros
	int = Depende del OS (32 o 64 bits)
	int8 = 8bits = -128 a 127
	int16 = 16bits = -2^15 a 2^15-1
	int32 = 32bits = -2^31 a 2^31-1
	int64 = 64bits = -2^63 a 2^63-1

	Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	uint = Depende del OS (32 o 64 bits)
	uint8 = 8bits = 0 a 127
	uint16 = 16bits = 0 a 2^15-1
	uint32 = 32bits = 0 a 2^31-1
	uint64 = 64bits = 0 a 2^63-1

	numeros decimales
	float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	textos y booleanos
	string = ""
	bool = true or false

	numeros complejos
	Complex64 = Real e Imaginario float32
	Complex128 = Real e Imaginario float64
	Ejemplo : c:=10 + 8i */

	//Declaración de variables

	helloMessage := "Hello"
	worldMessge := "world"

	//Println
	fmt.Println(helloMessage, worldMessge)
	fmt.Println(helloMessage, worldMessge)

	// Printf
	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)

	fmt.Println(message)

	//Tipo datos

	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("Cursos: %T\n", cursos)

	normalFuntion("Hola mundo")
	tripeArgument(1, 12, "Hola")
	value := returnValue(2)
	fmt.Println("Value: ", value)

	//value1, value2 := doubleReturn(2)
	//fmt.Println("Value1 y value2: ", value1, value2)

	value1, _ := doubleReturn(2)
	fmt.Println("Value1: ", value1)
	//Reto funciones

	value3 := areaTrapecio1(5, 10, 4)
	fmt.Println("El área de un trapecio", value3)

	//Ciclo For condition

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")
	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever

	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// 	if counterForever > 210213 {
	// 		break
	// 	}
	// }

	// RETO

	for decremento := 100; decremento >= 0; decremento-- {
		fmt.Println(decremento)
	}

	valor1 := 3
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With and

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Son iguales")
	} else {
		fmt.Println("No son iguales")
	}

	//  With or

	if valor1 == 1 || valor2 == 5 {
		fmt.Println("Se cumple una condición")
	} else {
		fmt.Println("No se cumple ninguna condición")
	}

	// Switch

	modulo := 4 % 2
	//switch modulo := 4 % 2; modulo
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	//Sin condicion
	value10 := 50
	switch {
	case value10 > 100:
		fmt.Println("Es mayor a 100")
	case value10 < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}

	// Defer
	// defer es un killword donde cierra el proceso
	//defer fmt.Println("Hello")
	fmt.Println("Mundo")

	// Continue y break

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue

		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//Break

		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice

	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
