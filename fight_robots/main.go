package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name      string
	stock     int
	hp        int
	damage    int
	isCooking bool
}

func NewRobot(name string, stock, hp, damage int) *Robot {
	return &Robot{
		name:   name,
		stock:  stock,
		hp:     hp,
		damage: damage,
	}
}

func (r *Robot) ThrowsPlate(target *Robot) { // добавил новый метод
	if r.isCooking {
		fmt.Printf("%s пропускает ход, чтобы сварить новый борщ!\n", r.name)
		r.isCooking = false
		return
	}

	if r.stock <= 0 {
		fmt.Printf("%s пытается сварить новый борщ\n", r.name)
		r.TryCook()
		return
	}

	if rand.Intn(2) == 0 {
		fmt.Printf("%s уклонился от тарелки с борщом\n", target.name)
		return
	}

	r.stock -= 1
	target.hp -= r.damage
	fmt.Printf("%s кинул тарелку борща в %s и нанес %d урона!\n", r.name, target.name, r.damage)
}

func (r *Robot) TryCook() {
	if rand.Intn(10) < 2 {
		fmt.Printf("%s взорвал кастрюлю и потерял 10 hp!\n", r.name)
		r.hp -= 10
	} else {
		r.stock += 5
		fmt.Printf("%s успешно сварил 5 литров борща!\n", r.name)
	}
}

func (r *Robot) IsAlive() bool {
	return r.hp > 0
}

func BattleRobots(robot1, robot2 *Robot) {
	rand.Seed(time.Now().UnixNano())

	turn := 0

	for robot1.IsAlive() && robot2.IsAlive() {
		turn++
		fmt.Printf("\n--- Ход %d ---\n", turn)

		if robot1.IsAlive() {
			robot1.ThrowsPlate(robot2)
			if !robot2.IsAlive() {
				fmt.Printf("%s побежден! %s побеждает!\n", robot2.name, robot1.name)
				break
			}
		}

		if robot2.IsAlive() {
			robot2.ThrowsPlate(robot1)
			if !robot1.IsAlive() {
				fmt.Printf("%s побежден! %s побеждает!\n", robot1.name, robot2.name)
				break
			}
		}

		fmt.Printf("%s: HP = %d, Борщ = %d\n", robot1.name, robot1.hp, robot1.stock)
		fmt.Printf("%s: HP = %d, Борщ = %d\n", robot2.name, robot2.hp, robot2.stock)
	}
}

func main() {
	robot1 := NewRobot("Senior400k/nanosek", 10, 100, 15)
	robot2 := NewRobot("Junior Developer", 8, 120, 12)

	fmt.Println("Битва роботов начинается!")
	BattleRobots(robot1, robot2)
}
