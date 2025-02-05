package main

import (
	"fmt"
	"math/rand"
)

const (
	boardSize  = 10
	fourShip   = 4
	thirdShip  = 3
	secondShip = 2
	firstShip  = 1
)

type Board [boardSize][boardSize]string

// создание поля
func createField() Board {
	var board Board
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board[i][j] = "."
		}
	}
	return board
}

// вывод поля
func printField(board Board) {
	fmt.Println(" A B C D E F G H I J")
	for i, row := range board {
		fmt.Printf("%2d ", i+1)
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
}

// генерация координат и проверка получится ли поставить корабль на поле
func placeShip(board *Board, size int) bool {
	orientation := rand.Intn(2)

	for attempts := 0; attempts < 100; attempts++ {
		x := rand.Intn(boardSize - (size-1)*(1-orientation))
		y := rand.Intn(boardSize - (size-1)*orientation)

		valid := true
		if orientation == 0 {
			for i := 0; i < size; i++ {
				if board[y][x+i] != "." {
					valid = false
					break
				}
			}
		} else {
			for i := 0; i < size; i++ {
				if board[y+i][x] != "." {
					valid = false
					break
				}
			}
		}

		if valid {
			if orientation == 0 {
				for i := 0; i < size; i++ {
					board[y][x+i] = "S"
				}
			} else {
				for i := 0; i < size; i++ {
					board[y+i][x] = "S"
				}
			}
			return true
		}
	}
	return false
}

// функция благодаря которой мы производим выстрел на поле, при этом проверяя его статус
func makeShoot(board *Board, x, y int) string {
	if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
		return "Ошибка координат"
	}
	cell := &board[x][y]

	if *cell == "." {
		*cell = "M"
		return "Промах!"
	} else if *cell == "S" {
		*cell = "X"
		return "Попал!"
	} else if *cell == "M" || *cell == "X" {
		return "Ты уже сюда стрелял"
	}
	return "Ошибка, попробуй снова!"
}

func main() {

}
