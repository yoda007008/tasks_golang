package main

import "fmt"

const (
	rowCount = 3
	colCount = 3
)

var Matrix [rowCount][colCount]int

// матрица, заполненная целыми числами
func main() {
	for row := 0; row < rowCount; row++ {
		for col := 0; col < colCount; col++ {
			Matrix[row][col] = row + col
			fmt.Print(Matrix[row][col], " ")
		}
		fmt.Println("")
	}
}
