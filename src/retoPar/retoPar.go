package main

import "fmt"

func par(number int) bool {
	return number%2 == 0
}

func main() {

	result := par(20)
	fmt.Println(result)
}
