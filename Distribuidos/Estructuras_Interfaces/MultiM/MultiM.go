package MultiM

import "fmt"

type ContenidoWeb struct {
	Vector []Multimedia
}

func (obj *ContenidoWeb) Mostrar() {
	if len(obj.Vector) == 0 {
		fmt.Println("No hay elementos")
		return
	}
	for _, elemento := range obj.Vector {
		elemento.Mostrar()
	}
}

type Multimedia interface {
	Mostrar()
}

type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

type Audio struct {
	Titulo   string
	Formato  string
	Duracion int //segundos
}

type Video struct {
	Titulo  string
	Formato string
	Frames  int
}

func (obj *Imagen) Mostrar() {
	fmt.Println("Título:\t" + obj.Titulo +
		"\nFormato:\t" + obj.Formato +
		"\nCanales:\t" + obj.Canales)
}
func (obj *Audio) Mostrar() {
	fmt.Println("Titulo:\t"+obj.Titulo+
		"\nFormato:\t"+obj.Formato+
		"\nDuración:\t", +obj.Duracion, " segundos")
}
func (obj *Video) Mostrar() {
	fmt.Println("Titulo:\t"+obj.Titulo+
		"\nFormato:\t"+obj.Formato+
		"\nFrames:\t", obj.Frames)
}
