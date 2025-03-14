package main

import "fmt"

type Game interface {
	Round()
	IsEnded() bool
}

type GameImpl struct {
	players []Player
	isEnded bool
}

func NewGame(playersCount int) Game {
	players := make([]Player, 0, playersCount)

	for i := 0; i < playersCount; i++ {
		board := NewBoard()
		players = append(players, &PlayerImpl{
			id:     i,
			board:  board,
			status: true,
		})
	}

	return &GameImpl{
		players: players,
		isEnded: false,
	}
}

func (g *GameImpl) Round() {
	for id, p := range g.players {
		// Создаем список других игроков, исключая текущего
		otherPlayers := make([]Player, 0, len(g.players)-1)
		otherPlayers = append(g.players[:id], g.players[id+1:]...)

		err := p.DoMove(otherPlayers)
		if err != nil {
			fmt.Printf("Игрок %d ошибся: %v\n", id, err)
		}
	}

	g.isEnded = g.IsEnded()
}

func (g *GameImpl) IsEnded() bool {
	alivePlayers := 0
	for _, p := range g.players {
		if p.GetStatusPlayer() == true {
			alivePlayers++
		}
	}
	return alivePlayers <= 1
}
