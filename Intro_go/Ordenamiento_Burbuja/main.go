package main

import "fmt"

func Burbuja(s []int64) {
	var i, j int
	for i = 0; i < len(s); i++ {
		for j = 1; j < len(s)-i; j++ {
			if s[j] < s[j-1] {
				temp := s[j]
				s[j] = s[j-1]
				s[j-1] = temp
			}
		}
	}
	fmt.Println("Despues de ordenarlo\t", s)
	return
}

func main() {
	var s []int64
	var i int64
	for i = 0; i < 10; i++ {
		if i%2 == 0 {
			s = append(s, 10-i)
		} else {
			s = append(s, 10+i)
		}
	}
	fmt.Println("Antes de la funcion\t", s)
	Burbuja(s)
	fmt.Println("Despues de la funcion\t", s)
	return
}
