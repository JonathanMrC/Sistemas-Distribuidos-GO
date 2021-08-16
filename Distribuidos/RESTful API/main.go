package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var registroM map[string]map[string]float64 //materia {alumnos, calificaciones
var registroA map[string][]string           //alumno:materias

//******************paquete de datos****************************

type PDatos struct {
	Materia      string  `json:"materia"`
	Alumno       string  `json:"alumno"`
	Calificacion float64 `json:"calificacion"`
}

//******************funciones de ayuda**************************
func DevolverAns(res http.ResponseWriter, ans []byte) {
	res.Header().Set(
		"Content-Type",
		"application/json",
	)
	res.Write(ans)
}

func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 2, 32)
}

//******************funciones basicas***************************

func Reg(pd PDatos) []byte {
	jsonData := []byte(`{"code": "agregado"}`)
	if _, existeM := registroM[pd.Materia]; existeM {
		if _, existeA := registroM[pd.Materia][pd.Alumno]; existeA {
			return []byte(`{"code": "El alumno, ya tiene una calificacion asignada para esta materia"}`)
		} else {
			registroA[pd.Alumno] = append(registroA[pd.Alumno], pd.Materia)
			registroM[pd.Materia][pd.Alumno] = pd.Calificacion
		}
	} else {
		registroA[pd.Alumno] = append(registroA[pd.Alumno], pd.Materia)
		temp := make(map[string]float64)
		temp[pd.Alumno] = pd.Calificacion
		registroM[pd.Materia] = temp
	}
	return jsonData
}

func ActualizaRegistro(pd PDatos) []byte {
	if _, existeM := registroM[pd.Materia]; !existeM {
		return []byte(`{"code": "No existe la materia"}`)
	}
	if _, existeA := registroM[pd.Materia][pd.Alumno]; !existeA {
		return []byte(`{"code": "El alumno no esta registrado en la materia"}`)
	}
	registroM[pd.Materia][pd.Alumno] = pd.Calificacion
	return []byte(`{"code": "actualizado"}`)
}

func EliminaDeRegistro(alumno string) []byte {
	if _, existeA := registroA[alumno]; !existeA {
		return []byte(`{"code": "El alumno no esta registrado"}`)
	}
	materias := registroA[alumno]
	for _, materia := range materias {
		delete(registroM[materia], alumno)
	}
	delete(registroA, alumno)
	return []byte(`{"code": "eliminado"}`)
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

func GetAll() ([]byte, error) { // Convierte registroM a un JSON
	jsonData, err := json.MarshalIndent(registroM, "", "    ")
	return jsonData, err
}

//*****************funciones del servidor***********************

func Registro(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method) // imprimimos el método usando del Request
	switch req.Method {
	case "POST": // si el método es POST
		var pd PDatos
		err := json.NewDecoder(req.Body).Decode(&pd)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		DevolverAns(res, Reg(pd))
	case "PUT": // si el método es POST
		var pd PDatos
		err := json.NewDecoder(req.Body).Decode(&pd)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		DevolverAns(res, ActualizaRegistro(pd))
	case "GET":
		reg, err := GetAll() //devolver todo
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		DevolverAns(res, reg)
	}
}

func Registro_(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method) // imprimimos el método usando del Request
	switch req.Method {
	case "DELETE": // si el método es POST
		alumno := req.URL.Path[len("/registro/"):]
		DevolverAns(res, EliminaDeRegistro(alumno))
	}
}

func Promedios(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method) // imprimimos el método usando del Request
	cadena := req.URL.Path[len("/promedio/"):]
	switch req.Method {
	case "GET":
		var ans []byte = []byte(`{"code": "peticion no reconocida, pruebe con: alumno/nombre | materia/nombre | general"}`)
		if strings.HasPrefix(cadena, "alumno/") {
			ans = []byte(`{"code": "El alumno no existe"}`)
			alumno := cadena[len("alumno/"):]
			prom := GetPromA(alumno)
			if prom != -1.0 {
				temp := make(map[string]float64)
				temp["Promedio de "+alumno+": "] = prom
				ansM, err := json.MarshalIndent(temp, "", "    ")
				if err != nil {
					ans = []byte(`{"code": "error al convertir JSON"}`)
				} else {
					ans = ansM
				}
			}
		} else if strings.HasPrefix(cadena, "materia/") {
			ans = []byte(`{"code": "La materia no existe"}`)
			materia := cadena[len("materia/"):]
			prom := GetPromM(materia)
			if prom != -1.0 {
				temp := make(map[string]float64)
				temp["Promedio de "+materia+": "] = prom
				ansM, err := json.MarshalIndent(temp, "", "    ")
				if err != nil {
					ans = []byte(`{"code": "error al convertir JSON"}`)
				} else {
					ans = ansM
				}
			}
		} else if strings.HasPrefix(cadena, "general") {
			temp := make(map[string]float64)
			temp["Promedio general: "] = GetPromG()
			ansM, err := json.MarshalIndent(temp, "", "    ")
			if err != nil {
				ans = []byte(`{"code": "error al convertir JSON"}`)
			} else {
				ans = ansM
			}
		}
		DevolverAns(res, ans)
	}
}

func Alumno(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	switch req.Method {
	case "GET":
		alumno := req.URL.Path[len("/alumno/"):]
		listaMaterias, existeA := registroA[alumno]
		if !existeA {
			DevolverAns(res, []byte(`{"code": "El alumno no esta registrado"}`))
		}
		temp := make(map[string]map[string]float64)
		temp[alumno] = make(map[string]float64)
		for _, materia := range listaMaterias {
			temp[alumno][materia] = registroM[materia][alumno]
		}
		ans, err := json.MarshalIndent(temp, "", "    ")
		if err != nil {
			ans = []byte(`{"code": "error al convertir JSON"}`)
		}
		DevolverAns(res, ans)
	}
}

func main() {
	//inicializar los registros
	registroA = make(map[string][]string)
	registroM = make(map[string]map[string]float64)

	http.HandleFunc("/registro", Registro)   // endpoint para registrar/actualizar/eliminar/obtener alumnos
	http.HandleFunc("/registro/", Registro_) // endpoint para registrar/actualizar/eliminar/obtener alumnos
	http.HandleFunc("/promedio/", Promedios) // endpoint para obtener promedio
	http.HandleFunc("/alumno/", Alumno)      // endpoint para "boleta" de un alumno
	fmt.Println("Corriendo RESTful API...")  // mensaje para empezar a probar
	http.ListenAndServe(":9000", nil)        // arrancamos el servidor en el puerto :9000
}
