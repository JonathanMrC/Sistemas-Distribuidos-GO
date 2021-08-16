package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

type Proceso struct {
	Id, Cont uint64
	Activo   bool
}

var p Proceso

func Ejecutar() {
	for p.Activo {
		fmt.Printf("id %d: %d\n", p.Id, p.Cont)
		p.Cont++
		time.Sleep(time.Millisecond * 500)
	}
}

func Peticion() {
	fmt.Println("\nPeticion de proceso...")
	conexion, err := net.Dial("tcp", ":9999") //Conecta
	defer conexion.Close()                    //cierra la conexion
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(conexion).Encode("0") //manda peticion de proceso al servidor
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewDecoder(conexion).Decode(&p) //recibe el proceso
	if err != nil {
		fmt.Println(err)
		return
	}
	p.Activo = true //activa el proceso
}
func Devolucion() {
	fmt.Println("\nPeticion para devolver el proceso...")
	conexion, err := net.Dial("tcp", ":9999") //Conecta
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conexion.Close()                     //cierra la conexion
	err = gob.NewEncoder(conexion).Encode("1") //Indica que tiene un proceso al servidor
	if err != nil {
		fmt.Println(err)
		return
	}
	p.Activo = false                          //detiene la ejecucion
	err = gob.NewEncoder(conexion).Encode(&p) //devuelve el proceso
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Proceso devuelto exitosamente!")
	return
}

func cliente() {
	Peticion()
	if !p.Activo {
		return
	}
	go Ejecutar()

	var input string
	fmt.Scanln(&input)

	Devolucion()
}

func main() {
	cliente()
}
