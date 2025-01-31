package main

import "fmt"

// факториал числа через рекурсию
/*
func main() {
	fmt.Print("Factorial = ", factorial(10))
}

func factorial(f uint) uint {
	if f == 0 {
		return 1
	}
	return f * factorial(f-1)
}
*/

// реализация через цикл

func main() {
	var x int
	fmt.Scan(&x)
	factorial := 1
	for i := 2; i <= x; i++ {
		factorial *= i
	}
	fmt.Println(factorial)
}
