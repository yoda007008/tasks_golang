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
		fmt.Errorf("Недостаточно топлива на заправке")
	}
	g.amountFuel -= amount
	fmt.Println("Топлива на заправке достаточно")
	return true
}

func CarRefueling(station *GazStation, car *Car, amount int) { // отдельный метод для логики заправки
	if car.volume < amount {
		res := amount - car.volume
		fmt.Printf("Помещается только %d литров, остальные %d литров лишние", car.volume, res)
		car.volume = amount
	}
	if station.GetFuel(amount) {
		car.Refuel(amount)
		fmt.Printf("Автомобилю %s удалось заправиться\n", car.name)
	} else {
		fmt.Println("На заправке недостаточно топлива")
	}
	fmt.Printf("Топливо в автомобиле %s: %d/%d\n", car.name, car.volume, car.maxVolume)
	fmt.Printf("Топливо на заправке: %d", station.amountFuel)
}

func main() {
	station := &GazStation{amountFuel: 100}
	car := NewCar("Kia", 40, 50)
	amount := 25
	CarRefueling(station, car, amount)
}
