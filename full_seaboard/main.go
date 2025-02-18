package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	boardSize  = 10
	fourShip   = 4
	thirdShip  = 3
	secondShip = 2
	firstShip  = 1
)

var errPlaceShip = errors.New("Не удалось разместить корабль")
var errCord = errors.New("Ошибка координат")

type orientation int

const (
	gorizontal = 0
	vertical   = 1
)

type Board [boardSize][boardSize]string

// создание поля
func createField() *Board {
	var board Board
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board[i][j] = "."
		}
	}
	return &board
}

// вывод поля
func printField(board Board) {
	fmt.Println("  1 2 3 4 5 6 7 8 9 10")
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
			x += k
		case vertical:
			y += k
		}
		if size < 0 || size > 4 {
			break
		}
		if x > boardSize-1 || y > boardSize-1 {
			return false
		}
		for row := x - 1; row < x-1+3; row++ {
			for col := y - 1; col < y-1+3; col++ {
				if row <= 0 || row > boardSize-1 || col <= 0 || col > boardSize-1 {
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
func PlaceShip(board *Board, size int) bool {
	o := orientation(rand.Intn(2))
	shipPlaced := false
	for !shipPlaced {
		x := rand.Intn(boardSize)
		y := rand.Intn(boardSize)
		//x, y := 6, 8
		if !isValid(x, y, size, o, *board) {
			continue
		}
		// todo x, y проверить все точки корабля, попадают они или нет (функция isValid)
		// конец игры нужно реализовать через проход по двумерному массиву
		for i := 0; i <= size; i++ {
			if size < 0 || size > 4 {
				break
			}
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
	return shipPlaced
}

// функция благодаря которой мы производим выстрел на поле, при этом проверяя его статус
func makeShoot(board *Board, x, y int) string {
	if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
		fmt.Errorf("%w", errCord)
	}
	cell := &board[x][y]

	if *cell == "." {
		*cell = "M"
		return "Промах!"
	}
	if *cell == "S" {
		*cell = "X"
		return "Попал!"
	}
	if *cell == "M" || *cell == "X" {
		return "Ты уже сюда стрелял"
	}
	return "Ошибка координат"
}

// главная функция где и запускается программа
func main() {
	rand.Seed(time.Now().UnixNano())
	ships := []int{fourShip, thirdShip, thirdShip, secondShip, secondShip, secondShip, firstShip, firstShip, firstShip, firstShip}

	board := createField()

	for _, shipSize := range ships {
		if !PlaceShip(board, shipSize) {
			fmt.Errorf("%w", errPlaceShip)
		}
	}

	for {
		var x, y int
		printField(*board)
		fmt.Println("Добро пожаловать в морской бой!")
		fmt.Println("Введите координату x: ")
		fmt.Scan(&x)
		fmt.Println("Введите координату y: ")
		fmt.Scan(&y)

		result := makeShoot(board, x, y)
		fmt.Println(result)

	}
}
