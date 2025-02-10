package main

type Field struct {
	Board [boardSize][boardSize]string
}

type Player struct {
}

type Cord struct {
	x int
	y int
}
type Ship struct {
	cords  []Cord
	size   int
	isDead bool
}

type Game struct {
	field  Field
	player Player
}

func (g *Game) start() {

}
