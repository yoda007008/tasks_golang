package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewGame(2) // 2 players
	fmt.Print(game)
	for !game.IsEnded() { // цикл игры
		game.Round()
	}

	for _, player := range game.(*GameImpl).players {
		if player.GetStatusPlayer() {
			fmt.Printf("Игрок %d победил!\n", player.(*PlayerImpl).id)
			break
		}
	}
}
