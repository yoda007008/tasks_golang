package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	hit    = true
	missed = false
)

type FirstRobot struct {
	name   string
	stock  int
	hp     int
	damage int
}

type SecondRobot struct {
	name   string
	stock  int
	hp     int
	damage int
}

type Game struct {
	f FirstRobot
	s SecondRobot
}

func (g *Game) BeginGame() { // начало игры
	fmt.Println("Добро пожаловать в битву роботов (один выстрел 0,1 литра борща)")
	fmt.Println("Введи имя первого робота ")
	fmt.Scan(&g.f.name)
	fmt.Println("Введи имя второго робота ")
	fmt.Scan(&g.s.name)
	fmt.Println("Введи литры для первого робота ")
	fmt.Scan(&g.f.stock)
	fmt.Println("Введи литры для второго робота ")
	fmt.Scan(&g.s.stock)
	fmt.Println("Введи hp для первого робота (hp не больше 100) ")
	fmt.Scan(&g.f.hp)
	fmt.Println("Введи hp для второго робота (hp не больше 100) ")
	fmt.Scan(&g.s.hp)

	if g.f.hp > 100 || g.s.hp > 100 || g.f.hp < 0 || g.s.hp < 0 {
		fmt.Println("Хп не может быть меньше 0 или меньше 100")
	}
}

func (g *Game) canShoot() bool {
	rand.Seed(time.Now().UnixNano())
	shoot := rand.Intn(2)
	return shoot == 1 // если выстрел успешен, возвращаем true
}

func (g *Game) DamagedRobots() { // основная логика игры
	if g.canShoot() {
		fmt.Printf("%s кидает %s!\n", g.f.name, g.s.name)
		g.s.hp -= 10
		fmt.Printf("%s получил урон! Осталось hp: %d\n", g.s.name, g.s.hp)
	} else {
		fmt.Printf("%s промахнулся!\n", g.f.name)
	}

	if g.canShoot() {
		fmt.Printf("%s кидает %s!\n", g.s.name, g.f.name)
		g.f.hp -= 10
		fmt.Printf("%s получил урон! Осталось hp: %d\n", g.f.name, g.f.hp)
	} else {
		fmt.Printf("%s промахнулся!\n", g.s.name)
	}
}

func (g *Game) StartBattle() { // процесс сражения роботов
	for g.f.hp > 0 && g.s.hp > 0 {
		g.DamagedRobots()
	}

	if g.f.hp <= 0 {
		fmt.Printf("%s побеждает, %s уничтожен\n", g.s.name, g.f.name)
	} else if g.s.hp <= 0 {
		fmt.Printf("%s побеждает, %s уничтожен\n", g.f.name, g.s.name)
	} else {
		fmt.Println("Ничья")
	}
}
func main() {
	for {
		game := Game{}
		game.BeginGame()
		game.StartBattle()
	}
}
