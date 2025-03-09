package main

type ShipStatus int

const (
	Alive ShipStatus = iota // цел
	Hurt                    // ранен
	Dead                    // мёртв
)

type DeckStatus int

const (
	AliveDeck DeckStatus = iota
	DeadDeck
)

type Ship interface {
	GetStatus() ShipStatus
	HandleShoot(x int, y int) bool
}

type StandardShipImpl struct {
	decks       []DeckStatus
	x           int // todo getX, координата начальной точки корабля
	y           int // todo getY, координата начальной точки корабля
	orientation int // todo iota
	status      ShipStatus
	size        int
}

func (s StandardShipImpl) GetStatus() ShipStatus {
	deadCount := 0
	for _, status := range s.decks {
		if status == DeadDeck {
			deadCount++
		}
	}
	if deadCount == s.size {
		return Dead
	} else if deadCount == 0 {
		return Alive
	}
	return Hurt
}

func (s StandardShipImpl) HandleShoot(x, y int) bool {
	index := -1 // определяем какая палуба была поражена
	if s.orientation == gorizontal {
		index = x - s.x
	} else {
		index = y - s.y
	}
	if index >= 0 && index <= s.size {
		s.decks[index] = DeadDeck
		return true
	}
	return false
}
