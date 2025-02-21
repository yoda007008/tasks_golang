package main

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
