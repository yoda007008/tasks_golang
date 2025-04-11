package main

import (
	"fmt"
	"strings"
)

type Game struct {
	rooms        map[string]Room
	player       Player
	doorState    bool
	kitchenFirst bool
}

type Room struct {
	description string
	tableItems  []string
	chairItems  []string
	exits       []string
}

type Player struct {
	state     string
	inventory []string
	backPack  bool // наличие рюкзака
}

func initGame() *Game {
	return &Game{
		rooms: map[string]Room{
			"кухня": {
				description: "ты находишься на кухне, надо собрать рюкзак и идти в универ",
				tableItems:  []string{"чай"},
				chairItems:  []string{},
				exits:       []string{"коридор"},
			},
			"комната": {
				description: "ты в своей комнате",
				tableItems:  []string{"ключи", "конспекты"},
				chairItems:  []string{"рюкзак"},
				exits:       []string{"коридор"},
			},
			"коридор": {
				description: "ничего интересного",
				tableItems:  []string{},
				chairItems:  []string{},
				exits:       []string{"кухня", "комната", "улица"},
			},
			"улица": {
				description: "на улице весна",
				tableItems:  []string{},
				chairItems:  []string{},
				exits:       []string{"домой"},
			},
		},
		player: Player{
			state:     "кухня", // должно быть в нижнем регистре
			inventory: []string{},
			backPack:  false,
		},
		doorState:    false,
		kitchenFirst: true,
	}
}

func lookAround(game *Game) string {
	loc := game.rooms[game.player.state]

	var itemsParts []string
	if len(loc.tableItems) > 0 {
		itemsParts = append(itemsParts, "на столе: "+strings.Join(loc.tableItems, ", "))
	}
	if len(loc.chairItems) > 0 {
		itemsParts = append(itemsParts, "на стуле: "+strings.Join(loc.chairItems, ", "))
	}

	itemsStr := strings.Join(itemsParts, ", ")
	exitsStr := "можно пройти - " + strings.Join(loc.exits, ", ")

	switch game.player.state {
	case "кухня":
		// если есть рюкзак - показываем полное описание с "надо идти в универ"
		if game.player.backPack {
			if itemsStr != "" {
				return fmt.Sprintf("ты находишься на кухне, %s, надо идти в универ. %s",
					itemsStr, exitsStr)
			}
			return fmt.Sprintf("ты находишься на кухне, надо идти в универ. %s",
				exitsStr)
		}

		// если это первое посещение кухни
		if game.kitchenFirst {
			game.kitchenFirst = false
			if itemsStr != "" {
				return fmt.Sprintf("ты находишься на кухне, %s, надо собрать рюкзак и идти в универ. %s",
					itemsStr, exitsStr)
			}
			return fmt.Sprintf("ты находишься на кухне, надо собрать рюкзак и идти в универ. %s",
				exitsStr)
		}

		// все остальные случаи - полное описание
		if itemsStr != "" {
			return fmt.Sprintf("ты находишься на кухне, %s, надо собрать рюкзак и идти в универ. %s",
				itemsStr, exitsStr)
		}
		return fmt.Sprintf("ты находишься на кухне, надо собрать рюкзак и идти в универ. %s",
			exitsStr)

	case "комната":
		if len(loc.tableItems) == 0 && len(loc.chairItems) == 0 {
			return fmt.Sprintf("пустая комната. %s", exitsStr)
		}
		if itemsStr != "" {
			return fmt.Sprintf("%s. %s", itemsStr, exitsStr)
		}
		return fmt.Sprintf("%s. %s", loc.description, exitsStr)

	default:
		if itemsStr != "" {
			return fmt.Sprintf("%s, %s. %s", loc.description, itemsStr, exitsStr)
		}
		return fmt.Sprintf("%s. %s", loc.description, exitsStr)
	}
}

func goTo(destination string, game *Game) string {
	loc := game.rooms[game.player.state]

	for _, exit := range loc.exits {
		if exit == destination {
			if destination == "улица" && !game.doorState {
				return "дверь закрыта"
			}
			game.player.state = destination

			// При входе в комнату - только описание
			if destination == "комната" {
				return fmt.Sprintf("%s. можно пройти - коридор",
					game.rooms[destination].description)
			}

			// При возвращении на кухню - всегда краткое описание
			if destination == "кухня" {
				return fmt.Sprintf("кухня, ничего интересного. можно пройти - коридор")
			}

			return lookAround(game)
		}
	}
	return "нет пути в " + destination
}

func takeItem(item string, game *Game) string {
	loc := game.rooms[game.player.state]

	// Проверяем предметы на столе
	for i, locItem := range loc.tableItems {
		if locItem == item {
			if !game.player.backPack && item != "рюкзак" {
				return "некуда класть"
			}

			loc.tableItems = append(loc.tableItems[:i], loc.tableItems[i+1:]...)
			game.rooms[game.player.state] = loc
			game.player.inventory = append(game.player.inventory, item)
			return "предмет добавлен в инвентарь: " + item
		}
	}

	// Проверяем предметы на стуле
	for i, locItem := range loc.chairItems {
		if locItem == item {
			if !game.player.backPack && item != "рюкзак" {
				return "некуда класть"
			}

			loc.chairItems = append(loc.chairItems[:i], loc.chairItems[i+1:]...)
			game.rooms[game.player.state] = loc
			game.player.inventory = append(game.player.inventory, item)
			return "предмет добавлен в инвентарь: " + item
		}
	}

	return "нет такого"
}

func wearItem(item string, game *Game) string {
	if item != "рюкзак" {
		return "можно надеть только рюкзак"
	}

	loc := game.rooms[game.player.state]

	// проверяем рюкзак на стуле
	for i, locItem := range loc.chairItems {
		if locItem == "рюкзак" {
			loc.chairItems = append(loc.chairItems[:i], loc.chairItems[i+1:]...)
			game.rooms[game.player.state] = loc
			game.player.backPack = true
			return "вы надели: рюкзак"
		}
	}

	// проверяем рюкзак на столе
	for i, locItem := range loc.tableItems {
		if locItem == "рюкзак" {
			loc.tableItems = append(loc.tableItems[:i], loc.tableItems[i+1:]...)
			game.rooms[game.player.state] = loc
			game.player.backPack = true
			return "вы надели: рюкзак"
		}
	}

	return "нет такого"
}

func useItem(item, target string, game *Game) string {
	hasItem := false
	for _, invItem := range game.player.inventory {
		if invItem == item {
			hasItem = true
			break
		}
	}

	if !hasItem {
		return "нет предмета в инвентаре - " + item
	}

	if item == "ключи" && target == "дверь" && game.player.state == "коридор" {
		game.doorState = true
		return "дверь открыта"
	}

	return "не к чему применить"
}

func handleCommand(command string, game *Game) string {
	parts := strings.Split(command, " ")
	switch parts[0] {
	case "осмотреться":
		return lookAround(game)
	case "идти":
		if len(parts) < 2 {
			return "укажите место"
		}
		return goTo(parts[1], game)
	case "взять":
		if len(parts) < 2 {
			return "укажите предмет"
		}
		return takeItem(parts[1], game)
	case "надеть":
		if len(parts) < 2 {
			return "укажите предмет"
		}
		return wearItem(parts[1], game)
	case "применить":
		if len(parts) < 3 {
			return "укажите предмет и объект"
		}
		return useItem(parts[1], parts[2], game)
	default:
		return "неизвестная команда"
	}
}

func main() {
	game := initGame()
	fmt.Println("Добро пожаловать в текстовую игру!")
	fmt.Println("Доступные команды: осмотреться, идти [место], взять [предмет], надеть [предмет], применить [предмет] [объект]")

	for {
		fmt.Print("> ")
		var input string
		fmt.Scanln(&input)

		if input == "выход" {
			break
		}

		response := handleCommand(input, game)
		fmt.Println(response)
	}
}
