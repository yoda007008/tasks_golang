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
	PrintField()
	PlaceShip(int) error
	HandleShoot(int, int) (string, error)
	Size() int
}

type orientation int

const (
	gorizontal = 0
	vertical   = 1
)

type BoardImpl struct {
	Board2d [boardSize][boardSize]Ship
	Ships   []Ship
}

func (b BoardImpl) PrintField() {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board := b.Board2d[i][j]
			if board == nil {
				fmt.Print(".")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
}

func (b *BoardImpl) canPlaced(x, y, size int, o orientation) bool { // данная функция проверяет, можем ли мы расположить корабль или нет
	if o == gorizontal && x+size > boardSize { // проверка, что мы не выходим за границы поля
		return false
	}
	if o == vertical && y+size > boardSize {
		return false
	}
	for i := -1; i <= size; i++ {
		for j := -1; j <= size; j++ {
			curX := x // координаты клетки для проверки
			curY := y

			if o == gorizontal {
				curX += i
				curY += j
			} else {
				curX += j
				curY += i
			}
			if curX >= 0 && curX < boardSize && curY >= 0 && curY < boardSize { // не выходим за пределы поля, проверяем на занятость полей
				if b.Board2d[curX][curY] != nil {
					return false
				}
			}
		}
	}
	return true
}
func (b *BoardImpl) PlaceShip(size int) error { // данная функция размещает корабли рандомно
	if size < 1 || size > 4 {
		fmt.Errorf("%d - данная длинна не походит", size)
	}

	o := orientation(rand.Intn(2))
	shipPlaced := false
	for !shipPlaced {
		x := rand.Intn(boardSize)
		y := rand.Intn(boardSize)

		if !b.canPlaced(x, y, size, o) {
			continue
		}

		// Проверка, можно ли разместить корабль
		canPlace := true
		for i := 0; i < size; i++ {
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

		// создание корабля
		ship := &StandardShipImpl{
			decks:       make([]DeckStatus, size),
			x:           x,
			y:           y,
			orientation: int(o),
			size:        size,
		}
		for j := 0; j < size; j++ {
			ship.decks[j] = AliveDeck
		}

		// размещение корабля на поле
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
		return "Промахнулся", nil
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
