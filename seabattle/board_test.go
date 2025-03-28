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
	// промах
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

	// уничтожение корабля
	//t.Run("Уничтожение корабля", func(t *testing.T) {
	//	ctrl := gomock.NewController(t)
	//	ctrl.Finish()
	//
	//	mockCell := mocks.NewMockCell(ctrl)
	//	mockShip := mocks.NewMockShip(ctrl)
	//
	//	board := realization.NewBoard().(*realization.BoardImpl)
	//	board.Board2d
	//	board.Board2d[0][0] = mockCell
	//
	//	mockCell.EXPECT().SetStatus(true).Times(1)
	//	mockCell.EXPECT().GetShip().Return(mockShip).AnyTimes()
	//	mockCell.EXPECT().HandleShoot(0, 0).Return(true)
	//	mockCell.EXPECT().GetStatus().Return(realization.Dead)
	//
	//	result := board.HandleShoot(0, 0)
	//	if result != "Корабль уничтожен" {
	//		t.Errorf("Expected 'Корабль уничтожен' got '%s'", result)
	//	}
	//})

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

		mockCell.EXPECT().GetShip().Return(mockShip).AnyTimes()

		mockCell.EXPECT().SetStatus(true).Times(1)
		mockShip.EXPECT().HandleShoot(0, 0).Return(true)
		mockShip.EXPECT().GetStatus().Return(realization.Alive)

		result := board.HandleShoot(0, 0)
		if result != "Попал" {
			t.Errorf("Expected 'Попал', got '%s'", result)
		}
	})

	// ошибка координат
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

func TestCanPlaced(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		gorizontal = iota
		vertical
	)

	// Создаем mock-ячейки для всей доски
	var mockCells [10][10]*mocks.MockCell
	var boardCells [10][10]realization.Cell

	for i := range mockCells {
		for j := range mockCells[i] {
			mockCells[i][j] = mocks.NewMockCell(ctrl)
			boardCells[i][j] = mockCells[i][j]
		}
	}

	board := &realization.BoardImpl{
		Board2d: boardCells,
	}

	t.Run("Горизонтальный корабль в пределах доски", func(t *testing.T) {
		for i := 0; i < 3; i++ {
			mockCells[0][i].EXPECT().GetShip().Return(nil).AnyTimes()
		}

		// Соседние ячейки:
		// Верхний ряд (y=1)
		for i := -1; i <= 3; i++ {
			if i >= 0 && i < 10 {
				mockCells[1][i].EXPECT().GetShip().Return(nil).AnyTimes()
			}
		}

		// боковые ячейки (x=-1 и x=3)
		if 3 < 10 {
			mockCells[0][3].EXPECT().GetShip().Return(nil).AnyTimes()
		}
		// x=-1 - за границей, не мокаем

		mockCells[1][1].EXPECT().GetShip().Return(nil).AnyTimes()
		mockCells[1][3].EXPECT().GetShip().Return(nil).AnyTimes()

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				// Пропускаем уже замоканые ячейки
				if (i == 0 && j >= 0 && j < 3) ||
					(i == 1 && j >= -1 && j <= 3) ||
					(i == 0 && j == 3) {
					continue
				}
				mockCells[i][j].EXPECT().GetShip().Return(nil).AnyTimes()
			}
		}

		result := board.CanPlaced(0, 0, 3, realization.Orientation(gorizontal))

		if !result {
			t.Error("Ожидалось успешное размещение корабля")
		}
	})

	t.Run("Вертикальный корабль в пределах доски", func(t *testing.T) {

		// основные ячейки корабля (2,2)-(4,2)
		for i := 2; i < 5; i++ {
			mockCells[i][2].EXPECT().GetShip().Return(nil).AnyTimes()
		}

		// левый ряд (x=1)
		for j := 1; j <= 3; j++ {
			mockCells[1][j].EXPECT().GetShip().Return(nil).AnyTimes()
		}

		// правый ряд (x=5)
		for j := 1; j <= 3; j++ {
			mockCells[5][j].EXPECT().GetShip().Return(nil).AnyTimes()
		}

		// верхняя граница (y=1)
		for i := 1; i <= 5; i++ {
			mockCells[i][1].EXPECT().GetShip().Return(nil).AnyTimes()
		}

		// нижняя граница (y=3)
		for i := 1; i <= 5; i++ {
			mockCells[i][3].EXPECT().GetShip().Return(nil).AnyTimes()
		}

		// диагональные углы
		mockCells[1][1].EXPECT().GetShip().Return(nil).AnyTimes()
		mockCells[1][3].EXPECT().GetShip().Return(nil).AnyTimes()
		mockCells[5][1].EXPECT().GetShip().Return(nil).AnyTimes()
		mockCells[5][3].EXPECT().GetShip().Return(nil).AnyTimes()

		// замокаем все остальные ячейки как свободные
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if (i >= 2 && i <= 4 && j == 2) ||
					(i >= 1 && i <= 5 && j >= 1 && j <= 3) {
					continue
				}
				mockCells[i][j].EXPECT().GetShip().Return(nil).AnyTimes()
			}
		}

		result := board.CanPlaced(2, 2, 3, realization.Orientation(vertical))

		if !result {
			t.Error("Ожидалось успешное вертикальное размещение корабля")
		}
	})

	t.Run("Горизонтальное размещение корабля у правого края", func(t *testing.T) { // запускаем подтест с описанием горизонтальное размещение корабля у правого края
		ctrl.Finish()                   // завершаем работу предыдущего контроллера
		ctrl := gomock.NewController(t) // инициализация мок-контроллера
		defer ctrl.Finish()             // данный вывов гаранитирует отчистку после завершения теста

		// полная переинициализация всех ячеек
		var mockCells [10][10]*mocks.MockCell   // инициализируем тестируемые клетки
		var boardCells [10][10]realization.Cell // инициализируем тестовое поле
		for i := range mockCells {              // данный цикл заполняет поле тестируемыми клетками
			for j := range mockCells[i] {
				mockCells[i][j] = mocks.NewMockCell(ctrl) // создаем mock-ячейку
				boardCells[i][j] = mockCells[i][j]        // присваиваем mock-ячейку ячейке доски
			}
		}

		// инициализируем готовое к тестам поле
		board := &realization.BoardImpl{Board2d: boardCells}

		x, y, size := 7, 0, 3 // координаты корабля, который находится у правого края (size 3, x = 7, y = 0)

		// настраиваем ожидания для основных ячеек корабля
		for i := x; i < x+size; i++ { // корабль горизонтальный, поэтому меняем x(i) в диапазоне от 7 до 9 (7+3=10)
			mockCells[0][i].EXPECT().GetShip().Return(nil).AnyTimes() // вызываем метод GetShip, который вернет статус nil, что корабля действительно нет
		}

		// тут тестирую соседние ячейки и проверяю не выодят ли они за границу поля
		for i := x - 1; i <= x+size; i++ { // прохожусь по ячейкам
			if i >= 0 && i < 10 { // проверка не выходит ли за границу поля
				mockCells[1][i].EXPECT().GetShip().Return(nil).AnyTimes() // вызываем метод GetShip, который вернет статус nil, что корабля действительно нет
			}
		}

		// настройка отдельных левых ячеек слева
		mockCells[0][6].EXPECT().GetShip().Return(nil).AnyTimes() // левый сосед основной ячейки [0][6]
		mockCells[1][6].EXPECT().GetShip().Return(nil).AnyTimes() // левый сосед нижней ячейки [1][6]

		// замокаем все остальные ячейки как свободные
		for i := 0; i < 10; i++ { // проходимся по row
			for j := 0; j < 10; j++ { // проходимся по col
				// пропуск ячеек, которые уже настроены
				if (i == 0 && j >= 7 && j <= 9) || // основные ячейки
					(i == 1 && j >= 6 && j <= 9) { // нижний ряд
					continue
				}
				mockCells[i][j].EXPECT().GetShip().Return(nil).AnyTimes() // для всех остальных ожидаем метод GetShip(), который вернет ni
			}
		}

		// тестируем CanPlaced, если результат не совпадает с ожидаемым, то выводим ошибку
		result := board.CanPlaced(x, y, size, realization.Orientation(gorizontal))

		if !result { // если ожидается не result, то выводим ошибку
			t.Error("Ожидалось успешное размещение у правого края") // вывод ошибки
		}
	})

	t.Run("Вертикальный корабль в левом верхнем углу", func(t *testing.T) {
		ctrl.Finish()                   // завершаем работу предыдущего контроллера
		ctrl := gomock.NewController(t) // инициализация мок-контроллера
		defer ctrl.Finish()             // данный вывов гаранитирует отчистку после завершения теста

		var mockCells [10][10]*mocks.MockCell   // инициализируем тестированные мок клетки
		var boardCells [10][10]realization.Cell // инициализируем тестируемое поле
		for i := range mockCells {              // данный цикл заполняет тестируемое поле мок-клетками
			for j := range mockCells[i] {
				mockCells[i][j] = mocks.NewMockCell(ctrl) // создаем фейковую клетку игрового поля
				boardCells[i][j] = mockCells[i][j]        // присваем ее игровому полю
			}
		}

		board := &realization.BoardImpl{Board2d: boardCells} // инициализация заполненнного тестируемого поля

		x, y, size := 0, 0, 2 // координаты корабля в левом верхнем углу (size = 2, x = 0, y = 0)

		// настройка размещения ключевых ячеек корабля y(i) от 0 до 1 (0+2=2)
		for i := y; i < y+size; i++ {
			mockCells[i][x].EXPECT().GetShip().Return(nil).AnyTimes() // вызов GetShip=клетка свободна nil
		}

		// тестируем соседние ячейки при помощи цикла
		for j := y; j < y+size; j++ {
			if j < 10 { // проверка, чтобы не выйти за нижнюю границу поля
				mockCells[j][1].EXPECT().GetShip().Return(nil).AnyTimes() // аналогично настройка соседних клеток справа
			}
		}

		// левый бок при y = 2
		mockCells[2][0].EXPECT().GetShip().Return(nil).AnyTimes() // левая клетка под кораблем (0, 2)
		mockCells[2][1].EXPECT().GetShip().Return(nil).AnyTimes() // правая клетка под кораблем (2, 1)

		// замокаем все остальные ячейки как свободные
		for i := 0; i < 10; i++ { // проходимся по row
			for j := 0; j < 10; j++ { // проходимся по col
				if (i >= 0 && i < 2 && j == 0) || // основные ячейки
					(i >= 0 && i <= 2 && j == 1) { // соседние
					continue
				}
				mockCells[i][j].EXPECT().GetShip().Return(nil).AnyTimes() // все остальные клетки тоже свободные
			}
		}

		// тестируем CanPlaced, если результат не совпадает с ожидаемым, то выводим ошибку
		result := board.CanPlaced(x, y, size, realization.Orientation(vertical))
		if !result { // если не result
			t.Error("Ожидалось размещение в левом верхнем углу") // возвращаем ошибку
		}
	})
}
