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
	PlaceShip(int) error
	HandleShoot(int, int) string
}

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
				switch board.GetStatus() {
				case Alive, Hurt:
					fmt.Print("X")
				case Dead:
					fmt.Print("D")
				}
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

func (b BoardImpl) HandleShoot(x, y int) string {
	if x < 0 || x > boardSize || y < 0 || y > boardSize {
		return "Ошибка координат"
	}

	ship := b.Board2d[x][y]
	if ship == nil {
		return "Мимо"
	}
	if ship.HandleShoot(x, y) {
		status := ship.GetStatus()
		switch status {
		case Alive:
			return "Попал"
		case Hurt:
			return "Корабль ранен"
		case Dead:
			return "Корабль уничтожен"
		}
	}
	return "Ошибка координат"
}

func NewBoard() Board {
	return &BoardImpl{}
}
