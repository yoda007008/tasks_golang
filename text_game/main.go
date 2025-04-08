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
	state     string // текущее состояние игрока
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

func takeItem(item string, game *Game) { // данный метод удаляет предмет с локации и перемещает его в инвентарь к игроку
	curLocation := game.player.state
	loc := game.rooms[curLocation]

	for i, locItem := range loc.items {
		if locItem == item {
			if !game.player.backPack && item == "рюкзак" {
				return "некуда класть"
			}
			game.rooms[game.player.state].items = append(loc.items[:i], loc.items[i+1:]...)

			game.player.inventory = append(game.player.inventory, item)

			return "предмет добавлен в инвентарь" + item
		}
	}
}

func wearItem(item string, game *Game) string { // данный метод проверяет наличие рюкзака и ищет его в текущей локации
	if item != "рюкзак" {
		return "можно надеть только рюкзак"
	}

	for i, lockItem := range game.rooms[game.player.state].items {
		if lockItem == "рюкзак" {
			game.rooms[game.player.state].items = append(game.rooms[game.player.state].items[:i], game.rooms[game.player.state].items[i+1:]...)
			game.player.backPack = true
			return "вы надели рюкзак"
		}
	}
	return "нет такого предмета"
}

func handleCommand(command string, game *Game) {

}

func main() {

}
