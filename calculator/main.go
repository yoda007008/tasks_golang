package main

import "fmt"

func main() {
	var x, y int
	var z string
	fmt.Scan(&x, &y, &z)

	if z == "+" {
		fmt.Println(x + y)
	} else if z == "-" {
		fmt.Println(x - y)
	} else if z == "/" {
		fmt.Println(x / y)
	} else if z == "*" {
		fmt.Println(x * y)
	} else {
		fmt.Println("Некоректная операция")
	}
}
