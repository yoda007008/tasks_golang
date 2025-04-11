package main

import (
	"fmt"
	"math/rand"
)

const (
	boardSize = 10
)

type orientation int

const (
	gorizontal = 0
	vertical   = 1
)

type Board interface {
	PrintField()
	PlaceShipRandom(int) error
	HandleShoot(int, int) string
	PlaceShipCoords(int, int, int, orientation) bool
}

type BoardImpl struct {
	Board2d [boardSize][boardSize]Cell
	Ships   []Ship
}

type Cell struct {
	status bool
	ship   Ship
}

func (b BoardImpl) PrintField() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			cell := b.Board2d[i][j]
			if cell.status {
				if cell.ship != nil {
					fmt.Print("X ") // попадание
				} else {
					fmt.Print("O ") // промах
				}
			} else if cell.ship != nil {
				fmt.Print("S ") // корабль (для отладки)
			} else {
				fmt.Print(". ") // пустая клетка
			}
		}
		fmt.Println()
	}
}

func (b *BoardImpl) canPlaced(x, y, size int, o orientation) bool {
	if o == gorizontal && x+size > boardSize {
		return false
	}
	if o == vertical && y+size > boardSize {
		return false
	}
	for i := -1; i <= size; i++ {
		for j := -1; j <= size; j++ {
			curX := x
			curY := y

			if o == gorizontal {
				curX += i
				curY += j
			} else {
				curX += j
				curY += i
			}
			if curX >= 0 && curX < boardSize && curY >= 0 && curY < boardSize {
				if b.Board2d[curX][curY].ship != nil {
					return false
				}
			}
		}
	}
	return true
}

func (b *BoardImpl) PlaceShipCoords(x, y, size int, o orientation) bool {
	if !b.canPlaced(x, y, size, o) {
		return false
	}

	ship := &StandardShipImpl{
		decks:       make([]ShipStatus, size),
		x:           x,
		y:           y,
		orientation: int(o),
		size:        size,
	}
	for i := 0; i < size; i++ {
		if o == gorizontal {
			b.Board2d[x+i][y].ship = ship
		} else {
			b.Board2d[x][y+i].ship = ship
		}
	}
	b.Ships = append(b.Ships, ship)
	return true
}

func (b *BoardImpl) PlaceShipRandom(size int) error {
	if size < 1 || size > 4 {
		return fmt.Errorf("%d - данная длина не подходит", size)
	}

	o := orientation(rand.Intn(2))
	shipPlaced := false
	for !shipPlaced {
		x := rand.Intn(boardSize)
		y := rand.Intn(boardSize)

		if !b.canPlaced(x, y, size, o) {
			continue
		}

		b.PlaceShipCoords(x, y, size, o)
		shipPlaced = true
	}
	return nil
}

func (b *BoardImpl) HandleShoot(x, y int) string {
	if x < 0 || x >= boardSize || y < 0 || y >= boardSize {
		return "Ошибка координат"
	}

	cell := &b.Board2d[x][y]
	cell.status = true

	if cell.ship == nil {
		return "Мимо"
	}

	if cell.ship.HandleShoot(x, y) {
		status := cell.ship.GetStatus()
		switch status {
		case Alive:
			return "Попал"
		case Hurt:
			return "Корабль ранен"
		case Dead:
			return "Корабль уничтожен"
		}
	}
	return "Ошибка, некорректный ввод"
}

func NewBoard() Board {
	return &BoardImpl{}
}
