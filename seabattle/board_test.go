// seabattle/board_test.go
package main

import (
	"github.com/golang/mock/gomock"
	"io"
	"os"
	"seabattle/seabattle/interfaces"
	"seabattle/seabattle/mocks"
	"strings"
	"testing"
)

func TestPrintField(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. Создаем полную доску с моками
	var boardArray [10][10]interfaces.Cell

	// 2. Заполняем ВСЕ ячейки моками с базовыми ожиданиями
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			mockCell := mocks.NewMockCell(ctrl)

			// Обязательные ожидания для всех ячеек
			mockCell.EXPECT().GetStatus().Return(false).AnyTimes() // Добавлено!
			mockCell.EXPECT().GetShip().Return(nil).AnyTimes()

			boardArray[i][j] = mockCell
		}
	}

	// 3. Настраиваем специальные ячейки:

	// Ячейка с подбитым кораблём (0,0)
	mockShip1 := mocks.NewMockShip(ctrl)
	mockCell1 := boardArray[0][0].(*mocks.MockCell)
	mockCell1.EXPECT().GetShip().Return(mockShip1).AnyTimes()
	mockCell1.EXPECT().GetStatus().Return(true).AnyTimes() // Подбитая

	// Ячейка с целым кораблём (0,2)
	mockShip2 := mocks.NewMockShip(ctrl)
	mockCell2 := boardArray[0][2].(*mocks.MockCell)
	mockCell2.EXPECT().GetShip().Return(mockShip2).AnyTimes()
	// GetStatus уже установлен в false для всех ячеек

	// 4. Создаем доску
	board := interfaces.BoardImpl{
		Board2d: boardArray,
	}

	// 5. Захватываем вывод
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// 6. Вызываем метод
	board.PrintField()

	// 7. Восстанавливаем и проверяем вывод
	w.Close()
	os.Stdout = oldStdout
	out, _ := io.ReadAll(r)

	// 8. Проверяем вывод
	expected := "X . S . . . . . . . \n" +
		strings.Repeat(". . . . . . . . . . \n", 9)

	if string(out) != expected {
		t.Errorf("Unexpected output:\nGot:\n%s\nExpected:\n%s", string(out), expected)
	}
}
func TestHandleShoot(t *testing.T) {

}

func TestPlaceShipCoords(t *testing.T) {

}

//ctrl := gomock.NewController(t)
//defer ctrl.Finish()
//
//// Создаем мок-объекты для Cell и Ship
//mockCell1 := mocks.NewMockCell(ctrl)
//mockCell2 := mocks.NewMockCell(ctrl)
//mockCell3 := mocks.NewMockCell(ctrl)
//
//mockShip1 := mocks.NewMockShip(ctrl)
//mockShip2 := mocks.NewMockShip(ctrl)
//
//// Настраиваем мок-объекты для Cell
//mockCell1.EXPECT().GetStatus().Return(true).AnyTimes()
//mockCell1.EXPECT().GetShip().Return(mockShip1).AnyTimes()
//
//mockCell2.EXPECT().GetStatus().Return(false).AnyTimes()
//mockCell2.EXPECT().GetShip().Return(nil).AnyTimes()
//
//mockCell3.EXPECT().GetStatus().Return(false).AnyTimes()
//mockCell3.EXPECT().GetShip().Return(mockShip2).AnyTimes()
//
//// Создаем доску и подставляем мок-объекты
//var boardArray [10][10]interfaces.Cell
//boardArray[0][0] = mockCell1
//boardArray[0][1] = mockCell2
//boardArray[0][2] = mockCell3
//// Заполните остальные ячейки по аналогии
//
//board := interfaces.BoardImpl{
//	Board2d: boardArray,
//}
//
//// Вызываем тестируемую функцию
//board.PrintField()
