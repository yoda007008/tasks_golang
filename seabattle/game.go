package main

type Game interface {
	//Start()
	Round()
	IsEnded() bool // todo как понять что игра окончена? У одного игрока все корабли остались в живых
}

type GameImpl struct {
	players []Player // todo когда игра окончена, у всех игроков, кроме победителя, все корабли мертвы
	ships   []Ship
	isEnded bool
}

func NewGame(playersCount int) Game {
	// todo init all entities
	//players := make([]Player, 0, playersCount)

	// todo slice basics
	// numbers := make([]int, 10, 10) // 0x01 [0, 0, 0, 0, 0, 0, 0, 0, 0, 0], len = 10, cap = 10
	// 0x02 [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5], len = 11, cap = 20

	// numbers2 := make([]int, 0, 10) // 0x01 [1, 2], len = 2, cap = 10

	//for i := 0; i < playersCount; i++ {
	//	board := NewBoard()
	//	players = append(players, &PlayerImpl{
	//		id:    i,
	//		board: board,
	//	})
	//}

	return &GameImpl{
		players: nil,
		ships:   nil,
		isEnded: false,
	}
}

//func (g *GameImpl) Start() {
//
//}

func (g *GameImpl) Round() {
	for id, p := range g.players { // [p0, p1, p2, p3]
		otherPlayers := make([]Player, 0, len(g.players)-1)
		if id == len(g.players)-1 {
			otherPlayers = append(otherPlayers, g.players[:id]...) // todo refactor duplicate code
		} else {
			otherPlayers = append(otherPlayers, g.players[:id]...)
			otherPlayers = append(otherPlayers, g.players[id+1:]...)
		}

		p.DoMove(otherPlayers) // если ходит игрок 1, то в otherPlayers будет игрок 2 например в случае если игрока 2
	}

	// todo в конце раунда проверяем состояние кораблей у игроков и определяем, окончена игра или нет, и если окончена то кто победитель(победитель тот, кто остался в живых) => у кого хотя бы один корабль жив, а у остальных все корабли мертвы
}

func (g *GameImpl) IsEnded() bool {
	return false
}
