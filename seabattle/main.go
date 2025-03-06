package main

import "fmt"

func main() {
	field := BoardImpl{}
	field.PlaceShip(1)
	err := field.PlaceShip(3)
	if err != nil {
		fmt.Println("Ошибка размещения корабля", err)
	}
	err = field.PlaceShip(4)
	if err != nil {
		fmt.Println("Ошибка размещения корабля", err)
	}
	err = field.PlaceShip(3)
	if err != nil {
		fmt.Println("Ошибка размещения корабля", err)
	}
	err = field.PlaceShip(2)
	if err != nil {
		fmt.Println("Ошибка размещения корабля", err)
	}
	err = field.PlaceShip(2)
	if err != nil {
		fmt.Println("Ошибка размещения корабля", err)
	}
	err = field.PlaceShip(1)
	if err != nil {
		fmt.Println("Ошибка размещения корабля", err)
	}
	err = field.PlaceShip(1)
	if err != nil {
		fmt.Println("Ошибка размещения корабля", err)
	}
	err = field.PlaceShip(1)
	if err != nil {
		fmt.Println("Ошибка размещения корабля", err)
	}
	err = field.PlaceShip(1)
	if err != nil {
		fmt.Println("Ошибка размещения корабля", err)
	}

	fmt.Println("Игровое поле")
	field.PrintField()
}
