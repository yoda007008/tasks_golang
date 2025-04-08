package main

import (
	"fmt"
	"strings"
)

type Game struct {
	rooms     map[string]Room
	player    Player
	doorState bool // состояние двери, false закрыто, true открыто
}

type Room struct {
	description string
	items       []string
	exits       []string
}

type Player struct {
	state     string
	inventory []string
	backPack  bool // наличие рюкзака
}

func initGame() *Game { // инициализация игры
	return &Game{
		rooms: map[string]Room{
			"кухня": {
				description: "ты находишься на кухне, надо собрать рюкзак и идти в универ",
				items:       []string{"чай"},
				exits:       []string{"коридор"},
			},
			"комната": {
				description: "ты находишься в комнате",
				items:       []string{"ключи", "конспекты", "рюкзак"},
				exits:       []string{"коридор"},
			},
			"коридор": {
				description: "ты находишься в коридоре, тут ничего интересеного",
				exits:       []string{"кухня", "комната", "улица"},
			},
			"улица": {
				description: "ты находишься на улице",
				exits:       []string{"домой"},
			},
		},
		player: Player{
			state:     "Кухня",
			inventory: []string{},
			backPack:  false,
		},
		doorState: false,
	}
}

func lookAround(game *Game) string { // осмотреться
	// получаем текущую локацию
	curLocation := game.player.state
	loc := game.rooms[curLocation]

	var itemsStr, exitsStr string

	if len(loc.items) > 0 {
		itemsStr = "на столе: " + strings.Join(loc.items, ", ")
	}

	exitsStr = "можно пройти - " + strings.Join(loc.exits, ", ")
	return fmt.Sprintf("%s, %s, %s", loc.description, itemsStr, exitsStr)
}

func goTo(destination string, game *Game) string { // перемещение игрока
	curLocation := game.player.state
	loc := game.rooms[curLocation]

	for _, exit := range loc.exits {
		if exit == destination {
			if destination == "улица" && game.doorState != false {
				return "Дверь закрыта"
			}
			game.player.state = destination
			return game.rooms[destination].description + "можно пройти - " + strings.Join(game.rooms[destination].exits, ", ")
		}
	}
	return "Нет пути в " + destination
}

func handleCommand(command string, game *Game) {

}

func main() {

}
