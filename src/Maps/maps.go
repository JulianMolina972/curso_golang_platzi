package main

import "fmt"

func main() {
	//Los mapas manejan operaciones de creación, inserción, eliminación,
	//-smodificación y recorrido.
	m := make(map[string]int) //make(map[key-type]val-type).

	m["Jose"] = 14
	m["David"] = 22

	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
