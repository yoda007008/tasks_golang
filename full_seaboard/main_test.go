package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	cases := []struct {
		name        string
		x           int
		y           int
		size        int
		orientation int
		expected    bool
	}{
		{
			// todo заполнить поле данными например точками, тестовое поле "сделано"
			name:        "Gorizontal",
			x:           3,
			y:           3,
			size:        3,
			orientation: gorizontal,
			expected:    true, // todo expected во всех кейсах
		}, {
			name:        "Vertical",
			x:           4,
			y:           4,
			size:        3,
			orientation: vertical,
			expected:    true,
		}, {
			name:        "Lower left corner",
			x:           0,
			y:           0,
			size:        1,
			orientation: gorizontal,
			expected:    true,
		}, {
			name:        "Upper right corner",
			x:           9,
			y:           9,
			size:        1,
			orientation: gorizontal,
			expected:    true,
		}, {
			name:        "Lower right corner",
			x:           9,
			y:           0,
			size:        1,
			orientation: gorizontal,
			expected:    true,
		},
	}
	for _, ts := range cases {
		ts := ts
		t.Run(ts.name, func(t *testing.T) {
			// todo заполнил поле точками
			var board Board
			for i := 0; i < boardSize; i++ {
				for j := 0; j < boardSize; j++ {
					board[i][j] = "."
				}
			}
			factRes := isValid(ts.x, ts.y, ts.size, orientation(ts.orientation), board)
			if factRes != ts.expected {
				t.Errorf("Expected %v but was %v", ts.expected, factRes)
			}
		})
	}
}

func TestPlaceShip(t *testing.T) {

	cases := []struct {
		name     string
		size     int
		expected bool
	}{
		{
			name:     "Center 1x1",
			size:     1,
			expected: true,
		},
		{
			name:     "Create ship 3x1 gorizont",
			size:     3,
			expected: true,
		},
		{
			name:     "Create ship 1x3 vertical",
			size:     3,
			expected: true,
		},
		{
			name:     "Place ship 10x1",
			size:     10,
			expected: true,
		},
		{
			name:     "Place ship 1x10",
			size:     10,
			expected: true,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var board Board
			for i := 0; i < boardSize; i++ {
				for j := 0; j < boardSize; j++ {
					board[i][j] = "."
				}
			}
			placeRes := PlaceShip(&board, tc.size)
			if placeRes != tc.expected {
				t.Errorf("Await %v, but was %v in test '%s'", tc.expected, placeRes, tc.name)
			}
		})
	}
}
