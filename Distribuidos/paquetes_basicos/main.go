package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	nombre string
	edad   uint64
}

type ByName []Persona

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].nombre < a[j].nombre }

type ByEdad []Persona

func (a ByEdad) Len() int           { return len(a) }
func (a ByEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByEdad) Less(i, j int) bool { return a[i].edad < a[j].edad }

func main() {
	personas := []Persona{
		Persona{nombre: "Jonathan", edad: 22},
		Persona{nombre: "Pedro", edad: 23},
		Persona{nombre: "Fernanda", edad: 21}}
	sort.Sort(ByName(personas))
	fmt.Println(personas)
	sort.Sort(ByEdad(personas))
	fmt.Println(personas)
}

/*package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{8, 2, 3, 5}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println("acendente", s)
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println("descendente", s)
}

/*package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l list.List

	l.PushBack(1)
	l.PushBack("bien locochon esta vaina")
	l.PushFront(3.5)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

/*package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println((strings.HasPrefix("distribuidos, apoco no", "dis")))
	fmt.Println((strings.HasSuffix("distribuidos, apoco no", "no")))
}
/*package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("prueba.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer file.Close()           //cierra el archivo al final de todo
	file.WriteString("apoco no") //escribe una cadena en el archivo
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	total := stat.Size()

	bs := make([]byte, total)

	count, err := file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}

	str := string(bs)

	fmt.Println(str, "bytes: ", count)
}
*/
