package main

import "fmt"

var Board [10][10]string

func main() {
	// координаты
	for {
		var x, y int
		fmt.Println("Введите координаты x и y и ширину корабля (size) от 1 до 4: ")
		fmt.Scan(&x, &y)

		if x < 0 || x >= 10 || y < 0 || y >= 10 {
			fmt.Println("Координаты не подходят")
			break
		}

		// создание самого поля
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				Board[i][j] = "."
			}
		}

		// создание квадрата 3 на 3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				Board[x+i][y+j] = "X"
			}
		}
		// вывод поля
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				fmt.Printf("%s ", Board[row][col])
			}
			fmt.Println()
		}
	}
}
