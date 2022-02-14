package main

import "fmt"

func par(number int) bool {
	return number%2 == 0
}

func main() {
	for i := 0; i < 100; i++ {
		var number_i int = i
		result := par(number_i)
		fmt.Println(result)
	}
}
