package main

import "fmt"

func main() {
	var n, temp int64
	var s []int64
	fmt.Scanln(&n)

	for ; n > 0; n-- {
		fmt.Scanln(&temp)
		s = append(s, temp)
	}
	temp = 0
	for e := 0; e < len(s); e++ {
		temp += s[e]
	}
	fmt.Println(temp)
}
