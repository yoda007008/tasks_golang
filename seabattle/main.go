package main

import (
	"math/rand"
	"time"
)

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

type GameRun struct {
	players []Player
	board   Board
	ships   []Ship
	isEnded bool
}

func (g *GameRun) Start() {
	rand.Seed(time.Now().UnixNano())

	shipSizes := []int{fourShip, thirdShip, thirdShip, secondShip, secondShip, secondShip, firstShip, firstShip, firstShip, firstShip}

	for _, p := range g.players {
		Player.DoMove(p)
	}

}
