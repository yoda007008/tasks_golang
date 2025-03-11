package main

import (
	"fmt"
)

type Player interface {
	DoMove(otherPlayers []Player) error
	TakeMove(int, int)
	GiveUp()
}

type PlayerImpl struct {
	id    int
	board Board
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
	fmt.Printf("Игрок %d сделал ход по координатам (%d, %d)", p.id, x, y)
	p.board.HandleShoot(x, y)
}

func (p *PlayerImpl) GiveUp() {
	fmt.Printf("Игрок %d сдался\n", p.id)
}

func NewPlayer(id int, board Board) Player {
	return &PlayerImpl{
		id:    id,
		board: board,
	}
}
