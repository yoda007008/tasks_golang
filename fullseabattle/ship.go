package main

type ShipStatus int

const (
	Alive ShipStatus = iota // цел
	Hurt                    // ранен
	Dead                    // мёртв
)

type DeckStatus int

type Ship interface {
	GetStatus() ShipStatus
	HandleShoot(x int, y int) bool
}

type StandardShipImpl struct {
	decks       []ShipStatus
	x           int // todo getX, координата начальной точки корабля
	y           int // todo getY, координата начальной точки корабля
	orientation int // todo iota
	status      ShipStatus
	size        int
}

func (s StandardShipImpl) GetStatus() ShipStatus {
	deadCount := 0
	for _, status := range s.decks {
		if status == Dead {
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

func (s *StandardShipImpl) HandleShoot(x, y int) bool {
	if s.orientation == gorizontal {
		if y != s.y || x < s.x || x >= s.x+s.size {
			return false
		}
		s.decks[x-s.x] = Dead
	} else {
		if x != s.x || y < s.y || y >= s.y+s.size {
			return false
		}
		s.decks[y-s.y] = Dead
	}
	return true
}

func NewShip() Ship {
	return &StandardShipImpl{}
}
