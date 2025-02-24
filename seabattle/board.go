package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	boardSize = 10
)

type Board interface {
	PlaceShips() error
	HandleShoot(int, int) (string, error)
	Size() int
}

type orientation int

const (
	gorizontal = 0
	vertical   = 1
)

type ShipStruct struct {
	Size int
}

type BoardImpl struct {
	Board2d [boardSize][boardSize]Ship
	Ships   []Ship // для быстрого доступа
}

func (b BoardImpl) PlaceShips(size int) error {
	if size < 0 || size > 4 {
		return fmt.Errorf("%w", "Invalid ship size")
	}

	o := orientation(rand.Intn(2))
	shipPlaced := false
	for !shipPlaced {
		x := rand.Intn(boardSize)
		y := rand.Intn(boardSize)

		// проверка, не выходит ли корабль за поле
		if o == gorizontal && x+size > boardSize {
			continue
		}
		if o == vertical && y+size > boardSize {
			continue
		}

		// проходимся по случаям, по которым нельзя разместить корабль
		canPlace := true
		for i := 0; i <= size; i++ {
			if o == gorizontal && b.Board2d[x+i][y] != nil {
				canPlace = false
				break
			}
			if o == vertical && b.Board2d[x][y+i] != nil {
				canPlace = false
				break
			}
		}
		if !canPlace {
			continue
		}

		// размещение корабля на поле
		ship := &ShipStruct{Size: size}
		for i := 0; i < size; i++ {
			if o == gorizontal {
				b.Board2d[x+i][y] = ship
			} else {
				b.Board2d[x][y+i] = ship
			}
		}
		b.Ships = append(b.Ships, ship)
		shipPlaced = true
	}
	return nil
}

func (b BoardImpl) HandleShoot(x, y int) (string, error) {
	// todo validate в отдельной функции выше по коду 1 раз
	if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
		fmt.Errorf("%w", errors.New("incorrect coords")) // todo custom error
	}
	ship := b.Board2d[x][y]
	if ship == nil {
		return "Miss shot", nil
	}

	// индекс палубы
	//var deskIndex int
	//switch o {
	//case gorizontal: // две ориентации
	//	deskIndex = x - ship.x
	//case vertical:
	//	deskIndex = y - ship.y
	//}

	switch ship.GetStatus() { // todo вычислить порядковый номер, перевести две координаты в одну. если горизонтально, то x - ship.x, если вертикально y - ship.y
	case Alive:
		ship.HandleShoot()
	case Hurt:
		ship.HandleShoot()
	case Dead:
		fmt.Println(false)
	}

	return "Ошибка координат", nil
	// Игровое поле
	// &Ship{}, &Ship{}, &Ship{}, nil, nil // [0, 0], [0, 1], [0, 2] => [0, 1, 2], пример выстрел 0 2, так как корабль горизонтальный, то 2-0 = 2 это координата палубы внутри корабля
	// nil, nil, nil, nil, nil
	// nil, nil, nil, nil, nil
	// nil, nil, nil, nil, nil
	// nil, nil, nil, nil, nil
}

func (b BoardImpl) Size() int {
	//TODO implement me
	panic("implement me")
}

func NewBoard() Board {
	return &BoardImpl{}
}
