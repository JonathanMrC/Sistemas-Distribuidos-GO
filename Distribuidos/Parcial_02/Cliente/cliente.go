package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"strconv"
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

type Datos struct {
	Usuario string
	Mensaje string
}

var terminar bool
var this Datos

func getline() string {
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')
	entrada = entrada[:len(entrada)-2]
	return entrada
}

func SubMenu(salas []Informacion) Informacion {
	if len(salas) == 0 {
		return Informacion{"No hay salas", 0, 0}
	}
	fmt.Println("Salas disponibles:")
	for i := 0; i < len(salas); i++ {
		fmt.Println(i+1, " -> ", salas[i].Tema, "\tUsuarios conectados: ", salas[i].CantUsuarios)
	}
	for {
		fmt.Println("Ingrese a que sala quiere conectarse:")
		var opc string
		fmt.Scanln(&opc)
		if opc == "1" || opc == "2" || opc == "3" {
			aux, _ := strconv.Atoi(opc)
			return salas[aux-1]
		} else {
			fmt.Println("La opcion ingresada no existe...")
		}
	}
}

func PeticionInfoSalas() Informacion {
	fmt.Println("Obteniendo informacion de las salas de chat")
	var salas []Informacion
	var puertomw string = ":4040"
	conexion, err := rpc.DialHTTP("tcp", puertomw)
	if err != nil {
		fmt.Println("Error no se pudo conectar con el middleware")
		return Informacion{"ERROR", -1, -1}
	}
	errc := conexion.Call("Estructura.ObtenerInfoSalas", true, &salas)
	if errc != nil {
		fmt.Println("Error al obtener la informacion de las salas", errc)
		return Informacion{"Error al obtener la informacion de las salas", -1, -1}
	}
	conexion.Close()
	return SubMenu(salas)
}

func PeticionDirSala(sala Informacion) Direccion {
	fmt.Println("Obteniendo direccion de la sala con tema: ", sala.Tema)
	var dir Direccion
	var puertomw string = ":4040"
	conexion, err := rpc.DialHTTP("tcp", puertomw)
	if err != nil {
		fmt.Println("Error no se pudo conectar con el middleware")
		return Direccion{"ERROR", "ERROR"}
	}
	errc := conexion.Call("Estructura.ObtenerDirSala", sala, &dir)
	if errc != nil {
		fmt.Println(errc)
		return Direccion{"ERROR", "ERROR"}
	}
	return dir
}

func Conectar(sala Informacion, dir Direccion) net.Conn {
	fmt.Println("Conectando al chat con tematica: ", sala.Tema)
	conexion, err := net.Dial("tcp", dir.Dir+dir.Puerto) //Conecta
	if err != nil {
		fmt.Println("No se pudo realizar la conexion", err)
		return conexion
	}
	fmt.Println("Conexion realizada...")
	return conexion
}

func Recibir(conexion net.Conn) {
	terminar = false
	for !terminar {
		var mensaje Datos
		err := gob.NewDecoder(conexion).Decode(&mensaje) //recibe la peticion
		if err != nil && !terminar {
			fmt.Println("Error al recibir mensajes: ", err)
			return
		}
		fmt.Println(mensaje.Usuario, ":\n", mensaje.Mensaje)
	}
}

func Chat(conexion net.Conn) {
	errR := gob.NewEncoder(conexion).Encode("C") //indica al servidor de chat que es un cliente
	if errR != nil {
		fmt.Println(errR)
		return
	}
	fmt.Println("Ingrese el nombre de usuario")
	this.Usuario = getline()
	fmt.Println("Para terminar el chat escriba:\n/salir")
	go Recibir(conexion) //se pone a escuchar
	fmt.Println("------------------------------------------------------------------")
	this.Mensaje = ""
	for {
		this.Mensaje = getline() //lee el mensaje
		if this.Mensaje == "" {
			continue
		}
		//tratar de poner un select de chan string
		errM := gob.NewEncoder(conexion).Encode(&this) //manda los datos {nombre/mensaje}
		if errM != nil {
			fmt.Println("Error al enviar", errM)
			return
		}
		if this.Mensaje == "/salir" {
			terminar = true
			fmt.Println("Sesión terminada")
			break
		}
	}
}

func main() {
	sala := PeticionInfoSalas()
	if sala.CantUsuarios != -1 {
		dir := PeticionDirSala(sala)
		if dir.Dir != "ERROR" {
			con := Conectar(sala, dir)
			Chat(con)
		}
	}
}
