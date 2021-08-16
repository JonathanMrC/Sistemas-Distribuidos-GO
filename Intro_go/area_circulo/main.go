package main

import "fmt"

func main() {
	var radio float64
	const PI float64 = 3.1416
	fmt.Print("Calcular el area de un circulo:\nIngrese el radio:\t")
	fmt.Scan(&radio)
	radio *= radio
	radio *= PI
	fmt.Print("Area:\t", radio)
}
