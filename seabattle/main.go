package main

import "fmt"

const (
	fourShip   = 4
	thirdShip  = 3
	secondShip = 2
	firstShip  = 1
)

type Field struct {
	Board [boardSize][boardSize]string
}

type Cord struct {
	x int
	y int
}

type CreateShip struct {
	size   int
	isDead bool
	cord   []Cord
}

type Human struct {
	Name string
	Age  int
}

var person = Human{Age: 18, Name: "Kirill"}

type HumanPlayer struct {
	board *Field
}

func NewField() *Field { // заполнять и создавать поля нужно также, как и в предыдущей программе морского боя
	f := &Field{}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			f.Board[i][j] = "."
		}
	}
	return f
}

func (f *Field) Print() { // вывод поля
	fmt.Println("  0 1 2 3 4 5 6 7 8 9 10")
	for i, row := range f.Board {
		fmt.Printf("%d", i+1)
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
}

func (f *Field) PlaceShip() { // функция располагает для двух игроков корабли на поле рандомно

}

func (f *Field) DoMove() { // функция DoMove делает ходы

}

func main() {
	//playersCount := 2
	//game := NewGame(playersCount)
	//
	//for !game.IsEnded() {
	//	game.Round()
	//}
	//var twoD [boardSize][boardSize]Ship
	//board := BoardImpl{Board2d: twoD, Ships: make([]Ship, 0)}
	//fmt.Println(board.String())

	s := make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
}
