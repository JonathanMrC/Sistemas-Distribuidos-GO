package main

import "fmt"

func intercambia(a *int64, b *int64) {
	temp := *a
	*a = *b
	*b = temp
	return
}

func main() {
	var a, b int64
	fmt.Println("Ingrese 2 valores separados por un espacio")
	fmt.Scanln(&a, &b)
	fmt.Println("a -> ", a, "\tb -> ", b)
	intercambia(&a, &b)
	fmt.Println("a -> ", a, "\tb -> ", b)
	return
}
