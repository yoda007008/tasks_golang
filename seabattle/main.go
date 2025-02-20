package main

const boardSize = 10

type Board [boardSize][boardSize]string
type FirstPlayer struct {
	Board  Board
	Result bool
}
type SecondPlayer struct {
	Board  Board
	Result bool
}
type ShipStatus struct {
	Wounded bool
	IsDead  bool
}
type FieldStatus struct {
	// тут filled - статус заполненного поля
	empty Board // пустое поле
}

func (p *FirstPlayer) MakeShoot(status ShipStatus) {

}

func (p *SecondPlayer) MakeShootSecond(status ShipStatus) {

}

func main() {

}
