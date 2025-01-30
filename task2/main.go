package main

import (
	"fmt"
)

// кубы чисел в диапазоне a и b

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c := b + 1
	for i := range c {
		if i >= a {
			fmt.Println(i, "cube =", i*i*i)
		}
	}

}
