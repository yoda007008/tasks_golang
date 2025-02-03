package main

import "fmt"

// количество четных и нечетных чисел

func main() {
	var a int
	fmt.Scan(&a)

	even := 0
	uneven := 0

	for a > 0 {
		if a%2 == 0 {
			even += 1
		} else {
			uneven += 1
		}
		a /= 10
	}
	fmt.Println("Четные числа:", even)
	fmt.Println("Нечетные числа:", uneven)
}
