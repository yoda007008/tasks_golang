// seabattle/board_test.go
package main

import (
	"github.com/golang/mock/gomock"
	"seabattle/seabattle/mocks"
	"seabattle/seabattle/realization"
	"testing"
)

func TestPrintField(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// создаем моки
	mockShip := mocks.NewMockShip(ctrl)
	mockCell1 := mocks.NewMockCell(ctrl)
	mockCell2 := mocks.NewMockCell(ctrl)
	mockCell3 := mocks.NewMockCell(ctrl)

	// настраиваем ожидания
	mockCell1.EXPECT().GetStatus().Return(true)
	mockCell1.EXPECT().GetShip().Return(mockShip)

	mockCell2.EXPECT().GetStatus().Return(false)
	mockCell2.EXPECT().GetShip().Return(mockShip)

	mockCell3.EXPECT().GetStatus().Return(true)
	mockCell3.EXPECT().GetShip().Return(nil)

	// создаем тестовую доску
	board := realization.NewBoard().(*realization.BoardImpl)

	// заполняем поле
	board.Board2d[0][0] = mockCell1
	board.Board2d[0][1] = mockCell2
	board.Board2d[0][2] = mockCell3

	// остальные клетки заполняем реальными CellImpl
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 0 && j < 3 {
				continue
			}
			board.Board2d[i][j] = &realization.CellImpl{}
		}
	}

	// тестируем PrintField()
	board.PrintField()
}
func TestHandleShoot(t *testing.T) {

}

func TestPlaceShipCoords(t *testing.T) {

}
