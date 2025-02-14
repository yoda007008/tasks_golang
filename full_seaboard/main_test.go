package main

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		fixture     Board
		name        string
		x           int
		y           int
		size        int
		orientation int
	}{
		{
			fixture:     Board{},
			name:        "Gorizontal",
			x:           3,
			y:           4,
			size:        3,
			orientation: gorizontal,
		}, {
			fixture:     Board{},
			name:        "Vertical",
			x:           4,
			y:           4,
			size:        3,
			orientation: vertical,
		}, {
			fixture:     Board{},
			name:        "Lower left corner",
			x:           0,
			y:           0,
			size:        1,
			orientation: gorizontal,
		}, {
			fixture:     Board{},
			name:        "Upper right corner",
			x:           10,
			y:           10,
			size:        1,
			orientation: gorizontal,
		}, {
			fixture:     Board{},
			name:        "Lower right corner",
			x:           10,
			y:           0,
			size:        1,
			orientation: gorizontal,
		},
	}
	for _, ts := range cases {
		ts := ts
		t.Run(ts.name, func(t *testing.T) {
			t.Parallel()
		})
	}
}
