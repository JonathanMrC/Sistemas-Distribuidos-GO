package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
)

type Registrar struct {
	Materia string
	Alumno  string
	Calif   float64
}

func getline() string {
	reader := bufio.NewReader(os.Stdin)
	entrada, _ := reader.ReadString('\n')
	entrada = entrada[:len(entrada)-2]
	return entrada
}

func regCalifMateria() {
	conexion, err := rpc.DialHTTP("tcp", ":4040")
	if err != nil {
		fmt.Println("No se realizó la conexión\nError: ", err)
		return
	}
	var alumno, materia string
	var calificacion float64
	fmt.Println("\nAgregar califición de una materia" +
		"\nIngrese el nombre del alumno:")
	alumno = getline()
	fmt.Println("\nIngrese la materia:")
	materia = getline()
	fmt.Println("\nIngrese la calificación:")
	fmt.Scanln(&calificacion)

	var reply bool
	args := Registrar{materia, alumno, calificacion}

	errc := conexion.Call("Estructura.RegistrarCalifAlumno", args, &reply)
	if reply {
		fmt.Println("Calificiación registrada correctamente")
	} else if errc != nil {
		fmt.Println("El alumno ya tiene una calificación asignada")
	}
}

func mostrarPromAlumno() {
	conexion, err := rpc.DialHTTP("tcp", ":4040")
	if err != nil {
		fmt.Println("No se realizó la conexión\nError: ", err)
		return
	}
	var alumno string
	fmt.Println("\nMostrar Promedio de Alumno" +
		"\nIngrese el nombre del alumno:")
	alumno = getline()

	var reply float64

	errc := conexion.Call("Estructura.ObtenerPromDeAlumno", alumno, &reply)
	if errc != nil {
		fmt.Println(errc)
		return
	}
	fmt.Println("\nPromedio: ", reply)
}

func mostrarPromGeneral() {
	conexion, err := rpc.DialHTTP("tcp", ":4040")
	if err != nil {
		fmt.Println("No se realizó la conexión\nError: ", err)
		return
	}
	var reply float64

	errc := conexion.Call("Estructura.ObtenerPromGen", "", &reply)
	if errc != nil {
		fmt.Println(errc)
		return
	}
	fmt.Println("\nPromedio General: ", reply)
}

func mostrarPromMateria() {
	conexion, err := rpc.DialHTTP("tcp", ":4040")
	if err != nil {
		fmt.Println("No se realizó la conexión\nError: ", err)
		return
	}
	var materia string
	fmt.Println("\nMostrar Promedio de Materia" +
		"\nIngrese el nombre de la materia:")
	materia = getline()

	var reply float64

	err2 := conexion.Call("Estructura.ObtenerPromDeMateria", materia, &reply)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("\nPromedio Materia: ", materia, " -> ", reply)
}

func main() {
	var opc string = "5"
	var menu string = "\n1 -> Agregar calificación de una materia" +
		"\n2 -> Mostrar el promedio de un alumno" +
		"\n3 -> Mostrar el promedio general" +
		"\n4 -> Mostrar el promedio de una materia" +
		"\n0 -> Salir"
	for opc != "0" {
		fmt.Println(menu)
		fmt.Scanln(&opc)
		if opc == "1" {
			regCalifMateria()
		} else if opc == "2" {
			mostrarPromAlumno()
		} else if opc == "3" {
			mostrarPromGeneral()
		} else if opc == "4" {
			mostrarPromMateria()
		} else if opc == "0" {
			fmt.Println("Terminando cliente")
		} else {
			fmt.Println("Ingreso una opción incorrecta...\nIntente de nuevo")
		}
	}
}
