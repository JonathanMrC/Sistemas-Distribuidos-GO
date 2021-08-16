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

func cliente(obj Persona) {
	conexion, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conexion.Close()
	err = gob.NewEncoder(conexion).Encode(obj)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	persona := Persona{
		Nombre: "Fer",
		Email:  []string{"fer@gmail.com"},
	}
	go cliente(persona)

	var input string
	fmt.Scanln(&input) //para poder terminar
}
