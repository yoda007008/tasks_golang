package main

import "fmt"

// сумма первой и последней цифры числа

func main() {
	var a int
	fmt.Scan(&a)
	b := a % 10 // последнее число
	for a >= 10 {
		a /= 10 // делим до тех пор, пока не дойдем до первой цифры
	}
	fmt.Print(a + b)
}
