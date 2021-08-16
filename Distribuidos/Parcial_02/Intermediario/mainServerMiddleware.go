package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Informacion struct {
	Tema         string
	CantUsuarios int64
	MaxUsuarios  int64
}

type Direccion struct {
	Dir    string
	Puerto string
}

type Paquete struct {
	Info Informacion
	Dir  Direccion
}

type Estructura map[string]Paquete

var Salas Estructura
var Puertos []string

func (t *Estructura) ObtenerInfoSalas(args *bool, reply *[]Informacion) error {
	ActualizaDatos()
	var ans []Informacion
	for _, paquete := range Salas {
		ans = append(ans, paquete.Info)
	}
	*reply = ans
	return nil
}

func (t *Estructura) ObtenerDirSala(tema *Informacion, reply *Direccion) error {
	if _, existe := Salas[tema.Tema]; existe {
		if Salas[tema.Tema].Info.CantUsuarios == Salas[tema.Tema].Info.MaxUsuarios {
			var cadena string = "Error: La sala con tema: " + tema.Tema + "\t esta llena"
			return errors.New(cadena)
		}
		*reply = Salas[tema.Tema].Dir
		return nil
	}
	var cadena string = "Error: " + tema.Tema + "\tno hay salas con ese tema"
	return errors.New(cadena)
}

func ActualizaDatos() {
	//manda peticion para registrar/actualizar los datos de cada servidor de salas de chat
	for _, puerto := range Puertos {
		conexion, errc := net.Dial("tcp", puerto) //Conecta
		if errc != nil {
			fmt.Println("No hay servidor activo en el puerto:", puerto)
			continue
		}
		errE := gob.NewEncoder(conexion).Encode("I") //manda peticion de que es registro de datos
		if errE != nil {
			fmt.Println(errE)
			conexion.Close()
			continue
		}
		var paquete Paquete
		errD := gob.NewDecoder(conexion).Decode(&paquete) //recibe los datos del servidor
		if errD != nil {
			fmt.Println(errD)
			conexion.Close()
			continue
		}
		Salas[paquete.Info.Tema] = paquete
		conexion.Close()
	}
}

func Middleware() {
	var puerto string = ":4040"
	estructura := new(Estructura)
	err := rpc.Register(estructura)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Ejecutando servidor middleware en puerto" + puerto)
	rpc.HandleHTTP()
	escuchar, err := net.Listen("tcp", puerto)
	if err != nil {
		fmt.Println(err)
		return
	}
	err2 := http.Serve(escuchar, nil)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	errc := escuchar.Close()
	if errc != nil {
		fmt.Println(err)
	}
}

func main() {
	Salas = make(map[string]Paquete)
	Puertos = make([]string, 0)
	Puertos = append(Puertos, ":4041")
	Puertos = append(Puertos, ":4042")
	Puertos = append(Puertos, ":4043")
	go Middleware()
	var t string
	fmt.Scanln(&t)
}
