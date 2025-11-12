package ch03

import "testing"

func TestStepFunction(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{0}, 0},
		{"test2", args{1}, 1},
		{"test3", args{-1}, 0},
		{"test4", args{2}, 1},
		{"test5", args{-2}, 0},
		{"test6", args{0.5}, 1},
		{"test7", args{-0.5}, 0},
		{"test8", args{1.5}, 1},
		{"test9", args{-1.5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StepFunction(tt.args.x); got != tt.want {
				t.Errorf("StepFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
