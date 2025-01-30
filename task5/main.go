package main

import "fmt"

// факториал числа через рекурсию
func main() {
	fmt.Print("Factorial = ", factorial(10))
}

func factorial(f uint) uint {
	if f == 0 {
		return 1
	}
	return f * factorial(f-1)
}
