package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	/*Reto: validar que si el texto contiene mayúsculas
	igual siga validando el palindromo*/
	var textReverse string
	text = strings.ToLower(text)               //pasamos el texto a minúsculas
	text = strings.ReplaceAll(" ", text, text) // Eliminamos los espacios

	for i := len(text) - 1; i >= 0; i-- { //Con el tamaño de text ir decrementando
		textReverse += string(text[i]) //Agregar elementos a la nueva variable
	}

	if text == textReverse {
		fmt.Println("Is palindromo")
	} else {
		fmt.Println("Is not palindromo")
	}
}

func main() {
	slice := []string{"Hello", "how", "are", "you"}

	for i := range slice {
		fmt.Println(i)
	}
	//ama
	//amo a roma

	isPalindromo("aMo A Roma")
}
