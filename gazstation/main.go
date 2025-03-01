package main

import (
	"fmt"
)

type Car struct {
	name      string
	volume    int
	maxVolume int
}

type GazStation struct {
	amountFuel int
}

//func (c *Car) AcceptCar() { // данная функция будет принимать машину
//	fmt.Println("Введи имя машины: ")
//	fmt.Scan(&c.name)
//	fmt.Println("Добро пожаловать на заправочную станцию,", c.name)
//	fmt.Println("Введи текущее количество топлива: ")
//	fmt.Scan(&c.volume)
//	fmt.Println("Введи максимальное значение бака: ")
//	fmt.Scan(&c.maxVolume)
//}

func NewCar(name string, volume int, maxVolume int) *Car { // конструктор для новой машины
	return &Car{
		name:      name,
		volume:    volume,
		maxVolume: maxVolume,
	}
}

func (c *Car) Refuel(g *GazStation) { // обновленная функция для заправки автомобиля
	if c.volume < 0 || c.maxVolume < 0 {
		fmt.Println("Ошибка: введены отрицательные значения.")
		return
	}

	if c.volume == c.maxVolume {
		fmt.Println("Ваша машина уже с полным баком.")
	} else if c.volume < c.maxVolume {
		neededFuel := c.maxVolume - c.volume
		fmt.Printf("Заправляю топлива на %d литров.\n", neededFuel)
		c.volume = c.maxVolume
		fmt.Println("Ваша машина успешно заправлена.")
		g.amountFuel -= neededFuel
		fmt.Printf("Топливо, которое осталось %d литров", g.amountFuel)
	} else {
		fmt.Printf("false\n")
	}
}

func main() {
	fuel := GazStation{amountFuel: 10}
	firstCar := NewCar("Kia", -2, -2)
	secondCar := NewCar("Ford", 2, 3)
	thirdCar := NewCar("Kia", 3, 2)
	fmt.Printf("Объем топлива %d литров\n", fuel)
	firstCar.Refuel(&GazStation{amountFuel: 10})
	secondCar.Refuel(&GazStation{amountFuel: 10})
	fmt.Println("")
	thirdCar.Refuel(&GazStation{amountFuel: 10})
}
