package main

import (
	"fmt"
	"math"
)

// вывод квадрата натуральных чисел начиная с единицы

func main() {
	var x, y, n float64
	fmt.Scan(&n)

	for y < n {
		if y > n {
			break
		} else if y < n {
			fmt.Print(y, " ")
			x += 1
			y = math.Pow(x, 2)
		}

	}
}
