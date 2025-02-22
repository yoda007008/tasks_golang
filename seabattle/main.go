package main

import "fmt"

const (
	boardSize  = 10
	fourShip   = 4
	thirdShip  = 3
	secondShip = 2
	firstShip  = 1
)

type Game interface {
	Start()
	IsEnded() bool
}

type Player interface {
	DoMove() (int, int, error)
	GiveUp()
}

type Board interface {
	PlaceShips() error
	TakeShoot() (string, error)
	Size() int
}

type Ship interface {
	GetStatus()
	TakeShoot()
}

type Field struct {
	Board [boardSize][boardSize]string
}

type Cord struct {
	x int
	y int
}

type GameRun struct {
	players []Player
	board   Board
	ships   []Ship
	isEnded bool
}

type CreateShip struct {
	size   int
	isDead bool
	cord   []Cord
}

type HumanPlayer struct {
	board *Field
}

func NewField() *Field { // заполнять и создавать поля нужно также, как и в предыдущей программе морского боя
	f := &Field{}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			f.Board[i][j] = "."
		}
	}
	return f
}

func (f *Field) Print() { // вывод поля
	fmt.Println("  0 1 2 3 4 5 6 7 8 9 10")
	for i, row := range f.Board {
		fmt.Printf("%d", i+1)
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
}

func (f *Field) PlaceShip() { // функция располагает для двух игроков корабли на поле рандомно

}

func (f *Field) DoMove() { // функция DoMove делает ходы

}
