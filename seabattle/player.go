package main

import (
	"errors"
	"fmt"
),
	"github.com/eiannone/keyboard"
)

type Player interface {
	DoMove(otherPlayers []Player) error
	TakeMove(x, y int)
	GiveUp()
}

// как вариант использовать интерфейсы: сделать мок реализацию игрока с методом DoMove(), которая будет просто выводить: "Игрок [id] делает ход"
type MockPlayer struct {
	id int
}

type PlayerImpl struct {
	id    int
	board Board
}

func (p *MockPlayer) DoMove(otherPlayers []Player) error { // todo не знаем про boardSize ? board.Size()
	//TODO implement me
	var x, y int
	fmt.Printf("Игрок %d ваш ход. Введите координату выстрела x\n", p.id)
	// todo выделить в отдельную структуру координату, на ней сделать метод New() error, и возвращать там ошибку
	// todo сделать пользовательский ввод, пока он не введёт корректные координаты
	fmt.Scan(&x, &y)
	fmt.Printf("Игрок %d ваш ход. Введите координату выстрела y\n", p.id)


	for _, p := range otherPlayers {
		p.TakeMove(x, y)
	}

	return errors.New("Ошибка координат")
}

// TakeMove принимает и обрабатывает выстрел
func (p *PlayerImpl) TakeMove(x, y int) {
	fmt.Printf("Игрок %d сделал ход по координатам (%d, %d)", p.id, x, y)
	p.board.HandleShoot(x, y)
}

func (p *PlayerImpl) GiveUp() {
	fmt.Printf("Игрок %d сдался", p.id)
}

func NewPlayer() Player {
	return &PlayerImpl{}
}
