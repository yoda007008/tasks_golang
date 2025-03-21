package interfaces

type Cell interface {
	GetStatus() bool
	SetStatus(status bool)
	GetShip() Ship
	SetShip(ship Ship)
}

type CellImpl struct {
	status bool
	ship   Ship
}

func (c *CellImpl) GetStatus() bool {
	return c.status
}

func (c *CellImpl) SetStatus(status bool) {
	c.status = status
}

func (c *CellImpl) GetShip() Ship {
	return c.ship
}

func (c *CellImpl) SetShip(ship Ship) {
	c.ship = ship
}
