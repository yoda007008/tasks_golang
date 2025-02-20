package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MatrixSize = 7

var Matrix [MatrixSize][MatrixSize]int
var rowSum [MatrixSize]int
var numMaxRow = 0

func main() {
	rand.Seed(time.Now().UnixNano())
	for rowNum := 0; rowNum < MatrixSize; rowNum++ {
		for colNum := 0; colNum < MatrixSize; colNum++ {
			Matrix[rowNum][colNum] = rand.Intn(91) + 10
		}
	}
	for rowNum := 0; rowNum < MatrixSize; rowNum++ {
		fmt.Println(" | ")
		for colNum := 0; colNum < MatrixSize; colNum++ {
			fmt.Printf("%3d ", Matrix[rowNum][colNum])
		}
	}

	// todo подсчет суммы отдельно в каждой строке
	for rowNum := 0; rowNum < MatrixSize; rowNum++ {
		for colNum := 0; colNum < MatrixSize; colNum++ {
			rowSum[rowNum] += Matrix[rowNum][colNum]
		}
	}
	fmt.Println("Сумма по строкам: ")
	for i := 0; i < MatrixSize; i++ {
		fmt.Println(i+1, "-я строка", rowSum[i])
	}

	// todo узнаем какая строка дает наибольшую сумму
	for i := 1; i < MatrixSize; i++ {
		if rowSum[i] > rowSum[numMaxRow] {
			numMaxRow = i
		}
	}
	fmt.Println("Строка, сумма элементов которых максимальна: ", numMaxRow)

}
