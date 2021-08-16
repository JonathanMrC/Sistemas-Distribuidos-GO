package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

type Archivo struct {
	Nombre string
	Datos  []byte
}

var Mensajes []string

func Recibir(conexion net.Conn) {
	var peticion string
	for {
		err := gob.NewDecoder(conexion).Decode(&peticion) //recibe la peticion
		if err != nil {
			fmt.Println(err)
			return
		}
		if peticion == "M" {
			var mensaje string
			err := gob.NewDecoder(conexion).Decode(&mensaje) //recibe la peticion
			if err != nil {
				fmt.Println(err)
				return
			}
			Mensajes = append(Mensajes, mensaje)
		} else if peticion == "A" {
			var aux Archivo
			err := gob.NewDecoder(conexion).Decode(&aux) //recibe la peticion
			if err != nil {
				fmt.Println(err)
				return
			}
			archivo, err2 := os.Create(aux.Nombre)
			if err2 != nil {
				fmt.Println(err2)
				return
			}
			archivo.Write(aux.Datos)
			archivo.Close()
			Mensajes = append(Mensajes, aux.Nombre)
		} else {
			fmt.Println("Peticion: ", peticion, " no reconocida")
		}
	}
}

func Peticion(conexion net.Conn, peticion string) {
	err := gob.NewEncoder(conexion).Encode(peticion) //manda peticion de proceso al servidor
	if err != nil {
		fmt.Println(err)
		return
	}
}

func EnviarA(aux Archivo, conexion net.Conn) {
	Peticion(conexion, "A")
	err := gob.NewEncoder(conexion).Encode(aux) //manda el nombre del archivo
	if err != nil {
		fmt.Println("Error al enviar nombre del archivo: ", err)
		return
	}
	fmt.Print("Archivo : ", aux.Nombre, " enviado")
}

func EnviarM(mensaje string, conexion net.Conn) {
	Peticion(conexion, "M")
	err := gob.NewEncoder(conexion).Encode(mensaje) //manda peticion de proceso al servidor
	if err != nil {
		fmt.Println(err)
		return
	}
}

func cliente() {
	conexion, err := net.Dial("tcp", ":9999") //Conecta
	if err != nil {
		fmt.Println(err)
		return
	}

	go Recibir(conexion) //se pone a escuchar
	var opc string = "0"
	for opc != "4" {
		fmt.Print("\n1 -> Enviar mensaje al servidor\n",
			"2 -> Enviar archivo al servidor\n",
			"3 -> Mostrar mensajes/nombre archivos recibidos del servidor\n",
			"4 -> Desconectar\n")
		fmt.Scanln(&opc)
		if opc == "1" {
			fmt.Println("\nEnviar mensaje al servidor\nEscriba el mensaje:")
			reader := bufio.NewReader(os.Stdin)
			mensaje, _ := reader.ReadString('\n')
			EnviarM(mensaje, conexion)
		} else if opc == "2" {
			var indice string
			var archivos []string
			directorio, err := os.ReadDir(".")
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Print("\nEnviar un archivo al servidor:\n",
				"Seleccione el archivo a mandar al servidor:\n")
			for i, archivo := range directorio {
				if !archivo.IsDir() {
					archivos = append(archivos, archivo.Name())
					fmt.Println(i+1, " -> ", archivo.Name())
				}
			}
			fmt.Scanln(&indice)
			i, err := strconv.Atoi(indice)
			if err != nil {
				fmt.Println(err)
				continue
			}
			datos, err2 := os.ReadFile(archivos[i-1])
			if err2 != nil {
				fmt.Println("Error al leer el archivo: ", err2)
				continue
			}
			aux := Archivo{Nombre: archivos[i-1], Datos: datos}
			EnviarA(aux, conexion)
		} else if opc == "3" {
			fmt.Print("\nMensajes recibidos:\n")
			for _, mensaje := range Mensajes {
				fmt.Println(mensaje)
			}
		} else if opc == "4" {
			fmt.Println("\nDesconectando del servidor...")
			Peticion(conexion, "D")
			time.Sleep(time.Millisecond * 500)
		} else {
			fmt.Print("\nOpcion ingresada incorrecta...\nPruebe de nuevo...\n\n")
		}
	}
}

func main() {
	cliente()
}
