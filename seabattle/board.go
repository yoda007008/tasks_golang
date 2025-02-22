package main

import (
	"errors"
	"fmt"
)

const (
	boardSize = 10
)

type Board interface {
	PlaceShips() error
	HandleShoot(int, int) (string, error)
	Size() int
}

type BoardImpl struct {
	Board2d [boardSize][boardSize]Ship
	Ships   []Ship // для быстрого доступа
}

func (b BoardImpl) PlaceShips() error {
	//TODO implement me
	panic("implement me")
}

func (b BoardImpl) HandleShoot(x, y int) (string, error) {
	// todo validate в отдельной функции выше по коду 1 раз
	if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
		fmt.Errorf("%w", errors.New("incorrect coords")) // todo custom error
	}

	ship := b.Board2d[x][y]

	// Игровое поле
	// &Ship{}, &Ship{}, &Ship{}, nil, nil // [0, 0], [0, 1], [0, 2] => [0, 1, 2], пример выстрел 0 2, так как корабль горизонтальный, то 2-0 = 2 это координата палубы внутри корабля
	// nil, nil, nil, nil, nil
	// nil, nil, nil, nil, nil
	// nil, nil, nil, nil, nil
	// nil, nil, nil, nil, nil

	switch ship.GetStatus() { // todo вычислить порядковый номер, перевести две координаты в одну. если горизонтально, то x - ship.x, если вертикально y - ship.y
	case Alive:
		ship.HandleShoot()
	case Hurt:
		ship.HandleShoot()
	case Dead:
		// нет смысла дальше по нему стрелять, это как в труп стрелять

	}
	//if ship.GetStatus() {
	//	*ship = "M"
	//	return "Промах!"
	//}
	//if *ship == "S" {
	//	*ship = "X"
	//	return "Попал!"
	//}
	//if *ship == "M" || *ship == "X" {
	//	return "Ты уже сюда стрелял"
	//}
	return "Ошибка координат", nil
}

func (b BoardImpl) Size() int {
	//TODO implement me
	panic("implement me")
}

func NewBoard() Board {
	return &BoardImpl{}
}
