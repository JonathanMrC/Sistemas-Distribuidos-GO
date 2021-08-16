package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func servidor() {
	serv, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conexion, err := serv.Accept() //espera por una conexion
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(conexion)
	}
}

func handleClient(conexion net.Conn) { //controlador conexion
	var mensaje string
	err := gob.NewDecoder(conexion).Decode(&mensaje)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", mensaje) //imprime el mensaje hasta donde se pudo leer
	}
	return
}

func main() {
	go servidor()
	var input string
	fmt.Scanln(&input)
}
