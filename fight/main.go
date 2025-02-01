package main

import (
	"fmt"
)

const (
	boardSize = 10
	shipSize  = 4
)

func SeaBattle() {
	my_board := [boardSize]string{"S", "S", "S", "S", "/", "/", "/", "/", "/", "/"}

	for {
		var cord int
		fmt.Print("Введите координату от 1 до 10: ")
		fmt.Scan(&cord)

		cord--
		if cord < 0 || cord >= boardSize {
			fmt.Println("Некорректные координаты для корабля")
			continue
		}

		if my_board[cord] == "S" {
			my_board[cord] = "X"
			fmt.Println("Попал!")
		}

		shipDown := true

		for _, cell := range my_board {
			if cell == "S" {
				shipDown = false
				break
			}
		}

		if shipDown {
			fmt.Println("Выигрыш!")
			for _, cell := range my_board {
				fmt.Printf("%s", cell)
			}
			fmt.Println()
			break
		} else if my_board[cord] == "/" {
			my_board[cord] = "@"
			fmt.Println("Промах, попробуй еще раз")
		}

	}
}

func main() {
	SeaBattle()
}
