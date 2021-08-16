package main

import "fmt"

func main() {
	var ans, valor float64
	fmt.Print("Convertir grados Fahrenheit a Celcius:\nIngrese los grados Fahrenheit:\t")
	fmt.Scan(&valor)
	ans = (valor - 32.0) * (5.0 / 9.0)
	fmt.Print(valor, " Fahrenheit = ", ans, " Celcius")
}
