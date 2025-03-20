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

func TestMockBoard_PrintField(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBoard := NewMockBoard(ctrl)
	mockBoard.EXPECT().PrintField()
	mockBoard.PrintField()
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
