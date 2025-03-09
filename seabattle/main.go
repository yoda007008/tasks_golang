package main

import "fmt"

func main() {
	board := NewBoard()
	board.PlaceShip(4) // Размещаем 4-палубный корабль
	board.PlaceShip(3) // Размещаем 3-палубный корабль

	board.PrintField()

	shoots := []struct {
		x, y int
	}{
		{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, // Промахи
		{5, 5}, {5, 6}, {5, 7}, {5, 8}, // Попадания
	}

	for _, shoot := range shoots {
		result := board.HandleShoot(shoot.x, shoot.y)
		fmt.Printf("Выстрел в (%d, %d): %s\n", shoot.x, shoot.y, result)
	}

	fmt.Println("Поле после стрельбы:")
	board.PrintField()
}
