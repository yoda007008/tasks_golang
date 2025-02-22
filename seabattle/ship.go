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
	//TODO implement me
	panic("implement me")
}

func (s StandardShipImpl) HandleShoot() {
	//TODO implement me
	panic("implement me")
}

// пример другой реализации корабля
type ArmoredShipImpl struct {
}

func (s ArmoredShipImpl) GetStatus() ShipStatus {
	//TODO implement me
	panic("implement me")
}

func (s ArmoredShipImpl) HandleShoot() {
	//TODO implement me
	panic("implement me")
}

func NewArmoredShip() Ship {
	return &ArmoredShipImpl{}
}
