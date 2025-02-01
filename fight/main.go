package main

import "fmt"

func main() {
	var index int
	fmt.Println("Введите число от 0 до 10, чтобы попытать удачу: ")
	fmt.Scan(&index)
	// y := [10]string{"@", "@", "@", "@", "X", "X", "X", "X", "@", "@"}
	/*
		for range y {
			fmt.Println("Введите число от 0 до 10, чтобы попытать удачу: ")
			if y[index] == "X" {
				fmt.Println("Попал!")
			} else if y[index] == "@" {
				fmt.Println("Мимо!")
			}
		}
	*/
}
