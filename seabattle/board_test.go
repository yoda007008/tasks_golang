package main

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMockBoard_HandleShoot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBoard := NewMockBoard(ctrl)
	mockBoard.EXPECT().HandleShoot(1, 1).Return("Miss")

	result := mockBoard.HandleShoot(1, 1)
	if result != "Miss" {
		fmt.Errorf("Ожидалось Miss, полученно %s", result)
	}
}

func TestMockBoard_PlaceShipCoords(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBoard := NewMockBoard(ctrl)
	mockBoard.EXPECT().PlaceShipCoords(2, 3, 4, orientation(0)).Return(true)

	success := mockBoard.PlaceShipCoords(2, 3, 4, orientation(0))
	if !success {
		t.Errorf("Ожидалось false, получено true")
	}
}

func TestPrintField(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// мок-объекты для клеток
	mockCell1 := NewMockCell(ctrl)
	mockCell2 := NewMockCell(ctrl)
	mockCell3 := NewMockCell(ctrl)

	mockShip1 := NewMockCell(ctrl)
	mockShip2 := NewMockCell(ctrl)

	// настраиваем мок-объекты
	mockCell1.EXPECT().GetStatus().Return(true).AnyTimes()
	mockCell1.EXPECT().GetShip().Return(mockShip1) // &CellImpl{}

	mockCell2.EXPECT().GetStatus().Return(false).AnyTimes()
	mockCell2.EXPECT().GetShip().Return(mockShip1).AnyTimes()

	mockCell3.EXPECT().GetStatus().Return(false).AnyTimes()
	mockCell3.EXPECT().GetShip().Return(mockShip2) // &CellImpl{}

	board := BoardImpl{
		Board2d: [10][10]Cell{
			{mockCell1, mockCell2, mockCell3},
			{mockCell2, mockCell3, mockCell1},
			{mockCell3, mockCell2, mockCell1},
		},
	}
	board.PrintField()
}

func TestMockCell_GetStatus_SetStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCell := NewMockCell(ctrl)

	// ожидаем, что SetStatus будет вызван с аргументом true
	mockCell.EXPECT().SetStatus(true)

	mockCell.SetStatus(true)

	// ожидаем, что GetStatus вернет true
	mockCell.EXPECT().GetStatus().Return(true)

	status := mockCell.GetStatus()
	if !status {
		t.Errorf("Ожидалось true, получено false")
	}
}
