package main

import (
	"fmt"
	"time"
)

var listaProcesos []bool
var mostrar bool

func Proceso(id uint64) {
	i := uint64(0)
	for listaProcesos[id] {
		if mostrar {
			fmt.Printf("id %d: %d\n", id, i)
		}
		i++
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	mostrar = false
	var opc int = -1
	var id uint64 = 0
	var menu string = "\nAdministrador" +
		"\n1) Agregar Proceso" +
		"\n2) Mostrar Procesos" +
		"\n3) Terminar Proceso" +
		"\n4) Salir\n->\t"
	for opc != 4 {
		fmt.Print(menu)
		fmt.Scanln(&opc)
		if opc == 1 {
			listaProcesos = append(listaProcesos, true)
			go Proceso(id)
			fmt.Println("Proceso: ", id, " creado")
			id++
		} else if opc == 2 {
			mostrar = true
			fmt.Print("Presione enter para terminar\n")
			var terminar string
			fmt.Scanln(&terminar)
			mostrar = false
		} else if opc == 3 {
			var aux uint64
			fmt.Println("Ingrese el id del proceso a eliminar:")
			fmt.Scanln(&aux)
			if aux < uint64(len(listaProcesos)) {
				if listaProcesos[aux] { //si el proceso no ha terminado
					listaProcesos[aux] = false
					fmt.Println("El proceso: ", aux, " ha terminado")
				}
			} else {
				fmt.Println("No existe ese proceso")
			}
		} else {
			if opc == 4 {
				for id--; id != 0; id-- {
					listaProcesos[id] = false
				}
			} else {
				fmt.Println("Ingreso una opcion inexistente\nIntente de nuevo")
			}
		}
	}
	fmt.Println("Programa terminado")
}
