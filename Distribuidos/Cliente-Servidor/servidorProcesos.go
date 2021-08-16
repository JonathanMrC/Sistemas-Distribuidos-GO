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

var listaProcesos []Proceso

func Ejecutar(id uint64) {
	for listaProcesos[id].Activo {
		fmt.Printf("id %d: %d\n", id, listaProcesos[id].Cont)
		listaProcesos[id].Cont++
		time.Sleep(time.Millisecond * 500)
	}
}

func buscarProceso() uint64 {
	var i uint64 = 0
	var cant uint64 = uint64(len(listaProcesos))
	for ; i < cant; i++ {
		if listaProcesos[i].Activo {
			return i
		}
	}
	return cant + 1
}

func servidor() {
	for i := uint64(0); i < 5; i++ { //crea los procesos
		p := Proceso{Id: i, Cont: 0, Activo: true}
		listaProcesos = append(listaProcesos, p)
		go Ejecutar(i)
	}

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
		go handleClient(conexion) //mueve al cliente aparte
	}
}

func handleClient(conexion net.Conn) { //controlador conexion
	var peticion string
	defer conexion.Close()
	err := gob.NewDecoder(conexion).Decode(&peticion) //recibe la peticion
	if err != nil {
		fmt.Println(err)
		return
	}
	if peticion == "0" { //peticion de un proceso
		fmt.Printf("\nPeticion de proceso...")
		id := buscarProceso()
		if id >= uint64(len(listaProcesos)) { //valida si hay procesos
			fmt.Println("\nNo hay mas procesos que asignar")
			return
		}
		listaProcesos[id].Activo = false                         //detiene la ejecucion del proceso
		err = gob.NewEncoder(conexion).Encode(listaProcesos[id]) //codifica y envia el proceso
		fmt.Println("\nProceso: ", id, "delegado")
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if peticion == "1" { //devolucion de proceso
		fmt.Printf("\nPeticion para devolver un proceso...")
		var temp Proceso
		err = gob.NewDecoder(conexion).Decode(&temp) //decodifica el proceso en temp
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("\nProceso: ", temp.Id, "devuelto")
		listaProcesos[temp.Id] = temp        //guardo el proceso
		listaProcesos[temp.Id].Activo = true //activo el proceso
		go Ejecutar(temp.Id)                 //ejecuto el proceso
	} else { //basura
		fmt.Println("\nPeticion no reconocida")
		return
	}
}

func main() {
	go servidor()
	var input string
	fmt.Scanln(&input)

}
