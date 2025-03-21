// seabattle/board_test.go
package main

import (
	"seabattle/seabattle/interfaces"
	"testing"

	"github.com/golang/mock/gomock"
	"seabattle/seabattle/mocks"
)

func TestPrintField(t *testing.T) {
	// Инициализация контроллера gomock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок-объекты для Cell и Ship
	mockCell1 := mocks.NewMockCell(ctrl)
	mockCell2 := mocks.NewMockCell(ctrl)
	mockCell3 := mocks.NewMockCell(ctrl)

	mockShip1 := mocks.NewMockShip(ctrl)
	mockShip2 := mocks.NewMockShip(ctrl)

	// Настраиваем мок-объекты для Cell
	mockCell1.EXPECT().GetStatus().Return(true).AnyTimes()
	mockCell1.EXPECT().GetShip().Return(mockShip1).AnyTimes()

	mockCell2.EXPECT().GetStatus().Return(false).AnyTimes()
	mockCell2.EXPECT().GetShip().Return(nil).AnyTimes()

	mockCell3.EXPECT().GetStatus().Return(false).AnyTimes()
	mockCell3.EXPECT().GetShip().Return(mockShip2).AnyTimes()

	// Создаем доску и подставляем мок-объекты
	var boardArray [10][10]interfaces.Cell
	boardArray[0][0] = mockCell1
	boardArray[0][1] = mockCell2
	boardArray[0][2] = mockCell3
	// Заполните остальные ячейки по аналогии

	board := interfaces.BoardImpl{
		Board2d: boardArray,
	}

	// Вызываем тестируемую функцию
	board.PrintField()
}
