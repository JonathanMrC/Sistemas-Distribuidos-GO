package main

import "fmt"

func Fibonacci(n int64) int64 {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	var n int64
	fmt.Println("Ingrese el n fibonacci a buscar")
	fmt.Scanln(&n)
	fmt.Print("Fibonacci en pos: ", n, " = ", Fibonacci(n))
	return
}
