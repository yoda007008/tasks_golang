package main

import (
	"fmt"
	"seabattle/seabattle/interfaces"
)

type Player interface {
	DoMove(otherPlayers []Player) error
	TakeMove(int, int)
	GiveUp()
	StatusPlayer() bool
}

type PlayerImpl struct {
	id     int
	board  interfaces.Board
	status bool
}

func (p *PlayerImpl) DoMove(otherPlayers []Player) error {
	var x, y int
	for {
		fmt.Printf("Игрок %d ваш ход. Введите координату выстрела x и y (через пробел):\n", p.id)
		_, err := fmt.Scan(&x, &y)
		if err == nil && x >= 0 && y >= 0 {
			break
		}
		fmt.Println("Ошибка ввода. Пожалуйста, введите корректные координаты.")
	}

	for _, player := range otherPlayers {
		player.TakeMove(x, y)
	}

	return nil
}

func (p *PlayerImpl) TakeMove(x, y int) {
	fmt.Printf("Игрок %d сделал ход по координатам (%d, %d)\n", p.id, x, y)
	result := p.board.HandleShoot(x, y)
	fmt.Println(result) // Выводим результат выстрела
}

func (p *PlayerImpl) GiveUp() {
	fmt.Printf("Игрок %d сдался\n", p.id)
	p.status = false // Игрок сдался, меняем его статус
}

func (p *PlayerImpl) StatusPlayer() bool {
	return p.status
}

func NewPlayer(id int, board interfaces.Board, status bool) Player {
	return &PlayerImpl{
		id:     id,
		board:  board,
		status: status,
	}
}
