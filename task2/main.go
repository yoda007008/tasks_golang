package main

import (
	"fmt"
)

// кубы чисел в диапазоне a и b

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a < b {
		for i := 0; i <= b; i++ {
			fmt.Println(i, "cube:", i*i*i)
		}
	} else {
		for i := a; i >= 1; i-- {
			fmt.Println(i, "cube:", i*i*i)
		}
	}
}
