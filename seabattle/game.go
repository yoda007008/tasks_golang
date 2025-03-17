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
		// Пропускаем ход, если игрок уже выбыл
		if !p.StatusPlayer() {
			continue
		}

		fmt.Printf("\n Ход игрока %d \n", id)

		// создаем список других игроков, исключая текущего
		otherPlayers := make([]Player, 0, len(g.players)-1)
		otherPlayers = append(g.players[:id], g.players[id+1:]...)

		// игрок делает ход
		err := p.DoMove(otherPlayers)
		if err != nil {
			fmt.Printf("Игрок %d ошибся: %v\n", id, err)
		}

		// выводим поле текущего игрока после хода
		fmt.Printf("\nПоле игрока %d после хода:\n", id)
		p.(*PlayerImpl).board.PrintField()

		// проверяем, завершилась ли игра после хода
		if g.IsEnded() {
			break
		}
	}
}
func (g *GameImpl) IsEnded() bool {
	alivePlayers := 0
	for _, p := range g.players {
		if p.StatusPlayer() == true {
			alivePlayers++
		}
	}
	return alivePlayers <= 1
}
