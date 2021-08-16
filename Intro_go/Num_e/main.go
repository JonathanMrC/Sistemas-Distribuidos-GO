package main

import "fmt"

func main() {
	var fact, e, i float64
	e = 1 //fact(0) = 1
	fact = 1.0
	for i = 1.0; i <= 15.0; i += 1.0 {
		fact *= i
		e += 1.0 / fact
	}
	fmt.Println("e = ", e)
}
