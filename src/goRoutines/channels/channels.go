package main

import "fmt"

func say(text string, c chan<- string) { //  Canal de entrada de datos
	c <- text //Para datos de entreada in
	//text = <- c       //Para datos out salida
}

func listen(c <-chan string) { //Canal de salida de datos
	var output string
	output = <-c
	fmt.Println(output)
}

func main() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	//fmt.Println(<-c)
	listen(c)
}
