package main

import "fmt"

func greater(args ...int) int {
	var ans int = args[0]
	for _, elemento := range args {
		if ans < elemento {
			ans = elemento
		}
	}
	return ans
}

func main() {
	fmt.Println("Elemento mayor en : ", 10, 1, 2, 3, 4,
		" -> ", greater(10, 1, 2, 3, 4))
	fmt.Println("Elemento mayor en : ", 2, 3, 124, 123, 13,
		" -> ", greater(2, 3, 124, 123, 13))
	return
}
