package main

import (
	"fmt"
	"os"
	"sort"
)

func Respaldar(nombreArchivo string, datos []string) {
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer archivo.Close() //cierro el archivo al final

	for i := 0; i < len(datos); i++ {
		archivo.WriteString(datos[i] + "\n")
	}
	return
}

func main() {
	var cant uint64
	var temp string
	var cadenas []string

	fmt.Print("Ingrese la cantidad de strings a guardar:\n-> ")
	fmt.Scanln(&cant)

	for ; cant != 0; cant-- {
		fmt.Scan(&temp)
		cadenas = append(cadenas, temp)
	}

	sort.Strings(cadenas) //ascendente
	Respaldar("ascendente.txt", cadenas)

	sort.Sort(sort.Reverse(sort.StringSlice(cadenas))) //descendente
	Respaldar("descendente.txt", cadenas)
}
