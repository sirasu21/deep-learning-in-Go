package ch02

import "testing"

func TestOrGate(t *testing.T) {
	type args struct {
		x1 float64
		x2 float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{0, 0}, 0},
		{"test2", args{0, 1}, 1},
		{"test3", args{1, 0}, 1},
		{"test4", args{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrGate(tt.args.x1, tt.args.x2); got != tt.want {
				t.Errorf("OrGate() = %v, want %v", got, tt.want)
			}
		})
	}
}
