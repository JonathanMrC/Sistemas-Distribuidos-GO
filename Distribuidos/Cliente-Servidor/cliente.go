package main

import (
	"fmt"
	"net"
)

func cliente() {
	conexion, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	mensaje := "Hola mundo"
	fmt.Println(mensaje)
	conexion.Write([]byte(mensaje))
	conexion.Close()
}

func main() {
	go cliente()

	var input string
	fmt.Scanln(&input) //para poder terminar
}
