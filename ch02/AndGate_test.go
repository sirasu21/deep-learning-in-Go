package ch02

import (
	"testing"
)

func TestAddGate(t *testing.T) {
	tests := []struct {
		x1, x2 float64
		want   int
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
	}
	for _, test := range tests {
		got := AndGate(test.x1, test.x2)
		if got != test.want {
			t.Errorf("AddGate(%f, %f) == %d, want %d", test.x1, test.x2, got, test.want)
		}
	}

}
