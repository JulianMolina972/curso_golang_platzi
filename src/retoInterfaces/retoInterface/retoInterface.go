package retointerface

import (
	"fmt"
	"math"
)

//Retos

// AreaFigurasGeo
type figuras2D interface {
	area() float64
}

//Rectangulo
type Rectangulo struct {
	Base   float64
	Altura float64
}

//Trapecio
type Trapecio struct {
	Basemin float64
	Basemax float64
	Altura  float64
}

//Circulo
type Circulo struct {
	Radio float64
}

func (R Rectangulo) area() float64 {
	return R.Base * R.Altura

}

func (T Trapecio) area() float64 {
	return ((T.Basemax + T.Basemin) / 2) * T.Altura
}

func (C Circulo) area() float64 {
	return math.Pi * math.Pow(C.Radio, 2)
}

//Calcular, para calcular el area de las figuras geométricas
func Calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}
