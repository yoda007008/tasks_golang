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
	gazVolume  int
}

func NewCar(name string, volume int, maxVolume int) *Car { // конструктор для новой машины
	if volume < 0 || maxVolume < 0 {
		fmt.Println("Ошибка: введены отрицательные значения.")
	}

	return &Car{
		name:      name,
		volume:    volume,
		maxVolume: maxVolume,
	}
}

func (c *Car) Refuel(amount int) { // добавляет топливо в автомобиль
	if c.volume+amount > c.maxVolume {
		c.volume = c.maxVolume // обновление до максимального объема, если сумма amount + c.volume превышает макс. объем
		fmt.Println("Бак заполнен")
	} else {
		c.volume += amount
	}
}

func (g *GazStation) GetFuel(amount int) bool { // проверяет достаточно ли топлива на заправке для автомобиля
	if g.amountFuel < amount {
		return false
	}
	g.amountFuel -= amount
	return true
}

func main() {
	station := &GazStation{amountFuel: 100}
	car := NewCar("Kia", 30, 50)
	amount := 25
	if station.GetFuel(amount) {
		car.Refuel(amount)
		fmt.Printf("Автомобилю %s удалось заправиться\n", car.name)
	} else {
		fmt.Println("На заправке недостаточно топлива")
	}
	fmt.Printf("Топливо в автомобиле %s: %d/%d\n", car.name, car.volume, car.maxVolume)
	fmt.Printf("Топливо на заправке: %d", station.amountFuel)
}
