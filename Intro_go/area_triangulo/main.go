package main

import "fmt"

func main() {
	var base float64
	var altura float64
	fmt.Print("Calcular el area del triangulo\nIngrese base:\t")
	fmt.Scan(&base)
	fmt.Print("Ingrese altura:\t")
	fmt.Scan(&altura)
	altura *= base / 2
	fmt.Print("Area:\t", altura)
}
