package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	board := NewBoard()
	board.PlaceShip(4) // Размещаем 4-палубный корабль
	board.PlaceShip(3) // Размещаем 3-палубный корабль

	fmt.Println("Начальное поле:")
	board.PrintField()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите координаты выстрела (x y): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		coords := strings.Split(input, " ")

		if len(coords) != 2 {
			fmt.Println("Ошибка: введите два числа через пробел.")
			continue
		}

		x, errX := strconv.Atoi(coords[0])
		y, errY := strconv.Atoi(coords[1])

		if errX != nil || errY != nil || x < 0 || x >= 10 || y < 0 || y >= 10 {
			fmt.Println("Ошибка: введите корректные координаты (числа от 0 до 9).")
			continue
		}

		result := board.HandleShoot(x, y)
		fmt.Printf("Выстрел в (%d, %d): %s\n", x, y, result)

		fmt.Println("Поле после стрельбы:")
		board.PrintField()
	}
}
