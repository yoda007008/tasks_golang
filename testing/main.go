package main

import "fmt"

var Board [10][10]string

type orientation int

const (
	gorizontal = 0
	vertical   = 1
)

func main() {
	var x, y, size int
	var o orientation
	fmt.Println("Введите координаты x и y")
	fmt.Scan(&x, &y)

	for k := 0; k < size; k++ {
		switch o {
		case gorizontal:
			x += k
		case vertical:
			y += k
		}
		for row := x - 1; row < 3; row++ {
			for col := y - 1; col < 3; col++ {
				if x < 0 || x >= 10 || y < 0 || y >= 10 {
					continue
				}
			}
		}
	}
	if Board[x][y] != "." {
		fmt.Print("true")
	} else {
		fmt.Print("false")
	}
}
func test() {
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

// горизонтальное расположение
// if !isValid(3, 4, 3, gorizontal, Board{}) {
// t.Error("Первый тест провален")
// }
// вертикальное расположение
// if !isValid(4, 4, 1, vertical, Board{}) {
// t.Error("Второй тест провален")
// }
// тестируем углы
// if !isValid(0, 0, 1, gorizontal, Board{}) {
// t.Error("Третий тест провален")
// }
// if !isValid(10, 10, 1, gorizontal, Board{}) {
// t.Error("Четвертый тест провален")
// }
// if !isValid(10, 0, 1, gorizontal, Board{}) {
// t.Error("Пятый тест провален")
// }
