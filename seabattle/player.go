package main

import (
	"fmt"
)

type Player interface {
	DoMove(otherPlayers []Player) error
	TakeMove(int, int)
	GiveUp()
	GetStatusPlayer() bool
}

type Status bool

type PlayerImpl struct {
	id     int
	board  Board
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
	fmt.Printf("Игрок %d сделал ход по координатам (%d, %d)", p.id, x, y)
	p.board.HandleShoot(x, y)
}

func (p *PlayerImpl) GiveUp() {
	fmt.Printf("Игрок %d сдался\n", p.id)
}

func (p *PlayerImpl) GetStatusPlayer() bool {
	if p.status == true {
		return true
	}
	return false
}

func NewPlayer(id int, board Board, status bool) Player {
	return &PlayerImpl{
		id:     id,
		board:  board,
		status: status,
	}
}
