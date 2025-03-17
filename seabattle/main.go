package main

import "fmt"

func main() {
	game := NewGame(2)

	shipSizes := []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}

	// Размещаем корабли на поле каждого игрока
	for _, player := range game.(*GameImpl).players {
		p := player.(*PlayerImpl)

		// Размещаем корабли из массива shipSizes
		for _, size := range shipSizes {
			err := p.board.PlaceShipRandom(size)
			if err != nil {
				fmt.Printf("Ошибка при размещении корабля размером %d: %v\n", size, err)
			}
		}

		fmt.Printf("\nПоле игрока %d после размещения кораблей:\n", p.id)
		p.board.PrintField()
	}

	round := 1
	for !game.IsEnded() {
		fmt.Printf("\n--- Раунд %d ---\n", round)
		game.Round()
		round++
	}

	var winnerID int
	for _, player := range game.(*GameImpl).players {
		if player.StatusPlayer() {
			winnerID = player.(*PlayerImpl).id
			break
		}
	}

	fmt.Printf("\nИгра окончена! Победил игрок %d\n", winnerID)
}
