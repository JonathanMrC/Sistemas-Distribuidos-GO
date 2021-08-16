package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Persona struct {
	Nombre string
	Email  []string
}

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
	var obj Persona
	err := gob.NewDecoder(conexion).Decode(&obj)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", obj) //imprime el mensaje hasta donde se pudo leer
	}
	return
}

func main() {
	go servidor()
	var input string
	fmt.Scanln(&input)
}
