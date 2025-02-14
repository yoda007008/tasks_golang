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

type orientation int

const (
	gorizontal = 0
	vertical   = 1
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
		fmt.Printf("%d ", i+1)
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
}

func isValid(x int, y int, size int, orientation orientation, board Board) bool {
	for k := 0; k < size; k++ {
		switch orientation {
		case gorizontal:
			x++
		case vertical:
			y++
		}
		for row := x - 1; row < 3; row++ {
			for col := y - 1; col < 3; col++ {
				if x <= 0 || x >= boardSize || y <= 0 || y >= boardSize {
					continue
				}
				if board[row][col] != "." {
					return false
				}
			}
		}
	}
	return true
}

// генерация координат и проверка получится ли поставить корабль на поле
func placeShip(board Board, size int) bool {
	o := orientation(rand.Intn(2))
	shipPlaced := false
	for !shipPlaced {
		x := rand.Intn(boardSize)
		y := rand.Intn(boardSize)
		if !isValid(x, y, size, o, board) {
			continue
		}
		// todo x, y проверить все точки корабля, попадают они или нет (функция isValid)
		// конец игры нужно реализовать через проход по двумерному массиву
		for i := 0; i < size; i++ {
			board[y][x] = "S"
			switch o {
			case gorizontal:
				x++
			case vertical:
				y++
			}
		}
		shipPlaced = true
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

// главная функция где и запускается программа
func main() {
	//var x, y int
	//fmt.Println("Добро пожаловать в морской бой!")
	//fmt.Println("Введите координату x: ")
	//fmt.Scan(&x)
	//fmt.Println("Введите координату y: ")
	//fmt.Scan(&y)
	//
	//rand.Seed(time.Now().UnixNano())
	//ships := []int{fourShip, thirdShip, thirdShip, secondShip, secondShip, secondShip, firstShip, firstShip, firstShip, firstShip}
	//board := createField()
	//
	//for _, shipSize := range ships {
	//	placeShip(&board, shipSize)
	//}
	//
	//result := makeShoot(&board, x, y)
	//fmt.Println(result)
	//printField(board)
	board := createField()
	placeShip(board, 2)
}
