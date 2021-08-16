package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"
)

type Archivo struct {
	Nombre string
	Datos  []byte
}

var MAEnviados []string      //Mensajes/Archivos
var ListaClientes []net.Conn //lista Clientes

func EnviarM(mensaje string, emisor net.Conn) { //envia mensaje a los clientes
	MAEnviados = append(MAEnviados, mensaje)
	var bandera string = "M"
	for _, cliente := range ListaClientes {
		err := gob.NewEncoder(cliente).Encode(&bandera) //prepara al cliente para mensaje
		if err != nil {
			fmt.Println(err)
			return
		}
		err = gob.NewEncoder(cliente).Encode(&mensaje) //envia mensaje
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func EnviarA(aux Archivo, emisor net.Conn) { //envia archivos
	MAEnviados = append(MAEnviados, aux.Nombre)
	for _, cliente := range ListaClientes {
		err := gob.NewEncoder(cliente).Encode("A") //manda peticion de proceso al servidor
		if err != nil {
			fmt.Println(err)
			return
		}
		err2 := gob.NewEncoder(cliente).Encode(&aux) //envia el archivo
		if err2 != nil {
			fmt.Println(err2)
			return
		}
	}
}

func handleClient(cliente net.Conn) { //controlador conexion
	var peticion string
	for {
		err := gob.NewDecoder(cliente).Decode(&peticion) //recibe la peticion
		if err != nil {
			fmt.Println(err)
			return
		}
		if peticion == "M" { //recibir un mensaje
			var mensaje string
			err := gob.NewDecoder(cliente).Decode(&mensaje) //recibe el mensaje
			if err != nil {
				fmt.Println(err)
				return
			}
			EnviarM(mensaje, cliente) //envia el mensaje a los demas clientes
		} else if peticion == "A" { //va a recibir un archivo
			var aux Archivo
			err := gob.NewDecoder(cliente).Decode(&aux) //recibe el archivo
			if err != nil {
				fmt.Println(err)
				return
			}
			EnviarA(aux, cliente)
		} else if peticion == "D" { //desconexion
			var pos int64 = 0
			for _, candidato := range ListaClientes {
				if candidato == cliente {
					break
				}
				pos++
			}
			if pos < int64(len(ListaClientes)) {
				ListaClientes = append(ListaClientes[:pos], ListaClientes[pos+1:]...)
			} else {
				ListaClientes = append(ListaClientes[:pos], nil)
			}
			cliente.Close()
			return
		} else {
			fmt.Println("Petición: ", peticion, " no reconocida")
		}
	}
}

func Escuchar() {
	serv, err := net.Listen("tcp", ":9999") //Escucha peticiones
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
		ListaClientes = append(ListaClientes, conexion)
		go handleClient(conexion) //mueve al cliente aparte
	}
}

func Respaldar(nombreArchivo string, datos []string) bool {
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer archivo.Close() //cierro el archivo al final

	for i := 0; i < len(datos); i++ {
		archivo.WriteString(datos[i] + "\n")
	}
	return true
}

func servidor() {
	go Escuchar()
	var opc string = "0"
	for opc != "3" {
		fmt.Println("\n Servidor\n",
			"1 -> Mostrar mensajes / nombre de archivos enviados\n",
			"2 -> Respaldar mensajes / nombre de archivos enviados\n",
			"3 -> Terminar servidor")
		fmt.Scanln(&opc)
		if opc == "1" {
			fmt.Print("\nMensajes / Archivos enviados\n")
			if len(MAEnviados) != 0 {
				for _, mensaje := range MAEnviados {
					fmt.Print("-> ", mensaje, "\n")
				}
			} else {
				fmt.Println("\nNo hay mensajes por mostrar")
			}
		} else if opc == "2" {
			fmt.Println("Respaldando...")
			if Respaldar("respaldoServidor.txt", MAEnviados) {
				fmt.Println("Respaldo creado exitosamente")
			} else {
				fmt.Println("No se pudo realizar el respaldo")
			}
		} else if opc == "3" {
			fmt.Println("Terminando el servidor...")
			time.Sleep(time.Millisecond * 500)
		} else {
			fmt.Print("\nOpcion ingresada incorrecta...\nPruebe de nuevo...\n\n")
		}
	}
}

func main() {
	servidor()
}
