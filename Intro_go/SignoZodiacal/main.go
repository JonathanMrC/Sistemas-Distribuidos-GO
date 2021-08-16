package main

import "fmt"

func main() {
	var dia, mes int64
	var signo string = "Ingreso una fecha erronea\nPruebe de nuevo"
	fmt.Scanln(&dia)
	fmt.Scanln(&mes)

	if mes >= 1 && mes <= 12 && dia > 0 {
		switch mes {
		case 1:
			if dia <= 19 {
				signo = "Capricornio"
			} else if dia <= 31 {
				signo = "Acuario"
			}
		case 2:
			if dia <= 18 {
				signo = "Acuario"
			} else if dia <= 29 {
				signo = "Piscis"
			}
		case 3:
			if dia <= 20 {
				signo = "Piscis"
			} else if dia <= 31 {
				signo = "Aries"
			}
		case 4:
			if dia <= 20 {
				signo = "Aries"
			} else if dia <= 30 {
				signo = "Tauro"
			}
		case 5:
			if dia <= 20 {
				signo = "Tauro"
			} else if dia <= 31 {
				signo = "Geminis"
			}
		case 6:
			if dia <= 20 {
				signo = "Geminis"
			} else if dia <= 30 {
				signo = "Cancer"
			}
		case 7:
			if dia <= 20 {
				signo = "Cancer"
			} else if dia <= 31 {
				signo = "Leo"
			}
		case 8:
			if dia <= 21 {
				signo = "Leo"
			} else if dia <= 31 {
				signo = "Virgo"
			}
		case 9:
			if dia <= 22 {
				signo = "Virgo"
			} else if dia <= 30 {
				signo = "Libra"
			}
		case 10:
			if dia <= 22 {
				signo = "Libra"
			} else if dia <= 31 {
				signo = "Escorpio"
			}
		case 11:
			if dia <= 22 {
				signo = "Escorpio"
			} else if dia <= 30 {
				signo = "Sagitario"
			}
		case 12:
			if dia <= 22 {
				signo = "Sagitario"
			} else if dia <= 31 {
				signo = "Capricornio"
			}
		}
	}
	fmt.Println(signo)
}
