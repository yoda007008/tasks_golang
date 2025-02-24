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
	HandleShoot()
}

type StandardShipImpl struct {
	decks       []DeckStatus
	x           int // todo getX, координата начальной точки корабля
	y           int // todo getY, координата начальной точки корабля
	orientation int // todo iota
	size        int
}

func (s StandardShipImpl) GetStatus() ShipStatus {
	for _, stat := range s.decks {
		switch stat {
		case AliveDeck:
			return Alive
		case DeadDeck:
			return Dead
		}
	}
	return Hurt // в противном случае hurt
}

func (s StandardShipImpl) HandleShoot() {
	//TODO implement me
	panic("implement me")
}
