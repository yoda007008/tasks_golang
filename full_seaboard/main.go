package main

import (
	"fmt"
)

const (
	boardSize  = 10
	fourShip   = 4
	thirdShip  = 3
	secondShip = 2
	firstShip  = 1
)

type Board [boardSize][boardSize]string

func createField() Board {
	var board Board
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board[j][i] = "."
		}
	}
	return board
}

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

func main() {

}
