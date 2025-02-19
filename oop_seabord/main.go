package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	boardSize  = 10
	fourShip   = 4
	thirdShip  = 3
	secondShip = 2
	thirstShip = 1
)

var errPlaceShip = errors.New("Не удалось разместить корабль")
var errCord = errors.New("Ошибка координат")

type orientation int

const (
	gorizontal = 0
	vertical   = 1
)

type Field struct {
	Board [boardSize][boardSize]string
}

type Cord struct {
	x int
	y int
}
type Ship struct {
	size   int
	isDead bool
}

func GetField() *Field {
	f := &Field{}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			f.Board[i][j] = "."
		}
	}
	return f
}

func Print(f *Field) {
	fmt.Println("  1 2 3 4 5 6 7 8 9 10")
	for i, row := range f.Board {
		fmt.Printf("%d ", i+1)
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
}

func (f *Field) IsValid(x int, y int, size int, orientation orientation) bool {
	for k := 0; k < boardSize; k++ {
		switch orientation {
		case gorizontal:
			x++
		case vertical:
			y++
		}
		if size < 0 || size > 4 {
			break
		}
		if x > boardSize-1 || y > boardSize-1 {
			return false
		}
		for row := x - 1; row < x-1+3; row++ {
			for col := y - 1; col < y-1+3; col++ {
				if row <= 0 || row > boardSize-1 || col <= 0 || col > boardSize-1 {
					continue
				}
				if f.Board[row][col] != "." {
					return false
				}
			}
		}
	}
	return true
}

func (f *Field) PlaceShip(ship *Ship) bool {
	o := orientation(rand.Intn(2))
	ship.isDead = false
	for !ship.isDead {
		x := rand.Intn(boardSize)
		y := rand.Intn(boardSize)
		if f.IsValid(x, y, ship.size, o) {
			continue
		}
		for i := 0; i < boardSize; i++ {
			if ship.size < 0 || ship.size > 4 {
				break
			}
			f.Board[x][y] = "S"
			switch o {
			case gorizontal:
				x++
			case vertical:
				y++
			}
		}
		ship.isDead = true
	}
	return ship.isDead
}

func (f *Field) MakeShoot(cord Cord) error {
	x := cord.x
	y := cord.y
	if cord.x < 0 || cord.x > boardSize || cord.y < 0 || cord.y > boardSize {
		fmt.Errorf("%w", errCord)
	}
	cell := &f.Board[x][y]

	if *cell == "." {
		*cell = "M"
		fmt.Println("Промах")
	}
	if *cell == "S" {
		*cell = "X"
		fmt.Println("Попал")
	}
	if *cell == "X" || *cell == "M" {
		fmt.Println("Ты уже сюда стрелял")
	}
	return errCord
}

func main() {
	rand.Seed(time.Now().UnixNano())
	shipsSizes := []int{fourShip, thirdShip, thirdShip, secondShip, secondShip, secondShip, thirstShip, thirstShip, thirstShip, thirstShip}
	board := GetField()

	ships := make([]Ship, len(shipsSizes))

	for i, shipSize := range shipsSizes {
		ships[i] = Ship{size: shipSize}
		if !board.PlaceShip(&ships[i]) {
			fmt.Errorf("%w", errPlaceShip)
			continue
		}
	}

	for {
		var x, y int
		Print(board)
		fmt.Println("Добро пожаловать в морской бой!")
		fmt.Print("Введите координату x: ")
		fmt.Scan(&x)
		fmt.Print("Введите координату y: ")
		fmt.Scan(&y)

		cord := Cord{x, y}
		err := board.MakeShoot(cord)
		if err != nil {
			fmt.Println(err)
		}

	}
}
