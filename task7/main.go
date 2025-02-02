package main

import "fmt"

// найти сумму четных чисел числа

func main() {
	var n int

	fmt.Scan(&n)

	sum := 0
	// найти сумму четных цифр числа
	for n > 0 {
		if n%10%2 == 0 {
			sum += n % 10
		}
		n /= 10
	}
	fmt.Print(sum)
}
