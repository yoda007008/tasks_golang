package main

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

func main() {

}
