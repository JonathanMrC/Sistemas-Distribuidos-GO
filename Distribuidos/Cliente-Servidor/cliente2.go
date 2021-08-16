package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func cliente() {
	conexion, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conexion.Close()
	mensaje := "Hola mundo"
	fmt.Println(mensaje)
	err = gob.NewEncoder(conexion).Encode(mensaje)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	go cliente()

	var input string
	fmt.Scanln(&input) //para poder terminar
}
