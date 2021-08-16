package main

import (
	"fmt"

	"./MultiM"
)

func main() {
	var opc int = 1
	cw := MultiM.ContenidoWeb{Vector: []MultiM.Multimedia{}}
	for opc != 0 {
		fmt.Print("Ingrese una opción para capturar:\n" +
			"1) Imagen\n" +
			"2) Audio\n" +
			"3) Video\n" +
			"4) Mostrar Todo\n" +
			"0) Terminar\n->\t")
		fmt.Scan(&opc)
		if opc == 1 {
			fmt.Println("Capturando datos de Imagen")
			im := MultiM.Imagen{Titulo: "", Formato: "", Canales: ""}
			fmt.Println("Ingrese el titulo:")
			fmt.Scan(&im.Titulo)
			fmt.Println("Ingrese el formato:")
			fmt.Scan(&im.Formato)
			fmt.Println("Ingrese los canales:")
			fmt.Scan(&im.Canales)
			fmt.Print("Datos de imagen capturados correctamente\n\n")
			cw.Vector = append(cw.Vector, &im)
		} else if opc == 2 {
			fmt.Println("Capturando datos de Audio")
			au := MultiM.Audio{Titulo: "", Formato: "", Duracion: 0}
			fmt.Println("Ingrese el titulo:")
			fmt.Scan(&au.Titulo)
			fmt.Println("Ingrese el formato:")
			fmt.Scan(&au.Formato)
			fmt.Println("Ingrese los duración:")
			fmt.Scan(&au.Duracion)
			fmt.Print("Datos de audio capturados correctamente\n\n")
			cw.Vector = append(cw.Vector, &au)
		} else if opc == 3 {
			fmt.Println("Capturando datos de Video")
			vd := MultiM.Video{Titulo: "", Formato: "", Frames: 0}
			fmt.Println("Ingrese el titulo:")
			fmt.Scan(&vd.Titulo)
			fmt.Println("Ingrese el formato:")
			fmt.Scan(&vd.Formato)
			fmt.Println("Ingrese los frames:")
			fmt.Scan(&vd.Frames)
			fmt.Print("Datos de video capturados correctamente\n\n")
			cw.Vector = append(cw.Vector, &vd)
		} else if opc == 4 {
			cw.Mostrar()
		}
	}
	return
}
