package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
)

type Registrar struct {
	Materia string
	Alumno  string
	Calif   float64
}

type Estructura map[string]map[string]float64

var Alumnos map[string][]string //nombre:materias
var Materias Estructura

func (t *Estructura) RegistrarCalifAlumno(args *Registrar, reply *bool) error {
	if _, existe := Materias[args.Materia]; existe {
		if _, existeA := Materias[args.Materia][args.Alumno]; existeA {
			var cadena string = "Error: " + args.Alumno + "\ttiene: " + strconv.FormatFloat(args.Calif, 'f', 2, 32) + "\tde calificación en: " + args.Materia
			return errors.New(cadena)
		} else {
			Alumnos[args.Alumno] = append(Alumnos[args.Alumno], args.Materia)
			Materias[args.Materia][args.Alumno] = args.Calif
		}
	} else {
		Alumnos[args.Alumno] = append(Alumnos[args.Alumno], args.Materia)
		alumnos := make(map[string]float64)
		alumnos[args.Alumno] = args.Calif
		Materias[args.Materia] = alumnos
	}
	*reply = true
	return nil
}

func (t *Estructura) ObtenerPromDeMateria(materia *string, reply *float64) error {
	var suma float64 = 0.0
	var prom float64
	if _, existe := Materias[*materia]; !existe {
		return errors.New("no existe la materia")
	}
	for _, alumno := range Materias[*materia] {
		suma += alumno
	}
	prom = suma / float64(len(Materias[*materia]))
	*reply = prom
	return nil
}

func (t *Estructura) ObtenerPromGen(vacio *string, reply *float64) error {
	var suma float64
	var suma_gen float64 = 0.0
	var prom float64
	for nombre, listaMaterias := range Alumnos {
		suma = 0.0
		for _, materia := range listaMaterias {
			suma += Materias[materia][nombre]
		}
		suma_gen += suma / float64(len(listaMaterias))
	}
	prom = suma_gen / float64(len(Alumnos))
	*reply = prom
	return nil
}

func (t *Estructura) ObtenerPromDeAlumno(alumno *string, reply *float64) error {
	var suma float64 = 0.0
	var prom float64
	var cont float64 = 0.0
	if _, existe := Alumnos[*alumno]; !existe {
		return errors.New("no existe el alumno")
	}
	for _, materia := range Materias {
		if _, existe := materia[*alumno]; existe {
			suma += materia[*alumno]
			cont += 1.0
		}
	}
	prom = suma / cont
	*reply = prom
	return nil
}

func servidor() {
	estructura := new(Estructura)
	err := rpc.Register(estructura)
	if err != nil {
		fmt.Println(err)
		return
	}
	rpc.HandleHTTP()
	escuchar, err := net.Listen("tcp", ":4040")
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

func Respaldar(nombreArchivo string, datos []byte) {
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer archivo.Close() //cierro el archivo al final
	archivo.Write(datos)
}

func main() {
	Alumnos = make(map[string][]string)
	Materias = make(map[string]map[string]float64)
	go servidor()
	var t string
	fmt.Scanln(&t)

	j, err := json.Marshal(Materias)
	if err != nil {
		fmt.Println(err)
	}
	Respaldar("Respaldo.txt", j)
}
