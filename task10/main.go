package main

import "fmt"

// удаление цифры из числа

func main() {
	var n, a, x, raz, newNum int
	fmt.Println("Введите натуральное число ")
	fmt.Scan(&n)
	fmt.Println("Введите число, которое хотите удалить ")
	fmt.Scan(&x)

	newNum = 0
	raz = 1

	for n > 0 {
		a = n % 10
		if x != a {
			newNum = newNum + a*raz
			raz = raz * 10
		}
		n = n / 10
	}
	fmt.Print("Ваше число ", newNum)
}
