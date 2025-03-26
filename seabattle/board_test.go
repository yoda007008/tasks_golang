// seabattle/board_test.go
package main

import (
	"github.com/golang/mock/gomock"
	"io"
	"os"
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
	//board := realization.NewBoard().(*realization.BoardImpl)
	board := realization.BoardImpl{}

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

	oldStdout := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	board.PrintField()

	w.Close()
	os.Stdout = oldStdout

	out, _ := io.ReadAll(r)
	actualOutput := string(out)

	// ожидаемый результат
	expectedOutput := `X S O . . . . . . . 
. . . . . . . . . . 
. . . . . . . . . . 
. . . . . . . . . . 
. . . . . . . . . . 
. . . . . . . . . . 
. . . . . . . . . . 
. . . . . . . . . . 
. . . . . . . . . . 
. . . . . . . . . . 
`
	// сравнение
	if actualOutput != expectedOutput {
		t.Errorf("\nОжидаемый вывод:\n%s\nФактический вывод:\n%s",
			expectedOutput, actualOutput)
	}
}

func TestHandleShoot(t *testing.T) {
	// Тест на промах
	t.Run("Промах", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCell := mocks.NewMockCell(ctrl)
		board := realization.NewBoard().(*realization.BoardImpl)
		board.Board2d[0][0] = mockCell

		mockCell.EXPECT().SetStatus(true)
		mockCell.EXPECT().GetShip().Return(nil) // Клетка пуста

		result := board.HandleShoot(0, 0)
		if result != "Мимо" {
			t.Errorf("Expected 'Мимо', got '%s'", result)
		}
	})

	// попадание
	t.Run("Попадание", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockCell := mocks.NewMockCell(ctrl)
		mockShip := mocks.NewMockShip(ctrl)

		// подмена клетки
		board := realization.NewBoard().(*realization.BoardImpl)
		board.Board2d = [10][10]realization.Cell{} // обнуление доски
		board.Board2d[0][0] = mockCell

		mockCell.EXPECT().GetShip().
			Return(mockShip).
			AnyTimes()

		mockCell.EXPECT().SetStatus(true).Times(1)
		mockShip.EXPECT().HandleShoot(0, 0).Return(true)
		mockShip.EXPECT().GetStatus().Return(realization.Alive)

		result := board.HandleShoot(0, 0)
		if result != "Попал" {
			t.Errorf("Expected 'Попал', got '%s'", result)
		}
	})

	t.Run("Ошибка координат", func(t *testing.T) {
		board := realization.NewBoard()

		testCases := []struct {
			x, y     int
			expected string
		}{
			{-1, 0, "Ошибка координат"},
			{10, 5, "Ошибка координат"},
			{0, -1, "Ошибка координат"},
			{5, 10, "Ошибка координат"},
		}

		for _, tc := range testCases {
			result := board.HandleShoot(tc.x, tc.y)
			if result != tc.expected {
				t.Errorf("Для (%d,%d) ожидалось '%s', получено '%s'",
					tc.x, tc.y, tc.expected, result)
			}
		}
	})
}

//func TestPlaceShipCoordsAndCanPlaced(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	// создание mock ячеек для доски
//	mockCells := make([][]*mocks.MockCell, 10)
//	for i, _ := range mockCells {
//		mockCells[i] = make([]*mocks.MockCell, 10)
//		for j, _ := range mockCells {
//			mockCells[i][j] = mocks.NewMockCell(ctrl)
//		}
//	}
//
//	// создание доски
//	board := &realization.BoardImpl{
//		Board2d: mockCells,
//		Ships:   make([]*realization.StandardShipImpl, 0),
//	}
//
//	// заменяем функцию canPlaced на более удобное применение
//	originalCanPlaced := realization.BoardImpl{}.CanPlaced
//	defer func() { realization.BoardImpl{}.CanPlaced() = originalCanPlaced}
//
//}
