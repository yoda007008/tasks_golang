package main

import "fmt"

type Player interface {
	DoMove(otherPlayers []Player) error
	TakeMove(x, y int)
	GiveUp()
}

// как вариант использовать интерфейсы: сделать мок реализацию игрока с методом DoMove(), которая будет просто выводить: "Игрок [id] делает ход"
type MockPlayer struct {
	id int
}

func (p *MockPlayer) DoMove(otherPlayers []Player) error {
	fmt.Println("Игрок ", p.id, " делает ход")
	return nil
}

type PlayerImpl struct {
	id    int
	board Board
}

func (p *PlayerImpl) DoMove(otherPlayers []Player) error { // todo не знаем про boardSize ? board.Size()
	//TODO implement me
	fmt.Printf("Игрок %d ваш ход. Введите координату выстрела x\n", p.id)

	// todo выделить в отдельную структуру координату, на ней сделать метод New() error, и возвращать там ошибку

	// todo сделать пользовательский ввод, пока он не введёт корректные координаты
	var x, y int

	fmt.Scanln(&x)

	fmt.Printf("Игрок %d ваш ход. Введите координату выстрела y\n", p.id)

	fmt.Scanln(&y)

	for _, p := range otherPlayers {
		p.TakeMove(x, y)
	}

	panic("implement me")
}

// TakeMove принимает и обрабатывает выстрел
//func (p *PlayerImpl) TakeMove(x, y int) {
//	//TODO implement me
//	p.board.HandleShoot(x, y)
//	panic("implement me")
//}

func (p *PlayerImpl) GiveUp() {
	//TODO implement me
	panic("implement me")
}

//func NewPlayer() Player {
//	return &PlayerImpl{}
//}

// todo
func validateInput() error {
	//TODO implement me
	panic("implement me")
}
