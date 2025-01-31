package main

// возведение числа в степень

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	for y >= 1 {
		result := x * y
		fmt.Println(result)
		break
	}
}
