package main

import "fmt"

const MatrixSize = 3

var Matrix [MatrixSize][MatrixSize]int

// todo главная функция main()
func main() {
	fmt.Println("Заполните матрицу", MatrixSize, "x", MatrixSize, "числами построчно")

	for rowNum := 0; rowNum < MatrixSize; rowNum++ {
		for colNum := 0; colNum < MatrixSize; colNum++ {
			fmt.Print(rowNum+1, "-я строка ", colNum+1, "-й столбец: ")
			fmt.Scan(&Matrix[rowNum][colNum])
		}
	}

	// todo Вывод матрицы на экран + подсчет суммы
	fmt.Println("Введенная матрица:")
	rowSum := [MatrixSize]int{}
	colSum := [MatrixSize]int{}

	for rowNum := 0; rowNum < MatrixSize; rowNum++ {
		fmt.Print(" | ")
		for colNum := 0; colNum < MatrixSize; colNum++ {
			fmt.Printf("%4d  ", Matrix[rowNum][colNum])
			rowSum[rowNum] += Matrix[rowNum][colNum]
			colSum[colNum] += Matrix[rowNum][colNum]
		}
		fmt.Println("|")
	}

	fmt.Println("Сумма по строкам:")
	for i := 0; i < MatrixSize; i++ {
		fmt.Printf("%d-я строка: %d\n", i+1, rowSum[i])
	}

	fmt.Println("Сумма по столбцам!")
	for i := 0; i < MatrixSize; i++ {
		fmt.Printf("%d-й столбец: %d\n", i+1, colSum[i])
	}
}
