package main

import "fmt"

// числа Фибоначчи

func main() {
	var n, a, b int
	fmt.Scan(&n)

	a = 0
	fmt.Print(a, " ")
	b = 1
	fmt.Print(b, " ")

	for i := 3; i <= n; i++ {
		fmt.Print(a+b, " ")
		c := b
		b = a + b
		a = c
	}
}
