package main

import "fmt"

func main() {
	var lado int64
	fmt.Print("Ingrese un lado del cuadrado:\t")
	fmt.Scan(&lado)
	lado *= lado
	fmt.Println("El area del cuadrado es:\t", lado)
}
