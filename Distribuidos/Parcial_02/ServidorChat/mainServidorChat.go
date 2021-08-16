package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
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

type Datos struct {
	Usuario string
	Mensaje string
}

var this Paquete
var Puertos []string
var ListaClientes []net.Conn //lista Clientes

func getline() string {
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')
	entrada = entrada[:len(entrada)-2]
	return entrada
}

func EnviarM(datos Datos, emisor net.Conn) {
	for _, cliente := range ListaClientes {
		if cliente != emisor { //excepto al mismo usuario
			err := gob.NewEncoder(cliente).Encode(&datos) //envia los datos a los demas {nombre/mensaje}
			if err != nil {
				fmt.Println("Error al enviar al cliente: ", err)
				return
			}
		}
	}
	fmt.Println("Información: ", datos, " enviada a los usuarios")
}

func Chat(cliente net.Conn) {
	for {
		var datos Datos
		err := gob.NewDecoder(cliente).Decode(&datos) //recibe los datos del usuario{nombre/mensaje}
		if err != nil {
			fmt.Println("Conexion perdida con el cliente: ", cliente, "\n", err)
			return
		}
		if datos.Mensaje == "/salir" {
			this.Info.CantUsuarios--
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
			break
		}
		go EnviarM(datos, cliente) //envia los datos al los demas usuarios {nombre/mensaje}
	}
}

func handleClient(cliente net.Conn) { //controlador conexion
	var peticion string = "VACIO" //la primera peticion indica si es un cliente o el middleware o servidor
	gob.NewDecoder(cliente).Decode(&peticion)
	fmt.Println("Peticion: ", peticion)
	if peticion == "I" { //es el intermediario
		err := gob.NewEncoder(cliente).Encode(&this) //devuelve su información actual
		if err != nil {
			fmt.Println("Error al enviar la información del servidor de chat al middleware", err)
			return
		}
		fmt.Println("Información enviada al intermediario")
		return
	} else if peticion == "S" { //es un test de otro servidor
		return
	} else if peticion == "C" {
		ListaClientes = append(ListaClientes, cliente)
		this.Info.CantUsuarios++
		Chat(cliente)
	}
}

func Ejecutar() {
	var ocupado bool = true
	for _, puerto := range Puertos {
		fmt.Println("Probando puerto: ", puerto)
		aux, err := net.Dial("tcp", puerto) //Conecta
		if err != nil {                     //el puerto esta libre
			this.Dir.Puerto = puerto
			ocupado = false
			Escuchar(puerto)
		} else { //puerto ocupado, intentar con otro puerto en la lista
			gob.NewEncoder(aux).Encode("S") //test
			aux.Close()
		}
	}
	if ocupado {
		fmt.Println("Todos los puertos disponibles estan ocupados")
	}
}

func Escuchar(puerto string) {
	fmt.Println("Ejecutando servidor de chat con tema:", this.Info.Tema, " en puerto", puerto)
	serv, err := net.Listen("tcp", puerto) //Escucha peticiones
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
		go handleClient(conexion) //mueve al cliente aparte
	}
}

func main() {
	Puertos = make([]string, 0)
	Puertos = append(Puertos, ":4041")
	Puertos = append(Puertos, ":4042")
	Puertos = append(Puertos, ":4043")
	this = Paquete{
		Info: Informacion{Tema: "", CantUsuarios: 0, MaxUsuarios: 0},
		Dir:  Direccion{Dir: "", Puerto: Puertos[0]},
	}
	fmt.Println("Ingrese el tema de este servidor de chat:")
	this.Info.Tema = getline()
	fmt.Println("Ingrese la cantidad maxima de usuarios:")
	fmt.Scanln(&this.Info.MaxUsuarios)
	go Ejecutar()
	var t string
	fmt.Scanln(&t)
}
