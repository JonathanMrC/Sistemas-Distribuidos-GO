package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var registroM map[string]map[string]float64 //materia {alumnos, calificaciones
var registroA map[string][]string           //alumno:materias

func Devolverhtml(res http.ResponseWriter, namehtml string, arg string) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	if arg != "" {
		fmt.Fprintf(
			res,
			cargarHtml(namehtml),
			arg,
		)
	} else {
		fmt.Fprintf(
			res,
			cargarHtml(namehtml),
		)
	}
}

func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 2, 32)
}

func Reg(alumno string, materia string, calif float64) bool {
	if _, existeM := registroM[materia]; existeM {
		if _, existeA := registroM[materia][alumno]; existeA {
			return false
		} else {
			registroA[alumno] = append(registroA[alumno], materia)
			registroM[materia][alumno] = calif
		}
	} else {
		registroA[alumno] = append(registroA[alumno], materia)
		temp := make(map[string]float64)
		temp[alumno] = calif
		registroM[materia] = temp
	}
	return true
}

func GetPromA(alumno string) float64 {
	if materiasA, existeA := registroA[alumno]; existeA {
		var ans float64 = 0.0
		for _, materia := range materiasA {
			ans += registroM[materia][alumno]
		}
		return ans / float64(len(materiasA))
	}
	return -1.0
}

func GetPromG() float64 {
	var ans float64 = 0.0
	for alumno := range registroA {
		ans += GetPromA(alumno)
	}
	return ans / float64(len(registroA))
}

func GetPromM(materia string) float64 {
	if listaA, existeM := registroM[materia]; existeM {
		var ans float64 = 0.0
		for _, calif := range listaA {
			ans += calif
		}
		return ans / float64(len(listaA))
	}
	return -1.0
}

func cargarHtml(a string) string {
	html, _ := ioutil.ReadFile(a)
	return string(html)
}

func pM(res http.ResponseWriter, req *http.Request) { //promedio de una materia
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		materia := strings.ToLower(req.FormValue("materia"))
		var namehtml string = "error.html"
		var cadena string
		if materia == "" {
			cadena = "Debe rellenar todos los campos"
		} else if prom := GetPromM(materia); prom == -1.0 {
			cadena = "La materia ingresada no esta registrada"
		} else {
			namehtml = "anspromMat.html"
			cadena = "Materia: " + materia + " tiene un promedio de: " + FloatToString(prom)
		}
		Devolverhtml(res, namehtml, cadena)
	case "GET":
		Devolverhtml(res, "promMateria.html", "")
	}
}

func pA(res http.ResponseWriter, req *http.Request) { //promedio de un alumno
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		alumno := strings.ToLower(req.FormValue("alumno"))
		var namehtml string = "error.html"
		var cadena string
		if alumno == "" {
			cadena = "Debe rellenar todos los campos"
		} else if prom := GetPromA(alumno); prom == -1.0 {
			cadena = "El alumno: " + alumno + " no esta registrado"
		} else {
			namehtml = "anspromAl.html"
			cadena = ": " + alumno + " tiene un promedio de: " + FloatToString(prom)
		}
		Devolverhtml(res, namehtml, cadena)
	case "GET":
		Devolverhtml(res, "promAlumno.html", "")
	}
}

func pG(res http.ResponseWriter, req *http.Request) { //promedio general
	switch req.Method {
	case "GET":
		Devolverhtml(res, "anspromGen.html", FloatToString(GetPromG()))
	}
}

func RegistrosM(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if len(registroA) == 0 {
			Devolverhtml(res, "tabla.html", "<tr><td>----</td><td>----</td><td>----</td></tr>")
		} else {
			var html string
			for nombre, listaA := range registroM {
				for alumno, calif := range listaA {
					html += "<tr>" +
						"<td>" + nombre + "</td>" +
						"<td>" + alumno + "</td>" +
						"<td>" + FloatToString(calif) + "</td>" +
						"</tr>"
				}
			}
			Devolverhtml(res, "tabla.html", html)
		}
	}
}

func regA(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		alumno := strings.ToLower(req.FormValue("alumno"))
		materia := strings.ToLower(req.FormValue("materia"))
		calif, err := strconv.ParseFloat(req.FormValue("calif"), 64)
		var namehtml string = "error.html"
		var cadena string
		if (alumno == "") || (materia == "") || err != nil || (req.FormValue("calif") == "") {
			cadena = "No se pudo registrar al alumno, debe rellenar todos los campos, la calificion debe ser un numero"
		} else if !Reg(alumno, materia, calif) {
			cadena = "El alumno: " + alumno + " ya tiene una calificion de: " + FloatToString(registroM[materia][alumno])
		} else {
			namehtml = "ansregAl.html"
			cadena = alumno
		}
		Devolverhtml(res, namehtml, cadena)
	case "GET":
		Devolverhtml(res, "regAlumno.html", "")
	}
}

func main() {
	//inicializar los registros
	registroA = make(map[string][]string)
	registroM = make(map[string]map[string]float64)

	//servidor
	http.HandleFunc("/promGeneral", pG) //dir:funcion
	http.HandleFunc("/promAlumno", pA)  //dir:funcion
	http.HandleFunc("/promMateria", pM) //dir:funcion

	http.HandleFunc("/regAlumno", regA) //dir:funcion

	http.HandleFunc("/registro", RegistrosM) //dir:funcion

	fmt.Println("Arrancando servidor...")
	http.ListenAndServe(":9000", nil)

}
