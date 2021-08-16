package main

import (
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
		conexion, err := serv.Accept() //espera por un conexion
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(conexion)
	}
}

func handleClient(conexion net.Conn) { //controlador conexion
	data := make([]byte, 100)      //crea un arreglo de bytes
	bs, err := conexion.Read(data) //lee el arreglo
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", string(data[:bs])) //imprime el mensaje hasta donde se pudo leer
		fmt.Println("Bytes: ", bs)                  //imprime cantidad de bytes leidos
	}
	return
}

func main() {
	go servidor()
	var input string
	fmt.Scanln(&input)
}
