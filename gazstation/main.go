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
}

func (c *Car) AcceptCar() { // данная функция будет принимать машину
	fmt.Println("Введи имя машины: ")
	fmt.Scan(&c.name)
	fmt.Println("Добро пожаловать на заправочную станцию,", c.name)
	fmt.Println("Введи текущее количество топлива: ")
	fmt.Scan(&c.volume)
	fmt.Println("Введи максимальное значение бака: ")
	fmt.Scan(&c.maxVolume)
}

func (g *GazStation) RefuelCar(c *Car) { // данная функция будет заправлять машину
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
	} else {
		fmt.Println("Ошибка: текущий объем топлива превышает максимальный объем бака.")
	}
}

func main() {
	for {
		myCar := Car{}
		myCar.AcceptCar()

		station := GazStation{}
		station.RefuelCar(&myCar)
	}
}
